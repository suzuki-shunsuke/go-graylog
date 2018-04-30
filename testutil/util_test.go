package testutil_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetNonAdminUser(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	if _, err := testutil.GetNonAdminUser(client); err != nil {
		t.Fatal(err)
	}
}

func TestGetRoleOrCreate(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	role, err := testutil.GetRoleOrCreate(client, "Admin")
	if err != nil {
		t.Fatal(err)
	}
	if role == nil {
		t.Fatal("role is nil")
	}
}
