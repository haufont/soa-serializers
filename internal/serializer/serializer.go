package serializer

import (
	"serialize-benchmarker/internal/format"
)

type Serializer[TestType any] interface {
	Serialize(v *TestType) (data []byte, err error)
	Deserialize([]byte, *TestType) error
}

func NewSerializer[TestType any](f format.Format) Serializer[TestType] {
	switch f {
	case format.Native:
		return NewNativeSerializer[TestType]()
	case format.XML:
		return NewXMLSerializer[TestType]()
	case format.JSON:
		return NewJSONSerializer[TestType]()
	case format.ProtoBuf:
		panic("Use NewProtoBufSerializer")
	case format.ApacheAvro:
		return NewAvroSerializer[TestType]()
	case format.YAML:
		return NewYAMLSerializer[TestType]()
	case format.MessagePack:
		return NewMessagePackSerializer[TestType]()
	}
	panic("unknown format")
}
