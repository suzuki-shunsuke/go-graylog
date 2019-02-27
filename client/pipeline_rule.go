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
