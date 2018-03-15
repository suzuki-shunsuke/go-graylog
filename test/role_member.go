package test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetRoleMembers(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	users, _, err := client.GetRoleMembers("Admin")
	if err != nil {
		t.Fatal("Failed to GetRoleMembers", err)
	}
	if len(users) != 0 {
		t.Fatalf("the number of Admin users is %d, wanted 0", len(users))
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
	defer server.Close()
	if _, err = client.AddUserToRole("admin", "Admin"); err != nil {
		t.Fatal("Failed to AddUserToRole", err)
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
	defer server.Close()
	if _, err = client.RemoveUserFromRole("admin", "Admin"); err != nil {
		t.Fatal("Failed to RemoveUserFromRole", err)
	}
	if _, err = client.RemoveUserFromRole("", "Admin"); err == nil {
		t.Fatal("user name is required")
	}
	if _, err = client.RemoveUserFromRole("admin", ""); err == nil {
		t.Fatal("role name is required")
	}
	if _, err = client.RemoveUserFromRole("h", "Admin"); err == nil {
		t.Fatal(`no user whose name is "h"`)
	}
	if _, err = client.RemoveUserFromRole("admin", "h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}
