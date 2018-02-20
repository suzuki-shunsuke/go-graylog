package graylog

// Graylog REST API http://docs.graylog.org/en/2.4/pages/configuration/rest_api.html
// Permission system http://docs.graylog.org/en/2.4/pages/users_and_roles/permission_system.html
// Acccess Token http://docs.graylog.org/en/2.4/pages/configuration/rest_api.html#creating-and-using-access-token
// Session Token http://docs.graylog.org/en/2.4/pages/configuration/rest_api.html#creating-and-using-session-token
// https://golang.org/pkg/net/http/#Request.SetBasicAuth
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
		"POST", client.endpoints.Roles, bytes.NewBuffer(b))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	req.SetBasicAuth(client.GetName(), client.GetPassword())
	req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to call POST /roles API")
	}
	defer resp.Body.Close()
	// read status code
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
