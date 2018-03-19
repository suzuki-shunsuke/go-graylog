package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

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
	b, err := json.Marshal(role)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(role)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPost, client.Endpoints.Roles, b, true)
	if err != nil {
		return ei, err
	}

	if err := json.Unmarshal(ei.ResponseBody, role); err != nil {
		return ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Role: %s",
				string(ei.ResponseBody)))
	}
	return ei, nil
}

// GetRoles returns all roles.
func (client *Client) GetRoles() ([]graylog.Role, *ErrorInfo, error) {
	return client.GetRolesContext(context.Background())
}

// GetRolesContext returns all roles with a context.
func (client *Client) GetRolesContext(ctx context.Context) (
	[]graylog.Role, *ErrorInfo, error,
) {
	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.Roles, nil, true)
	if err != nil {
		return nil, ei, err
	}

	roles := &graylog.RolesBody{}
	if err := json.Unmarshal(ei.ResponseBody, roles); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Roles: %s",
				string(ei.ResponseBody)))
	}
	return roles.Roles, ei, nil
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

	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.Role(name), nil, true)
	if err != nil {
		return nil, ei, err
	}

	role := &graylog.Role{}
	if err := json.Unmarshal(ei.ResponseBody, role); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Role: %s",
				string(ei.ResponseBody)))
	}
	return role, ei, nil
}

// UpdateRole updates a given role.
func (client *Client) UpdateRole(name string, role *graylog.Role) (
	*ErrorInfo, error,
) {
	return client.UpdateRoleContext(context.Background(), name, role)
}

// UpdateRoleContext updates a given role with a context.
func (client *Client) UpdateRoleContext(
	ctx context.Context, name string, role *graylog.Role,
) (*ErrorInfo, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	if role == nil {
		return nil, fmt.Errorf("role is nil")
	}
	b, err := json.Marshal(role)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(role)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPut, client.Endpoints.Role(name), b, true)
	if err != nil {
		return ei, err
	}

	if err := json.Unmarshal(ei.ResponseBody, role); err != nil {
		return ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Role: %s",
				string(ei.ResponseBody)))
	}
	return ei, nil
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

	return client.callReq(
		ctx, http.MethodDelete, client.Endpoints.Role(name), nil, false)
}
