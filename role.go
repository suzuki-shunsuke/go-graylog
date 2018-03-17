package graylog

// Acccess Token http://docs.graylog.org/en/2.4/pages/configuration/rest_api.html#creating-and-using-access-token
// Session Token http://docs.graylog.org/en/2.4/pages/configuration/rest_api.html#creating-and-using-session-token
// -u ADMIN:PASSWORD
// -u {token}:token
// -u {session}:session

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Role represents a role.
type Role struct {
	Name        string `json:"name,omitempty" v-create:"required" v-update:"required"`
	Description string `json:"description,omitempty"`
	// ex. ["clusterconfigentry:read", "users:edit"]
	Permissions []string `json:"permissions,omitempty" v-create:"required" v-update:"required"`
	ReadOnly    bool     `json:"read_only,omitempty"`
}

func CopyRole(src, dest *Role) {
	dest.Name = src.Name
	dest.Description = src.Description
	dest.Permissions = src.Permissions
	dest.ReadOnly = src.ReadOnly
}

// CreateRole creates a new role.
func (client *Client) CreateRole(role *Role) (*ErrorInfo, error) {
	return client.CreateRoleContext(context.Background(), role)
}

// CreateRoleContext creates a new role with a context.
func (client *Client) CreateRoleContext(
	ctx context.Context, role *Role,
) (*ErrorInfo, error) {
	b, err := json.Marshal(role)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(role)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPost, client.Endpoints.Roles, b, true)
	if err != nil {
		return ei, err
	}

	ret := &Role{}
	if err := json.Unmarshal(ei.ResponseBody, ret); err != nil {
		return ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Role: %s",
				string(ei.ResponseBody)))
	}
	CopyRole(ret, role)
	return ei, nil
}

type RolesBody struct {
	Roles []Role `json:"roles"`
	Total int    `json:"total"`
}

// GetRoles returns all roles.
func (client *Client) GetRoles() ([]Role, *ErrorInfo, error) {
	return client.GetRolesContext(context.Background())
}

// GetRolesContext returns all roles with a context.
func (client *Client) GetRolesContext(ctx context.Context) (
	[]Role, *ErrorInfo, error,
) {
	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.Roles, nil, true)
	if err != nil {
		return nil, ei, err
	}

	roles := RolesBody{}
	if err := json.Unmarshal(ei.ResponseBody, &roles); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Roles: %s",
				string(ei.ResponseBody)))
	}
	return roles.Roles, ei, nil
}

// GetRole returns a given role.
func (client *Client) GetRole(name string) (*Role, *ErrorInfo, error) {
	return client.GetRoleContext(context.Background(), name)
}

// GetRoleContext returns a given role with a context.
func (client *Client) GetRoleContext(
	ctx context.Context, name string,
) (*Role, *ErrorInfo, error) {
	if name == "" {
		return nil, nil, errors.New("name is empty")
	}

	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.Role(name), nil, true)
	if err != nil {
		return nil, ei, err
	}

	role := &Role{}
	if err := json.Unmarshal(ei.ResponseBody, role); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Role: %s",
				string(ei.ResponseBody)))
	}
	return role, ei, nil
}

// UpdateRole updates a given role.
func (client *Client) UpdateRole(name string, role *Role) (
	*ErrorInfo, error,
) {
	return client.UpdateRoleContext(context.Background(), name, role)
}

// UpdateRoleContext updates a given role with a context.
func (client *Client) UpdateRoleContext(
	ctx context.Context, name string, role *Role,
) (*ErrorInfo, error) {
	if name == "" {
		return nil, errors.New("name is empty")
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

	ret := &Role{}
	if err := json.Unmarshal(ei.ResponseBody, ret); err != nil {
		return ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Role: %s",
				string(ei.ResponseBody)))
	}
	CopyRole(ret, role)
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
