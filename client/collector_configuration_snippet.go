package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog/v9"
)

// CreateCollectorConfigurationSnippet creates a collector configuration snippet.
func (client *Client) CreateCollectorConfigurationSnippet(
	ctx context.Context, id string, snippet *graylog.CollectorConfigurationSnippet,
) (*ErrorInfo, error) {
	// POST /plugins/org.graylog.plugins.collector/configurations/{id}/snippets Create a configuration snippet
	if id == "" {
		return nil, errors.New("id is required")
	}
	if snippet == nil {
		return nil, errors.New("collector configuration is nil")
	}
	return client.callPost(
		ctx, client.Endpoints().CollectorConfigurationSnippets(id), snippet, nil)
}

// DeleteCollectorConfigurationSnippet deletes a collector configuration snippet.
func (client *Client) DeleteCollectorConfigurationSnippet(
	ctx context.Context, id, snippetID string,
) (*ErrorInfo, error) {
	// DELETE /plugins/org.graylog.plugins.collector/configurations/{id}/snippets/{snippetId} Delete snippet form configuration
	if id == "" {
		return nil, errors.New("id is required")
	}
	if snippetID == "" {
		return nil, errors.New("snippet id is required")
	}
	return client.callDelete(
		ctx, client.Endpoints().CollectorConfigurationSnippet(id, snippetID), nil, nil)
}

// UpdateCollectorConfigurationSnippet updates a collector configuration snippet.
func (client *Client) UpdateCollectorConfigurationSnippet(
	ctx context.Context, id, snippetID string,
	snippet *graylog.CollectorConfigurationSnippet,
) (*ErrorInfo, error) {
	// PUT /plugins/org.graylog.plugins.collector/configurations/{id}/snippets/{snippet_id} Update a configuration snippet
	if id == "" {
		return nil, errors.New("id is required")
	}
	if snippetID == "" {
		return nil, errors.New("snippet id is required")
	}
	if snippet == nil {
		return nil, errors.New("snippet is nil")
	}
	return client.callPut(
		ctx, client.Endpoints().CollectorConfigurationSnippet(id, snippetID), snippet, nil)
}
