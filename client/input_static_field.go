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
	u, err := client.Endpoints().InputStaticFields(inputID)
	if err != nil {
		return nil, err
	}
	return client.callPost(
		ctx, u.String(), map[string]string{
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
	u, err := client.Endpoints().InputStaticField(inputID, key)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
