package graylog

import (
	"reflect"
	"testing"
)

func dummyRole() *Role {
	return &Role{
		Name:        "Admin",
		Description: "Grants all permissions for Graylog administrators (built-in)",
		Permissions: []string{"*"},
		ReadOnly:    true}
}

func TestCreateRole(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	params := &Role{Name: "foo", Permissions: []string{"*"}}
	role, err := client.CreateRole(params)
	if err != nil {
		t.Fatal("Failed to CreateRole", err)
	}
	if role == nil {
		t.Fatal("client.CreateRole() == nil")
	}
	if role.Name != "foo" {
		t.Fatalf("role.Name == %s, wanted foo", role.Name)
	}
}

func TestGetRoles(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	admin := dummyRole()
	exp := []Role{*admin}
	server.Roles[admin.Name] = *admin
	roles, err := client.GetRoles()
	if err != nil {
		t.Fatal("Failed to GetRoles", err)
	}
	if !reflect.DeepEqual(roles, exp) {
		t.Fatalf("client.GetRoles() == %v, wanted %v", roles, exp)
	}
}

func TestGetRole(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	admin := dummyRole()
	server.Roles[admin.Name] = *admin
	role, err := client.GetRole(admin.Name)
	if err != nil {
		t.Fatal("Failed to GetRole", err)
	}
	if !reflect.DeepEqual(*role, *admin) {
		t.Fatalf("client.GetRole() == %v, wanted %v", role, admin)
	}
}

func TestUpdateRole(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	admin := dummyRole()
	server.Roles[admin.Name] = *admin
	admin.Description = "changed!"
	updatedRole, err := client.UpdateRole(admin.Name, admin)
	if err != nil {
		t.Fatal("Failed to UpdateRole", err)
	}
	if !reflect.DeepEqual(*updatedRole, *admin) {
		t.Fatalf("client.UpdateRole() == %v, wanted %v", updatedRole, admin)
	}
}

func TestDeleteRole(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	admin := dummyRole()
	server.Roles[admin.Name] = *admin
	err = client.DeleteRole(admin.Name)
	if err != nil {
		t.Fatal("Failed to DeleteRole", err)
	}
}
