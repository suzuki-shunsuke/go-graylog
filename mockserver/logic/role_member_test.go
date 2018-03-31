package logic_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

func TestRoleMembers(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("basic", func(t *testing.T) {
		if _, _, err := lgc.RoleMembers("Admin"); err != nil {
			t.Fatal(err)
		}
	})
	// TOTO check graylog REST API spec
	// t.Run("name is required", func(t *testing.T) {
	// 	if _, _, err := lgc.RoleMembers(""); err == nil {
	// 		t.Fatal("name is required")
	// 	}
	// })
	// t.Run("not found", func(t *testing.T) {
	// 	if _, _, err := lgc.RoleMembers("h"); err == nil {
	// 		t.Fatal(`no role with name "h"`)
	// 	}
	// })
}
