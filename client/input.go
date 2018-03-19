package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

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
	b, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(input)")
	}

	ei, err = client.callReq(
		ctx, http.MethodPost, client.Endpoints.Inputs, b, true)
	if err != nil {
		return ei, err
	}

	if err := json.Unmarshal(ei.ResponseBody, input); err != nil {
		return ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Input: %s",
				string(ei.ResponseBody)))
	}
	return ei, nil
}

// GetInputs returns all inputs.
func (client *Client) GetInputs() ([]graylog.Input, *ErrorInfo, error) {
	return client.GetInputsContext(context.Background())
}

// GetInputsContext returns all inputs with a context.
func (client *Client) GetInputsContext(ctx context.Context) (
	[]graylog.Input, *ErrorInfo, error,
) {
	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.Inputs, nil, true)
	if err != nil {
		return nil, ei, err
	}

	inputs := &graylog.InputsBody{}
	if err := json.Unmarshal(ei.ResponseBody, inputs); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as Inputs: %s",
				string(ei.ResponseBody)))
	}
	return inputs.Inputs, ei, nil
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

	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.Input(id), nil, true)
	if err != nil {
		return nil, ei, err
	}

	input := &graylog.Input{}
	if err := json.Unmarshal(ei.ResponseBody, input); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as Input: %s", string(ei.ResponseBody)))
	}
	return input, ei, nil
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
	b, err := json.Marshal(copiedInput)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(input)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPut, client.Endpoints.Input(input.ID), b, true)
	if err != nil {
		return ei, err
	}

	if err := json.Unmarshal(ei.ResponseBody, input); err != nil {
		return ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as Input: %s", string(ei.ResponseBody)))
	}
	return ei, nil
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

	return client.callReq(
		ctx, http.MethodDelete, client.Endpoints.Input(id), nil, false)
}
