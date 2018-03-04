package graylog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// GetRoleMembers returns a given role's members.
func (client *Client) GetRoleMembers(name string) ([]User, *ErrorInfo, error) {
	return client.GetRoleMembersContext(context.Background(), name)
}

// GetRoleMembersContext returns a given role's members with a context.
func (client *Client) GetRoleMembersContext(
	ctx context.Context, name string,
) ([]User, *ErrorInfo, error) {
	if name == "" {
		return nil, nil, errors.New("name is empty")
	}

	ei, err := client.callReq(
		ctx, http.MethodGet, client.endpoints.RoleMembers(name), nil, true)
	if err != nil {
		return nil, ei, err
	}
	users := usersBody{}
	err = json.Unmarshal(ei.ResponseBody, &users)
	if err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as Users: %s",
				string(ei.ResponseBody)))
	}
	return users.Users, ei, nil
}

// AddUserToRole adds a user to a role.
func (client *Client) AddUserToRole(userName, roleName string) (
	*ErrorInfo, error,
) {
	return client.AddUserToRoleContext(context.Background(), userName, roleName)
}

// AddUserToRoleContext adds a user to a role with a context.
func (client *Client) AddUserToRoleContext(
	ctx context.Context, userName, roleName string,
) (*ErrorInfo, error) {
	if userName == "" {
		return nil, errors.New("userName is empty")
	}
	if roleName == "" {
		return nil, errors.New("roleName is empty")
	}

	return client.callReq(
		ctx, http.MethodPut,
		client.endpoints.RoleMember(userName, roleName), nil, false)
}

// RemoveUserFromRole removes a user from a role.
func (client *Client) RemoveUserFromRole(
	userName, roleName string,
) (*ErrorInfo, error) {
	return client.RemoveUserFromRoleContext(
		context.Background(), userName, roleName)
}

// RemoveUserFromRoleContext removes a user from a role with a context.
func (client *Client) RemoveUserFromRoleContext(
	ctx context.Context, userName, roleName string,
) (*ErrorInfo, error) {
	if userName == "" {
		return nil, errors.New("userName is empty")
	}
	if roleName == "" {
		return nil, errors.New("roleName is empty")
	}

	return client.callReq(
		ctx, http.MethodDelete,
		client.endpoints.RoleMember(userName, roleName), nil, false)
}
