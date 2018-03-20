package client

import (
	"context"
	"fmt"

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
	return client.callPost(ctx, client.Endpoints.Users, user, nil)
}

// GetUsers returns all users.
func (client *Client) GetUsers() ([]graylog.User, *ErrorInfo, error) {
	return client.GetUsersContext(context.Background())
}

// GetUsersContext returns all users with a context.
func (client *Client) GetUsersContext(ctx context.Context) ([]graylog.User, *ErrorInfo, error) {
	users := &graylog.UsersBody{}
	ei, err := client.callGet(
		ctx, client.Endpoints.Users, nil, users)
	return users.Users, ei, err
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
	user := &graylog.User{}
	ei, err := client.callGet(
		ctx, client.Endpoints.User(name), nil, user)
	return user, ei, err
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
	return client.callPut(ctx, client.Endpoints.User(user.Username), user, nil)
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
	return client.callDelete(ctx, client.Endpoints.User(name), nil, nil)
}
