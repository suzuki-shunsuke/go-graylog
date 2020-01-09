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

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	id := "5b39838b08813b0000000000"

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Matcher: &flute.Matcher{
								Method: "DELETE",
								Path:   "/api/dashboards/" + id,
							},
							Tester: &flute.Tester{
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
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

	// id required
	if _, err := cl.DeleteDashboard(ctx, ""); err == nil {
		t.Fatal("id is required")
	}
	// invalid id
	if _, err := cl.DeleteDashboard(ctx, id); err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetDashboard(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/dashboard.json")
	require.Nil(t, err)

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
								Path:   "/api/dashboards/" + testdata.Dashboard().ID,
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
								BodyString: string(buf),
							},
						},
					},
				},
			},
		},
	})

	_, _, err = cl.GetDashboard(ctx, "")
	require.NotNil(t, err)

	db, _, err := cl.GetDashboard(ctx, testdata.Dashboard().ID)
	require.Nil(t, err)
	d := testdata.Dashboard()
	require.Equal(t, &d, db)
}

func TestClient_GetDashboards(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/dashboards.json")
	require.Nil(t, err)

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
								BodyString: string(buf),
							},
						},
					},
				},
			},
		},
	})

	dbs, total, _, err := cl.GetDashboards(ctx)
	require.Nil(t, err)
	require.Equal(t, testdata.Dashboards().Dashboards, dbs)
	require.Equal(t, testdata.Dashboards().Total, total)
}

func TestClient_UpdateDashboard(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	_, err = cl.UpdateDashboard(ctx, nil)
	require.NotNil(t, err, "dashboard should not be nil")

	buf, err := ioutil.ReadFile("../testdata/create_dashboard.json")
	require.Nil(t, err)

	ds := testdata.Dashboard()

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
								Path:   "/api/dashboards/" + ds.ID,
							},
							Tester: &flute.Tester{
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
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

	if _, err := cl.UpdateDashboard(ctx, &ds); err != nil {
		t.Fatal(err)
	}
	ds.ID = ""
	if _, err := cl.UpdateDashboard(ctx, &ds); err == nil {
		t.Fatal("id is required")
	}
	if _, err := cl.UpdateDashboard(ctx, nil); err == nil {
		t.Fatal("dashboard is nil")
	}
}
