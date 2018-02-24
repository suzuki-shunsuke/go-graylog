package graylog

import (
	"reflect"
	"testing"
)

func TestGetRoleMembers(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	user := dummyAdmin()
	server.Users[user.Username] = *user
	role := dummyRole()
	server.Roles[role.Name] = *role
	users, err := client.GetRoleMembers(role.Name)
	if err != nil {
		t.Error("Failed to GetRoleMembers", err)
		return
	}
	exp := []User{*user}
	if !reflect.DeepEqual(users, exp) {
		t.Errorf("client.GetRoleMembers() == %v, wanted %v", users, exp)
	}
}

func TestAddUserToRole(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	user := dummyAdmin()
	server.Users[user.Username] = *user
	role := dummyRole()
	server.Roles[role.Name] = *role
	err = client.AddUserToRole(user.Username, role.Name)
	if err != nil {
		t.Error("Failed to AddUserToRole", err)
		return
	}
}

func TestRemoveUserFromRole(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	user := dummyAdmin()
	server.Users[user.Username] = *user
	role := dummyRole()
	server.Roles[role.Name] = *role
	err = client.RemoveUserFromRole(user.Username, role.Name)
	if err != nil {
		t.Error("Failed to RemoveUserFromRole", err)
		return
	}
}
