package client

import (
	"context"

	"github.com/pkg/errors"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetRoleMembers returns a given role's members.
func (client *Client) GetRoleMembers(
	ctx context.Context, name string,
) ([]graylog.User, *ErrorInfo, error) {
	if name == "" {
		return nil, nil, errors.New("name is empty")
	}
	u, err := client.Endpoints().RoleMembers(name)
	if err != nil {
		return nil, nil, err
	}
	users := &graylog.UsersBody{}
	ei, err := client.callGet(ctx, u.String(), nil, users)
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
	u, err := client.Endpoints().RoleMember(userName, roleName)
	if err != nil {
		return nil, err
	}
	return client.callPut(ctx, u.String(), nil, nil)
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
	u, err := client.Endpoints().RoleMember(userName, roleName)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
