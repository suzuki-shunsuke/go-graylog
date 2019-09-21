package client_test

import (
	"context"
	"os"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/go-graylog/client"
)

func TestClient_GetLDAPGroups(t *testing.T) {
	ctx := context.Background()
	authName := os.Getenv("GRAYLOG_AUTH_NAME")
	authPass := os.Getenv("GRAYLOG_AUTH_PASSWORD")
	endpoint := os.Getenv("GRAYLOG_WEB_ENDPOINT_URI")

	if endpoint == "" {
		defer gock.Off()
		endpoint = "http://example.com/api"
		client, err := client.NewClient(endpoint, authName, authPass)
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
			gock.New("http://example.com").
				Get("/api/system/ldap/groups").
				MatchType("json").Reply(d.statusCode).BodyString(d.resp)
			m, _, err := client.GetLDAPGroups(ctx)
			if err != nil {
				require.Equal(t, d.groups, m)
			}
			d.checkErr(t, err)
		}
	}
}

func TestClient_GetLDAPGroupRoleMapping(t *testing.T) {
	ctx := context.Background()
	authName := os.Getenv("GRAYLOG_AUTH_NAME")
	authPass := os.Getenv("GRAYLOG_AUTH_PASSWORD")
	endpoint := os.Getenv("GRAYLOG_WEB_ENDPOINT_URI")

	if endpoint == "" {
		defer gock.Off()
		endpoint = "http://example.com/api"
		client, err := client.NewClient(endpoint, authName, authPass)
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
			gock.New("http://example.com").
				Get("/api/system/ldap/settings/groups").
				MatchType("json").Reply(d.statusCode).BodyString(d.resp)
			m, _, err := client.GetLDAPGroupRoleMapping(ctx)
			if err != nil {
				require.Equal(t, d.resp, m)
			}
			d.checkErr(t, err)
		}
	}
}

func TestClient_UpdateLDAPGroupRoleMapping(t *testing.T) {
	ctx := context.Background()
	authName := os.Getenv("GRAYLOG_AUTH_NAME")
	authPass := os.Getenv("GRAYLOG_AUTH_PASSWORD")
	endpoint := os.Getenv("GRAYLOG_WEB_ENDPOINT_URI")

	if endpoint == "" {
		defer gock.Off()
		endpoint = "http://example.com/api"
		client, err := client.NewClient(endpoint, authName, authPass)
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
			gock.New("http://example.com").
				Put("/api/system/ldap/settings/groups").
				MatchType("json").BodyString(d.body).Reply(d.statusCode)
			_, err := client.UpdateLDAPGroupRoleMapping(ctx, d.mapping)
			d.checkErr(t, err)
		}
	}
}
