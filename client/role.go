package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog"
)

// CreateRole creates a new role.
func (client *Client) CreateRole(
	ctx context.Context, role *graylog.Role,
) (*ErrorInfo, error) {
	if role == nil {
		return nil, errors.New("role is nil")
	}
	return client.callPost(ctx, client.Endpoints().Roles(), role, role)
}

// GetRoles returns all roles.
func (client *Client) GetRoles(ctx context.Context) (
	[]graylog.Role, int, *ErrorInfo, error,
) {
	roles := &graylog.RolesBody{}
	ei, err := client.callGet(ctx, client.Endpoints().Roles(), nil, roles)
	return roles.Roles, roles.Total, ei, err
}

// GetRole returns a given role.
func (client *Client) GetRole(
	ctx context.Context, name string,
) (*graylog.Role, *ErrorInfo, error) {
	if name == "" {
		return nil, nil, errors.New("name is empty")
	}
	role := &graylog.Role{}
	ei, err := client.callGet(ctx, client.Endpoints().Role(name), nil, role)
	return role, ei, err
}

// UpdateRole updates a given role.
func (client *Client) UpdateRole(
	ctx context.Context, name string, prms *graylog.RoleUpdateParams,
) (*graylog.Role, *ErrorInfo, error) {
	if name == "" {
		return nil, nil, errors.New("name is empty")
	}
	if prms == nil {
		return nil, nil, errors.New("role is nil")
	}
	role := &graylog.Role{}
	ei, err := client.callPut(ctx, client.Endpoints().Role(name), prms, role)
	return role, ei, err
}

// DeleteRole deletes a given role.
func (client *Client) DeleteRole(
	ctx context.Context, name string,
) (*ErrorInfo, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	return client.callDelete(ctx, client.Endpoints().Role(name), nil, nil)
}
