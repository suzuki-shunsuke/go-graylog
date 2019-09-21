package client_test

import (
	"context"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestClient_CreateRole(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	role := testutil.Role()
	client.DeleteRole(ctx, role.Name)
	// nil check
	if _, err := client.CreateRole(ctx, nil); err == nil {
		t.Fatal("role is nil")
	}
	if _, err := client.CreateRole(ctx, role); err != nil {
		t.Fatal(err)
	}
	if _, err := client.DeleteRole(ctx, role.Name); err != nil {
		t.Fatal(err)
	}
	// error check
	role.Name = ""
	if _, err := client.CreateRole(ctx, role); err == nil {
		t.Fatal("role name is empty")
	}
}

func TestClient_GetRoles(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	roles, _, _, err := client.GetRoles(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(roles) == 0 {
		t.Fatal("roles is empty")
	}
}

func TestClient_GetRole(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	role := testutil.Role()
	client.DeleteRole(ctx, role.Name)
	if _, _, err := client.GetRole(ctx, role.Name); err == nil {
		t.Fatal("role should be deleted")
	}
	if _, err := client.CreateRole(ctx, role); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteRole(ctx, role.Name)
	r, _, err := client.GetRole(ctx, role.Name)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("roles is nil")
	}
	if r.Name != role.Name {
		t.Fatalf(`role.Name = "%s", wanted "%s"`, r.Name, role.Name)
	}
	if _, _, err := client.GetRole(ctx, ""); err == nil {
		t.Fatal("role name is required")
	}
}

func TestClient_UpdateRole(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	role := testutil.Role()
	client.DeleteRole(ctx, role.Name)
	if _, _, err := client.UpdateRole(ctx, role.Name, role.NewUpdateParams()); err == nil {
		t.Fatal("role should be deleted")
	}
	if _, err := client.CreateRole(ctx, role); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteRole(ctx, role.Name)
	if _, _, err := client.UpdateRole(ctx, role.Name, role.NewUpdateParams()); err != nil {
		t.Fatal(err)
	}
	if _, _, err := client.UpdateRole(ctx, "", role.NewUpdateParams()); err == nil {
		t.Fatal("role name is required")
	}
	name := role.Name
	role.Name = ""
	if _, _, err := client.UpdateRole(ctx, name, role.NewUpdateParams()); err == nil {
		t.Fatal("role name is required")
	}
}

func TestClient_DeleteRole(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, err := client.DeleteRole(ctx, ""); err == nil {
		t.Fatal("role name is required")
	}
	if _, err := client.DeleteRole(ctx, "h"); err == nil {
		t.Fatal(`no role with name "h" is found`)
	}
}
