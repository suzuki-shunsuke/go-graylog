package client_test

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/assert"
	"github.com/suzuki-shunsuke/go-ptr"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
)

func TestCreateDashboardWidget(t *testing.T) {
	authName := os.Getenv("GRAYLOG_AUTH_NAME")
	authPass := os.Getenv("GRAYLOG_AUTH_PASSWORD")
	endpoint := os.Getenv("GRAYLOG_WEB_ENDPOINT_URI")
	dashboardID := "5b65868b08813b0001777af3"

	if endpoint == "" {
		defer gock.Off()
		endpoint = "http://example.com/api"
		client, err := client.NewClient(endpoint, authName, authPass)
		assert.Nil(t, err)
		data := []struct {
			body       string
			statusCode int
			resp       string
			widget     graylog.Widget
			checkErr   func(assert.TestingT, interface{}, ...interface{}) bool
		}{{
			body: `{
  "description": "Stream search result count",
  "type": "STREAM_SEARCH_RESULT_COUNT",
  "config": {
    "timerange": {
      "type": "relative",
			"range": 300
		}
  },
  "cache_time": 10
}`,
			statusCode: 201,
			resp:       `{"widget_id": "ee2532ce-6995-4b8b-8c2c-4de327c6cce4"}`,
			widget: graylog.Widget{
				Description: "Stream search result count",
				Type:        "STREAM_SEARCH_RESULT_COUNT",
				Config: &graylog.WidgetConfig{
					Timerange: &graylog.Timerange{
						Type:  "relative",
						Range: 300,
					},
				},
				CacheTime: ptr.PInt(10),
			},
			checkErr: assert.Nil,
		}}
		for _, d := range data {
			gock.New("http://example.com").
				Post(fmt.Sprintf("/api/dashboards/%s/widgets", dashboardID)).
				MatchType("json").BodyString(d.body).Reply(d.statusCode).
				BodyString(d.resp)
			w, _, err := client.CreateDashboardWidget(dashboardID, d.widget)
			if err != nil {
				assert.NotEqual(t, "", w.ID)
				d.widget.ID = w.ID
				assert.Equal(t, d.widget, w)
			}
			d.checkErr(t, err)
		}
	}
}
