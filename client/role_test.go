package client_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/suzuki-shunsuke/go-set/v6"

	"github.com/suzuki-shunsuke/go-graylog/v9/client"
	"github.com/suzuki-shunsuke/go-graylog/v9/testdata"
)

func TestClient_CreateRole(t *testing.T) {
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
								Test:   testRoleBody,
								Method: "POST",
								Path:   "/api/roles",
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 201,
								},
								BodyString: bodyStr,
							},
						},
					},
				},
			},
		},
	})

	if _, err := cl.CreateRole(ctx, nil); err == nil {
		t.Fatal("role is nil")
	}
	role := testdata.Role
	if _, err := cl.CreateRole(ctx, &role); err != nil {
		t.Fatal(err)
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
	require.Equal(t, testdata.Roles().Roles, roles)
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
	require.Equal(t, &testdata.Role, role)
}

func testRoleBody(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
	var body map[string]interface{}
	require.Nil(t, json.NewDecoder(req.Body).Decode(&body))

	perms := set.NewStrSet()
	for _, p := range body["permissions"].([]interface{}) {
		perms.Add(p.(string))
	}
	body["permissions"] = perms
	require.Equal(t, map[string]interface{}{
		"name":        "Views Manager",
		"description": "Allows reading and writing all views and extended searches (built-in)",
		"permissions": set.NewStrSet(
			"view:edit",
			"extendedsearch:use",
			"view:create",
			"extendedsearch:create",
			"view:read",
			"view:use"),
		"read_only": true,
	}, body)
}

func TestClient_UpdateRole(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/role.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	role := testdata.Role

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Test:   testRoleBody,
								Method: "PUT",
								Path:   "/api/roles/" + role.Name,
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

	if _, _, err := cl.UpdateRole(ctx, role.Name, role.NewUpdateParams()); err != nil {
		t.Fatal(err)
	}
	if _, _, err := cl.UpdateRole(ctx, "", role.NewUpdateParams()); err == nil {
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
