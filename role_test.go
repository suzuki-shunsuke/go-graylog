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
	server, err := GetMockServer()
	if err != nil {
		t.Error("Failed to Get Mock Server", err)
		return
	}
	defer server.Server.Close()
	client, err := NewClient(server.Endpoint, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	params := &Role{Name: "foo", Permissions: []string{"*"}}
	role, err := client.CreateRole(params)
	if err != nil {
		t.Error("Failed to CreateRole", err)
		return
	}
	if role == nil {
		t.Error("client.CreateRole() == nil")
		return
	}
	if role.Name != "foo" {
		t.Errorf("role.Name == %s, wanted foo", role.Name)
	}
}

func TestGetRoles(t *testing.T) {
	server, err := GetMockServer()
	if err != nil {
		t.Error("Failed to Get Mock Server", err)
		return
	}
	defer server.Server.Close()
	client, err := NewClient(server.Endpoint, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	admin := dummyRole()
	exp := []Role{*admin}
	server.Roles[admin.Name] = *admin
	roles, err := client.GetRoles()
	if err != nil {
		t.Error("Failed to GetRoles", err)
		return
	}
	if !reflect.DeepEqual(roles, exp) {
		t.Errorf("client.GetRoles() == %v, wanted %v", roles, exp)
	}
}

func TestGetRole(t *testing.T) {
	server, err := GetMockServer()
	if err != nil {
		t.Error("Failed to Get Mock Server", err)
		return
	}
	defer server.Server.Close()
	client, err := NewClient(server.Endpoint, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	admin := dummyRole()
	server.Roles[admin.Name] = *admin
	role, err := client.GetRole(admin.Name)
	if err != nil {
		t.Error("Failed to GetRole", err)
		return
	}
	if !reflect.DeepEqual(*role, *admin) {
		t.Errorf("client.GetRole() == %v, wanted %v", role, admin)
	}
}

func TestUpdateRole(t *testing.T) {
	server, err := GetMockServer()
	if err != nil {
		t.Error("Failed to Get Mock Server", err)
		return
	}
	defer server.Server.Close()
	client, err := NewClient(server.Endpoint, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	admin := dummyRole()
	server.Roles[admin.Name] = *admin
	admin.Description = "changed!"
	updatedRole, err := client.UpdateRole(admin.Name, admin)
	if err != nil {
		t.Error("Failed to UpdateRole", err)
		return
	}
	if !reflect.DeepEqual(*updatedRole, *admin) {
		t.Errorf("client.UpdateRole() == %v, wanted %v", updatedRole, admin)
	}
}

func TestDeleteRole(t *testing.T) {
	server, err := GetMockServer()
	if err != nil {
		t.Error("Failed to Get Mock Server", err)
		return
	}
	defer server.Server.Close()
	client, err := NewClient(server.Endpoint, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	admin := dummyRole()
	server.Roles[admin.Name] = *admin
	err = client.DeleteRole(admin.Name)
	if err != nil {
		t.Error("Failed to DeleteRole", err)
		return
	}
}
