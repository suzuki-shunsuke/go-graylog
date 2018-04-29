package client_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
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
	test.TestAddUserToRole(t)
}

func TestRemoveUserFromRole(t *testing.T) {
	test.TestRemoveUserFromRole(t)
}
