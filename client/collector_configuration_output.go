package client

import (
	"context"
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

// CreateCollectorConfigurationOutput creates a collector configuration output.
func (client *Client) CreateCollectorConfigurationOutput(
	id string, output *graylog.CollectorConfigurationOutput,
) (*ErrorInfo, error) {
	return client.CreateCollectorConfigurationOutputContext(
		context.Background(), id, output)
}

// CreateCollectorConfigurationOutputContext creates a collector configuration output with a context.
func (client *Client) CreateCollectorConfigurationOutputContext(
	ctx context.Context, id string, output *graylog.CollectorConfigurationOutput,
) (*ErrorInfo, error) {
	// POST /plugins/org.graylog.plugins.collector/configurations/{id}/outputs Create a configuration output
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if output == nil {
		return nil, fmt.Errorf("collector configuration is nil")
	}
	u, err := client.Endpoints().CollectorConfigurationOutputs(id)
	if err != nil {
		return nil, err
	}
	// 202 no content
	return client.callPost(ctx, u.String(), output, nil)
}

// DeleteCollectorConfigurationOutput deletes a collector configuration output.
func (client *Client) DeleteCollectorConfigurationOutput(id, outputID string) (*ErrorInfo, error) {
	return client.DeleteCollectorConfigurationOutputContext(
		context.Background(), id, outputID)
}

// DeleteCollectorConfigurationOutputContext deletes a collector configuration output with a context.
func (client *Client) DeleteCollectorConfigurationOutputContext(
	ctx context.Context, id, outputID string,
) (*ErrorInfo, error) {
	// DELETE /plugins/org.graylog.plugins.collector/configurations/{id}/outputs/{outputId} Delete output form configuration
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if outputID == "" {
		return nil, fmt.Errorf("output id is required")
	}
	u, err := client.Endpoints().CollectorConfigurationOutput(id, outputID)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}

// UpdateCollectorConfigurationOutput updates a collector configuration output.
func (client *Client) UpdateCollectorConfigurationOutput(
	id, outputID string, output *graylog.CollectorConfigurationOutput,
) (*ErrorInfo, error) {
	return client.UpdateCollectorConfigurationOutputContext(
		context.Background(), id, outputID, output)
}

// UpdateCollectorConfigurationOutputContext updates a collector configuration output with a context.
func (client *Client) UpdateCollectorConfigurationOutputContext(
	ctx context.Context, id, outputID string,
	output *graylog.CollectorConfigurationOutput,
) (*ErrorInfo, error) {
	// PUT /plugins/org.graylog.plugins.collector/configurations/{id}/outputs/{output_id} Update a configuration output
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if outputID == "" {
		return nil, fmt.Errorf("output id is required")
	}
	if output == nil {
		return nil, fmt.Errorf("output is nil")
	}
	u, err := client.Endpoints().CollectorConfigurationOutput(id, outputID)
	if err != nil {
		return nil, err
	}
	return client.callPut(ctx, u.String(), output, nil)
}
