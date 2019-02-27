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
