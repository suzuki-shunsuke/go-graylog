package client

import (
	"context"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetPipelineRules returns all pipeline rules.
func (client *Client) GetPipelineRules(ctx context.Context) (
	[]graylog.PipelineRule, *ErrorInfo, error,
) {
	rules := []graylog.PipelineRule{}
	ei, err := client.callGet(
		ctx, client.Endpoints().PipelineRules(), nil, &rules)
	return rules, ei, err
}

// GetPipelineRule returns a pipeline rule.
func (client *Client) GetPipelineRule(ctx context.Context, id string) (
	*graylog.PipelineRule, *ErrorInfo, error,
) {
	rule := &graylog.PipelineRule{}
	u, err := client.Endpoints().PipelineRule(id)
	if err != nil {
		return rule, nil, err
	}
	ei, err := client.callGet(ctx, u.String(), nil, rule)
	return rule, ei, err
}

// CreatePipelineRule creates a pipeline rule.
func (client *Client) CreatePipelineRule(
	ctx context.Context, rule *graylog.PipelineRule,
) (*ErrorInfo, error) {
	return client.callPost(
		ctx, client.Endpoints().PipelineRules(), rule, rule)
}

// UpdatePipelineRule updates a pipeline rule.
func (client *Client) UpdatePipelineRule(
	ctx context.Context, rule *graylog.PipelineRule,
) (*ErrorInfo, error) {
	u, err := client.Endpoints().PipelineRule(rule.ID)
	if err != nil {
		return nil, err
	}
	defer func(id string) {
		rule.ID = id
	}(rule.ID)
	rule.ID = ""
	return client.callPut(ctx, u.String(), rule, rule)
}

// DeletePipelineRule deletes a pipeline rule.
func (client *Client) DeletePipelineRule(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	u, err := client.Endpoints().PipelineRule(id)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
