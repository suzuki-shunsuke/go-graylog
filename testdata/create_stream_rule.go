package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v8"
)

func CreateStreamRule() graylog.StreamRule {
	return graylog.StreamRule{
		StreamID:    "5d84c1a92ab79c000d35d6ca",
		Field:       "tag",
		Value:       "4",
		Description: "test",
		Type:        1,
		Inverted:    false,
	}
}
