package handler_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestHandleGetUsers(t *testing.T) {
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
	if len(users) != 3 {
		t.Fatalf("len(users) == %d, wanted 3", len(users))
	}
	if users[0].Password != "" {
		t.Fatalf(
			"users[0].Password == %s, wanted empty", users[0].Password)
	}
}

func TestHandleGetUser(t *testing.T) {
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
	if user.Password != "" {
		t.Fatalf("user.Password = %s, wanted empty", user.Password)
	}
	if _, _, err := client.GetUser(""); err == nil {
		t.Fatal("username should be required.")
	}
	if _, _, err := client.GetUser("h"); err == nil {
		t.Fatal(`no user whoname name is "h"`)
	}
}

func TestHandleCreateUser(t *testing.T) {
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
	user.Roles = set.NewStrSet(roleName)
	if _, err := client.CreateUser(user); err == nil {
		t.Fatalf("No role found with name %s", roleName)
	}

	if _, err := client.CreateUser(nil); err == nil {
		t.Fatal("user is nil")
	}

	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPost, client.Endpoints().Users(), body)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth(client.Name(), client.Password())
	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Fatalf("resp.StatusCode == %d, wanted 400", resp.StatusCode)
	}
}

func TestHandleUpdateUser(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	user := testutil.DummyAdmin()
	user.Roles = nil
	userName := "foo"
	user.Username = userName
	if _, err := server.AddUser(user); err != nil {
		t.Fatal(err)
	}
	user.FullName = "changed!"
	if _, err := client.UpdateUser(user.NewUpdateParams()); err != nil {
		t.Fatal(err)
	}
	user.Username = ""
	if _, err := client.UpdateUser(user.NewUpdateParams()); err == nil {
		t.Fatal("username should be required.")
	}
	user.Username = "h"
	if _, err := client.UpdateUser(user.NewUpdateParams()); err == nil {
		t.Fatal(`no user with name is "h"`)
	}
	if _, err := client.UpdateUser(nil); err == nil {
		t.Fatal("user is nil")
	}

	body := bytes.NewBuffer([]byte("hoge"))
	u, err := client.Endpoints().User(userName)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPut, u.String(), body)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth(client.Name(), client.Password())
	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Fatalf("resp.StatusCode == %d, wanted 400", resp.StatusCode)
	}
}

func TestHandleDeleteUser(t *testing.T) {
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
		t.Fatal(err)
	}
	if _, err := client.DeleteUser(""); err == nil {
		t.Fatal("username should be required.")
	}
	if _, err := client.DeleteUser("h"); err == nil {
		t.Fatal(`no user whoname name is "h"`)
	}
}
