package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// KashMainnetPrivate is the version that is used for
// kaspa mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var KashMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// KashMainnetPublic is the version that is used for
// kaspa mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var KashMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// KashTestnetPrivate is the version that is used for
// kaspa testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var KashTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// KashTestnetPublic is the version that is used for
// kaspa testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var KashTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// KashDevnetPrivate is the version that is used for
// kaspa devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var KashDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// KashDevnetPublic is the version that is used for
// kaspa devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var KashDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// KashSimnetPrivate is the version that is used for
// kaspa simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var KashSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// KashSimnetPublic is the version that is used for
// kaspa simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var KashSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case KashMainnetPrivate:
		return KashMainnetPublic, nil
	case KashTestnetPrivate:
		return KashTestnetPublic, nil
	case KashDevnetPrivate:
		return KashDevnetPublic, nil
	case KashSimnetPrivate:
		return KashSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case KashMainnetPrivate:
		return true
	case KashTestnetPrivate:
		return true
	case KashDevnetPrivate:
		return true
	case KashSimnetPrivate:
		return true
	}

	return false
}
