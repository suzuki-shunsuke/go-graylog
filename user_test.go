package graylog

import (
	"reflect"
	"testing"
)

func dummyAdmin() *User {
	return &User{
		Id:          "local:admin",
		Username:    "admin",
		Email:       "",
		FullName:    "Administrator",
		Permissions: []string{"*"},
		Preferences: &Preferences{
			UpdateUnfocussed:  false,
			EnableSmartSearch: true,
		},
		Timezone:         "UTC",
		SessionTimeoutMs: 28800000,
		External:         false,
		Startpage:        nil,
		Roles:            []string{"Admin"},
		ReadOnly:         true,
		SessionActive:    true,
		LastActivity:     "2018-02-21T07:35:45.926+0000",
		ClientAddress:    "172.18.0.1",
	}
}

func TestCreateUser(t *testing.T) {
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
	params := &User{Username: "foo"}
	user, err := client.CreateUser(params)
	if err != nil {
		t.Error("Failed to CreateUser", err)
		return
	}
	if user == nil {
		t.Error("client.CreateUser() == nil")
		return
	}
	if user.Username != "foo" {
		t.Errorf("user.Username == %s, wanted foo", user.Username)
	}
}

func TestGetUsers(t *testing.T) {
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
	users, err := client.GetUsers()
	if err != nil {
		t.Error("Failed to GetUsers", err)
		return
	}
	admin := dummyAdmin()
	exp := []User{*admin}
	if !reflect.DeepEqual(users, exp) {
		t.Errorf("client.GetUsers() == %v, wanted %v", users, exp)
	}
}

func TestGetUser(t *testing.T) {
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
	user, err := client.GetUser("Admin")
	if err != nil {
		t.Error("Failed to GetUser", err)
		return
	}
	exp := dummyAdmin()
	if !reflect.DeepEqual(*user, *exp) {
		t.Errorf("client.GetUser() == %v, wanted %v", user, exp)
	}
}

func TestUpdateUser(t *testing.T) {
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
	user := dummyAdmin()
	updatedUser, err := client.UpdateUser("Admin", user)
	if err != nil {
		t.Error("Failed to UpdateUser", err)
		return
	}
	if !reflect.DeepEqual(*updatedUser, *user) {
		t.Errorf("client.UpdateUser() == %v, wanted %v", user, updatedUser)
	}
}

func TestDeleteUser(t *testing.T) {
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
	err = client.DeleteUser("Admin")
	if err != nil {
		t.Error("Failed to DeleteUser", err)
		return
	}
}
