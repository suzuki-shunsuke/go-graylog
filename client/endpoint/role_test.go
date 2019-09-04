package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/go-graylog/client/endpoint"
)

func TestRoles(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/roles", apiURL), ep.Roles())
}

func TestRole(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/roles/foo", apiURL), ep.Role("foo"))
}

func TestRoleMembers(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/roles/foo/members", apiURL), ep.RoleMembers("foo"))
}

func TestRoleMember(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(
		t, fmt.Sprintf("%s/roles/Admin/members/foo", apiURL), ep.RoleMember("foo", "Admin"))
}
