package client

import (
	"context"
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

// CreateCollectorConfigurationSnippet creates a collector configuration snippet.
func (client *Client) CreateCollectorConfigurationSnippet(
	ctx context.Context, id string, snippet *graylog.CollectorConfigurationSnippet,
) (*ErrorInfo, error) {
	// POST /plugins/org.graylog.plugins.collector/configurations/{id}/snippets Create a configuration snippet
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if snippet == nil {
		return nil, fmt.Errorf("collector configuration is nil")
	}
	u, err := client.Endpoints().CollectorConfigurationSnippets(id)
	if err != nil {
		return nil, err
	}
	return client.callPost(ctx, u.String(), snippet, nil)
}

// DeleteCollectorConfigurationSnippet deletes a collector configuration snippet.
func (client *Client) DeleteCollectorConfigurationSnippet(
	ctx context.Context, id, snippetID string,
) (*ErrorInfo, error) {
	// DELETE /plugins/org.graylog.plugins.collector/configurations/{id}/snippets/{snippetId} Delete snippet form configuration
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if snippetID == "" {
		return nil, fmt.Errorf("snippet id is required")
	}
	u, err := client.Endpoints().CollectorConfigurationSnippet(id, snippetID)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}

// UpdateCollectorConfigurationSnippet updates a collector configuration snippet.
func (client *Client) UpdateCollectorConfigurationSnippet(
	ctx context.Context, id, snippetID string,
	snippet *graylog.CollectorConfigurationSnippet,
) (*ErrorInfo, error) {
	// PUT /plugins/org.graylog.plugins.collector/configurations/{id}/snippets/{snippet_id} Update a configuration snippet
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if snippetID == "" {
		return nil, fmt.Errorf("snippet id is required")
	}
	if snippet == nil {
		return nil, fmt.Errorf("snippet is nil")
	}
	u, err := client.Endpoints().CollectorConfigurationSnippet(id, snippetID)
	if err != nil {
		return nil, err
	}
	return client.callPut(ctx, u.String(), snippet, nil)
}
