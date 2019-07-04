package client

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/suzuki-shunsuke/go-graylog"
)

// CreateUser creates a new user account.
func (client *Client) CreateUser(
	ctx context.Context, user *graylog.User,
) (*ErrorInfo, error) {
	if user == nil {
		return nil, fmt.Errorf("user is nil")
	}
	return client.callPost(ctx, client.Endpoints().Users(), user, nil)
}

// GetUsers returns all users.
func (client *Client) GetUsers(ctx context.Context) ([]graylog.User, *ErrorInfo, error) {
	users := &graylog.UsersBody{}
	ei, err := client.callGet(ctx, client.Endpoints().Users(), nil, users)
	return users.Users, ei, err
}

// GetUser returns a given user.
func (client *Client) GetUser(
	ctx context.Context, name string,
) (*graylog.User, *ErrorInfo, error) {
	if name == "" {
		return nil, nil, errors.New("name is empty")
	}
	u, err := client.Endpoints().User(name)
	if err != nil {
		return nil, nil, err
	}
	user := &graylog.User{}
	ei, err := client.callGet(ctx, u.String(), nil, user)
	return user, ei, err
}

// UpdateUser updates a given user.
func (client *Client) UpdateUser(
	ctx context.Context, prms *graylog.UserUpdateParams,
) (*ErrorInfo, error) {
	if prms == nil {
		return nil, fmt.Errorf("user is nil")
	}
	if prms.Username == "" {
		return nil, errors.New("name is empty")
	}
	u, err := client.Endpoints().User(prms.Username)
	if err != nil {
		return nil, err
	}
	return client.callPut(ctx, u.String(), prms, nil)
}

// DeleteUser deletes a given user.
func (client *Client) DeleteUser(
	ctx context.Context, name string,
) (*ErrorInfo, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	u, err := client.Endpoints().User(name)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
