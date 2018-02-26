package graylog

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// GetRoleMembers
// GET /roles/{rolename}/members Retrieve the role's members
func (client *Client) GetRoleMembers(name string) ([]User, error) {
	return client.GetRoleMembersContext(context.Background(), name)
}

// GetRoleMembersContext
// GET /roles/{rolename}/members Retrieve the role's members
func (client *Client) GetRoleMembersContext(
	ctx context.Context, name string,
) ([]User, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	req, err := http.NewRequest(
		http.MethodGet, client.RoleMembersEndpoint(name), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(
			err, "Failed to call GET /roles/{rolename}/members API")
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
	users := usersBody{}
	err = json.Unmarshal(b, &users)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Users: %s", string(b)))
	}
	return users.Users, nil
}

// AddUserToRole
// PUT /roles/{rolename}/members/{username} Add a user to a role
func (client *Client) AddUserToRole(userName, roleName string) error {
	return client.AddUserToRoleContext(context.Background(), userName, roleName)
}

// AddUserToRoleContext
// PUT /roles/{rolename}/members/{username} Add a user to a role
func (client *Client) AddUserToRoleContext(
	ctx context.Context, userName, roleName string,
) error {
	if userName == "" {
		return errors.New("userName is empty")
	}
	if roleName == "" {
		return errors.New("roleName is empty")
	}
	req, err := http.NewRequest(
		http.MethodPut, client.RoleMemberEndpoint(userName, roleName), nil)
	if err != nil {
		return errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return errors.Wrap(
			err, "Failed to call PUT /roles/{rolename}/members/{username} API")
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		e := Error{}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "Failed to read response body")
		}
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

// RemoveUserFromRole
// DELETE /roles/{rolename}/members/{username} Remove a user from a role
func (client *Client) RemoveUserFromRole(userName, roleName string) error {
	return client.RemoveUserFromRoleContext(context.Background(), userName, roleName)
}

// RemoveUserFromRoleContext
// DELETE /roles/{rolename}/members/{username} Remove a user from a role
func (client *Client) RemoveUserFromRoleContext(
	ctx context.Context, userName, roleName string,
) error {
	if userName == "" {
		return errors.New("userName is empty")
	}
	if roleName == "" {
		return errors.New("roleName is empty")
	}
	req, err := http.NewRequest(
		http.MethodDelete, client.RoleMemberEndpoint(userName, roleName), nil)
	if err != nil {
		return errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return errors.Wrap(
			err, "Failed to call DELETE /roles/{rolename}/members/{username} API")
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		e := Error{}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "Failed to read response body")
		}
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
