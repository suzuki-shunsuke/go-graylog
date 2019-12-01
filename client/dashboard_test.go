package client_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog/v8"
	"github.com/suzuki-shunsuke/go-graylog/v8/client"
	"github.com/suzuki-shunsuke/go-graylog/v8/testdata"
	"github.com/suzuki-shunsuke/go-graylog/v8/testutil"
)

func TestClient_CreateDashboard(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	_, err = cl.CreateDashboard(ctx, nil)
	require.NotNil(t, err, "dashboard should not be nil")

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Matcher: &flute.Matcher{
								Method: "POST",
								Path:   "/api/dashboards",
							},
							Tester: &flute.Tester{
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
								BodyJSONString: `{
								  "title": "dashboard title",
								  "description": "dashboard description"
								}`,
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 201,
									Header: http.Header{
										"Content-Type": []string{"application/json"},
									},
								},
								BodyString: `{
								  "dashboard_id": "5b39838b08813b0000000000"
								}`,
							},
						},
					},
				},
			},
		},
	})
	db := &graylog.Dashboard{
		Title:       "dashboard title",
		Description: "dashboard description",
	}
	_, err = cl.CreateDashboard(ctx, db)
	require.Nil(t, err)
	require.Equal(t, "5b39838b08813b0000000000", db.ID)
}

func TestClient_DeleteDashboard(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	// id required
	if _, err := client.DeleteDashboard(ctx, ""); err == nil {
		t.Fatal("id is required")
	}
	// invalid id
	if _, err := client.DeleteDashboard(ctx, "h"); err == nil {
		t.Fatal(`no dashboard with id "h" is found`)
	}
}

func TestClient_GetDashboard(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/dashboard.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method: "GET",
								Path:   "/api/dashboards/" + testdata.Dashboard.ID,
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: bodyStr,
							},
						},
					},
				},
			},
		},
	})

	_, _, err = cl.GetDashboard(ctx, "")
	require.NotNil(t, err)

	db, _, err := cl.GetDashboard(ctx, testdata.Dashboard.ID)
	require.Nil(t, err)
	require.Equal(t, testdata.Dashboard, db)
}

func TestClient_GetDashboards(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/dashboards.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method: "GET",
								Path:   "/api/dashboards",
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: bodyStr,
							},
						},
					},
				},
			},
		},
	})

	dbs, total, _, err := cl.GetDashboards(ctx)
	require.Nil(t, err)
	require.Equal(t, testdata.Dashboards.Dashboards, dbs)
	require.Equal(t, testdata.Dashboards.Total, total)
}

func TestClient_UpdateDashboard(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	dashboard := testutil.Dashboard()
	if _, err := client.CreateDashboard(ctx, dashboard); err != nil {
		t.Fatal(err)
	}
	// clean
	defer client.DeleteDashboard(ctx, dashboard.ID)

	dashboard.Description = "changed!"
	if _, err := client.UpdateDashboard(ctx, dashboard); err != nil {
		t.Fatal(err)
	}
	dashboard.ID = ""
	if _, err := client.UpdateDashboard(ctx, dashboard); err == nil {
		t.Fatal("id is required")
	}
	dashboard.ID = "h"
	if _, err := client.UpdateDashboard(ctx, dashboard); err == nil {
		t.Fatal(`no dashboard whose id is "h"`)
	}
	if _, err := client.UpdateDashboard(ctx, nil); err == nil {
		t.Fatal("dashboard is nil")
	}
}
