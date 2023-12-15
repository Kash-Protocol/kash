package server

import (
	"context"
	"github.com/Kash-Protocol/kashd/domain/consensus/model/externalapi"

	"github.com/Kash-Protocol/kashd/cmd/kashwallet/daemon/pb"
	"github.com/Kash-Protocol/kashd/cmd/kashwallet/libkashwallet"
)

type balancesType struct {
	available, pending uint64
}

type assetBalancesMapType map[*walletAddress]*balancesType

func (s *server) GetBalance(_ context.Context, _ *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	dagInfo, err := s.rpcClient.GetBlockDAGInfo()
	if err != nil {
		return nil, err
	}
	daaScore := dagInfo.VirtualDAAScore
	maturity := s.params.BlockCoinbaseMaturity

	assetBalancesMap := make(map[externalapi.AssetType]assetBalancesMapType)
	for _, entry := range s.utxosSortedByAmount {
		assetType := entry.UTXOEntry.AssetType()
		amount := entry.UTXOEntry.Amount()
		address := entry.address

		if _, ok := assetBalancesMap[assetType]; !ok {
			assetBalancesMap[assetType] = make(assetBalancesMapType)
		}

		balances, ok := assetBalancesMap[assetType][address]
		if !ok {
			balances = new(balancesType)
			assetBalancesMap[assetType][address] = balances
		}

		if isUTXOSpendable(entry, daaScore, maturity) {
			balances.available += amount
		} else {
			balances.pending += amount
		}
	}

	assetBalances := make([]*pb.AssetBalance, 0, len(assetBalancesMap))
	for assetType, balancesMap := range assetBalancesMap {
		addressBalances := make([]*pb.AddressBalances, 0, len(balancesMap))
		var available, pending uint64
		for walletAddress, balances := range balancesMap {
			address, err := libkashwallet.Address(s.params, s.keysFile.ExtendedPublicKeys, s.keysFile.MinimumSignatures, s.walletAddressPath(walletAddress), s.keysFile.ECDSA)
			if err != nil {
				return nil, err
			}
			addressBalances = append(addressBalances, &pb.AddressBalances{
				Address:   address.String(),
				AssetType: assetType.ToPbAssetType(),
				Available: balances.available,
				Pending:   balances.pending,
			})
			available += balances.available
			pending += balances.pending
		}
		assetBalances = append(assetBalances, &pb.AssetBalance{
			AssetType:       assetType.ToPbAssetType(),
			Available:       available,
			Pending:         pending,
			AddressBalances: addressBalances,
		})
	}

	return &pb.GetBalanceResponse{
		AssetBalances: assetBalances,
	}, nil
}

func isUTXOSpendable(entry *walletUTXO, virtualDAAScore uint64, coinbaseMaturity uint64) bool {
	if !entry.UTXOEntry.IsCoinbase() {
		return true
	}
	return entry.UTXOEntry.BlockDAAScore()+coinbaseMaturity < virtualDAAScore
}
