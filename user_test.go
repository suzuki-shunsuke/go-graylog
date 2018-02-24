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
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	admin := dummyAdmin()
	if err := client.CreateUser(admin); err != nil {
		t.Error("Failed to CreateUser", err)
		return
	}
}

func TestGetUsers(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	admin := dummyAdmin()
	server.Users[admin.Username] = *admin
	users, err := client.GetUsers()
	if err != nil {
		t.Error("Failed to GetUsers", err)
		return
	}
	exp := []User{*admin}
	if !reflect.DeepEqual(users, exp) {
		t.Errorf("client.GetUsers() == %v, wanted %v", users, exp)
	}
}

func TestGetUser(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	exp := dummyAdmin()
	server.Users[exp.Username] = *exp
	user, err := client.GetUser(exp.Username)
	if err != nil {
		t.Error("Failed to GetUser", err)
		return
	}
	if !reflect.DeepEqual(*user, *exp) {
		t.Errorf("client.GetUser() == %v, wanted %v", user, exp)
	}
}

func TestUpdateUser(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	user := dummyAdmin()
	server.Users[user.Username] = *user
	user.FullName = "changed!"
	updatedUser, err := client.UpdateUser(user.Username, user)
	if err != nil {
		t.Error("Failed to UpdateUser", err)
		return
	}
	if !reflect.DeepEqual(*updatedUser, *user) {
		t.Errorf("client.UpdateUser() == %v, wanted %v", user, updatedUser)
	}
}

func TestDeleteUser(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	user := dummyAdmin()
	server.Users[user.Username] = *user
	err = client.DeleteUser(user.Username)
	if err != nil {
		t.Error("Failed to DeleteUser", err)
		return
	}
}
