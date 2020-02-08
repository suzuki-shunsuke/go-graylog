package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
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

func CreateStreamRuleReqBodyMap() map[string]interface{} {
	return map[string]interface{}{
		"field":       "tag",
		"value":       "4",
		"description": "test",
		"type":        1,
		"inverted":    false,
	}
}

func UpdateStreamRuleReqBodyMap() map[string]interface{} {
	return map[string]interface{}{
		"field":       "tag",
		"value":       "4",
		"description": "updated description",
		"type":        1,
		"inverted":    false,
	}
}
