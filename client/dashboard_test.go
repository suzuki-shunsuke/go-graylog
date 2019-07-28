package client_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestCreateDashboard(t *testing.T) {
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
								Header: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
								},
								BodyJSONString: `{
								  "title": "dashboard title",
								  "description": "dashboard description"
								}`,
							},
							Response: &flute.Response{
								StatusCode: 201,
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

func TestDeleteDashboard(t *testing.T) {
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

func TestGetDashboard(t *testing.T) {
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
	defer client.DeleteDashboard(ctx, dashboard.ID)

	r, _, err := client.GetDashboard(ctx, dashboard.ID)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("dashboard is nil")
	}
	if r.ID != dashboard.ID {
		t.Fatalf(`dashboard.ID = "%s", wanted "%s"`, r.ID, dashboard.ID)
	}
	if _, _, err := client.GetDashboard(ctx, ""); err == nil {
		t.Fatal("id is required")
	}
	if _, _, err := client.GetDashboard(ctx, "h"); err == nil {
		t.Fatal("dashboard should not be found")
	}
}

func TestGetDashboards(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, _, _, err := client.GetDashboards(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestUpdateDashboard(t *testing.T) {
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
