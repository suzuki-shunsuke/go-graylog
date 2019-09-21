package client_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/go-jsoneq/jsoneq"
	"github.com/suzuki-shunsuke/go-ptr"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
)

func TestClient_CreateDashboardWidget(t *testing.T) {
	ctx := context.Background()
	authName := os.Getenv("GRAYLOG_AUTH_NAME")
	authPass := os.Getenv("GRAYLOG_AUTH_PASSWORD")
	endpoint := os.Getenv("GRAYLOG_WEB_ENDPOINT_URI")
	dashboardID := "5b65868b08813b0001777af3"

	if endpoint == "" {
		defer gock.Off()
		endpoint = "http://example.com/api"
		client, err := client.NewClient(endpoint, authName, authPass)
		require.Nil(t, err)
		data := []struct {
			statusCode int
			body       interface{}
			resp       string
			widget     graylog.Widget
			checkErr   func(require.TestingT, interface{}, ...interface{})
		}{{
			body: map[string]interface{}{
				"description": "Stream search result count",
				"type":        "STREAM_SEARCH_RESULT_COUNT",
				"cache_time":  10,
				"config": map[string]interface{}{
					"timerange": map[string]interface{}{
						"type":  "relative",
						"range": 300,
					},
					"query":           "",
					"lower_is_better": true,
					"trend":           true,
					"stream_id":       "000000000000000000000001",
				}},
			statusCode: 201,
			resp:       `{"widget_id": "ee2532ce-6995-4b8b-8c2c-4de327c6cce4"}`,
			widget: graylog.Widget{
				Description: "Stream search result count",
				Config: &graylog.WidgetConfigStreamSearchResultCount{
					Timerange: &graylog.Timerange{
						Type:  "relative",
						Range: 300,
					},
					LowerIsBetter: true,
					Trend:         true,
					StreamID:      "000000000000000000000001",
				},
				CacheTime: ptr.PInt(10),
			},
			checkErr: require.Nil,
		}}
		for _, d := range data {
			gock.New("http://example.com").
				Post(fmt.Sprintf("/api/dashboards/%s/widgets", dashboardID)).
				MatchType("json").JSON(d.body).Reply(d.statusCode).
				BodyString(d.resp)
			w, _, err := client.CreateDashboardWidget(ctx, dashboardID, d.widget)
			d.checkErr(t, err)
			if err == nil {
				require.NotEqual(t, "", w.ID)
				d.widget.ID = w.ID
				require.Equal(t, d.widget, w)
			}
		}
	}
}

func Test_client_UpdateDashboardWidget(t *testing.T) {
	widget := &graylog.Widget{
		Config: &graylog.WidgetConfigStreamSearchResultCount{
			Timerange: &graylog.Timerange{
				Type:  "relative",
				Range: 300,
			},
			LowerIsBetter: true,
			Trend:         true,
			StreamID:      "000000000000000000000001",
		},
	}

	b, err := jsoneq.Equal(
		widget.Config,
		[]byte(`{
  "timerange": {
    "type": "relative",
    "range": 300
  },
  "lower_is_better": true,
  "trend": true,
  "stream_id": "000000000000000000000001",
  "query": ""
}`))
	require.Nil(t, err)
	require.True(t, b)

	b, err = jsoneq.Equal(
		map[string]interface{}{
			"config": widget.Config,
		},
		[]byte(`{
  "config": {
    "timerange": {
      "type": "relative",
      "range": 300
    },
    "lower_is_better": true,
    "trend": true,
    "stream_id": "000000000000000000000001",
    "query": ""
  }
}`))
	require.Nil(t, err)
	require.True(t, b)
}
