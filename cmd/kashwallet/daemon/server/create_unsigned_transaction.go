package server

import (
	"context"
	"fmt"
	"github.com/Kash-Protocol/kashd/domain/consensus/model/externalapi"
	"time"

	"github.com/Kash-Protocol/kashd/cmd/kashwallet/daemon/pb"
	"github.com/Kash-Protocol/kashd/cmd/kashwallet/libkashwallet"
	"github.com/Kash-Protocol/kashd/domain/consensus/utils/constants"
	"github.com/Kash-Protocol/kashd/util"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

// TODO: Implement a better fee estimation mechanism
const feePerInput = 10000

func (s *server) CreateUnsignedTransactions(_ context.Context, request *pb.CreateUnsignedTransactionsRequest) (
	*pb.CreateUnsignedTransactionsResponse, error,
) {
	s.lock.Lock()
	defer s.lock.Unlock()

	unsignedTransactions, err := s.createUnsignedTransactions(request.Address,
		externalapi.AssetTypeFromUint32(request.AssetType), request.Amount, request.IsSendAll,
		request.From, request.UseExistingChangeAddress)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUnsignedTransactionsResponse{UnsignedTransactions: unsignedTransactions}, nil
}

func (s *server) createUnsignedTransactions(address string, assetType externalapi.AssetType,
	amount uint64, isSendAll bool, fromAddressesString []string, useExistingChangeAddress bool) ([][]byte, error) {
	if !s.isSynced() {
		return nil, errors.Errorf("wallet daemon is not synced yet, %s", s.formatSyncStateReport())
	}

	// make sure address string is correct before proceeding to a
	// potentially long UTXO refreshment operation
	toAddress, err := util.DecodeAddress(address, s.params.Prefix)
	if err != nil {
		return nil, err
	}

	err = s.refreshUTXOs()
	if err != nil {
		return nil, err
	}

	var fromAddresses []*walletAddress
	for _, from := range fromAddressesString {
		fromAddress, exists := s.addressSet[from]
		if !exists {
			return nil, fmt.Errorf("Specified from address %s does not exists", from)
		}
		fromAddresses = append(fromAddresses, fromAddress)
	}

	selectedUTXOs, spendValue, changeSompi, err := s.selectUTXOs(amount, isSendAll, feePerInput, fromAddresses, assetType)
	if err != nil {
		return nil, err
	}

	if len(selectedUTXOs) == 0 {
		return nil, errors.Errorf("couldn't find funds to spend")
	}

	changeAddress, changeWalletAddress, err := s.changeAddress(useExistingChangeAddress, fromAddresses)
	if err != nil {
		return nil, err
	}

	payments := []*libkashwallet.Payment{{
		Address:   toAddress,
		AssetType: assetType,
		Amount:    spendValue,
	}}
	if changeSompi > 0 {
		payments = append(payments, &libkashwallet.Payment{
			Address:   changeAddress,
			AssetType: assetType,
			Amount:    changeSompi,
		})
	}

	// TODO: Add support for MintKUSD, StakeKSH, RedeemKSH.
	TxTypeMapping := map[externalapi.AssetType]externalapi.DomainTransactionType{
		externalapi.KSH:  externalapi.TransferKSH,
		externalapi.KUSD: externalapi.TransferKUSD,
		externalapi.KRV:  externalapi.TransferKRV,
	}

	unsignedTransaction, err := libkashwallet.CreateUnsignedTransaction(s.keysFile.ExtendedPublicKeys,
		s.keysFile.MinimumSignatures,
		payments, selectedUTXOs, TxTypeMapping[assetType])
	if err != nil {
		return nil, err
	}

	unsignedTransactions, err := s.maybeAutoCompoundTransaction(unsignedTransaction, toAddress,
		changeAddress, changeWalletAddress, assetType)
	if err != nil {
		return nil, err
	}
	return unsignedTransactions, nil
}

func (s *server) selectUTXOs(spendAmount uint64, isSendAll bool, feePerInput uint64, fromAddresses []*walletAddress, assetType externalapi.AssetType) (
	selectedUTXOs []*libkashwallet.UTXO, totalReceived uint64, changeSompi uint64, err error) {

	selectedUTXOs = []*libkashwallet.UTXO{}
	totalValue := uint64(0)

	dagInfo, err := s.rpcClient.GetBlockDAGInfo()
	if err != nil {
		return nil, 0, 0, err
	}

	for _, utxo := range s.utxosSortedByAmount {
		if utxo.UTXOEntry.AssetType() != assetType ||
			(fromAddresses != nil && !slices.Contains(fromAddresses, utxo.address)) ||
			!isUTXOSpendable(utxo, dagInfo.VirtualDAAScore, s.params.BlockCoinbaseMaturity) {
			continue
		}

		if broadcastTime, ok := s.usedOutpoints[*utxo.Outpoint]; ok {
			if time.Since(broadcastTime) > time.Minute {
				delete(s.usedOutpoints, *utxo.Outpoint)
			} else {
				continue
			}
		}

		selectedUTXOs = append(selectedUTXOs, &libkashwallet.UTXO{
			Outpoint:       utxo.Outpoint,
			UTXOEntry:      utxo.UTXOEntry,
			DerivationPath: s.walletAddressPath(utxo.address),
		})

		totalValue += utxo.UTXOEntry.Amount()

		fee := feePerInput * uint64(len(selectedUTXOs))
		totalSpend := spendAmount + fee
		if !isSendAll && totalValue >= totalSpend {
			break
		}
	}

	fee := feePerInput * uint64(len(selectedUTXOs))
	var totalSpend uint64
	if isSendAll {
		totalSpend = totalValue
		totalReceived = totalValue - fee
	} else {
		totalSpend = spendAmount + fee
		totalReceived = spendAmount
	}
	if totalValue < totalSpend {
		return nil, 0, 0, errors.Errorf("Insufficient funds for send: %f required, while only %f available",
			float64(totalSpend)/constants.SompiPerKash, float64(totalValue)/constants.SompiPerKash)
	}

	return selectedUTXOs, totalReceived, totalValue - totalSpend, nil
}
