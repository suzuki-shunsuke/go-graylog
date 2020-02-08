package graylog

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/go-ptr"
)

func TestDashbaord_UnmarshalJSON(t *testing.T) {
	data := []struct {
		title string
		exp   *Dashboard
		body  string
	}{
		{
			title: "normal",
			exp: &Dashboard{
				ID:            "5d1aa378d497000000000000",
				Title:         "test",
				CreatorUserID: "admin",
				Description:   "test",
				CreatedAt:     "2019-07-02T00:21:12.623Z",
				Positions: []DashboardWidgetPosition{
					{
						WidgetID: "b3be33f5-0c60-409b-b1aa-000000000000",
						Width:    2,
						Col:      6,
						Row:      1,
						Height:   2,
					},
				},
				Widgets: []Widget{
					{
						CreatorUserID: "admin",
						CacheTime:     ptr.PInt(10),
						Description:   "Search result graph",
						ID:            "b3be33f5-0c60-409b-b1aa-000000000000",
						Config: &WidgetConfigSearchResultChart{
							Timerange: &Timerange{
								Type:  "relative",
								Range: 300,
							},
							Interval: "minute",
							StreamID: "000000000000000000000001",
							Query:    "",
						},
					},
				},
			},
			body: `{
      "creator_user_id": "admin",
      "description": "test",
      "created_at": "2019-07-02T00:21:12.623Z",
      "positions": {
        "b3be33f5-0c60-409b-b1aa-000000000000": {
          "width": 2,
          "col": 6,
          "row": 1,
          "height": 2
        }
      },
      "id": "5d1aa378d497000000000000",
      "title": "test",
      "widgets": [
        {
          "creator_user_id": "admin",
          "cache_time": 10,
          "description": "Search result graph",
          "id": "b3be33f5-0c60-409b-b1aa-000000000000",
          "type": "SEARCH_RESULT_CHART",
          "config": {
            "timerange": {
              "type": "relative",
              "range": 300
            },
            "interval": "minute",
            "stream_id": "000000000000000000000001",
            "query": ""
          }
        }
      ]
    }`,
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			dashboard := &Dashboard{}
			require.Nil(t, json.Unmarshal([]byte(d.body), dashboard))
			require.Equal(t, d.exp, dashboard)
		})
	}
}
