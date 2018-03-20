package client

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/suzuki-shunsuke/go-graylog"
)

// CreateInput creates an input.
func (client *Client) CreateInput(input *graylog.Input) (
	ei *ErrorInfo, err error,
) {
	return client.CreateInputContext(context.Background(), input)
}

// CreateInputContext creates an input with a context.
func (client *Client) CreateInputContext(
	ctx context.Context, input *graylog.Input,
) (ei *ErrorInfo, err error) {
	if input == nil {
		return nil, fmt.Errorf("input is nil")
	}
	return client.callPost(ctx, client.Endpoints.Inputs, input, input)
}

// GetInputs returns all inputs.
func (client *Client) GetInputs() ([]graylog.Input, *ErrorInfo, error) {
	return client.GetInputsContext(context.Background())
}

// GetInputsContext returns all inputs with a context.
func (client *Client) GetInputsContext(ctx context.Context) (
	[]graylog.Input, *ErrorInfo, error,
) {
	inputs := &graylog.InputsBody{}
	ei, err := client.callGet(
		ctx, client.Endpoints.Inputs, nil, inputs)
	return inputs.Inputs, ei, err
}

// GetInput returns a given input.
func (client *Client) GetInput(id string) (*graylog.Input, *ErrorInfo, error) {
	return client.GetInputContext(context.Background(), id)
}

// GetInputContext returns a given input with a context.
func (client *Client) GetInputContext(
	ctx context.Context, id string,
) (*graylog.Input, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	input := &graylog.Input{}
	ei, err := client.callGet(
		ctx, client.Endpoints.Input(id), nil, input)
	return input, ei, err
}

// UpdateInput updates an given input.
func (client *Client) UpdateInput(input *graylog.Input) (
	*ErrorInfo, error,
) {
	return client.UpdateInputContext(context.Background(), input)
}

// UpdateInputContext updates an given input with a context.
func (client *Client) UpdateInputContext(
	ctx context.Context, input *graylog.Input,
) (*ErrorInfo, error) {
	if input == nil {
		return nil, fmt.Errorf("input is nil")
	}
	if input.ID == "" {
		return nil, errors.New("id is empty")
	}

	copiedInput := *input
	copiedInput.ID = ""
	return client.callPut(ctx, client.Endpoints.Input(input.ID), &copiedInput, input)
}

// DeleteInput deletes an given input.
func (client *Client) DeleteInput(id string) (*ErrorInfo, error) {
	return client.DeleteInputContext(context.Background(), id)
}

// DeleteInputContext deletes an given input with a context.
func (client *Client) DeleteInputContext(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return client.callDelete(ctx, client.Endpoints.Input(id), nil, nil)
}
