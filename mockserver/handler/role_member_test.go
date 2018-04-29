package handler_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestHandleRoleMembers(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	users, _, err := client.GetRoleMembers("Admin")
	if err != nil {
		t.Fatal("failed to GetRoleMembers", err)
	}
	if len(users) != 0 {
		t.Fatalf("the number of Admin users is %d, wanted 0", len(users))
	}
	if _, _, err := client.GetRoleMembers("h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}

func TestHandleAddUserToRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	user, err := testutil.GetNonAdminUser(client)
	if err != nil {
		t.Fatal(err)
	}
	if user == nil {
		user = testutil.User()
		user.Username = "foo"
		if _, err := client.CreateUser(user); err != nil {
			t.Fatal(err)
		}
	}
	if _, err = client.AddUserToRole(user.Username, "Admin"); err != nil {
		// Cannot modify local root user, this is a bug.
		t.Fatal(err)
	}
	if _, err = client.AddUserToRole("h", "Admin"); err == nil {
		t.Fatal(`no user whose name is "h"`)
	}
	if _, err = client.AddUserToRole("admin", "h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}

func TestHandleRemoveUserFromRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	user, err := testutil.GetNonAdminUser(client)
	if err != nil {
		t.Fatal(err)
	}
	if user == nil {
		user = testutil.User()
		user.Username = "foo"
		if _, err := client.CreateUser(user); err != nil {
			t.Fatal(err)
		}
	}
	if _, err = client.RemoveUserFromRole(user.Username, "Admin"); err != nil {
		// Cannot modify local root user, this is a bug.
		t.Fatal(err)
	}
	if _, err = client.RemoveUserFromRole("h", "Admin"); err == nil {
		t.Fatal(`no user whose name is "h"`)
	}
	if _, err = client.RemoveUserFromRole(user.Username, "h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}
