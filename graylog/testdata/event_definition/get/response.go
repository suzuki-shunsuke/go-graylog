package get

import (
	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
)

func Response() *graylog.EventDefinition {
	return &graylog.EventDefinition{
		ID:       "5dea491ba1de18000d4bbcce",
		Title:    "new-event-definition",
		Priority: 1,
		Alert:    true,
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
		Notifications: []graylog.EventDefinitionNotification{},
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
					"expr": "<",
					"left": map[string]interface{}{
						"expr": "number-ref",
						"ref":  "9dfd012c-4f4d-417b-80d8-f7ebda2020a3",
					},
					"right": map[string]interface{}{
						"expr":  "number",
						"value": 0,
					},
				},
			},
			"execute_every_ms": 60000,
			"group_by": []interface{}{
				"alert",
			},
			"query":            "test",
			"search_within_ms": 60000,
			"series": []interface{}{
				map[string]interface{}{
					"field":    "alert",
					"function": "avg",
					"id":       "9dfd012c-4f4d-417b-80d8-f7ebda2020a3",
				},
			},
			"streams": []interface{}{
				"000000000000000000000001",
			},
			"type": "aggregation-v1",
		},
	}
}
