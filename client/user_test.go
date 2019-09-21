package client_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestClient_DeleteUser(t *testing.T) {
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
								Path: "/api/users/foo",
							},
							Tester: &flute.Tester{
								Method: "DELETE",
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
						{
							Matcher: &flute.Matcher{
								Path: "/api/users/h",
							},
							Tester: &flute.Tester{
								Method: "DELETE",
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
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
  "message": "Couldn't find user h"
}`,
							},
						},
					},
				},
			},
		},
	})

	if _, err := cl.DeleteUser(ctx, ""); err == nil {
		t.Fatal("username is required")
	}
	if _, err := cl.DeleteUser(ctx, "foo"); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.DeleteUser(ctx, "h"); err == nil {
		t.Fatal(`no user with name "h" is found`)
	}
}

func TestClient_CreateUser(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	_, err = cl.CreateUser(ctx, nil)
	require.NotNil(t, err, "user should not be nil")

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
								Path:   "/api/users",
							},
							Tester: &flute.Tester{
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
								BodyJSONString: `{
								  "username": "test",
								  "email": "test@example.com",
								  "full_name": "test test",
									"password": "password",
									"roles": ["Reader"],
									"permissions": []
								}`,
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 201,
								},
							},
						},
					},
				},
			},
		},
	})
	user := &graylog.User{
		Username: "test",
		Email:    "test@example.com",
		FullName: "test test",
		Password: "password",
		Roles:    set.NewStrSet("Reader"),
		External: true,
	}
	_, err = cl.CreateUser(ctx, user)
	require.Nil(t, err)
}

func TestClient_GetUsers(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	users, _, err := client.GetUsers(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if users == nil {
		t.Fatal("users is nil")
	}
	if len(users) == 0 {
		t.Fatal("users is empty")
	}
}

func TestClient_GetUser(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	user := testutil.User()
	client.DeleteUser(ctx, user.Username)
	if _, _, err := client.GetUser(ctx, user.Username); err == nil {
		t.Fatal("user should be deleted")
	}
	if _, err := client.CreateUser(ctx, user); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteUser(ctx, user.Username)
	u, _, err := client.GetUser(ctx, user.Username)
	if err != nil {
		t.Fatal(err)
	}
	if u == nil {
		t.Fatal("user is nil")
	}
	if u.Username != user.Username {
		t.Fatalf(`user.Username = "%s", wanted "%s"`, u.Username, user.Username)
	}
	if _, _, err := client.GetUser(ctx, ""); err == nil {
		t.Fatal("user name is required")
	}
}

func TestClient_UpdateUser(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	_, err = cl.UpdateUser(ctx, nil)
	require.NotNil(t, err, "user should not be nil")

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
								Path:   "/api/users/test",
							},
							Tester: &flute.Tester{
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
								BodyJSONString: `{
								  "username": "test",
									"roles": ["Reader"]
								}`,
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
	user := &graylog.UserUpdateParams{
		Username: "test",
		Roles:    set.NewStrSet("Reader"),
	}
	_, err = cl.UpdateUser(ctx, user)
	require.Nil(t, err)
}
