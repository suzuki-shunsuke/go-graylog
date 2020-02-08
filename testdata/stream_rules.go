package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v10"
)

func StreamRules() *graylog.StreamRulesBody {
	return &graylog.StreamRulesBody{
		Total: 1,
		StreamRules: []graylog.StreamRule{
			{
				ID:          "5d84c1a92ab79c000d35d6d7",
				StreamID:    "5d84c1a92ab79c000d35d6ca",
				Field:       "tag",
				Value:       "4",
				Description: "test",
				Type:        1,
				Inverted:    false,
			},
		},
	}
}
