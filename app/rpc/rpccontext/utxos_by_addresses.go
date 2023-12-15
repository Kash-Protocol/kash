package rpccontext

import (
	"encoding/hex"
	"github.com/Kash-Protocol/kashd/domain/consensus/model/externalapi"
	"github.com/Kash-Protocol/kashd/domain/consensus/utils/txscript"
	"github.com/Kash-Protocol/kashd/util"
	"github.com/pkg/errors"

	"github.com/Kash-Protocol/kashd/app/appmessage"
	"github.com/Kash-Protocol/kashd/domain/utxoindex"
)

// ConvertUTXOOutpointEntryPairsToUTXOsByAddressesEntries converts
// UTXOOutpointEntryPairs to a slice of UTXOsByAddressesEntry
func ConvertUTXOOutpointEntryPairsToUTXOsByAddressesEntries(address string, pairs utxoindex.UTXOOutpointEntryPairs) []*appmessage.UTXOsByAddressesEntry {
	utxosByAddressesEntries := make([]*appmessage.UTXOsByAddressesEntry, 0, len(pairs))
	for outpoint, utxoEntry := range pairs {
		utxosByAddressesEntries = append(utxosByAddressesEntries, &appmessage.UTXOsByAddressesEntry{
			Address: address,
			Outpoint: &appmessage.RPCOutpoint{
				TransactionID: outpoint.TransactionID.String(),
				Index:         outpoint.Index,
			},
			UTXOEntry: &appmessage.RPCUTXOEntry{
				Amount:          utxoEntry.Amount(),
				ScriptPublicKey: &appmessage.RPCScriptPublicKey{Script: hex.EncodeToString(utxoEntry.ScriptPublicKey().Script), Version: utxoEntry.ScriptPublicKey().Version},
				BlockDAAScore:   utxoEntry.BlockDAAScore(),
				IsCoinbase:      utxoEntry.IsCoinbase(),
			},
		})
	}
	return utxosByAddressesEntries
}

// ConvertAddressStringsToUTXOsChangedNotificationAddresses converts address strings
// and asset types to UTXOsChangedNotificationAddresses.
func (ctx *Context) ConvertAddressStringsToUTXOsChangedNotificationAddresses(
	addressStrings []string, assetTypes []uint32) ([]*UTXOsChangedNotificationAddress, error) {

	// Ensure addressStrings and assetTypes have the same length
	if len(addressStrings) != len(assetTypes) {
		return nil, errors.New("addressStrings and assetTypes must have the same length")
	}

	addresses := make([]*UTXOsChangedNotificationAddress, len(addressStrings))
	for i := range addressStrings {
		addressString := addressStrings[i]
		assetType := assetTypes[i]

		address, err := util.DecodeAddress(addressString, ctx.Config.ActiveNetParams.Prefix)
		if err != nil {
			return nil, errors.Errorf("Could not decode address '%s': %s", addressString, err)
		}

		// Use the corresponding assetType for each address
		scriptPublicKey, err := txscript.PayToAddrScript(address, externalapi.AssetTypeFromUint32(assetType))
		if err != nil {
			return nil, errors.Errorf("Could not create a scriptPublicKey for address '%s': %s", addressString, err)
		}

		scriptPublicKeyString := utxoindex.ScriptPublicKeyString(scriptPublicKey.String())
		addresses[i] = &UTXOsChangedNotificationAddress{
			Address:               addressString,
			ScriptPublicKeyString: scriptPublicKeyString,
		}
	}
	return addresses, nil
}
