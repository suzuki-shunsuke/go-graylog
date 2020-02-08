package terraform

import (
	"encoding/json"
	"errors"
	"io"
	"testing"

	"github.com/suzuki-shunsuke/go-set/v6"
)

func TestAccRole(t *testing.T) {
	setEnv()

	tc := &testCase{
		t:          t,
		Name:       "role",
		CreatePath: "/api/roles",
		GetPath:    "/api/roles/foo",

		ConvertReqBody: func(body io.Reader) (map[string]interface{}, error) {
			m := map[string]interface{}{}
			if err := json.NewDecoder(body).Decode(&m); err != nil {
				return nil, err
			}
			if perms, ok := m["permissions"]; ok {
				arr, ok := perms.([]interface{})
				if !ok {
					return nil, errors.New("permissions should be list")
				}
				perms := make([]string, len(arr))
				for i, p := range arr {
					a, ok := p.(string)
					if !ok {
						return nil, errors.New("permissions should be list of string")
					}
					perms[i] = a
				}
				m["permissions"] = set.NewStrSet(perms...)
			}
			return m, nil
		},

		CreateReqBodyMap: map[string]interface{}{
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
		},
		UpdateReqBodyMap: map[string]interface{}{
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
		},
		CreatedDataPath:    "role/role.json",
		UpdatedDataPath:    "role/update_role.json",
		CreateRespBodyPath: "role/role.json",
		UpdateRespBodyPath: "role/role.json",
		CreateTFPath:       "role/role.tf",
		UpdateTFPath:       "role/update_role.tf",
	}
	tc.Test()
}
