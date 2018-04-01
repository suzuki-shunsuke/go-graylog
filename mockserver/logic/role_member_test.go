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
	t.Run("name is required", func(t *testing.T) {
		if _, _, err := lgc.RoleMembers(""); err == nil {
			t.Fatal("name is required")
		}
	})
	t.Run("not found", func(t *testing.T) {
		if _, _, err := lgc.RoleMembers("h"); err == nil {
			t.Fatal(`no role with name "h"`)
		}
	})
}

func TestAddUserToRole(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("user name is required", func(t *testing.T) {
		if _, err := lgc.AddUserToRole("", "Admin"); err == nil {
			t.Fatal("user name is required")
		}
	})
	t.Run("role name is required", func(t *testing.T) {
		if _, err := lgc.AddUserToRole("admin", ""); err == nil {
			t.Fatal("role name is required")
		}
	})
}

func TestRemoveUserFromRole(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("user name is required", func(t *testing.T) {
		if _, err := lgc.RemoveUserFromRole("", "Admin"); err == nil {
			t.Fatal("user name is required")
		}
	})
	t.Run("role name is required", func(t *testing.T) {
		if _, err := lgc.RemoveUserFromRole("admin", ""); err == nil {
			t.Fatal("role name is required")
		}
	})
}
