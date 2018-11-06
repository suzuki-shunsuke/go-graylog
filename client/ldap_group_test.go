package client_test

import (
	"os"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/assert"

	"github.com/suzuki-shunsuke/go-graylog/client"
)

func TestGetLDAPGroups(t *testing.T) {
	authName := os.Getenv("GRAYLOG_AUTH_NAME")
	authPass := os.Getenv("GRAYLOG_AUTH_PASSWORD")
	endpoint := os.Getenv("GRAYLOG_WEB_ENDPOINT_URI")

	if endpoint == "" {
		defer gock.Off()
		endpoint = "http://example.com/api"
		client, err := client.NewClient(endpoint, authName, authPass)
		assert.Nil(t, err)
		data := []struct {
			statusCode int
			resp       string
			groups     []string
			checkErr   func(assert.TestingT, interface{}, ...interface{}) bool
		}{{
			statusCode: 200,
			resp:       `["foo"]`,
			groups:     []string{"foo"},
			checkErr:   assert.Nil,
		}}
		for _, d := range data {
			gock.New("http://example.com").
				Get("/api/system/ldap/groups").
				MatchType("json").Reply(d.statusCode).BodyString(d.resp)
			m, _, err := client.GetLDAPGroups()
			if err != nil {
				assert.Equal(t, d.groups, m)
			}
			d.checkErr(t, err)
		}
	}
}

func TestGetLDAPGroupRoleMapping(t *testing.T) {
	authName := os.Getenv("GRAYLOG_AUTH_NAME")
	authPass := os.Getenv("GRAYLOG_AUTH_PASSWORD")
	endpoint := os.Getenv("GRAYLOG_WEB_ENDPOINT_URI")

	if endpoint == "" {
		defer gock.Off()
		endpoint = "http://example.com/api"
		client, err := client.NewClient(endpoint, authName, authPass)
		assert.Nil(t, err)
		data := []struct {
			statusCode int
			resp       string
			checkErr   func(assert.TestingT, interface{}, ...interface{}) bool
		}{{
			statusCode: 200,
			resp:       `{"foo": "Reader"}`,
			checkErr:   assert.Nil,
		}}
		for _, d := range data {
			gock.New("http://example.com").
				Get("/api/system/ldap/settings/groups").
				MatchType("json").Reply(d.statusCode).BodyString(d.resp)
			m, _, err := client.GetLDAPGroupRoleMapping()
			if err != nil {
				assert.Equal(t, d.resp, m)
			}
			d.checkErr(t, err)
		}
	}
}

func TestUpdateLDAPGroupRoleMapping(t *testing.T) {
	authName := os.Getenv("GRAYLOG_AUTH_NAME")
	authPass := os.Getenv("GRAYLOG_AUTH_PASSWORD")
	endpoint := os.Getenv("GRAYLOG_WEB_ENDPOINT_URI")

	if endpoint == "" {
		defer gock.Off()
		endpoint = "http://example.com/api"
		client, err := client.NewClient(endpoint, authName, authPass)
		assert.Nil(t, err)
		data := []struct {
			statusCode int
			body       string
			mapping    map[string]string
			checkErr   func(assert.TestingT, interface{}, ...interface{}) bool
		}{{
			statusCode: 204,
			body:       `{"foo": "Reader"}`,
			mapping:    map[string]string{"foo": "Reader"},
			checkErr:   assert.Nil,
		}}
		for _, d := range data {
			gock.New("http://example.com").
				Put("/api/system/ldap/settings/groups").
				MatchType("json").BodyString(d.body).Reply(d.statusCode)
			_, err := client.UpdateLDAPGroupRoleMapping(d.mapping)
			d.checkErr(t, err)
		}
	}
}
