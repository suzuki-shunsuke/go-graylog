package client

import (
	"context"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetPipelineRules returns all pipeline rules.
func (client *Client) GetPipelineRules() ([]graylog.PipelineRule, *ErrorInfo, error) {
	return client.GetPipelineRulesContext(context.Background())
}

// GetPipelineRulesContext returns all pipeline rules with a context.
func (client *Client) GetPipelineRulesContext(ctx context.Context) (
	[]graylog.PipelineRule, *ErrorInfo, error,
) {
	rules := []graylog.PipelineRule{}
	ei, err := client.callGet(
		ctx, client.Endpoints().PipelineRules(), nil, &rules)
	return rules, ei, err
}

// GetPipelineRule returns a pipeline rule.
func (client *Client) GetPipelineRule(id string) (*graylog.PipelineRule, *ErrorInfo, error) {
	return client.GetPipelineRuleContext(context.Background(), id)
}

// GetPipelineRuleContext returns a pipeline rule with a context.
func (client *Client) GetPipelineRuleContext(ctx context.Context, id string) (
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
	rule *graylog.PipelineRule,
) (*ErrorInfo, error) {
	return client.CreatePipelineRuleContext(context.Background(), rule)
}

// CreatePipelineRuleContext creates a pipeline rule with a context.
func (client *Client) CreatePipelineRuleContext(
	ctx context.Context, rule *graylog.PipelineRule,
) (*ErrorInfo, error) {
	return client.callPost(
		ctx, client.Endpoints().PipelineRules(), rule, &rule)
}

// UpdatePipelineRule updates a pipeline rule.
func (client *Client) UpdatePipelineRule(
	rule *graylog.PipelineRule,
) (*ErrorInfo, error) {
	return client.UpdatePipelineRuleContext(context.Background(), rule)
}

// UpdatePipelineRuleContext updates a pipeline rule with a context.
func (client *Client) UpdatePipelineRuleContext(
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
func (client *Client) DeletePipelineRule(id string) (*ErrorInfo, error) {
	return client.DeletePipelineRuleContext(context.Background(), id)
}

// DeletePipelineRuleContext deletes a pipeline rule with a context.
func (client *Client) DeletePipelineRuleContext(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	u, err := client.Endpoints().PipelineRule(id)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
