package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
)

// CreateCollectorConfigurationOutput creates a collector configuration output.
func (client *Client) CreateCollectorConfigurationOutput(
	ctx context.Context, id string, output *graylog.CollectorConfigurationOutput,
) (*ErrorInfo, error) {
	// POST /plugins/org.graylog.plugins.collector/configurations/{id}/outputs Create a configuration output
	if id == "" {
		return nil, errors.New("id is required")
	}
	if output == nil {
		return nil, errors.New("collector configuration is nil")
	}
	// 202 no content
	return client.callPost(
		ctx, client.Endpoints().CollectorConfigurationOutputs(id), output, nil)
}

// DeleteCollectorConfigurationOutput deletes a collector configuration output.
func (client *Client) DeleteCollectorConfigurationOutput(
	ctx context.Context, id, outputID string,
) (*ErrorInfo, error) {
	// DELETE /plugins/org.graylog.plugins.collector/configurations/{id}/outputs/{outputId} Delete output form configuration
	if id == "" {
		return nil, errors.New("id is required")
	}
	if outputID == "" {
		return nil, errors.New("output id is required")
	}
	return client.callDelete(
		ctx, client.Endpoints().CollectorConfigurationOutput(id, outputID), nil, nil)
}

// UpdateCollectorConfigurationOutput updates a collector configuration output.
func (client *Client) UpdateCollectorConfigurationOutput(
	ctx context.Context, id, outputID string,
	output *graylog.CollectorConfigurationOutput,
) (*ErrorInfo, error) {
	// PUT /plugins/org.graylog.plugins.collector/configurations/{id}/outputs/{output_id} Update a configuration output
	if id == "" {
		return nil, errors.New("id is required")
	}
	if outputID == "" {
		return nil, errors.New("output id is required")
	}
	if output == nil {
		return nil, errors.New("output is nil")
	}
	return client.callPut(
		ctx, client.Endpoints().CollectorConfigurationOutput(id, outputID), output, nil)
}
