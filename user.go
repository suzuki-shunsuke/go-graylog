package graylog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// User represents a user.
type User struct {
	// ex. "local:admin"
	Username    string   `json:"username,omitempty" v-create:"required" v-update:"required"`
	Email       string   `json:"email,omitempty" v-create:"required"`
	Permissions []string `json:"permissions,omitempty" v-create:"required"`
	FullName    string   `json:"full_name,omitempty" v-create:"required"`
	Password    string   `json:"password,omitempty" v-create:"required"`

	Id          string       `json:"id,omitempty"`
	Preferences *Preferences `json:"preferences,omitempty"`
	// ex. "UTC"
	Timezone string `json:"timezone,omitempty"`
	// ex. 28800000
	SessionTimeoutMs int        `json:"session_timeout_ms,omitempty"`
	External         bool       `json:"external,omitempty"`
	Startpage        *Startpage `json:"startpage,omitempty"`
	// ex. ["Admin"]
	Roles         []string `json:"roles,omitempty"`
	ReadOnly      bool     `json:"read_only,omitempty"`
	SessionActive bool     `json:"session_active,omitempty"`
	// ex. "2018-03-02T06:32:01.841+0000"
	LastActivity string `json:"last_activity,omitempty"`
	// ex. "192.168.192.1"
	ClientAddress string `json:"client_address,omitempty"`
}

// Preferences represents user's preferences.
type Preferences struct {
	UpdateUnfocussed  bool `json:"updateUnfocussed,omitempty"`
	EnableSmartSearch bool `json:"enableSmartSearch,omitempty"`
}

// Startpage represents a user's startpage.
type Startpage struct {
	Type string `json:"type,omitempty"`
	Id   string `json:"id,omitempty"`
}

// CreateUser creates a new user account.
func (client *Client) CreateUser(user *User) (*ErrorInfo, error) {
	return client.CreateUserContext(context.Background(), user)
}

// CreateUserContext creates a new user account with a context.
func (client *Client) CreateUserContext(
	ctx context.Context, user *User,
) (*ErrorInfo, error) {
	b, err := json.Marshal(user)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(user)")
	}

	return client.callReq(
		ctx, http.MethodPost, client.endpoints.Users, b, false)
}

type usersBody struct {
	Users []User `json:"users"`
}

// GetUsers returns all users.
func (client *Client) GetUsers() ([]User, *ErrorInfo, error) {
	return client.GetUsersContext(context.Background())
}

// GetUsersContext returns all users with a context.
func (client *Client) GetUsersContext(ctx context.Context) ([]User, *ErrorInfo, error) {
	ei, err := client.callReq(
		ctx, http.MethodGet, client.endpoints.Users, nil, true)
	if err != nil {
		return nil, ei, err
	}

	users := usersBody{}
	err = json.Unmarshal(ei.ResponseBody, &users)
	if err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Users: %s",
				string(ei.ResponseBody)))
	}
	return users.Users, ei, nil
}

// GetUser returns a given user.
func (client *Client) GetUser(name string) (*User, *ErrorInfo, error) {
	return client.GetUserContext(context.Background(), name)
}

// GetUserContext returns a given user with a context.
func (client *Client) GetUserContext(
	ctx context.Context, name string,
) (*User, *ErrorInfo, error) {
	if name == "" {
		return nil, nil, errors.New("name is empty")
	}

	ei, err := client.callReq(
		ctx, http.MethodGet, client.endpoints.User(name), nil, true)
	if err != nil {
		return nil, ei, err
	}
	user := &User{}
	err = json.Unmarshal(ei.ResponseBody, user)
	if err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as User: %s",
				string(ei.ResponseBody)))
	}
	return user, ei, nil
}

// UpdateUser updates a given user.
func (client *Client) UpdateUser(user *User) (*ErrorInfo, error) {
	return client.UpdateUserContext(context.Background(), user)
}

// UpdateUserContext updates a given user with a context.
func (client *Client) UpdateUserContext(
	ctx context.Context, user *User,
) (*ErrorInfo, error) {
	if user.Username == "" {
		return nil, errors.New("name is empty")
	}
	b, err := json.Marshal(user)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(user)")
	}

	return client.callReq(
		ctx, http.MethodPut, client.endpoints.User(user.Username), b, false)
}

// DeleteUser deletes a given user.
func (client *Client) DeleteUser(name string) (*ErrorInfo, error) {
	return client.DeleteUserContext(context.Background(), name)
}

// DeleteUserContext deletes a given user with a context.
func (client *Client) DeleteUserContext(
	ctx context.Context, name string,
) (*ErrorInfo, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}

	return client.callReq(
		ctx, http.MethodDelete, client.endpoints.User(name), nil, false)
}
