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
	0x3b, 0x6b, 0xab, 0x14, 0x54, 0xe3, 0x52, 0xc5,
	0x87, 0x0e, 0x4c, 0xf7, 0xc8, 0xce, 0x3b, 0x87,
	0x6c, 0xb4, 0xec, 0xf3, 0xfb, 0x69, 0xdb, 0x72,
	0x7d, 0x92, 0x2e, 0x49, 0xca, 0x4c, 0x71, 0x8a,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x35, 0x2b, 0xc4, 0xdd, 0xe7, 0x94, 0x6b, 0x4d,
	0xfc, 0x7d, 0x5e, 0xfe, 0xfe, 0x0f, 0x38, 0xdd,
	0x55, 0x11, 0xd0, 0x06, 0x92, 0x5b, 0x0e, 0x9c,
	0x87, 0x54, 0x0b, 0xa2, 0x60, 0x6d, 0x9f, 0x86,
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
	0x97, 0xa6, 0xb7, 0xa8, 0xe8, 0x5a, 0xc0, 0x67,
	0xb8, 0x11, 0xac, 0x75, 0x30, 0x06, 0xd8, 0xd4,
	0x53, 0x63, 0xa3, 0xa3, 0x17, 0xe6, 0x2e, 0x3f,
	0x32, 0x7b, 0xb1, 0xb3, 0x50, 0x1d, 0x6e, 0x28,
})

// devnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var devnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x58, 0xab, 0xf2, 0x03, 0x21, 0xd7, 0x07, 0x16,
	0x16, 0x2b, 0x6b, 0xf8, 0xd9, 0xf5, 0x89, 0xca,
	0x33, 0xae, 0x6e, 0x32, 0xb3, 0xb1, 0x9a, 0xbb,
	0x7f, 0xa6, 0x5d, 0x11, 0x41, 0xa3, 0xf9, 0x4d,
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
	0xfa, 0x59, 0xaf, 0xd7, 0xb2, 0xe8, 0x0d, 0x55,
	0x2c, 0x75, 0x7b, 0xbb, 0x42, 0xbe, 0x42, 0x23,
	0x27, 0x99, 0x78, 0x98, 0x39, 0x2a, 0xd2, 0x6c,
	0xb5, 0xc6, 0x8d, 0xaf, 0xa9, 0xc6, 0x7c, 0x73,
})

// simnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var simnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x19, 0x46, 0xd6, 0x29, 0xf7, 0xe9, 0x22, 0xa7,
	0xbc, 0xed, 0x59, 0x19, 0x05, 0x21, 0xc3, 0x77,
	0x1f, 0x73, 0xd3, 0x52, 0xdd, 0xbb, 0xb6, 0x86,
	0x56, 0x4a, 0xd7, 0xfd, 0x56, 0x85, 0x7c, 0x1b,
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
	0x83, 0x82, 0x5b, 0x56, 0x6d, 0x9e, 0x82, 0x9c,
	0xd7, 0x5a, 0x63, 0x75, 0xa8, 0x5e, 0xb7, 0x40,
	0x2a, 0x24, 0xcb, 0xea, 0x19, 0xca, 0xf7, 0xca,
	0xe6, 0x28, 0x7b, 0xef, 0xbd, 0x02, 0x82, 0x5f,
})

// testnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for testnet.
var testnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x3b, 0x31, 0xb3, 0xe2, 0x9a, 0xf1, 0xc9, 0x67,
	0x5d, 0xd1, 0x48, 0xa4, 0x8a, 0x91, 0x74, 0xd4,
	0x6b, 0x1a, 0x3a, 0xed, 0xa5, 0x73, 0x2b, 0x0c,
	0x6d, 0xf7, 0xd2, 0xb1, 0x1f, 0x54, 0x52, 0x2a,
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
