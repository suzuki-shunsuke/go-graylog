package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v8"
)

func SlackStreamAlarmCallback() graylog.AlarmCallback {
	return graylog.AlarmCallback{
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
	}
}
