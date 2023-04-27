package serializer

import "github.com/vmihailenco/msgpack/v5"

type MessagePackSerializer[TestType any] struct {
}

func (*MessagePackSerializer[TestType]) Serialize(v *TestType) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (*MessagePackSerializer[TestType]) Deserialize(data []byte, v *TestType) error {
	return msgpack.Unmarshal(data, v)
}

func NewMessagePackSerializer[TestType any]() *MessagePackSerializer[TestType] {
	return &MessagePackSerializer[TestType]{}
}
