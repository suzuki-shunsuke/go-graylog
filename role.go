package graylog

// Acccess Token http://docs.graylog.org/en/2.4/pages/configuration/rest_api.html#creating-and-using-access-token
// Session Token http://docs.graylog.org/en/2.4/pages/configuration/rest_api.html#creating-and-using-session-token
// -u ADMIN:PASSWORD
// -u {token}:token
// -u {session}:session

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type Role struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	ReadOnly    bool     `json:"read_only,omitempty"`
}

func callRequest(
	req *http.Request, client *Client, ctx context.Context,
) (*http.Response, error) {
	req.SetBasicAuth(client.GetName(), client.GetPassword())
	req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	hc := &http.Client{}
	return hc.Do(req)
}

// CreateRole create a new role
// POST /roles
func (client *Client) CreateRole(params *Role) (*Role, error) {
	return client.CreateRoleContext(context.Background(), params)
}

// CreateRoleContext create a new role
// POST /roles
func (client *Client) CreateRoleContext(
	ctx context.Context, params *Role,
) (*Role, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(params)")
	}
	req, err := http.NewRequest(
		http.MethodPost, client.endpoints.Roles, bytes.NewBuffer(b))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to call POST /roles API")
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read response body")
	}
	if resp.StatusCode >= 400 {
		e := Error{}
		err = json.Unmarshal(b, &e)
		if err != nil {
			return nil, errors.Wrap(
				err, fmt.Sprintf(
					"Failed to parse response body as Error: %s", string(b)))
		}
		return nil, errors.New(e.Message)
	}
	role := Role{}
	err = json.Unmarshal(b, &role)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Role: %s", string(b)))
	}
	return &role, nil
}

type rolesBody struct {
	Roles []Role `json:"roles"`
	Total int    `json:"total"`
}

// GetRoles List all roles
// GET /roles
func (client *Client) GetRoles() ([]Role, error) {
	return client.GetRolesContext(context.Background())
}

// GetRolesContext List all roles
// GET /roles
func (client *Client) GetRolesContext(
	ctx context.Context,
) ([]Role, error) {
	req, err := http.NewRequest(http.MethodGet, client.endpoints.Roles, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to call GET /roles API")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read response body")
	}
	if resp.StatusCode >= 400 {
		e := Error{}
		err = json.Unmarshal(b, &e)
		if err != nil {
			return nil, errors.Wrap(
				err, fmt.Sprintf(
					"Failed to parse response body as Error: %s", string(b)))
		}
		return nil, errors.New(e.Message)
	}
	roles := rolesBody{}
	err = json.Unmarshal(b, &roles)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Roles: %s", string(b)))
	}
	return roles.Roles, nil
}

// GetRole
// GET /roles/{rolename} Retrieve permissions for a single role
func (client *Client) GetRole(name string) (*Role, error) {
	return client.GetRoleContext(context.Background(), name)
}

// GetRoleContext
// GET /roles/{rolename} Retrieve permissions for a single role
func (client *Client) GetRoleContext(
	ctx context.Context, name string,
) (*Role, error) {
	req, err := http.NewRequest(
		http.MethodGet, fmt.Sprintf("%s/%s", client.endpoints.Roles, name), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to call GET /roles API")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read response body")
	}
	if resp.StatusCode >= 400 {
		e := Error{}
		err = json.Unmarshal(b, &e)
		if err != nil {
			return nil, errors.Wrap(
				err, fmt.Sprintf(
					"Failed to parse response body as Error: %s", string(b)))
		}
		return nil, errors.New(e.Message)
	}
	role := Role{}
	err = json.Unmarshal(b, &role)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Role: %s", string(b)))
	}
	return &role, nil
}

// UpdateRole
// PUT /roles/{rolename} Update an existing role
func (client *Client) UpdateRole(name string, params *Role) (*Role, error) {
	return client.UpdateRoleContext(context.Background(), name, params)
}

// UpdateRoleContext
// PUT /roles/{rolename} Update an existing role
func (client *Client) UpdateRoleContext(
	ctx context.Context, name string, params *Role,
) (*Role, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(params)")
	}
	req, err := http.NewRequest(
		http.MethodPut, fmt.Sprintf("%s/%s", client.endpoints.Roles, name),
		bytes.NewBuffer(b))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to call PUT /roles/{rolename} API")
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read response body")
	}
	if resp.StatusCode >= 400 {
		e := Error{}
		err = json.Unmarshal(b, &e)
		if err != nil {
			return nil, errors.Wrap(
				err, fmt.Sprintf(
					"Failed to parse response body as Error: %s", string(b)))
		}
		return nil, errors.New(e.Message)
	}
	role := Role{}
	err = json.Unmarshal(b, &role)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Role: %s", string(b)))
	}
	return &role, nil
}

// DeleteRole
// DELETE /roles/{rolename} Remove the named role and dissociate any users from it
func (client *Client) DeleteRole(name string) error {
	return client.DeleteRoleContext(context.Background(), name)
}

// DeleteRoleContext
// DELETE /roles/{rolename} Remove the named role and dissociate any users from it
func (client *Client) DeleteRoleContext(
	ctx context.Context, name string,
) error {
	req, err := http.NewRequest(
		http.MethodDelete, fmt.Sprintf("%s/%s", client.endpoints.Roles, name), nil)
	if err != nil {
		return errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to call DELETE /roles API")
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "Failed to read response body")
		}
		e := Error{}
		err = json.Unmarshal(b, &e)
		if err != nil {
			return errors.Wrap(
				err, fmt.Sprintf(
					"Failed to parse response body as Error: %s", string(b)))
		}
		return errors.New(e.Message)
	}
	return nil
}
