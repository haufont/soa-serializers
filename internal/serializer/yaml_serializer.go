package serializer

import "gopkg.in/yaml.v2"

type YAMLSerializer[TestType any] struct {
}

func (*YAMLSerializer[TestType]) Serialize(v *TestType) ([]byte, error) {
	return yaml.Marshal(v)
}

func (*YAMLSerializer[TestType]) Deserialize(data []byte, v *TestType) error {
	return yaml.Unmarshal(data, v)
}

func NewYAMLSerializer[TestType any]() *YAMLSerializer[TestType] {
	return &YAMLSerializer[TestType]{}
}
