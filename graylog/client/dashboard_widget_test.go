package client_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/suzuki-shunsuke/go-ptr"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/client"
	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/testdata"
)

func TestClient_CreateDashboardWidget(t *testing.T) {
	ctx := context.Background()
	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	dashboardID := "5b65868b08813b0001777af3"

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
		cl.SetHTTPClient(&http.Client{
			Transport: &flute.Transport{
				T: t,
				Services: []flute.Service{
					{
						Endpoint: "http://example.com",
						Routes: []flute.Route{
							{
								Tester: &flute.Tester{
									Method:       "POST",
									Path:         fmt.Sprintf("/api/dashboards/%s/widgets", dashboardID),
									PartOfHeader: getTestHeader(),
									BodyJSON:     d.body,
								},
								Response: &flute.Response{
									Base: http.Response{
										StatusCode: d.statusCode,
									},
									BodyString: d.resp,
								},
							},
						},
					},
				},
			},
		})

		w, _, err := cl.CreateDashboardWidget(ctx, dashboardID, d.widget)
		d.checkErr(t, err)
		if err == nil {
			require.NotEqual(t, "", w.ID)
			d.widget.ID = w.ID
			require.Equal(t, d.widget, w)
		}
	}
}

func TestClient_UpdateDashboardWidget(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)
	dashboardID := "5b65868b08813b0001777af3"
	widgetID := "5b65868b0881300000000000"

	_, err = cl.UpdateDashboardWidget(ctx, "", graylog.Widget{
		ID: widgetID,
	})
	require.NotNil(t, err, "dashboard id is required")

	_, err = cl.UpdateDashboardWidget(ctx, dashboardID, graylog.Widget{})
	require.NotNil(t, err, "dashboard widget id is required")

	buf, err := ioutil.ReadFile("../testdata/dashboard_widget/stacked_chart/update.json")
	require.Nil(t, err)

	widget := testdata.UpdateDashboardWidgetStackedChart()
	widget.ID = widgetID

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Matcher: &flute.Matcher{
								Method: "PUT",
							},
							Tester: &flute.Tester{
								Path:           "/api/dashboards/" + dashboardID + "/widgets/" + widgetID,
								PartOfHeader:   getTestHeader(),
								BodyJSONString: string(buf),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 204,
									Header: http.Header{
										"Content-Type": []string{"application/json"},
									},
								},
							},
						},
					},
				},
			},
		},
	})

	_, err = cl.UpdateDashboardWidget(ctx, dashboardID, widget)
	require.Nil(t, err)
}
