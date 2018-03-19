package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/suzuki-shunsuke/go-graylog"
)

// CreateUser creates a new user account.
func (client *Client) CreateUser(user *graylog.User) (*ErrorInfo, error) {
	return client.CreateUserContext(context.Background(), user)
}

// CreateUserContext creates a new user account with a context.
func (client *Client) CreateUserContext(
	ctx context.Context, user *graylog.User,
) (*ErrorInfo, error) {
	if user == nil {
		return nil, fmt.Errorf("user is nil")
	}
	b, err := json.Marshal(user)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(user)")
	}

	return client.callReq(
		ctx, http.MethodPost, client.Endpoints.Users, b, false)
}

// GetUsers returns all users.
func (client *Client) GetUsers() ([]graylog.User, *ErrorInfo, error) {
	return client.GetUsersContext(context.Background())
}

// GetUsersContext returns all users with a context.
func (client *Client) GetUsersContext(ctx context.Context) ([]graylog.User, *ErrorInfo, error) {
	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.Users, nil, true)
	if err != nil {
		return nil, ei, err
	}

	users := &graylog.UsersBody{}
	if err := json.Unmarshal(ei.ResponseBody, users); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Users: %s",
				string(ei.ResponseBody)))
	}
	return users.Users, ei, nil
}

// GetUser returns a given user.
func (client *Client) GetUser(name string) (*graylog.User, *ErrorInfo, error) {
	return client.GetUserContext(context.Background(), name)
}

// GetUserContext returns a given user with a context.
func (client *Client) GetUserContext(
	ctx context.Context, name string,
) (*graylog.User, *ErrorInfo, error) {
	if name == "" {
		return nil, nil, errors.New("name is empty")
	}

	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.User(name), nil, true)
	if err != nil {
		return nil, ei, err
	}
	user := &graylog.User{}
	if err := json.Unmarshal(ei.ResponseBody, user); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as User: %s",
				string(ei.ResponseBody)))
	}
	return user, ei, nil
}

// UpdateUser updates a given user.
func (client *Client) UpdateUser(user *graylog.User) (*ErrorInfo, error) {
	return client.UpdateUserContext(context.Background(), user)
}

// UpdateUserContext updates a given user with a context.
func (client *Client) UpdateUserContext(
	ctx context.Context, user *graylog.User,
) (*ErrorInfo, error) {
	if user == nil {
		return nil, fmt.Errorf("user is nil")
	}
	if user.Username == "" {
		return nil, errors.New("name is empty")
	}
	b, err := json.Marshal(user)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(user)")
	}

	return client.callReq(
		ctx, http.MethodPut, client.Endpoints.User(user.Username), b, false)
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
		ctx, http.MethodDelete, client.Endpoints.User(name), nil, false)
}
