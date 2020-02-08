package client_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/client"
)

func TestClient_GetRoleMembers(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/role_member/role_members.json")
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
							Matcher: &flute.Matcher{
								Path: "/api/roles/Admin/members",
							},
							Tester: &flute.Tester{
								Method:       "GET",
								PartOfHeader: getTestHeader(),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: bodyStr,
							},
						},
						{
							Matcher: &flute.Matcher{
								Path: "/api/roles/h/members",
							},
							Tester: &flute.Tester{
								Method:       "GET",
								PartOfHeader: getTestHeader(),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 404,
									Header: http.Header{
										"Content-Type": []string{"application/json"},
									},
								},
								BodyString: `{
  "type": "ApiError",
  "message": "Couldn't find role h"
}`,
							},
						},
					},
				},
			},
		},
	})

	if _, _, err := cl.GetRoleMembers(ctx, "Admin"); err != nil {
		t.Fatal("Failed to GetRoleMembers", err)
	}
	if _, _, err := cl.GetRoleMembers(ctx, ""); err == nil {
		t.Fatal("name is required")
	}
	if _, _, err := cl.GetRoleMembers(ctx, "h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}

func TestClient_AddUserToRole(t *testing.T) {
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
							Matcher: &flute.Matcher{
								Path: "/api/roles/Admin/members/test",
							},
							Tester: &flute.Tester{
								Method:       "PUT",
								PartOfHeader: getTestHeader(),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 204,
								},
							},
						},
						{
							Matcher: &flute.Matcher{
								Path: "/api/roles/h/members/test",
							},
							Tester: &flute.Tester{
								Method:       "PUT",
								PartOfHeader: getTestHeader(),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 404,
									Header: http.Header{
										"Content-Type": []string{"application/json"},
									},
								},
								BodyString: `{
  "type": "ApiError",
  "message": "Couldn't find role h"
}`,
							},
						},
					},
				},
			},
		},
	})

	if _, err = cl.AddUserToRole(ctx, "test", "Admin"); err != nil {
		// Cannot modify local root user, this is a bug.
		t.Fatal(err)
	}
	if _, err = cl.AddUserToRole(ctx, "", "Admin"); err == nil {
		t.Fatal("user name is required")
	}
	if _, err = cl.AddUserToRole(ctx, "admin", ""); err == nil {
		t.Fatal("role name is required")
	}
	if _, err = cl.AddUserToRole(ctx, "test", "h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}

func TestClient_RemoveUserFromRole(t *testing.T) {
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
							Matcher: &flute.Matcher{
								Path: "/api/roles/Admin/members/test",
							},
							Tester: &flute.Tester{
								Method:       "DELETE",
								PartOfHeader: getTestHeader(),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 204,
								},
							},
						},
						{
							Matcher: &flute.Matcher{
								Path: "/api/roles/h/members/test",
							},
							Tester: &flute.Tester{
								Method:       "DELETE",
								PartOfHeader: getTestHeader(),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 404,
									Header: http.Header{
										"Content-Type": []string{"application/json"},
									},
								},
								BodyString: `{
  "type": "ApiError",
  "message": "Couldn't find role h"
}`,
							},
						},
					},
				},
			},
		},
	})

	if _, err = cl.RemoveUserFromRole(ctx, "test", "Admin"); err != nil {
		// Cannot modify local root user, this is a bug.
		t.Fatal(err)
	}
	if _, err = cl.RemoveUserFromRole(ctx, "", "Admin"); err == nil {
		t.Fatal("user name is required")
	}
	if _, err = cl.RemoveUserFromRole(ctx, "test", ""); err == nil {
		t.Fatal("role name is required")
	}
	if _, err = cl.RemoveUserFromRole(ctx, "test", "h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}
