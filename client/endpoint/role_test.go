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
	require.Equal(t, fmt.Sprintf("%s/%s", apiURL, "roles"), ep.Roles())
}

func TestRole(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/%s", apiURL, "roles/foo"), ep.Role("foo"))
}

func TestRoleMembers(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/%s", apiURL, "roles/foo/members"), ep.RoleMembers("foo"))
}

func TestRoleMember(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(
		t, fmt.Sprintf("%s/%s", apiURL, "roles/Admin/members/foo"), ep.RoleMember("foo", "Admin"))
}
