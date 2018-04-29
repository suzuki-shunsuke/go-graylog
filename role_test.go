package graylog_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestRoleNewUpdateParams(t *testing.T) {
	role := testutil.Role()
	prms := role.NewUpdateParams()
	if role.Name != prms.Name {
		t.Fatalf(`prms.Name = "%s", wanted "%s"`, prms.Name, role.Name)
	}
}
