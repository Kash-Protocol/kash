package binaryserialization

import (
	"bytes"
	"encoding/binary"
	"github.com/Kash-Protocol/kashd/domain/consensus/model/externalapi"
	"github.com/pkg/errors"
)

// SerializeCirculatingSupply serializes the CirculatingSupply struct.
func SerializeCirculatingSupply(supply *externalapi.CirculatingSupply) ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, supply)
	if err != nil {
		return nil, errors.Wrap(err, "failed to serialize circulating supply")
	}
	return buf.Bytes(), nil
}

// DeserializeCirculatingSupply deserializes bytes into a CirculatingSupply struct.
func DeserializeCirculatingSupply(data []byte) (*externalapi.CirculatingSupply, error) {
	var supply externalapi.CirculatingSupply
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.LittleEndian, &supply)
	if err != nil {
		return nil, errors.Wrap(err, "failed to deserialize circulating supply")
	}
	return &supply, nil
}
