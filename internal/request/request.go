package request

import (
	"errors"
	"serialize-benchmarker/internal/format"
	"strings"
)

const (
	GetResult    = "get_result"
	GetResultAll = "get_result all"
)

func ParseRequestWithFormatOrAll(request string) (format.Format, bool, error) {
	if len(request) < len(GetResult)+1 || !strings.HasPrefix(request, GetResult) {
		return format.Unknown, false, errors.New(UnknownRequest(request, "get_result {format or all}"))
	}
	if request == GetResultAll {
		return format.Unknown, true, nil
	}
	request = request[len(GetResult)+1:]
	f := format.ParseFormat(request)
	if f == format.Unknown {
		return format.Unknown, false, errors.New(strings.Join([]string{"Unknown format: \"", request, "\""}, ""))
	}
	return f, false, nil
}

func UnknownRequest(actually string, expected string) string {
	return strings.Join([]string{"Unknown request \"", actually, "\". Use \"", expected, "\""}, "")
}

func UnknownSimpleRequest(request string) string {
	return UnknownRequest(request, GetResult)
}
