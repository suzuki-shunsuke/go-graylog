package handler_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
	"github.com/suzuki-shunsuke/go-set"
)

func TestServerHandleCreateRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPost, client.Endpoints.Roles, body)
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

func TestServerHandleUpdateRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	admin := testutil.Role()
	server.AddRole(admin)
	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPut, client.Endpoints.Role(admin.Name), body)
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

// Same as client test

func TestCreateRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	role := &graylog.Role{Name: "foo", Permissions: set.NewStrSet("*")}
	if _, err := client.CreateRole(role); err != nil {
		t.Fatal("Failed to CreateRole", err)
	}
	if role.Name != "foo" {
		t.Fatalf("role.Name == %s, wanted foo", role.Name)
	}
	ei, err := client.CreateRole(role)
	if err == nil {
		t.Fatal("user name must be unique")
	}
	if ei.Response.StatusCode != 400 {
		t.Fatal("status code must be 400")
	}

	role.Name = ""
	if _, err := client.CreateRole(role); err == nil {
		t.Fatal("user name is required")
	}

	role.Name = "bar"
	role.Permissions = nil
	if _, err := client.CreateRole(role); err == nil {
		t.Fatal("user permissions are required")
	}
}

func TestGetRoles(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	roles, _, err := client.GetRoles()
	if err != nil {
		t.Fatal("Failed to GetRoles", err)
	}
	if roles == nil {
		t.Fatal("client.GetRoles() is nil")
	}
	if len(roles) != 1 {
		t.Fatalf("len(roles) == %d, wanted 1", len(roles))
	}
}

func TestGetRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	role, _, err := client.GetRole("Admin")
	if err != nil {
		t.Fatal("Failed to GetRole", err)
	}
	if role.Name != "Admin" {
		t.Fatalf(`role name is "%s", wanted "Admin"`, role.Name)
	}
	if _, _, err := client.GetRole(""); err == nil {
		t.Fatal("role name is required")
	}
	if _, _, err := client.GetRole("h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}

func TestUpdateRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	role, err := testutil.GetRoleOrCreate(client, "foo")
	if err != nil {
		t.Fatal(err)
	}
	role.Description += " changed!"
	if _, err := client.UpdateRole(role.Name, role); err != nil {
		t.Fatal(err)
	}
	if _, err := client.UpdateRole("", role); err == nil {
		t.Fatal("role name is required")
	}
	if _, err := client.UpdateRole("h", role); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}

	cRole := *role
	role.Name = ""
	if _, err := client.UpdateRole(cRole.Name, role); err == nil {
		t.Fatal("role name is required")
	}
	role.Name = cRole.Name
	role.Permissions = nil
	if _, err := client.UpdateRole(role.Name, role); err == nil {
		t.Fatal("role permissions is required")
	}
}

func TestDeleteRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	role, err := testutil.GetRoleOrCreate(client, "foo")
	if err != nil {
		t.Fatal(err)
	}
	if _, err = client.DeleteRole(role.Name); err != nil {
		t.Fatal("Failed to DeleteRole", err)
	}
	if _, err = client.DeleteRole(""); err == nil {
		t.Fatal("role name is required")
	}
	if _, err = client.DeleteRole("h"); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}
}

func TestHasRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	nobody, _, err := client.GetUser("nobody")
	if err != nil {
		t.Fatal(err)
	}
	sc, err := server.Authorize(nobody, "roles:read", "Admin")
	if err == nil {
		t.Fatal("nobody should have a permission to read role")
	}
	if sc != 403 {
		t.Fatalf("sc = %d, wanted 403", sc)
	}
}
