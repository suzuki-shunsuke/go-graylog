package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v10"
)

func Stream() *graylog.Stream {
	return &graylog.Stream{
		ID:              "000000000000000000000003",
		Title:           "test",
		IndexSetID:      "5d84bfbe2ab79c000d35d4a9",
		CreatedAt:       "2019-09-20T12:02:06.078Z",
		CreatorUserID:   "admin",
		Description:     "test",
		MatchingType:    "AND",
		Outputs:         []graylog.Output{},
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

func CreateStreamReqBodyMap() map[string]interface{} {
	return map[string]interface{}{
		"title":                              "test",
		"description":                        "test",
		"index_set_id":                       "5d84bfbe2ab79c000d35d4a9",
		"matching_type":                      "AND",
		"remove_matches_from_default_stream": true,
	}
}

func UpdateStreamReqBodyMap() map[string]interface{} {
	return map[string]interface{}{
		"title":                              "updated title",
		"description":                        "updated description",
		"index_set_id":                       "5d84bfbe2ab79c000d35d4a9",
		"matching_type":                      "AND",
		"remove_matches_from_default_stream": true,
	}
}
