package client

import (
	"context"
	"errors"
)

// CreateInputStaticField creates an input static field.
func (client *Client) CreateInputStaticField(
	ctx context.Context, inputID, key, value string,
) (ei *ErrorInfo, err error) {
	if inputID == "" {
		return nil, errors.New("input id is required")
	}
	if key == "" {
		return nil, errors.New("key is required")
	}
	if value == "" {
		return nil, errors.New("value is required")
	}
	return client.callPost(
		ctx, client.Endpoints().InputStaticFields(inputID), map[string]string{
			"key":   key,
			"value": value,
		}, nil)
}

// DeleteInputStaticField deletes an given input static field.
func (client *Client) DeleteInputStaticField(
	ctx context.Context, inputID, key string,
) (*ErrorInfo, error) {
	if inputID == "" {
		return nil, errors.New("id is empty")
	}
	if key == "" {
		return nil, errors.New("key is empty")
	}
	return client.callDelete(ctx, client.Endpoints().InputStaticField(inputID, key), nil, nil)
}
