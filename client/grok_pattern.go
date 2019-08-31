package client

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/suzuki-shunsuke/go-graylog"
)

// CreateGrokPattern creates a new grok pattern.
func (client *Client) CreateGrokPattern(
	ctx context.Context, grokPattern *graylog.GrokPattern,
) (*ErrorInfo, error) {
	// required: name, pattern
	// allowed: pattern, name, content_pack, id
	if grokPattern == nil {
		return nil, fmt.Errorf("grok pattern is nil")
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
	u, err := client.Endpoints().GrokPattern(id)
	if err != nil {
		return nil, nil, err
	}
	grokPattern := &graylog.GrokPattern{}
	ei, err := client.callGet(ctx, u.String(), nil, grokPattern)
	return grokPattern, ei, err
}

// UpdateGrokPattern updates a given grok pattern.
func (client *Client) UpdateGrokPattern(
	ctx context.Context, grokPattern *graylog.GrokPattern,
) (*ErrorInfo, error) {
	if grokPattern == nil {
		return nil, fmt.Errorf("grok pattern is nil")
	}
	if grokPattern.ID == "" {
		return nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().GrokPattern(grokPattern.ID)
	if err != nil {
		return nil, err
	}
	return client.callPut(ctx, u.String(), grokPattern, grokPattern)
}

// DeleteGrokPattern deletes a given grok pattern.
func (client *Client) DeleteGrokPattern(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().GrokPattern(id)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
