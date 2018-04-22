package plain_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/plain"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestHasRole(t *testing.T) {
	store := plain.NewStore("")
	ok, err := store.HasRole("foo")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("role foo should not exist")
	}
}

func TestGetRole(t *testing.T) {
	store := plain.NewStore("")
	role, err := store.GetRole("foo")
	if err != nil {
		t.Fatal(err)
	}
	if role != nil {
		t.Fatal("role foo should not exist")
	}
}

func TestGetRoles(t *testing.T) {
	store := plain.NewStore("")
	roles, _, err := store.GetRoles()
	if err != nil {
		t.Fatal(err)
	}
	if len(roles) != 0 {
		t.Fatal("roles should be nil or empty array")
	}
}

func TestAddRole(t *testing.T) {
	store := plain.NewStore("")
	role := testutil.Role()
	if err := store.AddRole(role); err != nil {
		t.Fatal(err)
	}
	r, err := store.GetRole(role.Name)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("role is nil")
	}
}

func TestUpdateRole(t *testing.T) {
	store := plain.NewStore("")
	role := testutil.Role()
	if err := store.AddRole(role); err != nil {
		t.Fatal(err)
	}
	r, err := store.GetRole(role.Name)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("role is nil")
	}
	role.Description += " changed"
	if _, err := store.UpdateRole(role.Name, role.NewUpdateParams()); err != nil {
		t.Fatal(err)
	}
	r, err = store.GetRole(role.Name)
	if err != nil {
		t.Fatal(err)
	}
	if role.Description != r.Description {
		t.Fatalf(`role.Description = "%s", wanted "%s"`, r.Description, role.Description)
	}
}

func TestDeleteRole(t *testing.T) {
	store := plain.NewStore("")
	if err := store.DeleteRole("foo"); err != nil {
		t.Fatal(err)
	}
	role := testutil.Role()
	if err := store.AddRole(role); err != nil {
		t.Fatal(err)
	}
	if err := store.DeleteRole(role.Name); err != nil {
		t.Fatal(err)
	}
	ok, err := store.HasRole(role.Name)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("role should be deleted")
	}
}
