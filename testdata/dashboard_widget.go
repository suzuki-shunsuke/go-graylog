package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v9"
)

func UpdateDashboardWidgetStackedChart() graylog.Widget {
	return graylog.Widget{
		Description: "updated description",
		Config: &graylog.WidgetConfigUnknownType{
			T: "STACKED_CHART",
			Fields: map[string]interface{}{
				"interval": "minute",
				"timerange": map[string]interface{}{
					"type":  "relative",
					"range": 86400,
				},
				"renderer":      "area",
				"interpolation": "linear",
				"stream_id":     "000000000000000000000003",
				"series": []interface{}{
					map[string]interface{}{
						"query":                "labels_app: nginx-ingress AND response:[200 TO 399]",
						"field":                "response",
						"statistical_function": "count",
					},
					map[string]interface{}{
						"query":                "labels_app: nginx-ingress AND response:[500 TO 599]",
						"field":                "response",
						"statistical_function": "count",
					},
					map[string]interface{}{
						"query":                "labels_app: nginx-ingress AND response:[400 TO 499]",
						"field":                "response",
						"statistical_function": "count",
					},
				},
			},
		},
	}
}
