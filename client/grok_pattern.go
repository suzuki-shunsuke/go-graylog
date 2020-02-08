package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog/v10"
)

// CreateGrokPattern creates a new grok pattern.
func (client *Client) CreateGrokPattern(
	ctx context.Context, grokPattern *graylog.GrokPattern,
) (*ErrorInfo, error) {
	// required: name, pattern
	// allowed: pattern, name, content_pack, id
	if grokPattern == nil {
		return nil, errors.New("grok pattern is nil")
	}
	return client.callPost(ctx, client.Endpoints().GrokPatterns(), grokPattern, grokPattern)
}

// GetGrokPatterns returns all grok patterns.
func (client *Client) GetGrokPatterns(ctx context.Context) (
	[]graylog.GrokPattern, *ErrorInfo, error,
) {
	grokPatterns := &graylog.GrokPatternsBody{}
	ei, err := client.callGet(ctx, client.Endpoints().GrokPatterns(), nil, grokPatterns)
	return grokPatterns.Patterns, ei, err
}

// GetGrokPattern returns a given grok pattern.
func (client *Client) GetGrokPattern(
	ctx context.Context, id string,
) (*graylog.GrokPattern, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	grokPattern := &graylog.GrokPattern{}
	ei, err := client.callGet(ctx, client.Endpoints().GrokPattern(id), nil, grokPattern)
	return grokPattern, ei, err
}

// UpdateGrokPattern updates a given grok pattern.
func (client *Client) UpdateGrokPattern(
	ctx context.Context, grokPattern *graylog.GrokPattern,
) (*ErrorInfo, error) {
	if grokPattern == nil {
		return nil, errors.New("grok pattern is nil")
	}
	if grokPattern.ID == "" {
		return nil, errors.New("id is empty")
	}
	return client.callPut(
		ctx, client.Endpoints().GrokPattern(grokPattern.ID), grokPattern, grokPattern)
}

// DeleteGrokPattern deletes a given grok pattern.
func (client *Client) DeleteGrokPattern(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return client.callDelete(ctx, client.Endpoints().GrokPattern(id), nil, nil)
}
