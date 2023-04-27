// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: schema/schema.proto

package proto_models

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// По размеру сериализованных данных repeated KVType = map<...,...>
// поэтому можно считать эквивалентом
type KVString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	L string `protobuf:"bytes,1,opt,name=L,proto3" json:"L,omitempty"`
	R string `protobuf:"bytes,2,opt,name=R,proto3" json:"R,omitempty"`
}

func (x *KVString) Reset() {
	*x = KVString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVString) ProtoMessage() {}

func (x *KVString) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVString.ProtoReflect.Descriptor instead.
func (*KVString) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{0}
}

func (x *KVString) GetL() string {
	if x != nil {
		return x.L
	}
	return ""
}

func (x *KVString) GetR() string {
	if x != nil {
		return x.R
	}
	return ""
}

type NestedTestType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoolField        bool        `protobuf:"varint,1,opt,name=BoolField,proto3" json:"BoolField,omitempty"`
	ByteField        int32       `protobuf:"varint,2,opt,name=ByteField,proto3" json:"ByteField,omitempty"`
	IntField         int32       `protobuf:"varint,3,opt,name=IntField,proto3" json:"IntField,omitempty"`
	Int64Field       int64       `protobuf:"varint,4,opt,name=Int64Field,proto3" json:"Int64Field,omitempty"`
	UInt32Field      uint32      `protobuf:"varint,5,opt,name=UInt32Field,proto3" json:"UInt32Field,omitempty"`
	FloatField       float32     `protobuf:"fixed32,6,opt,name=FloatField,proto3" json:"FloatField,omitempty"`
	DoubleField      float64     `protobuf:"fixed64,7,opt,name=DoubleField,proto3" json:"DoubleField,omitempty"`
	NullStringField  string      `protobuf:"bytes,8,opt,name=NullStringField,proto3" json:"NullStringField,omitempty"`
	StringField      string      `protobuf:"bytes,9,opt,name=StringField,proto3" json:"StringField,omitempty"`
	BytesField       []byte      `protobuf:"bytes,10,opt,name=BytesField,proto3" json:"BytesField,omitempty"`
	StringArrayField []string    `protobuf:"bytes,11,rep,name=StringArrayField,proto3" json:"StringArrayField,omitempty"`
	IntArrayField    []int32     `protobuf:"varint,12,rep,packed,name=IntArrayField,proto3" json:"IntArrayField,omitempty"`
	MapField         []*KVString `protobuf:"bytes,13,rep,name=MapField,proto3" json:"MapField,omitempty"`
}

func (x *NestedTestType) Reset() {
	*x = NestedTestType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedTestType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedTestType) ProtoMessage() {}

func (x *NestedTestType) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedTestType.ProtoReflect.Descriptor instead.
func (*NestedTestType) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{1}
}

func (x *NestedTestType) GetBoolField() bool {
	if x != nil {
		return x.BoolField
	}
	return false
}

func (x *NestedTestType) GetByteField() int32 {
	if x != nil {
		return x.ByteField
	}
	return 0
}

func (x *NestedTestType) GetIntField() int32 {
	if x != nil {
		return x.IntField
	}
	return 0
}

func (x *NestedTestType) GetInt64Field() int64 {
	if x != nil {
		return x.Int64Field
	}
	return 0
}

func (x *NestedTestType) GetUInt32Field() uint32 {
	if x != nil {
		return x.UInt32Field
	}
	return 0
}

func (x *NestedTestType) GetFloatField() float32 {
	if x != nil {
		return x.FloatField
	}
	return 0
}

func (x *NestedTestType) GetDoubleField() float64 {
	if x != nil {
		return x.DoubleField
	}
	return 0
}

func (x *NestedTestType) GetNullStringField() string {
	if x != nil {
		return x.NullStringField
	}
	return ""
}

func (x *NestedTestType) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

func (x *NestedTestType) GetBytesField() []byte {
	if x != nil {
		return x.BytesField
	}
	return nil
}

func (x *NestedTestType) GetStringArrayField() []string {
	if x != nil {
		return x.StringArrayField
	}
	return nil
}

func (x *NestedTestType) GetIntArrayField() []int32 {
	if x != nil {
		return x.IntArrayField
	}
	return nil
}

func (x *NestedTestType) GetMapField() []*KVString {
	if x != nil {
		return x.MapField
	}
	return nil
}

type KVNestedTestType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	L string          `protobuf:"bytes,1,opt,name=L,proto3" json:"L,omitempty"`
	R *NestedTestType `protobuf:"bytes,2,opt,name=R,proto3" json:"R,omitempty"`
}

func (x *KVNestedTestType) Reset() {
	*x = KVNestedTestType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVNestedTestType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVNestedTestType) ProtoMessage() {}

func (x *KVNestedTestType) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVNestedTestType.ProtoReflect.Descriptor instead.
func (*KVNestedTestType) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{2}
}

func (x *KVNestedTestType) GetL() string {
	if x != nil {
		return x.L
	}
	return ""
}

func (x *KVNestedTestType) GetR() *NestedTestType {
	if x != nil {
		return x.R
	}
	return nil
}

type TestType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoolField       bool                `protobuf:"varint,1,opt,name=BoolField,proto3" json:"BoolField,omitempty"`
	ByteField       int32               `protobuf:"varint,2,opt,name=ByteField,proto3" json:"ByteField,omitempty"`
	IntField        int32               `protobuf:"varint,3,opt,name=IntField,proto3" json:"IntField,omitempty"`
	Int64Field      int64               `protobuf:"varint,4,opt,name=Int64Field,proto3" json:"Int64Field,omitempty"`
	UInt32Field     uint32              `protobuf:"varint,5,opt,name=UInt32Field,proto3" json:"UInt32Field,omitempty"`
	FloatField      float32             `protobuf:"fixed32,6,opt,name=FloatField,proto3" json:"FloatField,omitempty"`
	DoubleField     float64             `protobuf:"fixed64,7,opt,name=DoubleField,proto3" json:"DoubleField,omitempty"`
	NullStringField string              `protobuf:"bytes,8,opt,name=NullStringField,proto3" json:"NullStringField,omitempty"`
	StringField     string              `protobuf:"bytes,9,opt,name=StringField,proto3" json:"StringField,omitempty"`
	BytesField      []byte              `protobuf:"bytes,10,opt,name=BytesField,proto3" json:"BytesField,omitempty"`
	NestedField     *NestedTestType     `protobuf:"bytes,11,opt,name=NestedField,proto3" json:"NestedField,omitempty"`
	ArrayField      []*NestedTestType   `protobuf:"bytes,12,rep,name=ArrayField,proto3" json:"ArrayField,omitempty"`
	MapField        []*KVNestedTestType `protobuf:"bytes,13,rep,name=MapField,proto3" json:"MapField,omitempty"`
}

func (x *TestType) Reset() {
	*x = TestType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestType) ProtoMessage() {}

func (x *TestType) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestType.ProtoReflect.Descriptor instead.
func (*TestType) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{3}
}

func (x *TestType) GetBoolField() bool {
	if x != nil {
		return x.BoolField
	}
	return false
}

func (x *TestType) GetByteField() int32 {
	if x != nil {
		return x.ByteField
	}
	return 0
}

func (x *TestType) GetIntField() int32 {
	if x != nil {
		return x.IntField
	}
	return 0
}

func (x *TestType) GetInt64Field() int64 {
	if x != nil {
		return x.Int64Field
	}
	return 0
}

func (x *TestType) GetUInt32Field() uint32 {
	if x != nil {
		return x.UInt32Field
	}
	return 0
}

func (x *TestType) GetFloatField() float32 {
	if x != nil {
		return x.FloatField
	}
	return 0
}

func (x *TestType) GetDoubleField() float64 {
	if x != nil {
		return x.DoubleField
	}
	return 0
}

func (x *TestType) GetNullStringField() string {
	if x != nil {
		return x.NullStringField
	}
	return ""
}

func (x *TestType) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

func (x *TestType) GetBytesField() []byte {
	if x != nil {
		return x.BytesField
	}
	return nil
}

func (x *TestType) GetNestedField() *NestedTestType {
	if x != nil {
		return x.NestedField
	}
	return nil
}

func (x *TestType) GetArrayField() []*NestedTestType {
	if x != nil {
		return x.ArrayField
	}
	return nil
}

func (x *TestType) GetMapField() []*KVNestedTestType {
	if x != nil {
		return x.MapField
	}
	return nil
}

var File_schema_schema_proto protoreflect.FileDescriptor

var file_schema_schema_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x08, 0x4b, 0x56, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x0c, 0x0a, 0x01, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x4c, 0x12,
	0x0c, 0x0a, 0x01, 0x52, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x52, 0x22, 0xd1, 0x03,
	0x0a, 0x0e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x42, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x42, 0x79, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x42, 0x79, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x49, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x49, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a,
	0x46, 0x6c, 0x6f, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0b, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x28, 0x0a, 0x0f,
	0x4e, 0x75, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x4e, 0x75, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x41, 0x72, 0x72, 0x61, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x10, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x72, 0x61, 0x79, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0d, 0x49, 0x6e, 0x74,
	0x41, 0x72, 0x72, 0x61, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x4d, 0x61,
	0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4b,
	0x56, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x22, 0x3f, 0x0a, 0x10, 0x4b, 0x56, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x65, 0x73,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x01, 0x4c, 0x12, 0x1d, 0x0a, 0x01, 0x52, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x01, 0x52, 0x22, 0xe5, 0x03, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x42, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x42, 0x79, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x42, 0x79, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x49,
	0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x49,
	0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x46,
	0x6c, 0x6f, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x4e,
	0x75, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x4e, 0x75, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x31, 0x0a, 0x0b, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2f, 0x0a, 0x0a, 0x41, 0x72,
	0x72, 0x61, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x41, 0x72, 0x72, 0x61, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2d, 0x0a, 0x08, 0x4d,
	0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x4b, 0x56, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_schema_schema_proto_rawDescOnce sync.Once
	file_schema_schema_proto_rawDescData = file_schema_schema_proto_rawDesc
)

func file_schema_schema_proto_rawDescGZIP() []byte {
	file_schema_schema_proto_rawDescOnce.Do(func() {
		file_schema_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_schema_proto_rawDescData)
	})
	return file_schema_schema_proto_rawDescData
}

var file_schema_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_schema_schema_proto_goTypes = []interface{}{
	(*KVString)(nil),         // 0: KVString
	(*NestedTestType)(nil),   // 1: NestedTestType
	(*KVNestedTestType)(nil), // 2: KVNestedTestType
	(*TestType)(nil),         // 3: TestType
}
var file_schema_schema_proto_depIdxs = []int32{
	0, // 0: NestedTestType.MapField:type_name -> KVString
	1, // 1: KVNestedTestType.R:type_name -> NestedTestType
	1, // 2: TestType.NestedField:type_name -> NestedTestType
	1, // 3: TestType.ArrayField:type_name -> NestedTestType
	2, // 4: TestType.MapField:type_name -> KVNestedTestType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_schema_schema_proto_init() }
func file_schema_schema_proto_init() {
	if File_schema_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVString); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedTestType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVNestedTestType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_schema_proto_goTypes,
		DependencyIndexes: file_schema_schema_proto_depIdxs,
		MessageInfos:      file_schema_schema_proto_msgTypes,
	}.Build()
	File_schema_schema_proto = out.File
	file_schema_schema_proto_rawDesc = nil
	file_schema_schema_proto_goTypes = nil
	file_schema_schema_proto_depIdxs = nil
}