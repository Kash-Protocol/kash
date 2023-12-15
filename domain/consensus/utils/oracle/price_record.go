package oracle

import (
	"encoding/json"
	"fmt"
)

// PriceRecord stores the prices and moving averages of KSH, KRV, and KUSD against USDT.
type PriceRecord struct {
	KSH       float64
	KRV       float64
	KUSD      float64
	KSHMA     float64
	KRVMA     float64
	KUSDMA    float64
	Timestamp int64
}

// ToBytes serializes the PriceRecord into a byte slice.
func (pr *PriceRecord) ToBytes() ([]byte, error) {
	data, err := json.Marshal(pr)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize PriceRecord: %w", err)
	}
	return data, nil
}

// FromBytes deserializes a byte slice into a PriceRecord.
func FromBytes(data []byte) (*PriceRecord, error) {
	var pr PriceRecord
	if err := json.Unmarshal(data, &pr); err != nil {
		return nil, fmt.Errorf("failed to deserialize byte slice to PriceRecord: %w", err)
	}
	return &pr, nil
}
