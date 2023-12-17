package externalapi

// CirculatingSupply stores the circulating supply for each asset type.
type CirculatingSupply struct {
	KSH  uint64
	KUSD uint64
	KRV  uint64
}

// AssetSupplyChange represents the change in supply for each asset type.
// Separate variables for 'toAdd' and 'toRemove' are used instead of a single int64 variable
// to avoid the risks associated with converting between int64 and uint64.
type AssetSupplyChange struct {
	KSHtoAdd     uint64
	KSHtoRemove  uint64
	KUSDtoAdd    uint64
	KUSDtoRemove uint64
	KRVtoAdd     uint64
	KRVtoRemove  uint64
}

// Add adds the amount to the corresponding asset type.
func (asc *AssetSupplyChange) Add(assetType AssetType, amount uint64) {
	switch assetType {
	case KSH:
		asc.KSHtoAdd += amount
	case KUSD:
		asc.KUSDtoAdd += amount
	case KRV:
		asc.KRVtoAdd += amount
	}
}

// Subtract subtracts the amount from the corresponding asset type.
func (asc *AssetSupplyChange) Subtract(assetType AssetType, amount uint64) {
	switch assetType {
	case KSH:
		asc.KSHtoRemove += amount
	case KUSD:
		asc.KUSDtoRemove += amount
	case KRV:
		asc.KRVtoRemove += amount
	}
}
