package client_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/suzuki-shunsuke/go-set/v6"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/client"
	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
)

func TestClient_GetLDAPSetting(t *testing.T) {
	ctx := context.Background()
	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		resp       string
		setting    *graylog.LDAPSetting
		checkErr   func(require.TestingT, interface{}, ...interface{})
	}{{
		statusCode: 200,
		resp: `{
  "enabled": true,
  "system_username": "CN=admin",
  "system_password": "***",
  "ldap_uri": "ldap://ldap.example.com:389/",
  "use_start_tls": false,
  "trust_all_certificates": false,
  "active_directory": false,
  "search_base": "OU=user,OU=foo,DC=example,DC=com",
  "search_pattern": "(cn={0})",
  "display_name_attribute": "displayname",
  "default_group": "Reader",
  "group_mapping": {
    "foo": "Reader"
  },
  "group_search_base": "",
  "group_id_attribute": "",
  "additional_default_groups": [],
  "group_search_pattern": ""
}`,
		setting: &graylog.LDAPSetting{
			Enabled:                 true,
			SystemUsername:          "admin",
			SystemPassword:          "***",
			LDAPURI:                 "ldap://ldap.example.com:389/",
			SearchBase:              "OU=user,OU=foo,DC=example,DC=com",
			SearchPattern:           "(cn={0})",
			DisplayNameAttribute:    "displayname",
			DefaultGroup:            "Reader",
			GroupMapping:            map[string]string{"foo": "Reader"},
			AdditionalDefaultGroups: set.StrSet{},
		},
		checkErr: require.Nil,
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
									Path:         "/api/system/ldap/settings",
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

		m, _, err := cl.GetLDAPSetting(ctx)
		if err != nil {
			require.Equal(t, d.setting, m)
		}
		d.checkErr(t, err)
	}
}

func TestClient_UpdateLDAPSetting(t *testing.T) {
	ctx := context.Background()
	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		body       string
		setting    *graylog.LDAPSetting
		checkErr   func(require.TestingT, interface{}, ...interface{})
	}{{
		statusCode: 204,
		body: `{
  "enabled": true,
  "use_start_tls": false,
  "trust_all_certificates": false,
  "active_directory": false,
  "system_username": "admin",
  "system_password": "***",
  "ldap_uri": "ldap://ldap.example.com:389/",
  "search_base": "OU=user,OU=foo,DC=example,DC=com",
  "search_pattern": "(cn={0})",
  "display_name_attribute": "displayname",
  "default_group": "Reader",
  "group_mapping": {
    "foo": "Reader"
  }
}`,
		setting: &graylog.LDAPSetting{
			Enabled:              true,
			SystemUsername:       "admin",
			SystemPassword:       "***",
			LDAPURI:              "ldap://ldap.example.com:389/",
			SearchBase:           "OU=user,OU=foo,DC=example,DC=com",
			SearchPattern:        "(cn={0})",
			DisplayNameAttribute: "displayname",
			DefaultGroup:         "Reader",
			GroupMapping:         map[string]string{"foo": "Reader"},
		},
		checkErr: require.Nil,
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
									Path:           "/api/system/ldap/settings",
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

		_, err := cl.UpdateLDAPSetting(ctx, d.setting)
		d.checkErr(t, err)
	}
}
