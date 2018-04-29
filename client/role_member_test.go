package client_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetRoleMembers(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, _, err := client.GetRoleMembers("Admin"); err != nil {
		t.Fatal("Failed to GetRoleMembers", err)
	}
	if _, _, err := client.GetRoleMembers(""); err == nil {
		t.Fatal("name is required")
	}
	if _, _, err := client.GetRoleMembers("h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}

func TestAddUserToRole(t *testing.T) {
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
	if _, err = client.AddUserToRole("", "Admin"); err == nil {
		t.Fatal("user name is required")
	}
	if _, err = client.AddUserToRole("admin", ""); err == nil {
		t.Fatal("role name is required")
	}
	if _, err = client.AddUserToRole("h", "Admin"); err == nil {
		t.Fatal(`no user whose name is "h"`)
	}
	if _, err = client.AddUserToRole("admin", "h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}

func TestRemoveUserFromRole(t *testing.T) {
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
	if _, err = client.RemoveUserFromRole("", "Admin"); err == nil {
		t.Fatal("user name is required")
	}
	if _, err = client.RemoveUserFromRole(user.Username, ""); err == nil {
		t.Fatal("role name is required")
	}
	if _, err = client.RemoveUserFromRole("h", "Admin"); err == nil {
		t.Fatal(`no user whose name is "h"`)
	}
	if _, err = client.RemoveUserFromRole(user.Username, "h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}
