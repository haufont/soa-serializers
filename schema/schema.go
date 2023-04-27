package schema

import (
	_ "embed"
	proto_models "serialize-benchmarker/schema/proto.models"
)

//go:embed schema.avro
var AvroSchema string

type KVPair[LType any, RType any] struct {
	L LType
	R RType
}

type NestedTestType struct {
	BoolField        bool
	ByteField        byte
	IntField         int
	Int64Field       int64
	UInt32Field      uint32
	FloatField       float32
	DoubleField      float64
	NullStringField  string
	StringField      string
	BytesField       []byte
	StringArrayField []string
	IntArrayField    []int32
	// xml: unsupported type: map[string]string
	MapField []KVPair[string, string]
}

type TestType struct {
	BoolField       bool
	ByteField       byte
	IntField        int
	Int64Field      int64
	UInt32Field     uint32
	FloatField      float32
	DoubleField     float64
	NullStringField string
	StringField     string
	BytesField      []byte
	NestedField     NestedTestType
	ArrayField      []NestedTestType
	// xml: unsupported type: map[string]NestedTestType
	MapField []KVPair[string, NestedTestType]
}

func InitNestedTestType() NestedTestType {
	return NestedTestType{
		BoolField:   false,
		ByteField:   0x72,
		IntField:    -987654321,
		Int64Field:  -987654321987654321,
		UInt32Field: 987654321,
		FloatField:  0.987654321,
		DoubleField: 987.654321987654321,
		StringField: "abacabadabacaba",
		BytesField:  []byte{0x61, 0x62, 0x63, 0x64, 0x65},
		StringArrayField: []string{
			"a", "ba", "cab", "adab", "acaba",
		},
		IntArrayField: []int32{9, 8, 7, 6, 5, 4, 3, 2, 1},
		MapField: []KVPair[string, string]{
			{L: "a", R: "acaba"},
			{L: "ba", R: "adab"},
			{L: "cab", R: "cab"},
			{L: "adab", R: "ba"},
			{L: "acaba", R: "a"},
		},
	}
}

func InitTestType() TestType {
	return TestType{
		BoolField:   false,
		ByteField:   0x71,
		IntField:    -123456789,
		Int64Field:  123456789123456789,
		UInt32Field: 123456789,
		FloatField:  0.123456789,
		DoubleField: 123.456789123456789,
		StringField: "abacabadabacaba",
		BytesField:  []byte{0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56},
		NestedField: InitNestedTestType(),
		ArrayField: []NestedTestType{
			InitNestedTestType(), InitNestedTestType(), InitNestedTestType(),
		},
		MapField: []KVPair[string, NestedTestType]{
			{L: "aba", R: InitNestedTestType()},
			{L: "abacaba", R: InitNestedTestType()},
			{L: "abacabadabacaba", R: InitNestedTestType()},
		},
	}
}

func KVStringsToProtoKVs(kvs []KVPair[string, string]) (res []*proto_models.KVString) {
	for i := 0; i < len(kvs); i++ {
		res = append(res, &proto_models.KVString{
			L: kvs[i].L,
			R: kvs[i].R,
		})
	}
	return
}

func NestedTestTypeToProto(nestedTestType *NestedTestType) *proto_models.NestedTestType {
	return &proto_models.NestedTestType{
		BoolField:        nestedTestType.BoolField,
		ByteField:        int32(nestedTestType.ByteField),
		IntField:         int32(nestedTestType.IntField),
		Int64Field:       nestedTestType.Int64Field,
		UInt32Field:      nestedTestType.UInt32Field,
		FloatField:       nestedTestType.FloatField,
		DoubleField:      nestedTestType.DoubleField,
		StringField:      nestedTestType.StringField,
		BytesField:       nestedTestType.BytesField,
		StringArrayField: nestedTestType.StringArrayField,
		IntArrayField:    nestedTestType.IntArrayField,
		MapField:         KVStringsToProtoKVs(nestedTestType.MapField),
	}
}

func KVNestedTestTypesToProtoKVs(kvs []KVPair[string, NestedTestType]) (res []*proto_models.KVNestedTestType) {
	for i := 0; i < len(kvs); i++ {
		res = append(res, &proto_models.KVNestedTestType{
			L: kvs[i].L,
			R: NestedTestTypeToProto(&kvs[i].R),
		})
	}
	return
}

func NestedTestTypesToProto(nestedTestTypes []NestedTestType) (res []*proto_models.NestedTestType) {
	for i := 0; i < len(nestedTestTypes); i++ {
		res = append(res, NestedTestTypeToProto(&nestedTestTypes[i]))
	}
	return
}

func TestTypeToProto(testType *TestType) *proto_models.TestType {
	return &proto_models.TestType{
		BoolField:   testType.BoolField,
		ByteField:   int32(testType.ByteField),
		IntField:    int32(testType.IntField),
		Int64Field:  testType.Int64Field,
		UInt32Field: testType.UInt32Field,
		FloatField:  testType.FloatField,
		DoubleField: testType.DoubleField,
		StringField: testType.StringField,
		BytesField:  testType.BytesField,
		NestedField: NestedTestTypeToProto(&testType.NestedField),
		ArrayField:  NestedTestTypesToProto(testType.ArrayField),
		MapField:    KVNestedTestTypesToProtoKVs(testType.MapField),
	}
}
