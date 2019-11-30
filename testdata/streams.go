package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v8"
)

var (
	Streams = &graylog.StreamsBody{
		Total: 4,
		Streams: []graylog.Stream{
			{
				ID:            "5d84c1a92ab79c000d35d6ca",
				Title:         "test",
				IndexSetID:    "5d84bf242ab79c000d691b7f",
				CreatedAt:     "2019-09-20T12:10:17.506Z",
				CreatorUserID: "admin",
				Description:   "test",
				MatchingType:  "AND",
				Outputs:       []graylog.Output{}, // p0
				Rules: []graylog.StreamRule{
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
				AlertConditions: []graylog.AlertCondition{
					{
						ID:            "56f9f507-601d-4a54-a2f4-4bda93bb8492",
						CreatorUserID: "admin",
						CreatedAt:     "2019-09-20T12:10:17.792+0000",
						Title:         "test",
						InGrace:       false,
						Parameters: graylog.FieldContentAlertConditionParameters{
							Grace:               0,
							Backlog:             2,
							RepeatNotifications: false,
							Field:               "message",
							Value:               "hoge hoge",
							Query:               "*",
						},
					},
				},
				AlertReceivers: &graylog.AlertReceivers{
					Emails: []string{},
					Users:  []string{},
				},
				Disabled:                       true,
				RemoveMatchesFromDefaultStream: false,
				IsDefault:                      false,
			},
			{
				ID:              "000000000000000000000002",
				Title:           "All events",
				IndexSetID:      "5d84bfbe2ab79c000d35d4a6",
				CreatedAt:       "2019-09-20T12:02:06.058Z",
				CreatorUserID:   "admin",
				Description:     "Stream containing all events created by Graylog",
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
			},
			{
				ID:              "000000000000000000000001",
				Title:           "All messages",
				IndexSetID:      "5d84bf242ab79c000d691b7f",
				CreatedAt:       "2019-09-20T11:59:32.311Z",
				CreatorUserID:   "local:admin",
				Description:     "Stream containing all messages",
				MatchingType:    "AND",
				Outputs:         []graylog.Output{},
				Rules:           []graylog.StreamRule{},
				AlertConditions: []graylog.AlertCondition{},
				AlertReceivers: &graylog.AlertReceivers{
					Emails: []string{},
					Users:  []string{},
				},
				Disabled:                       false,
				RemoveMatchesFromDefaultStream: false,
				IsDefault:                      true,
			},
			{
				ID:              "000000000000000000000003",
				Title:           "All system events",
				IndexSetID:      "5d84bfbe2ab79c000d35d4a9",
				CreatedAt:       "2019-09-20T12:02:06.078Z",
				CreatorUserID:   "admin",
				Description:     "Stream containing all system events created by Graylog",
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
			},
		},
	}
)
