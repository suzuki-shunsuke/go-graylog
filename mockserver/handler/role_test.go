package handler_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
	"github.com/suzuki-shunsuke/go-set"
)

func TestHandleGetRole(t *testing.T) {
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

func TestHandleGetRoles(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	roles, _, _, err := client.GetRoles()
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

func TestHandleCreateRole(t *testing.T) {
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

	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPost, client.Endpoints().Roles(), body)
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

func TestHandleUpdateRole(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	name := "foo"
	role, err := testutil.GetRoleOrCreate(client, name)
	if err != nil {
		t.Fatal(err)
	}
	role.Description += " changed!"
	if _, _, err := client.UpdateRole(name, role.NewUpdateParams()); err != nil {
		t.Fatal(err)
	}
	if _, _, err := client.UpdateRole("", role.NewUpdateParams()); err == nil {
		t.Fatal("role name is required")
	}
	if _, _, err := client.UpdateRole("h", role.NewUpdateParams()); err == nil {
		t.Fatal(`no role whose name is "h"`)
	}

	role.Name = ""
	if _, _, err := client.UpdateRole(name, role.NewUpdateParams()); err == nil {
		t.Fatal("role name is required")
	}
	role.Name = name
	role.Permissions = nil
	if _, _, err := client.UpdateRole(name, role.NewUpdateParams()); err == nil {
		t.Fatal("role permissions is required")
	}

	body := bytes.NewBuffer([]byte("hoge"))
	u, err := client.Endpoints().Role(name)
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

func TestHandleDeleteRole(t *testing.T) {
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
