package graylog_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/testdata"
)

func TestRoleNewUpdateParams(t *testing.T) {
	role := testdata.Role
	prms := role.NewUpdateParams()
	if role.Name != prms.Name {
		t.Fatalf(`prms.Name = "%s", wanted "%s"`, prms.Name, role.Name)
	}
}
