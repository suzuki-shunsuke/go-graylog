package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
)

// CreateCollectorConfigurationInput creates a collector configuration input.
func (client *Client) CreateCollectorConfigurationInput(
	ctx context.Context, id string, input *graylog.CollectorConfigurationInput,
) (*ErrorInfo, error) {
	// POST /plugins/org.graylog.plugins.collector/configurations/{id}/inputs Create a configuration input
	if id == "" {
		return nil, errors.New("id is required")
	}
	if input == nil {
		return nil, errors.New("collector configuration is nil")
	}
	// 202 no content
	return client.callPost(
		ctx, client.Endpoints().CollectorConfigurationInputs(id), input, nil)
}

// DeleteCollectorConfigurationInput deletes a collector configuration input.
func (client *Client) DeleteCollectorConfigurationInput(
	ctx context.Context, id, inputID string,
) (*ErrorInfo, error) {
	// DELETE /plugins/org.graylog.plugins.collector/configurations/{id}/inputs/{inputId} Delete input form configuration
	if id == "" {
		return nil, errors.New("id is required")
	}
	if inputID == "" {
		return nil, errors.New("input id is required")
	}
	return client.callDelete(
		ctx, client.Endpoints().CollectorConfigurationInput(id, inputID), nil, nil)
}

// UpdateCollectorConfigurationInput updates a collector configuration input.
func (client *Client) UpdateCollectorConfigurationInput(
	ctx context.Context, id, inputID string,
	input *graylog.CollectorConfigurationInput,
) (*ErrorInfo, error) {
	// PUT /plugins/org.graylog.plugins.collector/configurations/{id}/inputs/{input_id} Update a configuration input
	if id == "" {
		return nil, errors.New("id is required")
	}
	if inputID == "" {
		return nil, errors.New("input id is required")
	}
	if input == nil {
		return nil, errors.New("input is nil")
	}
	return client.callPut(
		ctx, client.Endpoints().CollectorConfigurationInput(id, inputID), input, nil)
}
