package serializer

import (
	"bytes"
	"encoding/gob"
)

type NativeSerializer[TestType any] struct {
	buffer  bytes.Buffer
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func (serializer *NativeSerializer[TestType]) Serialize(v *TestType) (data []byte, err error) {
	err = serializer.encoder.Encode(v)
	data = serializer.buffer.Bytes()
	return
}

func (serializer *NativeSerializer[TestType]) Deserialize(data []byte, v *TestType) error {
	return serializer.decoder.Decode(v)
}

func NewNativeSerializer[TestType any]() *NativeSerializer[TestType] {
	serializer := NativeSerializer[TestType]{bytes.Buffer{}, nil, nil}
	serializer.encoder = gob.NewEncoder(&serializer.buffer)
	serializer.decoder = gob.NewDecoder(&serializer.buffer)
	return &serializer
}
