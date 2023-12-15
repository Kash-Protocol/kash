package rpchandlers

import (
	"github.com/Kash-Protocol/kashd/app/appmessage"
	"github.com/Kash-Protocol/kashd/app/rpc/rpccontext"
	"github.com/Kash-Protocol/kashd/domain/consensus/model/externalapi"
	"github.com/Kash-Protocol/kashd/domain/consensus/utils/txscript"
	"github.com/Kash-Protocol/kashd/infrastructure/network/netadapter/router"
	"github.com/Kash-Protocol/kashd/util"
)

// HandleGetMempoolEntriesByAddresses handles the respectively named RPC command
func HandleGetMempoolEntriesByAddresses(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	getMempoolEntriesByAddressesRequest := request.(*appmessage.GetMempoolEntriesByAddressesRequestMessage)
	mempoolEntriesByAddresses := make([]*appmessage.MempoolEntryByAddress, 0)

	sendingInTransactionPool, receivingInTransactionPool, sendingInOrphanPool, receivingInOrphanPool, err := context.Domain.MiningManager().GetTransactionsByAddresses(!getMempoolEntriesByAddressesRequest.FilterTransactionPool, getMempoolEntriesByAddressesRequest.IncludeOrphanPool)
	if err != nil {
		return nil, err
	}

	for _, addressString := range getMempoolEntriesByAddressesRequest.Addresses {
		address, err := util.DecodeAddress(addressString, context.Config.NetParams().Prefix)
		if err != nil {
			errorMessage := &appmessage.GetMempoolEntriesByAddressesResponseMessage{}
			errorMessage.Error = appmessage.RPCErrorf("Could not decode address '%s': %s", addressString, err)
			return errorMessage, nil
		}

		for _, assetType := range []externalapi.AssetType{externalapi.KSH, externalapi.KUSD, externalapi.KRV} {
			scriptPublicKey, err := txscript.PayToAddrScript(address, assetType)
			if err != nil {
				errorMessage := &appmessage.GetMempoolEntriesByAddressesResponseMessage{}
				errorMessage.Error = appmessage.RPCErrorf("Could not extract scriptPublicKey from address '%s' for asset type '%d': %s", addressString, assetType, err)
				return errorMessage, nil
			}

			sending := make([]*appmessage.MempoolEntry, 0)
			receiving := make([]*appmessage.MempoolEntry, 0)

			// Process transactions in the transaction pool
			if !getMempoolEntriesByAddressesRequest.FilterTransactionPool {
				if transaction, found := sendingInTransactionPool[scriptPublicKey.String()]; found {
					rpcTransaction := appmessage.DomainTransactionToRPCTransaction(transaction)
					err := context.PopulateTransactionWithVerboseData(rpcTransaction, nil)
					if err != nil {
						return nil, err
					}
					sending = append(sending, &appmessage.MempoolEntry{
						Fee:         transaction.Fee,
						Transaction: rpcTransaction,
						IsOrphan:    false,
					})
				}

				if transaction, found := receivingInTransactionPool[scriptPublicKey.String()]; found {
					rpcTransaction := appmessage.DomainTransactionToRPCTransaction(transaction)
					err := context.PopulateTransactionWithVerboseData(rpcTransaction, nil)
					if err != nil {
						return nil, err
					}
					receiving = append(receiving, &appmessage.MempoolEntry{
						Fee:         transaction.Fee,
						Transaction: rpcTransaction,
						IsOrphan:    false,
					})
				}
			}

			// Process transactions in the orphan pool
			if getMempoolEntriesByAddressesRequest.IncludeOrphanPool {
				if transaction, found := sendingInOrphanPool[scriptPublicKey.String()]; found {
					rpcTransaction := appmessage.DomainTransactionToRPCTransaction(transaction)
					err := context.PopulateTransactionWithVerboseData(rpcTransaction, nil)
					if err != nil {
						return nil, err
					}
					sending = append(sending, &appmessage.MempoolEntry{
						Fee:         transaction.Fee,
						Transaction: rpcTransaction,
						IsOrphan:    true,
					})
				}

				if transaction, found := receivingInOrphanPool[scriptPublicKey.String()]; found {
					rpcTransaction := appmessage.DomainTransactionToRPCTransaction(transaction)
					err := context.PopulateTransactionWithVerboseData(rpcTransaction, nil)
					if err != nil {
						return nil, err
					}
					receiving = append(receiving, &appmessage.MempoolEntry{
						Fee:         transaction.Fee,
						Transaction: rpcTransaction,
						IsOrphan:    true,
					})
				}
			}

			if len(sending) > 0 || len(receiving) > 0 {
				mempoolEntriesByAddresses = append(
					mempoolEntriesByAddresses,
					&appmessage.MempoolEntryByAddress{
						Address:   address.String(),
						Sending:   sending,
						Receiving: receiving,
					},
				)
			}
		}
	}

	return appmessage.NewGetMempoolEntriesByAddressesResponseMessage(mempoolEntriesByAddresses), nil
}
