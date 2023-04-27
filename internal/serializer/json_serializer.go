package serializer

import (
	"encoding/json"
)

type JSONSerializer[TestType any] struct {
}

func (*JSONSerializer[TestType]) Serialize(v *TestType) ([]byte, error) {
	return json.Marshal(v)
}

func (*JSONSerializer[TestType]) Deserialize(data []byte, v *TestType) error {
	return json.Unmarshal(data, v)
}

func NewJSONSerializer[TestType any]() *JSONSerializer[TestType] {
	return &JSONSerializer[TestType]{}
}
