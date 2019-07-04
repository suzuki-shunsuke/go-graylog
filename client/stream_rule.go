package client

import (
	"context"

	"github.com/pkg/errors"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetStreamRuleTypes returns all available stream types
// GET /streams/{streamid}/rules/types Get all available stream types

type streamRuleIDBody struct {
	StreamRuleID string `json:"streamrule_id"`
}

// GetStreamRules returns a list of all stream rules.
func (client *Client) GetStreamRules(
	ctx context.Context, streamID string,
) (streamRules []graylog.StreamRule, total int, ei *ErrorInfo, err error) {
	// GET /streams/{streamid}/rules Get a list of all stream rules
	u, err := client.Endpoints().StreamRules(streamID)
	if err != nil {
		return nil, 0, nil, err
	}
	body := &graylog.StreamRulesBody{}
	ei, err = client.callGet(ctx, u.String(), nil, body)
	return body.StreamRules, body.Total, ei, err
}

// CreateStreamRule creates a stream.
func (client *Client) CreateStreamRule(
	ctx context.Context, rule *graylog.StreamRule,
) (*ErrorInfo, error) {
	// POST /streams/{streamid}/rules Create a stream rule
	if rule == nil {
		return nil, errors.New("rule is required")
	}
	u, err := client.Endpoints().StreamRules(rule.StreamID)
	if err != nil {
		return nil, err
	}

	cr := *rule
	cr.StreamID = ""
	body := &streamRuleIDBody{}
	ei, err := client.callPost(ctx, u.String(), &cr, body)
	rule.ID = body.StreamRuleID
	return ei, err
}

// UpdateStreamRule updates a stream rule
func (client *Client) UpdateStreamRule(
	ctx context.Context, rule *graylog.StreamRule,
) (*ErrorInfo, error) {
	// PUT /streams/{streamid}/rules/{streamRuleID} Update a stream rule
	if rule == nil {
		return nil, errors.New("rule is required")
	}
	if rule.StreamID == "" {
		return nil, errors.New("streamID is empty")
	}
	if rule.ID == "" {
		return nil, errors.New("streamRuleID is empty")
	}
	u, err := client.Endpoints().StreamRule(rule.StreamID, rule.ID)
	if err != nil {
		return nil, err
	}
	cr := *rule
	cr.StreamID = ""
	cr.ID = ""
	return client.callPut(ctx, u.String(), &cr, nil)
}

// DeleteStreamRule deletes a stream rule.
func (client *Client) DeleteStreamRule(
	ctx context.Context, streamID, ruleID string,
) (*ErrorInfo, error) {
	// DELETE /streams/{streamid}/rules/{streamRuleID} Delete a stream rule
	if streamID == "" {
		return nil, errors.New("stream id is required")
	}
	if ruleID == "" {
		return nil, errors.New("stream rule id is required")
	}
	u, err := client.Endpoints().StreamRule(streamID, ruleID)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}

// GetStreamRule returns a stream rule.
func (client *Client) GetStreamRule(
	ctx context.Context, streamID, ruleID string,
) (*graylog.StreamRule, *ErrorInfo, error) {
	// GET /streams/{streamid}/rules/{streamRuleID} Get a single stream rules
	if streamID == "" {
		return nil, nil, errors.New("stream id is required")
	}
	if ruleID == "" {
		return nil, nil, errors.New("stream rule id is required")
	}
	u, err := client.Endpoints().StreamRule(streamID, ruleID)
	if err != nil {
		return nil, nil, err
	}
	rule := &graylog.StreamRule{}
	ei, err := client.callGet(ctx, u.String(), nil, rule)
	return rule, ei, err
}
