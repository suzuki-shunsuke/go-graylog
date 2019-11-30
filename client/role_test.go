package client_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog/client/v8"
	"github.com/suzuki-shunsuke/go-graylog/testdata"
	"github.com/suzuki-shunsuke/go-graylog/testutil/v8"
)

func TestClient_CreateRole(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	role := testutil.Role()
	client.DeleteRole(ctx, role.Name)
	// nil check
	if _, err := client.CreateRole(ctx, nil); err == nil {
		t.Fatal("role is nil")
	}
	if _, err := client.CreateRole(ctx, role); err != nil {
		t.Fatal(err)
	}
	if _, err := client.DeleteRole(ctx, role.Name); err != nil {
		t.Fatal(err)
	}
	// error check
	role.Name = ""
	if _, err := client.CreateRole(ctx, role); err == nil {
		t.Fatal("role name is empty")
	}
}

func TestClient_GetRoles(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/roles.json")
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
								Path:   "/api/roles",
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

	roles, _, _, err := cl.GetRoles(ctx)
	require.Nil(t, err)
	require.Equal(t, testdata.Roles.Roles, roles)
}

func TestClient_GetRole(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/role.json")
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
								Path:   "/api/roles/Views Manager",
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

	_, _, err = cl.GetRole(ctx, "")
	require.NotNil(t, err)

	role, _, err := cl.GetRole(ctx, testdata.Role.Name)
	require.Nil(t, err)
	require.Equal(t, testdata.Role, role)
}

func TestClient_UpdateRole(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	role := testutil.Role()
	client.DeleteRole(ctx, role.Name)
	if _, _, err := client.UpdateRole(ctx, role.Name, role.NewUpdateParams()); err == nil {
		t.Fatal("role should be deleted")
	}
	if _, err := client.CreateRole(ctx, role); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteRole(ctx, role.Name)
	if _, _, err := client.UpdateRole(ctx, role.Name, role.NewUpdateParams()); err != nil {
		t.Fatal(err)
	}
	if _, _, err := client.UpdateRole(ctx, "", role.NewUpdateParams()); err == nil {
		t.Fatal("role name is required")
	}
	name := role.Name
	role.Name = ""
	if _, _, err := client.UpdateRole(ctx, name, role.NewUpdateParams()); err == nil {
		t.Fatal("role name is required")
	}
}

func TestClient_DeleteRole(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
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
								Method: "DELETE",
								Path:   "/api/roles/foo",
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 204,
								},
							},
						},
					},
				},
			},
		},
	})

	_, err = cl.DeleteRole(ctx, "")
	require.NotNil(t, err)
	_, err = cl.DeleteRole(ctx, "foo")
	require.Nil(t, err)
}
