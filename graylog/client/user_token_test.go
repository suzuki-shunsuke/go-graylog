package client_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/client"
	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
)

func TestClient_DeleteUserToken(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	token := "76niusu98vg2229l3r4kobllu6g3fd83p195hkehp64n5pfsuj3"

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Matcher: &flute.Matcher{
								Path: "/api/users/foo/tokens/" + token,
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
								Path: "/api/users/foo/tokens/h",
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
  "message": "Couldn't find user token h"
}`,
							},
						},
					},
				},
			},
		},
	})

	if _, err := cl.DeleteUserToken(ctx, "", token); err == nil {
		t.Fatal("username is required")
	}
	if _, err := cl.DeleteUserToken(ctx, "foo", ""); err == nil {
		t.Fatal("user token is required")
	}
	if _, err := cl.DeleteUserToken(ctx, "foo", token); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.DeleteUserToken(ctx, "foo", "h"); err == nil {
		t.Fatal(`no user token with name "h" is found`)
	}
}

func TestClient_CreateUserToken(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	_, _, err = cl.CreateUserToken(ctx, "", "test")
	require.NotNil(t, err, "user name is required")
	_, _, err = cl.CreateUserToken(ctx, "foo", "")
	require.NotNil(t, err, "token name is required")

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
								Path:   "/api/users/foo/tokens/test",
							},
							Tester: &flute.Tester{
								PartOfHeader:   getTestHeader(),
								BodyJSONString: `{}`,
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 201,
								},
								BodyString: `{
  "name": "test",
  "token": "76niusu98vg2229l3r4kobllu6g3fd83p195hkehp64n5pfsuj3",
  "last_access": "1970-01-01T00:00:00.000Z"
}`,
							},
						},
					},
				},
			},
		},
	})
	token, _, err := cl.CreateUserToken(ctx, "foo", "test")
	require.Nil(t, err)
	require.Equal(t, &graylog.UserToken{
		Name:       "test",
		Token:      "76niusu98vg2229l3r4kobllu6g3fd83p195hkehp64n5pfsuj3",
		LastAccess: "1970-01-01T00:00:00.000Z",
	}, token)
}

func TestClient_GetUserTokens(t *testing.T) {
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
								Path: "/api/users/foo/tokens",
							},
							Tester: &flute.Tester{
								Method:       "GET",
								PartOfHeader: getTestHeader(),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: `{
  "tokens": [
    {
      "name": "test",
      "token": "76niusu98vg2229l3r4kobllu6g3fd83p195hkehp64n5pfsuj3",
      "last_access": "1970-01-01T00:00:00.000Z"
    }
	]
}`,
							},
						},
					},
				},
			},
		},
	})

	tokens, _, err := cl.GetUserTokens(ctx, "foo")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, []graylog.UserToken{
		{
			Name:       "test",
			Token:      "76niusu98vg2229l3r4kobllu6g3fd83p195hkehp64n5pfsuj3",
			LastAccess: "1970-01-01T00:00:00.000Z",
		},
	}, tokens)
}
