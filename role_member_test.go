package graylog

import (
	"reflect"
	"testing"
)

func TestGetRoleMembers(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	user := dummyAdmin()
	server.Users[user.Username] = *user
	role := dummyRole()
	server.Roles[role.Name] = *role
	users, err := client.GetRoleMembers(role.Name)
	if err != nil {
		t.Fatal("Failed to GetRoleMembers", err)
	}
	exp := []User{*user}
	if !reflect.DeepEqual(users, exp) {
		t.Fatalf("client.GetRoleMembers() == %v, wanted %v", users, exp)
	}
}

func TestAddUserToRole(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	user := dummyAdmin()
	server.Users[user.Username] = *user
	role := dummyRole()
	server.Roles[role.Name] = *role
	err = client.AddUserToRole(user.Username, role.Name)
	if err != nil {
		t.Fatal("Failed to AddUserToRole", err)
	}
}

func TestRemoveUserFromRole(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	user := dummyAdmin()
	server.Users[user.Username] = *user
	role := dummyRole()
	server.Roles[role.Name] = *role
	err = client.RemoveUserFromRole(user.Username, role.Name)
	if err != nil {
		t.Fatal("Failed to RemoveUserFromRole", err)
	}
}
