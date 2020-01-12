package graylog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/suzuki-shunsuke/go-set/v6"
)

func genTestRoleBody(role map[string]interface{}, bodyString string, store *bodyStore) func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
	return func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
		body := map[string]interface{}{}
		require.Nil(t, json.NewDecoder(req.Body).Decode(&body))

		perms := set.NewStrSet()
		for _, p := range body["permissions"].([]interface{}) {
			perms.Add(p.(string))
		}
		body["permissions"] = perms
		assert.Equal(t, role, body)
		store.Set(bodyString)
	}
}

func TestAccRole(t *testing.T) {
	setEnv()

	roleBody, err := ioutil.ReadFile("../../testdata/role/role.json")
	require.Nil(t, err)

	updateRoleBody, err := ioutil.ReadFile("../../testdata/role/update_role.json")
	require.Nil(t, err)

	roleTF, err := ioutil.ReadFile("../../testdata/role/role.tf")
	require.Nil(t, err)

	updateRoleTF, err := ioutil.ReadFile("../../testdata/role/update_role.tf")
	require.Nil(t, err)

	store := newBodyStore("")

	defaultTransport := http.DefaultClient.Transport
	defer func() {
		http.DefaultClient.Transport = defaultTransport
	}()
	http.DefaultClient.Transport = &flute.Transport{
		T: t,
		Services: []flute.Service{
			{
				Endpoint: "http://example.com",
				Routes: []flute.Route{
					{
						Matcher: &flute.Matcher{
							Method: "GET",
						},
						Tester: &flute.Tester{
							Path:         "/api/roles/foo",
							PartOfHeader: getTestHeader(),
						},
						Response: &flute.Response{
							Response: func(req *http.Request) (*http.Response, error) {
								return &http.Response{
									StatusCode: 200,
									Body:       ioutil.NopCloser(strings.NewReader(store.Get())),
								}, nil
							},
						},
					},
					{
						Matcher: &flute.Matcher{
							Method: "POST",
						},
						Tester: &flute.Tester{
							Path:         "/api/roles",
							PartOfHeader: getTestHeader(),
							Test: genTestRoleBody(map[string]interface{}{
								"name":        "foo",
								"description": "Allows reading and writing all views and extended searches (built-in)",
								"permissions": set.NewStrSet(
									"extendedsearch:create",
									"extendedsearch:use",
									"view:create",
									"view:edit",
									"view:read",
									"view:use",
								),
								"read_only": true,
							}, string(roleBody), store),
						},
						Response: &flute.Response{
							Base: http.Response{
								StatusCode: 201,
							},
							BodyString: string(roleBody),
						},
					},
					{
						Matcher: &flute.Matcher{
							Method: "PUT",
						},
						Tester: &flute.Tester{
							Path:         "/api/roles/foo",
							PartOfHeader: getTestHeader(),
							Test: genTestRoleBody(map[string]interface{}{
								"name":        "foo",
								"description": "updated description",
								"permissions": set.NewStrSet(
									"extendedsearch:create",
									"extendedsearch:use",
									"view:edit",
									"view:read",
									"view:use",
								),
								"read_only": false,
							}, string(updateRoleBody), store),
						},
						Response: &flute.Response{
							Base: http.Response{
								StatusCode: 201,
							},
							BodyString: string(roleBody),
						},
					},
					{
						Matcher: &flute.Matcher{
							Method: "DELETE",
						},
						Tester: &flute.Tester{
							Path:         "/api/roles/foo",
							PartOfHeader: getTestHeader(),
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
	}

	resource.Test(t, resource.TestCase{
		Providers: getTestProviders(),
		Steps: []resource.TestStep{
			{
				Config: string(roleTF),
			},
			{
				Config: string(updateRoleTF),
			},
		},
	})
}
