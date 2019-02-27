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

// GetPipelineRule returns a pipeline rules.
func (client *Client) GetPipelineRule(id string) (*graylog.PipelineRule, *ErrorInfo, error) {
	return client.GetPipelineRuleContext(context.Background(), id)
}

// GetPipelineRuleContext returns a pipeline rules with a context.
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
	ei, err := client.callPost(
		ctx, client.Endpoints().PipelineRules(), rule, &rule)
	return ei, err
}
