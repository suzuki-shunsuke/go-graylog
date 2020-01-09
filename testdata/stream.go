package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v8"
)

func Stream() *graylog.Stream {
	return &graylog.Stream{
		ID:              "000000000000000000000003",
		Title:           "All system events",
		IndexSetID:      "5d84bfbe2ab79c000d35d4a9",
		CreatedAt:       "2019-09-20T12:02:06.078Z",
		CreatorUserID:   "admin",
		Description:     "Stream containing all system events created by Graylog",
		MatchingType:    "AND",
		Outputs:         []graylog.Output{}, // p0
		Rules:           []graylog.StreamRule{},
		AlertConditions: []graylog.AlertCondition{},
		AlertReceivers: &graylog.AlertReceivers{
			Emails: []string{},
			Users:  []string{},
		},
		Disabled:                       false,
		RemoveMatchesFromDefaultStream: true,
		IsDefault:                      false,
	}
}
