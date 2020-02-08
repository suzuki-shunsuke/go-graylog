package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/go-graylog/v10/client/endpoint"
)

func TestEndpoints_Roles(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/roles", apiURL), ep.Roles())
}

func TestEndpoints_Role(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/roles/foo", apiURL), ep.Role("foo"))
}

func TestEndpoints_RoleMembers(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/roles/foo/members", apiURL), ep.RoleMembers("foo"))
}

func TestEndpoints_RoleMember(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(
		t, fmt.Sprintf("%s/roles/Admin/members/foo", apiURL), ep.RoleMember("foo", "Admin"))
}
