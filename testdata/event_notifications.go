package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v10"
)

func EventNotificationsBody() *graylog.EventNotificationsBody {
	return &graylog.EventNotificationsBody{
		EventNotifications: []graylog.EventNotification{
			{
				ID:          "5de59d56a1de18000cdfd770",
				Title:       "test",
				Description: "Migrated legacy alarm callback",
				Config: map[string]interface{}{
					"callback_type": "org.graylog2.alarmcallbacks.HTTPAlarmCallback",
					"configuration": map[string]interface{}{
						"url": "https://example.com",
					},
					"type": "legacy-alarm-callback-notification-v1",
				},
			},
			{
				ID:          "5de59d56a1de18000cdfd772",
				Title:       "test",
				Description: "Migrated legacy alarm callback",
				Config: map[string]interface{}{
					"callback_type": "org.graylog2.alarmcallbacks.EmailAlarmCallback",
					"configuration": map[string]interface{}{
						"body": "##########\\nAlert Description: ${check_result.resultDescription}\\nDate: ${check_result.triggeredAt}\\nStream ID: ${stream.id}\\nStream title: ${stream.title}\\nStream description: ${stream.description}\\nAlert Condition Title: ${alertCondition.title}\\n${if stream_url}Stream URL: ${stream_url}${end}\\n\\nTriggered condition: ${check_result.triggeredCondition}\\n##########\\n\\n${if backlog}Last messages accounting for this alert:\\n${foreach backlog message}${message}\\n\\n${end}${else}<No backlog>\\n${end}\\n",
						"email_receivers": []interface{}{
							"graylog@example.com",
						},
						"sender":  "graylog@example.org",
						"subject": "Graylog alert for stream: ${stream.title}: ${check_result.resultDescription}",
						"user_receivers": []interface{}{
							"username",
						},
					},
					"type": "legacy-alarm-callback-notification-v1",
				},
			},
			{
				ID:          "5de59d56a1de18000cdfd774",
				Title:       "test",
				Description: "Migrated legacy alarm callback",
				Config: map[string]interface{}{
					"callback_type": "org.graylog2.plugins.slack.callback.SlackAlarmCallback",
					"configuration": map[string]interface{}{
						"backlog_items":  5,
						"channel":        "#general",
						"color":          "#FF0000",
						"custom_message": "${alert_condition.title}\\n\\n${foreach backlog message}\\n<https://graylog.example.com/streams/${stream.id}/search?rangetype=absolute&from=${message.timestamp}&to=${message.timestamp} | link> ${message.message}\\n${end}",
						"graylog2_url":   "https://graylog.example.com",
						"link_names":     true,
						"notify_channel": false,
						"user_name":      "Graylog",
						"webhook_url":    "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX",
					},
					"type": "legacy-alarm-callback-notification-v1",
				},
			},
		},
		Total:      3,
		Page:       1,
		PerPage:    50,
		Count:      3,
		GrandTotal: 3,
		Query:      "",
	}
}
