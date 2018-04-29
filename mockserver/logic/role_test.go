package logic_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestAddRole(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("basic", func(t *testing.T) {
		role := testutil.Role()
		if _, err := lgc.AddRole(role); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("name is required", func(t *testing.T) {
		role := testutil.Role()
		role.Name = ""
		if _, err := lgc.AddRole(role); err == nil {
			t.Fatal("role name is required")
		}
	})
}

func TestGetRoles(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	roles, _, _, err := lgc.GetRoles()
	if err != nil {
		t.Fatal(err)
	}
	if len(roles) == 0 {
		t.Fatal("len(roles) == 0")
	}
}

func TestGetRole(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("basic", func(t *testing.T) {
		if _, _, err := lgc.GetRole("Admin"); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("name is required", func(t *testing.T) {
		if _, _, err := lgc.GetRole(""); err == nil {
			t.Fatal("name is required")
		}
	})
	t.Run("not found", func(t *testing.T) {
		if _, _, err := lgc.GetRole("h"); err == nil {
			t.Fatal(`no role with name "h"`)
		}
	})
}

func TestUpdateRole(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	role := testutil.Role()
	if _, err := lgc.AddRole(role); err != nil {
		t.Fatal(err)
	}
	name := role.Name

	t.Run("basic", func(t *testing.T) {
		role.Description += " changed!"
		if _, _, err := lgc.UpdateRole(name, role.NewUpdateParams()); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("name is required", func(t *testing.T) {
		if _, _, err := lgc.UpdateRole("", role.NewUpdateParams()); err == nil {
			t.Fatal("name is required")
		}
	})
	t.Run("name is required", func(t *testing.T) {
		role.Name = ""
		if _, _, err := lgc.UpdateRole(name, role.NewUpdateParams()); err == nil {
			t.Fatal("name is required")
		}
	})
	t.Run("not found", func(t *testing.T) {
		role.Name = name
		if _, _, err := lgc.UpdateRole("h", role.NewUpdateParams()); err == nil {
			t.Fatal("not found")
		}
	})
	t.Run("permissions is required", func(t *testing.T) {
		role.Permissions = nil
		if _, _, err := lgc.UpdateRole(name, role.NewUpdateParams()); err == nil {
			t.Fatal("permissions is required")
		}
	})
	t.Run("nil", func(t *testing.T) {
		if _, _, err := lgc.UpdateRole(name, nil); err == nil {
			t.Fatal("role is nil")
		}
	})
}

func TestDeleteRole(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	role := testutil.Role()
	if _, err := lgc.AddRole(role); err != nil {
		t.Fatal(err)
	}
	t.Run("basic", func(t *testing.T) {
		if _, err := lgc.DeleteRole(role.Name); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("name is required", func(t *testing.T) {
		if _, err := lgc.DeleteRole(""); err == nil {
			t.Fatal("name is required")
		}
	})
	t.Run("not found", func(t *testing.T) {
		if _, err := lgc.DeleteRole("h"); err == nil {
			t.Fatal("not found")
		}
	})
}
