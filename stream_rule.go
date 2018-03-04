package graylog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// GetStreamRuleTypes returns all available stream types
// GET /streams/{streamid}/rules/types Get all available stream types
// GetStreamRule returns a stream rule
// GET /streams/{streamid}/rules/{streamRuleId} Get a single stream rules
// UpdateStreamRule updates a stream rule
// PUT /streams/{streamid}/rules/{streamRuleId} Update a stream rule
// DeleteStreamRule deletes a stream rule
// DELETE /streams/{streamid}/rules/{streamRuleId} Delete a stream rule

// StreamRule represents a stream rule.
type StreamRule struct {
	// ex. "5a9b53c7c006c6000127f965"
	Id    string `json:"id,omitempty"`
	Field string `json:"field,omitempty"`
	// ex. "5a94abdac006c60001f04fc1"
	StreamId    string `json:"stream_id,omitempty"`
	Description string `json:"description,omitempty"`
	Type        int    `json:"type,omitempty"`
	Inverted    bool   `json:"inverted,omitempty"`
	Value       string `json:"value,omitempty"`
}

// GetStreamRules returns a list of all stream rules
func (client *Client) GetStreamRules(streamId string) (
	streamRules []StreamRule, total int, ei *ErrorInfo, err error,
) {
	return client.GetStreamRulesContext(context.Background(), streamId)
}

type streamRulesBody struct {
	Total       int          `json:"total"`
	StreamRules []StreamRule `json:"stream_rules"`
}

// GetStreamRulesContext returns a list of all stream rules with a context.
func (client *Client) GetStreamRulesContext(
	ctx context.Context, streamId string,
) (streamRules []StreamRule, total int, ei *ErrorInfo, err error) {
	// GET /streams/{streamid}/rules Get a list of all stream rules
	ei, err = client.callReq(
		ctx, http.MethodGet, client.endpoints.StreamRules(streamId), nil, true)
	if err != nil {
		return nil, 0, ei, err
	}

	body := &streamRulesBody{}
	if err := json.Unmarshal(ei.ResponseBody, body); err != nil {
		return nil, 0, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as streamRulesBody: %s",
				string(ei.ResponseBody)))
	}
	return body.StreamRules, body.Total, ei, nil
}

// CreateStreamRule creates a stream
func (client *Client) CreateStreamRule(streamId string, rule *StreamRule) (
	string, *ErrorInfo, error,
) {
	return client.CreateStreamRuleContext(context.Background(), streamId, rule)
}

type streamRuleIdBody struct {
	StreamRuleId string `json:"streamrule_id"`
}

// CreateStreamRuleContext creates a stream with a context
func (client *Client) CreateStreamRuleContext(
	ctx context.Context, streamId string, rule *StreamRule,
) (ruleId string, ei *ErrorInfo, err error) {
	// POST /streams/{streamid}/rules Create a stream rule
	if rule == nil {
		return "", nil, errors.New("rule is required")
	}
	if streamId == "" {
		return "", nil, errors.New("streamId is required")
	}
	b, err := json.Marshal(rule)
	if err != nil {
		return "", nil, errors.Wrap(err, "Failed to json.Marshal(stream)")
	}

	ei, err = client.callReq(
		ctx, http.MethodPost, client.endpoints.StreamRules(streamId), b, true)
	if err != nil {
		return "", ei, err
	}

	body := &streamRuleIdBody{}
	if err = json.Unmarshal(ei.ResponseBody, body); err != nil {
		return "", ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body: %s", string(ei.ResponseBody)))
	}
	return body.StreamRuleId, ei, nil
}
