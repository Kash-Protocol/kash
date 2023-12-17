// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dagconfig

import (
	"encoding/hex"
	"fmt"
	"github.com/Kash-Protocol/kashd/domain/consensus/model/externalapi"
	"log"
	"testing"

	"github.com/Kash-Protocol/kashd/domain/consensus/utils/consensushashing"
)

// TestGenesisBlock tests the genesis block of the main network for validity by
// checking the encoded hash.
func TestGenesisBlock(t *testing.T) {
	// Check hash of the block against expected hash.
	hash := consensushashing.BlockHash(MainnetParams.GenesisBlock)
	if !MainnetParams.GenesisHash.Equal(hash) {
		t.Fatalf("TestGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", hash, MainnetParams.GenesisHash)
	}
}

// TestTestnetGenesisBlock tests the genesis block of the test network for
// validity by checking the hash.
func TestTestnetGenesisBlock(t *testing.T) {
	// Check hash of the block against expected hash.
	hash := consensushashing.BlockHash(TestnetParams.GenesisBlock)
	if !TestnetParams.GenesisHash.Equal(hash) {
		t.Fatalf("TestTestnetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", hash,
			TestnetParams.GenesisHash)
	}
}

// TestSimnetGenesisBlock tests the genesis block of the simulation test network
// for validity by checking the hash.
func TestSimnetGenesisBlock(t *testing.T) {
	// Check hash of the block against expected hash.
	hash := consensushashing.BlockHash(SimnetParams.GenesisBlock)
	if !SimnetParams.GenesisHash.Equal(hash) {
		t.Fatalf("TestSimnetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", hash,
			SimnetParams.GenesisHash)
	}
}

// TestDevnetGenesisBlock tests the genesis block of the development network
// for validity by checking the encoded hash.
func TestDevnetGenesisBlock(t *testing.T) {
	// Check hash of the block against expected hash.
	hash := consensushashing.BlockHash(DevnetParams.GenesisBlock)
	if !DevnetParams.GenesisHash.Equal(hash) {
		t.Fatalf("TestDevnetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", hash,
			DevnetParams.GenesisHash)
	}
}

// TestPrintMerkleHash is a scratchpad for printing the genesis merkle root hash
func TestPrintMerkleHash(t *testing.T) {
	//// Check hash of the block against expected hash.
	//hash := consensushashing.BlockHash(MainnetParams.GenesisBlock)
	//t.Logf("TestPrintHash: Genesis block hash: %v", hash)
	//
	//hash = consensushashing.BlockHash(TestnetParams.GenesisBlock)
	//t.Logf(Logf"TestPrintHash: Genesis block hash: %v", hash)

	// Replace with the actual transactions for genesis and testnet blocks
	genesisTransactions := []*externalapi.DomainTransaction{genesisCoinbaseTx}
	testnetTransactions := []*externalapi.DomainTransaction{testnetGenesisCoinbaseTx}
	devnetTransactions := []*externalapi.DomainTransaction{devnetGenesisCoinbaseTx}
	simnetTransactions := []*externalapi.DomainTransaction{simnetGenesisCoinbaseTx}

	fmt.Printf("Genesis Merkle Root: %s\n", consensushashing.TransactionHash(genesisTransactions[0]))
	fmt.Printf("Testnet Merkle Root: %s\n", consensushashing.TransactionHash(testnetTransactions[0]))
	fmt.Printf("Devnet Merkle Root: %s\n", consensushashing.TransactionHash(devnetTransactions[0]))
	fmt.Printf("Simnet Merkle Root: %s\n", consensushashing.TransactionHash(simnetTransactions[0]))
}

// TestPrintsGenesisHash is a scratchpad for printing the genesis block hash
func TestPrintsGenesisHash(t *testing.T) {
	mainhashStr := consensushashing.BlockHash(MainnetParams.GenesisBlock).String()
	testhashStr := consensushashing.BlockHash(TestnetParams.GenesisBlock).String()
	simhashStr := consensushashing.BlockHash(SimnetParams.GenesisBlock).String()
	devhashStr := consensushashing.BlockHash(DevnetParams.GenesisBlock).String()

	for _, hashStr := range []string{mainhashStr, testhashStr, simhashStr, devhashStr} {
		hashBytes, err := hex.DecodeString(hashStr)
		if err != nil {
			log.Fatalf("Failed to decode hash: %v", err)
		}

		domainHashSize := 32

		if len(hashBytes) != domainHashSize {
			log.Fatalf("Hash size (%d) does not match DomainHashSize (%d)", len(hashBytes), domainHashSize)
		}

		fmt.Println("Hash bytes:")
		for i, b := range hashBytes {
			fmt.Printf("0x%02x", b)
			if i < len(hashBytes)-1 {
				fmt.Print(", ")
			}
			if (i+1)%8 == 0 {
				fmt.Println()
			}
		}
	}
}

// TestPrintsHashStrToReadableBytes is a scratchpad for printing the genesis block hash
func TestPrintsHashStrToReadableBytes(t *testing.T) {
	hashstr := "a2b71eec8405f10304db58000d51388689b716cfa975e3c404e2df73095767b6"
	hashBytes, err := hex.DecodeString(hashstr)
	if err != nil {
		log.Fatalf("Failed to decode hash: %v", err)
	}

	domainHashSize := 32

	if len(hashBytes) != domainHashSize {
		log.Fatalf("Hash size (%d) does not match DomainHashSize (%d)", len(hashBytes), domainHashSize)
	}

	fmt.Println("Hash bytes:")
	for i, b := range hashBytes {
		fmt.Printf("0x%02x", b)
		if i < len(hashBytes)-1 {
			fmt.Print(", ")
		}
		if (i+1)%8 == 0 {
			fmt.Println()
		}
	}
}
