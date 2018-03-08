package graylog

import (
	"reflect"
	"testing"
)

func dummyNewUser() *User {
	return &User{
		Username:    "admin",
		Email:       "hoge@example.com",
		FullName:    "Administrator",
		Password:    "password",
		Permissions: []string{"*"},
	}
}

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
		t.Fatal(err)
	}
	defer server.Close()
	admin := dummyNewUser()
	if _, err := client.CreateUser(admin); err != nil {
		t.Fatal("Failed to CreateUser", err)
	}
	if _, err := client.CreateUser(admin); err == nil {
		t.Fatal("User name must be unique.")
	}

	userName := admin.Username
	admin.Username = ""
	if _, err := client.CreateUser(admin); err == nil {
		t.Fatal("Username is required.")
	}
	admin.Username = userName
	roleName := "no roles"
	admin.Roles = []string{roleName}
	if _, err := client.CreateUser(admin); err == nil {
		t.Fatalf("No role found with name %s", roleName)
	}
}

func TestGetUsers(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	admin := dummyAdmin()
	server.Users[admin.Username] = *admin
	users, _, err := client.GetUsers()
	if err != nil {
		t.Fatal("Failed to GetUsers", err)
	}
	exp := []User{*admin}
	if !reflect.DeepEqual(users, exp) {
		t.Fatalf("client.GetUsers() == %v, wanted %v", users, exp)
	}
}

func TestGetUser(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	exp := dummyAdmin()
	server.Users[exp.Username] = *exp
	user, _, err := client.GetUser(exp.Username)
	if err != nil {
		t.Fatal("Failed to GetUser", err)
	}
	if !reflect.DeepEqual(*user, *exp) {
		t.Fatalf("client.GetUser() == %v, wanted %v", user, exp)
	}
	if _, _, err := client.GetUser(""); err == nil {
		t.Fatal("username should be required.")
	}
	if _, _, err := client.GetUser("h"); err == nil {
		t.Fatal(`no user whoname name is "h"`)
	}
}

func TestUpdateUser(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	user := dummyAdmin()
	server.Users[user.Username] = *user
	user.FullName = "changed!"
	if _, err := client.UpdateUser(user); err != nil {
		t.Fatal("Failed to UpdateUser", err)
	}
	user.Username = ""
	if _, err := client.UpdateUser(user); err == nil {
		t.Fatal("username should be required.")
	}
	user.Username = "h"
	if _, err := client.UpdateUser(user); err == nil {
		t.Fatal(`no user whoname name is "h"`)
	}
}

func TestDeleteUser(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	user := dummyAdmin()
	server.Users[user.Username] = *user
	if _, err := client.DeleteUser(user.Username); err != nil {
		t.Fatal("Failed to DeleteUser", err)
	}
	if _, err := client.DeleteUser(""); err == nil {
		t.Fatal("username should be required.")
	}
	if _, err := client.DeleteUser("h"); err == nil {
		t.Fatal(`no user whoname name is "h"`)
	}
}
