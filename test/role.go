package test

import (
	"reflect"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestCreateRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	params := &graylog.Role{Name: "foo", Permissions: []string{"*"}}
	role, _, err := client.CreateRole(params)
	if err != nil {
		t.Fatal("Failed to CreateRole", err)
	}
	if role == nil {
		t.Fatal("client.CreateRole() == nil")
	}
	if role.Name != "foo" {
		t.Fatalf("role.Name == %s, wanted foo", role.Name)
	}
	_, ei, err := client.CreateRole(params)
	if err == nil {
		t.Fatal("user name must be unique")
	}
	if ei.Response.StatusCode != 400 {
		t.Fatal("status code must be 400")
	}

	params.Name = ""
	if _, _, err := client.CreateRole(params); err == nil {
		t.Fatal("user name is required")
	}

	params.Name = "bar"
	params.Permissions = nil
	if _, _, err := client.CreateRole(params); err == nil {
		t.Fatal("user permissions are required")
	}
}

func TestGetRoles(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	roles, _, err := client.GetRoles()
	if err != nil {
		t.Fatal("Failed to GetRoles", err)
	}
	if roles == nil {
		t.Fatal("client.GetRoles() is nil")
	}
	if len(roles) != 1 {
		t.Fatalf("len(roles) == %d, wanted 1", len(roles))
	}
}

func TestGetRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	role, _, err := client.GetRole("Admin")
	if err != nil {
		t.Fatal("Failed to GetRole", err)
	}
	if role.Name != "Admin" {
		t.Fatalf(`role name is "%s", wanted "Admin"`, role.Name)
	}
	if _, _, err := client.GetRole(""); err == nil {
		t.Fatal("role name is required")
	}
	if _, _, err := client.GetRole("h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}

func TestUpdateRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	admin, err := server.GetRole("Admin")
	if err != nil {
		t.Fatal(err)
	}
	admin.Description = "changed!"
	updatedRole, _, err := client.UpdateRole(admin.Name, admin)
	if err != nil {
		t.Fatal("Failed to UpdateRole", err)
	}
	if !reflect.DeepEqual(*updatedRole, *admin) {
		t.Fatalf("client.UpdateRole() == %v, wanted %v", updatedRole, admin)
	}
	if _, _, err := client.UpdateRole("", admin); err == nil {
		t.Fatal("role name is required")
	}
	if _, _, err := client.UpdateRole("h", admin); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}

	copiedAdmin := *admin
	admin.Name = ""
	if _, _, err := client.UpdateRole(copiedAdmin.Name, admin); err == nil {
		t.Fatal("role name is required")
	}
	admin.Name = copiedAdmin.Name
	admin.Permissions = nil
	if _, _, err := client.UpdateRole(copiedAdmin.Name, admin); err == nil {
		t.Fatal("role permissions is required")
	}
}

func TestDeleteRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	if _, err = client.DeleteRole("Admin"); err != nil {
		t.Fatal("Failed to DeleteRole", err)
	}
	if _, err = client.DeleteRole(""); err == nil {
		t.Fatal("role name is required")
	}
	if _, err = client.DeleteRole("h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}