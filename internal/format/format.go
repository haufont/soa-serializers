package format

import "strings"

type Format int64

const (
	Unknown Format = iota
	Native
	XML
	JSON
	ProtoBuf
	ApacheAvro
	YAML
	MessagePack
)

func (f Format) Name() string {
	switch f {
	case Native:
		return "Native(gob)"
	case XML:
		return "XML"
	case JSON:
		return "JSON"
	case ProtoBuf:
		return "GoogleProtocolBuffers"
	case ApacheAvro:
		return "ApacheAvro"
	case YAML:
		return "YAML"
	case MessagePack:
		return "MessagePack"
	}
	panic("unknown format")
}

func (f Format) String() string {
	switch f {
	case Native:
		return "native"
	case XML:
		return "xml"
	case JSON:
		return "json"
	case ProtoBuf:
		return "protobuf"
	case ApacheAvro:
		return "avro"
	case YAML:
		return "yaml"
	case MessagePack:
		return "message_pack"
	}
	return "unknown format"
}

func ParseFormat(s string) Format {
	switch s {
	case Native.String():
		return Native
	case XML.String():
		return XML
	case JSON.String():
		return JSON
	case ProtoBuf.String():
		return ProtoBuf
	case ApacheAvro.String():
		return ApacheAvro
	case YAML.String():
		return YAML
	case MessagePack.String():
		return MessagePack
	}
	return Unknown
}

func ParseFormatOrPanic(s string) Format {
	f := ParseFormat(s)
	if f == Unknown {
		panic("unknown format")
	}
	return f
}

func AllEnumToString() string {
	enums := []string{
		Native.String(),
		XML.String(),
		JSON.String(),
		ProtoBuf.String(),
		ApacheAvro.String(),
		YAML.String(),
		MessagePack.String(),
	}
	return strings.Join([]string{"[", strings.Join(enums, ", "), "]"}, "")
}
