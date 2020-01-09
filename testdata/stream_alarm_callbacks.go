package testdata

import (
	"github.com/suzuki-shunsuke/go-set/v6"

	"github.com/suzuki-shunsuke/go-graylog/v8"
)

var (
	StreamAlarmCallbacks = &graylog.AlarmCallbacksBody{
		AlarmCallbacks: []graylog.AlarmCallback{
			{
				ID:            "5d84c1a92ab79c000d35d6d5",
				StreamID:      "5d84c1a92ab79c000d35d6ca",
				Title:         "test",
				CreatorUserID: "admin",
				CreatedAt:     "2019-09-20T12:10:17.793+0000",
				Configuration: &graylog.SlackAlarmCallbackConfiguration{
					Color:         "#FF0000",
					WebhookURL:    "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX",
					Channel:       "#general",
					IconURL:       "",
					Graylog2URL:   "https://graylog.example.com",
					IconEmoji:     "",
					UserName:      "Graylog",
					ProxyAddress:  "",
					CustomMessage: "${alert_condition.title}\\n\\n${foreach backlog message}\\n<https://graylog.example.com/streams/${stream.id}/search?rangetype=absolute&from=${message.timestamp}&to=${message.timestamp} | link> ${message.message}\\n${end}",
					BacklogItems:  5,
					LinkNames:     true,
					NotifyChannel: false,
				},
			},
			{
				ID:            "5d84c1a92ab79c000d35d6d4",
				StreamID:      "5d84c1a92ab79c000d35d6ca",
				Title:         "test",
				CreatorUserID: "admin",
				CreatedAt:     "2019-09-20T12:10:17.792+0000",
				Configuration: &graylog.HTTPAlarmCallbackConfiguration{
					URL: "https://example.com",
				},
			},
			{
				ID:            "5d84ca7f2ab79c000d35e083",
				StreamID:      "5d84c1a92ab79c000d35d6ca",
				Title:         "test",
				CreatorUserID: "admin",
				CreatedAt:     "2019-09-20T12:47:59.170+0000",
				Configuration: &graylog.EmailAlarmCallbackConfiguration{
					Sender:  "graylog@example.org",
					Subject: "Graylog alert for stream: ${stream.title}: ${check_result.resultDescription}",
					Body:    "##########\\nAlert Description: ${check_result.resultDescription}\\nDate: ${check_result.triggeredAt}\\nStream ID: ${stream.id}\\nStream title: ${stream.title}\\nStream description: ${stream.description}\\nAlert Condition Title: ${alertCondition.title}\\n${if stream_url}Stream URL: ${stream_url}${end}\\n\\nTriggered condition: ${check_result.triggeredCondition}\\n##########\\n\\n${if backlog}Last messages accounting for this alert:\\n${foreach backlog message}${message}\\n\\n${end}${else}<No backlog>\\n${end}\\n",
					UserReceivers: set.StrSet{
						"username": struct{}{},
					},
					EmailReceivers: set.StrSet{
						"graylog@example.com": struct{}{},
					},
				},
			},
		},
		Total: 3,
	}
)
