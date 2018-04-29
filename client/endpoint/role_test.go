package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/client/endpoint"
)

func TestRoles(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s", apiURL, "roles")
	if ep.Roles() != exp {
		t.Fatalf(`ep.Roles() = "%s", wanted "%s"`, ep.Roles(), exp)
	}
}

func TestRole(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s", apiURL, "roles/foo")
	act, err := ep.Role("foo")
	if err != nil {
		t.Fatal(err)
	}
	if act.String() != exp {
		t.Fatalf(`ep.Role("foo") = "%s", wanted "%s"`, act.String(), exp)
	}
}

func TestRoleMembers(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s", apiURL, "roles/foo/members")
	act, err := ep.RoleMembers("foo")
	if err != nil {
		t.Fatal(err)
	}
	if act.String() != exp {
		t.Fatalf(`ep.RoleMembers("foo") = "%s", wanted "%s"`, act.String(), exp)
	}
}

func TestRoleMember(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s", apiURL, "roles/Admin/members/foo")
	act, err := ep.RoleMember("foo", "Admin")
	if err != nil {
		t.Fatal(err)
	}
	if act.String() != exp {
		t.Fatalf(`ep.RoleMember("foo", "Admin") = "%s", wanted "%s"`, act.String(), exp)
	}
}
