package create

import (
	"github.com/suzuki-shunsuke/go-graylog/v8"
)

var Request = &graylog.EventDefinition{
	ID:          "",
	Title:       "new-event-definition",
	Description: "",
	Priority:    2,
	Alert:       true,
	FieldSpec: map[string]graylog.EventDefinitionFieldSpec{
		"test": {
			DataType: "string",
			Providers: []interface{}{
				map[string]interface{}{
					"require_values": false,
					"template":       "test",
					"type":           "template-v1",
				},
			},
		},
	},
	NotificationSettings: graylog.EventDefinitionNotificationSettings{
		GracePeriodMS: 0,
		BacklogSize:   0,
	},
	Notifications: []graylog.EventDefinitionNotification{
		{
			NotificationID: "5de5a365a1de18000cdfdf49",
		},
	},
	Storage: []interface{}{
		map[string]interface{}{
			"streams": []interface{}{
				"000000000000000000000002",
			},
			"type": "persist-to-streams-v1",
		},
	},
	Config: map[string]interface{}{
		"conditions": map[string]interface{}{
			"expression": nil,
		},
		"execute_every_ms": 60000,
		"group_by":         []interface{}{}, // p0
		"query":            "test",
		"search_within_ms": 60000,
		"series":           []interface{}{},
		"streams": []interface{}{
			"000000000000000000000001",
		},
		"type": "aggregation-v1",
	},
}
