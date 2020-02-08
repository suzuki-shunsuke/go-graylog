package testdata

import (
	"github.com/suzuki-shunsuke/go-set/v6"

	"github.com/suzuki-shunsuke/go-graylog/v10"
)

func EmailStreamAlarmCallback() graylog.AlarmCallback {
	return graylog.AlarmCallback{
		ID:            "5d84ca7f2ab79c000d35e083",
		StreamID:      "5d84c1a92ab79c000d35d6ca",
		Title:         "test",
		CreatorUserID: "admin",
		CreatedAt:     "2019-09-20T12:47:59.170+0000",
		Configuration: &graylog.EmailAlarmCallbackConfiguration{
			Sender:         "graylog@example.org",
			Subject:        "Graylog alert for stream: ${stream.title}: ${check_result.resultDescription}",
			Body:           "##########\\nAlert Description: ${check_result.resultDescription}\\nDate: ${check_result.triggeredAt}\\nStream ID: ${stream.id}\\nStream title: ${stream.title}\\nStream description: ${stream.description}\\nAlert Condition Title: ${alertCondition.title}\\n${if stream_url}Stream URL: ${stream_url}${end}\\n\\nTriggered condition: ${check_result.triggeredCondition}\\n##########\\n\\n${if backlog}Last messages accounting for this alert:\\n${foreach backlog message}${message}\\n\\n${end}${else}<No backlog>\\n${end}\\n",
			UserReceivers:  set.NewStrSet("username"),
			EmailReceivers: set.NewStrSet("graylog@example.com"),
		},
	}
}
