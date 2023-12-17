package rpchandlers

import (
	"github.com/Kash-Protocol/kashd/app/appmessage"
	"github.com/Kash-Protocol/kashd/app/rpc/rpccontext"
	"github.com/Kash-Protocol/kashd/domain/consensus/model/externalapi"
	"github.com/Kash-Protocol/kashd/domain/consensus/utils/txscript"
	"github.com/Kash-Protocol/kashd/infrastructure/network/netadapter/router"
	"github.com/Kash-Protocol/kashd/util"
	"github.com/pkg/errors"
)

// HandleGetBalanceByAddress handles the respectively named RPC command
func HandleGetBalanceByAddress(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetUTXOsByAddressesResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when kashd is run without --utxoindex")
		return errorMessage, nil
	}

	getBalanceByAddressRequest := request.(*appmessage.GetBalanceByAddressRequestMessage)

	kshBalance, kusdBalance, krvBalance, err := getBalanceByAddress(context, getBalanceByAddressRequest.Address)

	if err != nil {
		rpcError := &appmessage.RPCError{}
		if !errors.As(err, &rpcError) {
			return nil, err
		}
		errorMessage := &appmessage.GetUTXOsByAddressesResponseMessage{}
		errorMessage.Error = rpcError
		return errorMessage, nil
	}

	response := appmessage.NewGetBalanceByAddressResponse(kshBalance, kusdBalance, krvBalance)
	return response, nil
}

func getBalanceByAddress(context *rpccontext.Context, addressString string) (uint64, uint64, uint64, error) {
	address, err := util.DecodeAddress(addressString, context.Config.ActiveNetParams.Prefix)
	if err != nil {
		return 0, 0, 0, appmessage.RPCErrorf("Couldn't decode address '%s': %s", addressString, err)
	}

	scriptPublicKey, err := txscript.PayToAddrScript(address)
	if err != nil {
		return 0, 0, 0, appmessage.RPCErrorf("Could not create a scriptPublicKey for address '%s': %s", addressString, err)
	}
	utxoOutpointEntryPairs, err := context.UTXOIndex.UTXOs(scriptPublicKey)
	if err != nil {
		return 0, 0, 0, err
	}

	kshBalance := uint64(0)
	kusdBalance := uint64(0)
	krvBalance := uint64(0)

	for _, utxoOutpointEntryPair := range utxoOutpointEntryPairs {
		switch utxoOutpointEntryPair.AssetType() {
		case externalapi.KSH:
			kshBalance += utxoOutpointEntryPair.Amount()
		case externalapi.KUSD:
			kusdBalance += utxoOutpointEntryPair.Amount()
		case externalapi.KRV:
			krvBalance += utxoOutpointEntryPair.Amount()
		}
	}
	return kshBalance, kusdBalance, krvBalance, nil
}
