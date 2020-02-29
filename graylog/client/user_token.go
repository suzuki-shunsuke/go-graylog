package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
)

// CreateUserToken generates a new access token for a user
func (client *Client) CreateUserToken(
	ctx context.Context, userName, tokenName string,
) (*graylog.UserToken, *ErrorInfo, error) {
	if userName == "" {
		return nil, nil, errors.New("user name is required")
	}
	if tokenName == "" {
		return nil, nil, errors.New("token name is required")
	}
	token := &graylog.UserToken{}

	ei, err := client.callPost(ctx, client.Endpoints().UserToken(userName, tokenName), map[string]interface{}{}, token)
	return token, ei, err
}

// GetUserTokens returns the list of access tokens for a user.
func (client *Client) GetUserTokens(ctx context.Context, name string) ([]graylog.UserToken, *ErrorInfo, error) {
	tokens := map[string][]graylog.UserToken{}
	ei, err := client.callGet(ctx, client.Endpoints().UserTokens(name), nil, &tokens)
	if a, ok := tokens["tokens"]; ok {
		return a, ei, err
	}
	return nil, ei, err
}

// DeleteUserToken removes a token for a user.
func (client *Client) DeleteUserToken(
	ctx context.Context, name, token string,
) (*ErrorInfo, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	if token == "" {
		return nil, errors.New("name is empty")
	}
	return client.callDelete(ctx, client.Endpoints().UserToken(name, token), nil, nil)
}
