package client

import (
	"context"
	"errors"
)

// CreateInputStaticFields creates an input static field.
func (client *Client) CreateInputStaticField(inputID, key, value string) (
	ei *ErrorInfo, err error,
) {
	return client.CreateInputStaticFieldContext(
		context.Background(), inputID, key, value)
}

// CreateInputStaticFieldContext creates an input static field with a context.
func (client *Client) CreateInputStaticFieldContext(
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
func (client *Client) DeleteInputStaticField(inputID, key string) (*ErrorInfo, error) {
	return client.DeleteInputStaticFieldContext(context.Background(), inputID, key)
}

// DeleteInputStaticFieldContext deletes an given input static field with a context.
func (client *Client) DeleteInputStaticFieldContext(
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
