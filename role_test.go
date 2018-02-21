package graylog

import (
	"reflect"
	"testing"
)

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
	roles, err := client.GetRoles()
	if err != nil {
		t.Error("Failed to GetRoles", err)
		return
	}
	exp := []Role{
		{
			Name:        "Admin",
			Description: "Grants all permissions for Graylog administrators (built-in)",
			Permissions: []string{"*"},
			ReadOnly:    true},
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
	role, err := client.GetRole("Admin")
	if err != nil {
		t.Error("Failed to GetRole", err)
		return
	}
	exp := &Role{
		Name:        "Admin",
		Description: "Grants all permissions for Graylog administrators (built-in)",
		Permissions: []string{"*"},
		ReadOnly:    true,
	}
	if !reflect.DeepEqual(role, exp) {
		t.Errorf("client.GetRole() == %v, wanted %v", role, exp)
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
	role := Role{
		Name:        "foo",
		Description: "",
		Permissions: []string{"users:edit"},
		ReadOnly:    false,
	}
	updatedRole, err := client.UpdateRole("Admin", &role)
	if err != nil {
		t.Error("Failed to UpdateRole", err)
		return
	}
	if !reflect.DeepEqual(*updatedRole, role) {
		t.Errorf("client.UpdateRole() == %v, wanted %v", role, updatedRole)
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
	err = client.DeleteRole("Admin")
	if err != nil {
		t.Error("Failed to DeleteRole", err)
		return
	}
}
