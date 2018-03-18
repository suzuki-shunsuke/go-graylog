package graylog_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestCreateRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	role := testutil.Role()
	client.DeleteRole(role.Name)
	// nil check
	if _, err := client.CreateRole(nil); err == nil {
		t.Fatal("role is nil")
	}
	if _, err := client.CreateRole(role); err != nil {
		t.Fatal(err)
	}
	if _, err := client.DeleteRole(role.Name); err != nil {
		t.Fatal(err)
	}
	// error check
	role.Name = ""
	if _, err := client.CreateRole(role); err == nil {
		t.Fatal("role name is empty")
	}
}

func TestGetRoles(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	roles, _, err := client.GetRoles()
	if err != nil {
		t.Fatal(err)
	}
	if roles == nil {
		t.Fatal("roles is nil")
	}
	if len(roles) == 0 {
		t.Fatal("roles is empty")
	}
}

func TestGetRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	role := testutil.Role()
	client.DeleteRole(role.Name)
	if _, _, err := client.GetRole(role.Name); err == nil {
		t.Fatal("role should be deleted")
	}
	if _, err := client.CreateRole(role); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteRole(role.Name)
	r, _, err := client.GetRole(role.Name)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("roles is nil")
	}
	if r.Name != role.Name {
		t.Fatalf(`role.Name = "%s", wanted "%s"`, r.Name, role.Name)
	}
	if _, _, err := client.GetRole(""); err == nil {
		t.Fatal("role name is required")
	}
}

func TestUpdateRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	role := testutil.Role()
	client.DeleteRole(role.Name)
	if _, err := client.UpdateRole(role.Name, role); err == nil {
		t.Fatal("role should be deleted")
	}
	if _, err := client.CreateRole(role); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteRole(role.Name)
	if _, err := client.UpdateRole(role.Name, role); err != nil {
		t.Fatal(err)
	}
	if _, err := client.UpdateRole("", role); err == nil {
		t.Fatal("role name is required")
	}
	name := role.Name
	role.Name = ""
	if _, err := client.UpdateRole(name, role); err == nil {
		t.Fatal("role name is required")
	}
}

func TestDeleteRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, err := client.DeleteRole(""); err == nil {
		t.Fatal("role name is required")
	}
	if _, err := client.DeleteRole("h"); err == nil {
		t.Fatal(`no role with name "h" is found`)
	}
}
