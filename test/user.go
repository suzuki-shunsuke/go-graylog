package test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestCreateUser(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	user := testutil.User()
	user.Username += "foo"
	if _, err := client.CreateUser(user); err != nil {
		t.Fatal("Failed to CreateUser", err)
	}
	if _, err := client.CreateUser(user); err == nil {
		t.Fatal("User name must be unique.")
	}

	userName := user.Username
	user.Username = ""
	if _, err := client.CreateUser(user); err == nil {
		t.Fatal("Username is required.")
	}
	user.Username = userName
	roleName := "no roles"
	user.Roles = []string{roleName}
	if _, err := client.CreateUser(user); err == nil {
		t.Fatalf("No role found with name %s", roleName)
	}

	if _, err := client.CreateUser(nil); err == nil {
		t.Fatal("user is nil")
	}
}

func TestGetUsers(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	user := testutil.DummyAdmin()
	user.Roles = nil
	user.Username = "foo"
	if _, err := server.AddUser(user); err != nil {
		t.Fatal(err)
	}
	users, _, err := client.GetUsers()
	if err != nil {
		t.Fatal("Failed to GetUsers", err)
	}
	if users == nil {
		t.Fatal("client.GetUsers() returns nil")
	}
	if len(users) != 2 {
		t.Fatalf("len(users) == %d, wanted 2", len(users))
	}
	if users[0].Password != user.Password {
		t.Fatalf(
			"users[0].Password == %v, wanted %v", users[0].Password, user.Password)
	}
}

func TestGetUser(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	exp := testutil.DummyAdmin()
	exp.Roles = nil
	exp.Username = "foo"
	if _, err := server.AddUser(exp); err != nil {
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
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	user := testutil.DummyAdmin()
	user.Roles = nil
	user.Username = "foo"
	if _, err := server.AddUser(user); err != nil {
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
		t.Fatal(`no user with name is "h"`)
	}
	if _, err := client.UpdateUser(nil); err == nil {
		t.Fatal("user is nil")
	}
}

func TestDeleteUser(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	user := testutil.DummyAdmin()
	user.Username = "foo"
	user.Roles = nil
	if _, err := server.AddUser(user); err != nil {
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
