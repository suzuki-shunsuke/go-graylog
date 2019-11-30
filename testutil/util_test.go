package testutil_test

import (
	"context"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil/v8"
)

func TestGetNonAdminUser(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	if _, err := testutil.GetNonAdminUser(ctx, client); err != nil {
		t.Fatal(err)
	}
}

func TestGetRoleOrCreate(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	role, err := testutil.GetRoleOrCreate(ctx, client, "Admin")
	if err != nil {
		t.Fatal(err)
	}
	if role == nil {
		t.Fatal("role is nil")
	}
}
