package externalapi

import (
	"errors"
	"strings"
)

// AssetType defines a type for different asset types.
type AssetType uint32

// Enumeration of asset types.
const (
	KSH AssetType = iota
	KUSD
	KRV
	UNKNOWN
)

// Opcode values for asset types.
const (
	OpAssetKSH  = 0xb2 // 178
	OpAssetKUSD = 0xb3 // 179
	OpAssetKRV  = 0xb4 // 180
)

// ToUint32 converts an asset.AssetType to a uint32.
func (t AssetType) ToUint32() uint32 {
	return uint32(t)
}

func (t AssetType) String() string {
	switch t {
	case KSH:
		return "KSH"
	case KRV:
		return "KRV"
	case KUSD:
		return "KUSD"
	}
	return "UNKNOWN"
}

// AssetTypeFromString creates an asset.AssetType from a string.
func AssetTypeFromString(s string) AssetType {
	switch strings.ToUpper(s) {
	case "KSH":
		return KSH
	case "KRV":
		return KRV
	case "KUSD":
		return KUSD
	default:
		return UNKNOWN
	}
}

// AssetTypeFromUint32 creates an asset.AssetType from a uint32.
func AssetTypeFromUint32(u uint32) AssetType {
	switch u {
	case uint32(KSH):
		return KSH
	case uint32(KRV):
		return KRV
	case uint32(KUSD):
		return KUSD
	default:
		return UNKNOWN
	}
}

// ExtractAssetType extracts the asset type from a given PubKeyAddress.
// This function analyzes the provided script to determine the asset type.
// It returns the asset type and any error encountered during the extraction.
func ExtractAssetType(pubKeyAddress []byte) (AssetType, error) {
	if len(pubKeyAddress) == 0 {
		return UNKNOWN, errors.New("empty PubKeyAddress provided")
	}

	// The first byte of the PubKeyAddress typically determines the asset type.
	switch pubKeyAddress[0] {
	case OpAssetKSH: // Example value for KSH
		return KSH, nil
	case OpAssetKUSD: // Example value for KUSD
		return KUSD, nil
	case OpAssetKRV: // Example value for KRV
		return KRV, nil
	default:
		return UNKNOWN, nil
	}
}

// ConvertAssetTypeSliceToUint32Slice converts a slice of AssetType to a slice of uint32.
func ConvertAssetTypeSliceToUint32Slice(assetTypes []AssetType) []uint32 {
	uint32Slice := make([]uint32, len(assetTypes))
	for i, v := range assetTypes {
		uint32Slice[i] = uint32(v)
	}
	return uint32Slice
}

// GetAssetTypeFromDomainTransactionType returns the asset type from a given DomainTransactionType.
func GetAssetTypeFromDomainTransactionType(domainTransactionType DomainTransactionType) (AssetType, AssetType) {
	switch domainTransactionType {
	case TransferKSH:
		return KSH, KSH
	case TransferKUSD:
		return KUSD, KUSD
	case TransferKRV:
		return KRV, KRV
	case MintKUSD:
		return KSH, KUSD
	case StakeKSH:
		return KSH, KRV
	case RedeemKSH:
		return KRV, KSH
	default:
		return UNKNOWN, UNKNOWN
	}
}
