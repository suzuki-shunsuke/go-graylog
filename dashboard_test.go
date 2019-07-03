package graylog

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/go-jsoneq/jsoneq"
	"github.com/suzuki-shunsuke/go-ptr"
)

func TestWidget_MarshalJSON(t *testing.T) {
	widget := &Widget{
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
	}
	d, err := widget.MarshalJSON()
	require.Nil(t, err)
	b, err := jsoneq.Equal(d, []byte(`{
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
}`))
	require.Nil(t, err)
	require.True(t, b)
}
