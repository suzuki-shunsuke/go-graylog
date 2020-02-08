package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog/v10"
)

// CreateOutput creates a new output.
func (client *Client) CreateOutput(
	ctx context.Context, output *graylog.Output,
) (*ErrorInfo, error) {
	// required: title, type, configuration
	if output == nil {
		return nil, errors.New("output is nil")
	}
	return client.callPost(
		ctx, client.Endpoints().Outputs(),
		map[string]interface{}{
			"title":         output.Title,
			"type":          output.Type,
			"configuration": output.Configuration,
		}, output)
}

// GetOutputs returns all outputs.
func (client *Client) GetOutputs(ctx context.Context) (
	[]graylog.Output, int, *ErrorInfo, error,
) {
	outputs := &graylog.OutputsBody{}
	ei, err := client.callGet(ctx, client.Endpoints().Outputs(), nil, outputs)
	return outputs.Outputs, outputs.Total, ei, err
}

// GetOutput returns a given output.
func (client *Client) GetOutput(
	ctx context.Context, id string,
) (*graylog.Output, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	output := &graylog.Output{}
	ei, err := client.callGet(ctx, client.Endpoints().Output(id), nil, output)
	return output, ei, err
}

// UpdateOutput updates a given output.
func (client *Client) UpdateOutput(
	ctx context.Context, output *graylog.Output,
) (*ErrorInfo, error) {
	if output == nil {
		return nil, errors.New("output is nil")
	}
	if output.ID == "" {
		return nil, errors.New("id is empty")
	}
	ei, err := client.callPut(
		ctx, client.Endpoints().Output(output.ID), map[string]interface{}{
			"title":         output.Title,
			"type":          output.Type,
			"configuration": output.Configuration,
		}, output)
	return ei, err
}

// DeleteOutput deletes a given output.
func (client *Client) DeleteOutput(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return client.callDelete(ctx, client.Endpoints().Output(id), nil, nil)
}
