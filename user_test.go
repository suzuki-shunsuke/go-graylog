package graylog

import (
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
		ID:          "local:admin",
		Username:    "admin",
		Email:       "hoge@example.com",
		FullName:    "Administrator",
		Password:    "password",
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
	admin.Roles = nil
	if _, _, err := server.AddUser(admin); err != nil {
		t.Fatal(err)
	}
	users, _, err := client.GetUsers()
	if err != nil {
		t.Fatal("Failed to GetUsers", err)
	}
	if users == nil {
		t.Fatal("client.GetUsers() returns nil")
	}
	if len(users) != 1 {
		t.Fatalf("len(users) == %d, wanted 1", len(users))
	}
	if users[0].Password != admin.Password {
		t.Fatalf(
			"users[0].Password == %v, wanted %v", users[0].Password, admin.Password)
	}
}

func TestGetUser(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	exp := dummyAdmin()
	exp.Roles = nil
	if _, _, err := server.AddUser(exp); err != nil {
		t.Fatal(err)
	}
	user, _, err := client.GetUser(exp.Username)
	if err != nil {
		t.Fatal("Failed to GetUser", err)
	}
	if user.Password != exp.Password {
		t.Fatalf("user.Password %v, wanted %v", user.Password, exp.Password)
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
	user.Roles = nil
	if _, _, err := server.AddUser(user); err != nil {
		t.Fatal(err)
	}
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
	user.Roles = nil
	if _, _, err := server.AddUser(user); err != nil {
		t.Fatal(err)
	}
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
