// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package txscript_test

import (
	"encoding/hex"
	"fmt"
	"github.com/Kash-Protocol/kashd/domain/consensus/model/externalapi"
	"github.com/Kash-Protocol/kashd/domain/consensus/utils/txscript"
	"github.com/Kash-Protocol/kashd/domain/dagconfig"
	"github.com/Kash-Protocol/kashd/util"
)

// This example demonstrates creating a script which pays to a kaspa address.
// It also prints the created script hex and uses the DisasmString function to
// display the disassembled script.
func ExamplePayToAddrScript() {
	// Parse the address to send the coins to into a util.Address
	// which is useful to ensure the accuracy of the address and determine
	// the address type. It is also required for the upcoming call to
	// PayToAddrScript.
	addressStr := "kash:qr35ennsep3hxfe7lnz5ee7j5jgmkjswsn35ennsep3hxfe7ln35c24fvumj0"
	address, err := util.DecodeAddress(addressStr, util.Bech32PrefixKash)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a public key script that pays to the address.
	script, err := txscript.PayToAddrScript(address, externalapi.KSH)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Script Hex: %x\n", script.Script)

	disasm, err := txscript.DisasmString(script.Version, script.Script)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Script Disassembly:", disasm)

	// Output:
	// Script Hex: 20e34cce70c86373273efcc54ce7d2a491bb4a0e84e34cce70c86373273efce34cac
	// Script Disassembly: e34cce70c86373273efcc54ce7d2a491bb4a0e84e34cce70c86373273efce34c OP_CHECKSIG
}

// This example demonstrates extracting information from a standard public key
// script.
func ExampleExtractScriptPubKeyAddress() {
	// Start with a standard pay-to-pubkey script.
	scriptHex := "2089ac24ea10bb751af4939623ccc5e550d96842b64e8fca0f63e94b4373fd555eac"
	script, err := hex.DecodeString(scriptHex)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Extract and print details from the script.
	scriptClass, address, _, err := txscript.ExtractScriptPubKeyAddress(
		&externalapi.ScriptPublicKey{
			Script:  script,
			Version: 0,
		}, &dagconfig.MainnetParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Script Class:", scriptClass)
	fmt.Println("Address:", address)

	// Output:
	// Script Class: pubkey
	// Address: kash:qzy6cf82zzah2xh5jwtz8nx9u4gdj6zzke8gljs0v055ksmnl424uas2l9cqt
}
