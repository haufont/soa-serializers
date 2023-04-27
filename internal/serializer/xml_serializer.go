package serializer

import "encoding/xml"

type XMLSerializer[TestType any] struct {
}

func (*XMLSerializer[TestType]) Serialize(v *TestType) ([]byte, error) {
	return xml.Marshal(v)
}

func (*XMLSerializer[TestType]) Deserialize(data []byte, v *TestType) error {
	return xml.Unmarshal(data, v)
}

func NewXMLSerializer[TestType any]() *XMLSerializer[TestType] {
	return &XMLSerializer[TestType]{}
}
