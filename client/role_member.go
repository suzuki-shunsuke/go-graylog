package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog/v9"
)

// GetRoleMembers returns a given role's members.
func (client *Client) GetRoleMembers(
	ctx context.Context, name string,
) ([]graylog.User, *ErrorInfo, error) {
	if name == "" {
		return nil, nil, errors.New("name is empty")
	}
	users := &graylog.UsersBody{}
	ei, err := client.callGet(ctx, client.Endpoints().RoleMembers(name), nil, users)
	return users.Users, ei, err
}

// AddUserToRole adds a user to a role.
func (client *Client) AddUserToRole(
	ctx context.Context, userName, roleName string,
) (*ErrorInfo, error) {
	if userName == "" {
		return nil, errors.New("userName is empty")
	}
	if roleName == "" {
		return nil, errors.New("roleName is empty")
	}
	return client.callPut(ctx, client.Endpoints().RoleMember(userName, roleName), nil, nil)
}

// RemoveUserFromRole removes a user from a role.
func (client *Client) RemoveUserFromRole(
	ctx context.Context, userName, roleName string,
) (*ErrorInfo, error) {
	if userName == "" {
		return nil, errors.New("userName is empty")
	}
	if roleName == "" {
		return nil, errors.New("roleName is empty")
	}
	return client.callDelete(ctx, client.Endpoints().RoleMember(userName, roleName), nil, nil)
}
