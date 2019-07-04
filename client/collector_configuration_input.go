package client

import (
	"context"
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

// CreateCollectorConfigurationInput creates a collector configuration input.
func (client *Client) CreateCollectorConfigurationInput(
	ctx context.Context, id string, input *graylog.CollectorConfigurationInput,
) (*ErrorInfo, error) {
	// POST /plugins/org.graylog.plugins.collector/configurations/{id}/inputs Create a configuration input
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if input == nil {
		return nil, fmt.Errorf("collector configuration is nil")
	}
	u, err := client.Endpoints().CollectorConfigurationInputs(id)
	if err != nil {
		return nil, err
	}
	// 202 no content
	return client.callPost(ctx, u.String(), input, nil)
}

// DeleteCollectorConfigurationInput deletes a collector configuration input.
func (client *Client) DeleteCollectorConfigurationInput(
	ctx context.Context, id, inputID string,
) (*ErrorInfo, error) {
	// DELETE /plugins/org.graylog.plugins.collector/configurations/{id}/inputs/{inputId} Delete input form configuration
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if inputID == "" {
		return nil, fmt.Errorf("input id is required")
	}
	u, err := client.Endpoints().CollectorConfigurationInput(id, inputID)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}

// UpdateCollectorConfigurationInput updates a collector configuration input.
func (client *Client) UpdateCollectorConfigurationInput(
	ctx context.Context, id, inputID string,
	input *graylog.CollectorConfigurationInput,
) (*ErrorInfo, error) {
	// PUT /plugins/org.graylog.plugins.collector/configurations/{id}/inputs/{input_id} Update a configuration input
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if inputID == "" {
		return nil, fmt.Errorf("input id is required")
	}
	if input == nil {
		return nil, fmt.Errorf("input is nil")
	}
	u, err := client.Endpoints().CollectorConfigurationInput(id, inputID)
	if err != nil {
		return nil, err
	}
	return client.callPut(ctx, u.String(), input, nil)
}
