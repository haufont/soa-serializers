package serializer

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ProtoBufSerializer[TestType protoreflect.ProtoMessage] struct {
}

func (*ProtoBufSerializer[TestType]) Serialize(v *TestType) ([]byte, error) {
	return proto.Marshal(*v)
}

func (*ProtoBufSerializer[TestType]) Deserialize(data []byte, v *TestType) error {
	return proto.Unmarshal(data, *v)
}

func NewProtoBufSerializer[TestType protoreflect.ProtoMessage]() *ProtoBufSerializer[TestType] {
	return &ProtoBufSerializer[TestType]{}
}
