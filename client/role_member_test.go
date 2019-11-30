package client_test

import (
	"context"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil/v8"
)

func TestClient_GetRoleMembers(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, _, err := client.GetRoleMembers(ctx, "Admin"); err != nil {
		t.Fatal("Failed to GetRoleMembers", err)
	}
	if _, _, err := client.GetRoleMembers(ctx, ""); err == nil {
		t.Fatal("name is required")
	}
	if _, _, err := client.GetRoleMembers(ctx, "h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}

func TestClient_AddUserToRole(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	user, err := testutil.GetNonAdminUser(ctx, client)
	if err != nil {
		t.Fatal(err)
	}
	if user == nil {
		user = testutil.User()
		user.Username = "foo"
		if _, err := client.CreateUser(ctx, user); err != nil {
			t.Fatal(err)
		}
	}
	if _, err = client.AddUserToRole(ctx, user.Username, "Admin"); err != nil {
		// Cannot modify local root user, this is a bug.
		t.Fatal(err)
	}
	if _, err = client.AddUserToRole(ctx, "", "Admin"); err == nil {
		t.Fatal("user name is required")
	}
	if _, err = client.AddUserToRole(ctx, "admin", ""); err == nil {
		t.Fatal("role name is required")
	}
	if _, err = client.AddUserToRole(ctx, "h", "Admin"); err == nil {
		t.Fatal(`no user whose name is "h"`)
	}
	if _, err = client.AddUserToRole(ctx, "admin", "h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}

func TestClient_RemoveUserFromRole(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	user, err := testutil.GetNonAdminUser(ctx, client)
	if err != nil {
		t.Fatal(err)
	}
	if user == nil {
		user = testutil.User()
		user.Username = "foo"
		if _, err := client.CreateUser(ctx, user); err != nil {
			t.Fatal(err)
		}
	}
	if _, err = client.RemoveUserFromRole(ctx, user.Username, "Admin"); err != nil {
		// Cannot modify local root user, this is a bug.
		t.Fatal(err)
	}
	if _, err = client.RemoveUserFromRole(ctx, "", "Admin"); err == nil {
		t.Fatal("user name is required")
	}
	if _, err = client.RemoveUserFromRole(ctx, user.Username, ""); err == nil {
		t.Fatal("role name is required")
	}
	if _, err = client.RemoveUserFromRole(ctx, "h", "Admin"); err == nil {
		t.Fatal(`no user whose name is "h"`)
	}
	if _, err = client.RemoveUserFromRole(ctx, user.Username, "h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}
