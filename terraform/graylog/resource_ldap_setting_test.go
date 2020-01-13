package graylog

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/v9/testdata"
)

func TestAccLDAPSetting(t *testing.T) {
	setEnv()

	tc := &testCase{
		t:       t,
		Name:    "ldap setting",
		GetPath: "/api/system/ldap/settings",

		CreateReqBodyMap: map[string]interface{}{},
		UpdateReqBodyMap: testdata.CreateLDAPSettingMap(),
		CreatedDataPath:  "ldap_setting/create.json",
		UpdatedDataPath:  "ldap_setting/create.json",
		CreateTFPath:     "ldap_setting/create.tf",
		UpdateTFPath:     "ldap_setting/create.tf",
	}
	tc.Test()
}
