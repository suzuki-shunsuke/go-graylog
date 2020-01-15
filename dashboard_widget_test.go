package graylog

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/go-ptr"
)

func TestWidget_MarshalJSON(t *testing.T) {
	data := []struct {
		title  string
		widget *Widget
		exp    string
	}{
		{
			title: "stream search result count",
			widget: &Widget{
				Description: "Stream search result count",
				Config: &WidgetConfigStreamSearchResultCount{
					Timerange: &Timerange{
						Type:  "relative",
						Range: 300,
					},
					LowerIsBetter: true,
					Trend:         true,
					StreamID:      "000000000000000000000001",
				},
				CacheTime: ptr.PInt(10),
			},
			exp: `{
  "type": "STREAM_SEARCH_RESULT_COUNT",
  "description": "Stream search result count",
  "config": {
		"timerange": {
      "type": "relative",
      "range": 300
    },
    "query": "",
    "lower_is_better": true,
    "trend": true,
    "stream_id": "000000000000000000000001"
  },
  "cache_time": 10
}`,
		},
		{
			title: "stacked chart",
			widget: &Widget{
				Description: "stacked chart",
				Config: &WidgetConfigUnknownType{
					T: "STACKED_CHART",
					Fields: map[string]interface{}{
						"interval": "hour",
						"timerange": map[string]interface{}{
							"type":  "relative",
							"range": 86400,
						},
						"renderer":      "bar",
						"interpolation": "linear",
						"series": []map[string]interface{}{
							{
								"query":                "",
								"field":                "AccessMask",
								"statistical_function": "count",
							},
						},
					},
				},
				CacheTime: ptr.PInt(10),
			},
			exp: `{
  "type": "STACKED_CHART",
  "description": "stacked chart",
  "config": {
    "interval": "hour",
    "timerange": {
      "type": "relative",
      "range": 86400
    },
    "renderer": "bar",
    "interpolation": "linear",
    "series": [
      {
        "query": "",
        "field": "AccessMask",
        "statistical_function": "count"
      }
    ]
  },
  "cache_time": 10
}`,
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			w, err := d.widget.MarshalJSON()
			require.Nil(t, err)
			require.JSONEq(t, d.exp, string(w))
		})
	}
}
