package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v9"
)

func StreamRule() *graylog.StreamRule {
	return &graylog.StreamRule{
		ID:          "5d84c1a92ab79c000d35d6d7",
		StreamID:    "5d84c1a92ab79c000d35d6ca",
		Field:       "tag",
		Value:       "4",
		Description: "test",
		Type:        1,
		Inverted:    false,
	}
}
