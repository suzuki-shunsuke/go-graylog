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
	if _, err := client.GetRoleMembers(""); err == nil {
		t.Fatal("name is required")
	}
	if _, err := client.GetRoleMembers("h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
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
	if err = client.AddUserToRole(user.Username, role.Name); err != nil {
		t.Fatal("Failed to AddUserToRole", err)
	}
	if err = client.AddUserToRole("", role.Name); err == nil {
		t.Fatal("user name is required")
	}
	if err = client.AddUserToRole(user.Username, ""); err == nil {
		t.Fatal("role name is required")
	}
	if err = client.AddUserToRole("h", role.Name); err == nil {
		t.Fatal(`no user whose name is "h"`)
	}
	if err = client.AddUserToRole(user.Username, "h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
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
	if err = client.RemoveUserFromRole(user.Username, role.Name); err != nil {
		t.Fatal("Failed to RemoveUserFromRole", err)
	}
	if err = client.RemoveUserFromRole("", role.Name); err == nil {
		t.Fatal("user name is required")
	}
	if err = client.RemoveUserFromRole(user.Username, ""); err == nil {
		t.Fatal("role name is required")
	}
	if err = client.RemoveUserFromRole("h", role.Name); err == nil {
		t.Fatal(`no user whose name is "h"`)
	}
	if err = client.RemoveUserFromRole(user.Username, "h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}
