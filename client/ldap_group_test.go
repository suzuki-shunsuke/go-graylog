package client_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/suzuki-shunsuke/go-graylog/v10/client"
)

func TestClient_GetLDAPGroups(t *testing.T) {
	ctx := context.Background()
	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		resp       string
		groups     []string
		checkErr   func(require.TestingT, interface{}, ...interface{})
	}{{
		statusCode: 200,
		resp:       `["foo"]`,
		groups:     []string{"foo"},
		checkErr:   require.Nil,
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
									Method:       "GET",
									Path:         "/api/system/ldap/groups",
									PartOfHeader: getTestHeader(),
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

		m, _, err := cl.GetLDAPGroups(ctx)
		if err != nil {
			require.Equal(t, d.groups, m)
		}
		d.checkErr(t, err)
	}
}

func TestClient_GetLDAPGroupRoleMapping(t *testing.T) {
	ctx := context.Background()
	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		resp       string
		checkErr   func(require.TestingT, interface{}, ...interface{})
	}{{
		statusCode: 200,
		resp:       `{"foo": "Reader"}`,
		checkErr:   require.Nil,
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
									Method:       "GET",
									Path:         "/api/system/ldap/settings/groups",
									PartOfHeader: getTestHeader(),
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

		m, _, err := cl.GetLDAPGroupRoleMapping(ctx)
		if err != nil {
			require.Equal(t, d.resp, m)
		}
		d.checkErr(t, err)
	}
}

func TestClient_UpdateLDAPGroupRoleMapping(t *testing.T) {
	ctx := context.Background()
	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		body       string
		mapping    map[string]string
		checkErr   func(require.TestingT, interface{}, ...interface{})
	}{{
		statusCode: 204,
		body:       `{"foo": "Reader"}`,
		mapping:    map[string]string{"foo": "Reader"},
		checkErr:   require.Nil,
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
									Method:         "PUT",
									Path:           "/api/system/ldap/settings/groups",
									PartOfHeader:   getTestHeader(),
									BodyJSONString: d.body,
								},
								Response: &flute.Response{
									Base: http.Response{
										StatusCode: d.statusCode,
									},
								},
							},
						},
					},
				},
			},
		})

		_, err := cl.UpdateLDAPGroupRoleMapping(ctx, d.mapping)
		d.checkErr(t, err)
	}
}
