package serializer

import (
	"serialize-benchmarker/schema"

	"github.com/hamba/avro/v2"
)

type AvroSerializer[TestType any] struct {
	schema avro.Schema
}

func (s *AvroSerializer[TestType]) Serialize(v *TestType) ([]byte, error) {
	return avro.Marshal(s.schema, v)
}

func (s *AvroSerializer[TestType]) Deserialize(data []byte, v *TestType) error {
	return avro.Unmarshal(s.schema, data, v)
}

func NewAvroSerializer[TestType any]() *AvroSerializer[TestType] {
	avroSchema, err := avro.Parse(schema.AvroSchema)
	if err != nil {
		panic(err)
	}
	return &AvroSerializer[TestType]{avroSchema}
}
