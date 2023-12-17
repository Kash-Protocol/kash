// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dagconfig

import (
	"github.com/Kash-Protocol/kashd/domain/consensus/model/externalapi"
	"github.com/Kash-Protocol/kashd/domain/consensus/utils/blockheader"
	"github.com/Kash-Protocol/kashd/domain/consensus/utils/subnetworks"
	"github.com/Kash-Protocol/kashd/domain/consensus/utils/transactionhelper"
	"github.com/kaspanet/go-muhash"
	"math/big"
)

var genesisTxOuts = []*externalapi.DomainTransactionOutput{}

var genesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, //script version
	0x01,                                                                   // Varint
	0x00,                                                                   // OP-FALSE
	0x4b, 0x61, 0x73, 0x68, 0x20, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, // "Kash Genesis Block"
	0x20, 0x2d, 0x20, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64,
	0x20, 0x74, 0x6f, 0x20, 0x44, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x63, 0x6f, 0x69, 0x6e, 0x20, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73,
}

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network.
var genesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0, []*externalapi.DomainTransactionInput{}, genesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, genesisTxPayload)

// genesisHash is the hash of the first block in the block DAG for the main
// network (genesis block).
var genesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x9b, 0xc6, 0x44, 0x22, 0xee, 0x9b, 0x91, 0x91,
	0xea, 0x02, 0x4f, 0x43, 0x3a, 0x31, 0xdf, 0xde,
	0xb7, 0xc9, 0xe4, 0xa9, 0x6f, 0xda, 0xa8, 0x8a,
	0xa3, 0x1a, 0x02, 0xbf, 0x07, 0x87, 0xd8, 0xa0,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xda, 0x62, 0x38, 0xc4, 0xdc, 0x5b, 0x16, 0x4c,
	0xc2, 0xb6, 0x60, 0x02, 0xb7, 0xc4, 0x83, 0x96,
	0x02, 0xb5, 0x09, 0x8d, 0x56, 0xb6, 0xa0, 0x7b,
	0x18, 0xed, 0x51, 0x2a, 0x0b, 0x2e, 0x04, 0x82,
})

// genesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the main network.
var genesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		genesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1637609671037,
		0x1f2c83cd,
		0x22352,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{genesisCoinbaseTx},
}

var devnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var devnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                   // Varint
	0x00,                                                                   // OP-FALSE
	0x6b, 0x61, 0x73, 0x70, 0x61, 0x2d, 0x64, 0x65, 0x76, 0x6e, 0x65, 0x74, // kash-devnet
}

// devnetGenesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the development network.
var devnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, devnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, devnetGenesisTxPayload)

// devGenesisHash is the hash of the first block in the block DAG for the development
// network (genesis block).
var devnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x4f, 0xfa, 0x52, 0x0c, 0xf3, 0xc2, 0x59, 0x3c,
	0x64, 0xad, 0xe3, 0xa0, 0xc4, 0xc8, 0x3e, 0x45,
	0x36, 0x87, 0x39, 0xb2, 0x7f, 0x9d, 0xc9, 0xc7,
	0xf7, 0x1d, 0x7b, 0x7a, 0x6f, 0xcd, 0xee, 0x21,
})

// devnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var devnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x45, 0x2b, 0xe9, 0xf0, 0x4b, 0x96, 0xa6, 0x5a,
	0x7a, 0xd0, 0x06, 0x88, 0x7e, 0xbb, 0xb3, 0xa6,
	0xfc, 0xa7, 0x57, 0xdb, 0x2c, 0xc0, 0x1f, 0x4d,
	0x42, 0xfb, 0xfc, 0x27, 0x52, 0x15, 0xed, 0x99,
})

// devnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var devnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		devnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x11e9db49828,
		0x200b20f3,
		0x48e5e,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{devnetGenesisCoinbaseTx},
}

var simnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var simnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                   // Varint
	0x00,                                                                   // OP-FALSE
	0x6b, 0x61, 0x73, 0x70, 0x61, 0x2d, 0x73, 0x69, 0x6d, 0x6e, 0x65, 0x74, // kash-simnet
}

// simnetGenesisCoinbaseTx is the coinbase transaction for the simnet genesis block.
var simnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, simnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, simnetGenesisTxPayload)

// simnetGenesisHash is the hash of the first block in the block DAG for
// the simnet (genesis block).
var simnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xae, 0xcd, 0x84, 0xad, 0xa4, 0x95, 0xe0, 0xc6,
	0x59, 0x08, 0x18, 0x68, 0x21, 0x07, 0x9c, 0x09,
	0x3c, 0x66, 0x80, 0x7f, 0xb9, 0x47, 0x75, 0xe7,
	0x00, 0xc6, 0xa6, 0xc3, 0xd5, 0xc4, 0x43, 0x7e,
})

// simnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var simnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xcf, 0x11, 0x0c, 0x54, 0xbb, 0x1c, 0x4a, 0x0c,
	0x4d, 0x80, 0xcf, 0x43, 0x12, 0x0d, 0xa7, 0x22,
	0x9f, 0xac, 0x69, 0xda, 0xfd, 0xa2, 0x1c, 0xab,
	0xf7, 0xd1, 0x15, 0x90, 0xe3, 0x4c, 0xf9, 0xed,
})

// simnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var simnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		simnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x17c5f62fbb6,
		0x206f497e,
		0x2,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{simnetGenesisCoinbaseTx},
}

var testnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var testnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                           // Varint
	0x00,                                           // OP-FALSE
	0x6b, 0x61, 0x73, 0x68, 0x2d, 0x74, 0x65, 0x73, // kash-testnet
	0x74, 0x6e, 0x65, 0x74,
}

// testnetGenesisCoinbaseTx is the coinbase transaction for the testnet genesis block.
var testnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, testnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, testnetGenesisTxPayload)

// testnetGenesisHash is the hash of the first block in the block DAG for the test
// network (genesis block).
var testnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x07, 0x64, 0xe1, 0x1f, 0xe4, 0x33, 0xd3, 0x6b,
	0xe7, 0x9a, 0x7b, 0xe5, 0x7b, 0x5a, 0x58, 0x33,
	0xe9, 0xfc, 0x7f, 0xd3, 0xd5, 0xf9, 0x09, 0x1a,
	0x87, 0x79, 0xdc, 0x85, 0x01, 0x2a, 0x51, 0x1b,
})

// testnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for testnet.
var testnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xf1, 0x43, 0x08, 0x08, 0x80, 0xdc, 0xb5, 0x71,
	0x0d, 0xb3, 0x82, 0x36, 0x7d, 0x52, 0x9f, 0x1a,
	0x15, 0xa1, 0xc1, 0x9c, 0xd5, 0x9e, 0x9e, 0x08,
	0x5c, 0x07, 0x3a, 0x6e, 0x30, 0xc8, 0x9f, 0xda,
})

// testnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for testnet.
var testnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		testnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x17c5f62fbb6,
		0x2001641e,
		0x14582,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{testnetGenesisCoinbaseTx},
}
