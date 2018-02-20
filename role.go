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
	"net/http"
)

type Role struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	ReadOnly    bool     `json:"read_only,omitempty"`
}

// CreateRole create a new role
func (client *Client) CreateRole(params *Role) (*Role, error) {
	return client.CreateRoleContext(context.Background(), params)
}

// CreateRoleContext create a new role
func (client *Client) CreateRoleContext(
	ctx context.Context, params *Role,
) (*Role, error) {
	// POST /roles
	hc := &http.Client{}
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	u, err := client.getUrl("/roles")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"POST", u, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(client.GetName(), client.GetPassword())
	req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return nil, nil
}
