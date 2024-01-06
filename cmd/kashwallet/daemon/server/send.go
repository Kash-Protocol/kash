package server

import (
	"context"
	"github.com/Kash-Protocol/kashd/cmd/kashwallet/daemon/pb"
	"github.com/Kash-Protocol/kashd/domain/consensus/model/externalapi"
	"github.com/pkg/errors"
)

func (s *server) Send(_ context.Context, request *pb.SendRequest) (*pb.SendResponse, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	// TODO: This mapping of AssetType to DomainTransactionType is a temporary solution.
	// It allows wallet usage of the 'send' command for transfers. Currently, AssetType
	// directly corresponds to a specific Transfer TxType. This may evolve with more complex
	// transaction types in the future.
	var txType externalapi.DomainTransactionType
	switch externalapi.AssetTypeFromUint32(request.AssetType) {
	case externalapi.KSH:
		txType = externalapi.TransferKSH
	case externalapi.KUSD:
		txType = externalapi.TransferKUSD
	case externalapi.KRV:
		txType = externalapi.TransferKRV
	default:
		return nil, errors.Errorf("Unknown asset type %d", request.AssetType)
	}

	unsignedTransactions, err := s.createUnsignedTransactions(request.ToAddress, txType,
		request.Amount, request.IsSendAll, request.From, request.UseExistingChangeAddress)

	if err != nil {
		return nil, err
	}

	signedTransactions, err := s.signTransactions(unsignedTransactions, request.Password)
	if err != nil {
		return nil, err
	}

	txIDs, err := s.broadcast(signedTransactions, false)
	if err != nil {
		return nil, err
	}

	return &pb.SendResponse{TxIDs: txIDs, SignedTransactions: signedTransactions}, nil
}
