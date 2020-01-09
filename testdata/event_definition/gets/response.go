package gets

import (
	"github.com/suzuki-shunsuke/go-graylog/v8"
)

func Response() *graylog.EventDefinitionsBody {
	return &graylog.EventDefinitionsBody{
		EventDefinitions: []graylog.EventDefinition{
			{
				ID:          "5de5a9e7a1de18000cdfe192",
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
					"group_by":         []interface{}{},
					"query":            "test",
					"search_within_ms": 60000,
					"series":           []interface{}{},
					"streams": []interface{}{
						"000000000000000000000001",
					},
					"type": "aggregation-v1",
				},
			},
			{
				ID:          "5de5aac1a1de18000cdfe2b3",
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
					"group_by":         []interface{}{},
					"query":            "test",
					"search_within_ms": 60000,
					"series":           []interface{}{},
					"streams": []interface{}{
						"000000000000000000000001",
					},
					"type": "aggregation-v1",
				},
			},
			{
				ID:          "5de59d56a1de18000cdfd776",
				Title:       "test",
				Description: "Migrated message count alert condition",
				Priority:    2,
				Alert:       true,
				FieldSpec:   map[string]graylog.EventDefinitionFieldSpec{},
				NotificationSettings: graylog.EventDefinitionNotificationSettings{
					GracePeriodMS: 0,
					BacklogSize:   2,
				},
				Notifications: []graylog.EventDefinitionNotification{
					{
						NotificationID: "5de59d56a1de18000cdfd774",
					},
					{
						NotificationID: "5de59d56a1de18000cdfd770",
					},
					{
						NotificationID: "5de59d56a1de18000cdfd772",
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
						"expression": map[string]interface{}{
							"expr": ">",
							"left": map[string]interface{}{
								"expr": "number-ref",
								"ref":  "40aea221-88d2-492a-a3f0-62da7aa68f56",
							},
							"right": map[string]interface{}{
								"expr":  "number",
								"value": 0,
							},
						},
					},
					"execute_every_ms": 60000,
					"group_by":         []interface{}{},
					"query":            "message:\"hoge hoge\"",
					"search_within_ms": 60000,
					"series": []interface{}{
						map[string]interface{}{
							"field":    nil,
							"function": "count",
							"id":       "40aea221-88d2-492a-a3f0-62da7aa68f56",
						},
					},
					"streams": []interface{}{
						"5de4fcf7a1de1800127e2fbe",
					},
					"type": "aggregation-v1",
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
