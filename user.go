package graylog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type User struct {
	Id               string       `json:"id,omitempty"`
	Username         string       `json:"username,omitempty"`
	Email            string       `json:"email,omitempty"`
	FullName         string       `json:"full_name,omitempty"`
	Permissions      []string     `json:"permissions,omitempty"`
	Preferences      *Preferences `json:"preferences,omitempty"`
	Timezone         string       `json:"timezone,omitempty"`
	SessionTimeoutMs int          `json:"session_timeout_ms,omitempty"`
	External         bool         `json:"external,omitempty"`
	Startpage        *Startpage   `json:"startpage,omitempty"`
	Roles            []string     `json:"roles,omitempty"`
	ReadOnly         bool         `json:"read_only,omitempty"`
	SessionActive    bool         `json:"session_active,omitempty"`
	LastActivity     string       `json:"last_activity,omitempty"`
	ClientAddress    string       `json:"client_address,omitempty"`

	Password string `json:"password,omitempty"`
}

type Preferences struct {
	UpdateUnfocussed  bool `json:"updateUnfocussed,omitempty"`
	EnableSmartSearch bool `json:"enableSmartSearch,omitempty"`
}

type Startpage struct {
	Type string `json:"type,omitempty"`
	Id   string `json:"id,omitempty"`
}

// CreateUser
// POST /users Create a new user account.
func (client *Client) CreateUser(user *User) error {
	return client.CreateUserContext(context.Background(), user)
}

// CreateUserContext
// POST /users Create a new user account.
func (client *Client) CreateUserContext(
	ctx context.Context, user *User,
) error {
	b, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "Failed to json.Marshal(user)")
	}
	req, err := http.NewRequest(
		http.MethodPost, client.endpoints.Users, bytes.NewBuffer(b))
	if err != nil {
		return errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to call POST /users API")
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

type usersBody struct {
	Users []User `json:"users"`
}

// GetUsers
// GET /users List all users
func (client *Client) GetUsers() ([]User, error) {
	return client.GetUsersContext(context.Background())
}

// GetUsersContext
// GET /users List all users
func (client *Client) GetUsersContext(
	ctx context.Context,
) ([]User, error) {
	req, err := http.NewRequest(http.MethodGet, client.endpoints.Users, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to call GET /users API")
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

// GetUser
// GET /users/{username} Get user details
func (client *Client) GetUser(name string) (*User, error) {
	return client.GetUserContext(context.Background(), name)
}

// GetUserContext
// GET /users/{username} Get user details
func (client *Client) GetUserContext(
	ctx context.Context, name string,
) (*User, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	req, err := http.NewRequest(
		http.MethodGet, fmt.Sprintf("%s/%s", client.endpoints.Users, name), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to call GET /users API")
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
	user := User{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as User: %s", string(b)))
	}
	return &user, nil
}

// UpdateUser
// PUT /users/{username} Modify user details.
func (client *Client) UpdateUser(name string, user *User) error {
	return client.UpdateUserContext(context.Background(), name, user)
}

// UpdateUserContext
// PUT /users/{username} Modify user details.
func (client *Client) UpdateUserContext(
	ctx context.Context, name string, user *User,
) error {
	if name == "" {
		return errors.New("name is empty")
	}
	b, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "Failed to json.Marshal(user)")
	}
	req, err := http.NewRequest(
		http.MethodPut, fmt.Sprintf("%s/%s", client.endpoints.Users, name),
		bytes.NewBuffer(b))
	if err != nil {
		return errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to call PUT /users/{username} API")
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

// DeleteUser
// DELETE /users/{username} Removes a user account
func (client *Client) DeleteUser(name string) error {
	return client.DeleteUserContext(context.Background(), name)
}

// DeleteUserContext
// DELETE /users/{username} Removes a user account
func (client *Client) DeleteUserContext(
	ctx context.Context, name string,
) error {
	if name == "" {
		return errors.New("name is empty")
	}
	req, err := http.NewRequest(
		http.MethodDelete, fmt.Sprintf("%s/%s", client.endpoints.Users, name), nil)
	if err != nil {
		return errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to call DELETE /users API")
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
