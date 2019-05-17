package client

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/suzuki-shunsuke/go-graylog"
)

// CreateRole creates a new role.
func (client *Client) CreateRole(role *graylog.Role) (*ErrorInfo, error) {
	return client.CreateRoleContext(context.Background(), role)
}

// CreateRoleContext creates a new role with a context.
func (client *Client) CreateRoleContext(
	ctx context.Context, role *graylog.Role,
) (*ErrorInfo, error) {
	if role == nil {
		return nil, fmt.Errorf("role is nil")
	}
	return client.callPost(ctx, client.Endpoints().Roles(), role, role)
}

// GetRoles returns all roles.
func (client *Client) GetRoles() ([]graylog.Role, int, *ErrorInfo, error) {
	return client.GetRolesContext(context.Background())
}

// GetRolesContext returns all roles with a context.
func (client *Client) GetRolesContext(ctx context.Context) (
	[]graylog.Role, int, *ErrorInfo, error,
) {
	roles := &graylog.RolesBody{}
	ei, err := client.callGet(ctx, client.Endpoints().Roles(), nil, roles)
	return roles.Roles, roles.Total, ei, err
}

// GetRole returns a given role.
func (client *Client) GetRole(name string) (*graylog.Role, *ErrorInfo, error) {
	return client.GetRoleContext(context.Background(), name)
}

// GetRoleContext returns a given role with a context.
func (client *Client) GetRoleContext(
	ctx context.Context, name string,
) (*graylog.Role, *ErrorInfo, error) {
	if name == "" {
		return nil, nil, errors.New("name is empty")
	}
	u, err := client.Endpoints().Role(name)
	if err != nil {
		return nil, nil, err
	}
	role := &graylog.Role{}
	ei, err := client.callGet(ctx, u.String(), nil, role)
	return role, ei, err
}

// UpdateRole updates a given role.
func (client *Client) UpdateRole(name string, role *graylog.RoleUpdateParams) (
	*graylog.Role, *ErrorInfo, error,
) {
	return client.UpdateRoleContext(context.Background(), name, role)
}

// UpdateRoleContext updates a given role with a context.
func (client *Client) UpdateRoleContext(
	ctx context.Context, name string, prms *graylog.RoleUpdateParams,
) (*graylog.Role, *ErrorInfo, error) {
	if name == "" {
		return nil, nil, errors.New("name is empty")
	}
	if prms == nil {
		return nil, nil, fmt.Errorf("role is nil")
	}
	u, err := client.Endpoints().Role(name)
	if err != nil {
		return nil, nil, err
	}
	role := &graylog.Role{}
	ei, err := client.callPut(ctx, u.String(), prms, role)
	return role, ei, err
}

// DeleteRole deletes a given role.
func (client *Client) DeleteRole(name string) (*ErrorInfo, error) {
	return client.DeleteRoleContext(context.Background(), name)
}

// DeleteRoleContext deletes a given role with a context.
func (client *Client) DeleteRoleContext(
	ctx context.Context, name string,
) (*ErrorInfo, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	u, err := client.Endpoints().Role(name)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
