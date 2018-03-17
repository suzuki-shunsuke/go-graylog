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

// StreamRule represents a stream rule.
type StreamRule struct {
	// ex. "5a9b53c7c006c6000127f965"
	ID    string `json:"id,omitempty" v-create:"isdefault" v-update:"required"`
	Field string `json:"field,omitempty" v-create:"required" v-update:"required"`
	// ex. "5a94abdac006c60001f04fc1"
	StreamID    string `json:"stream_id,omitempty" v-create:"required" v-update:"required"`
	Description string `json:"description,omitempty"`
	Type        int    `json:"type,omitempty"`
	Inverted    bool   `json:"inverted,omitempty"`
	Value       string `json:"value,omitempty" v-create:"required" v-update:"required"`
}

type StreamRulesBody struct {
	Total       int          `json:"total"`
	StreamRules []StreamRule `json:"stream_rules"`
}

// GetStreamRules returns a list of all stream rules
func (client *Client) GetStreamRules(streamID string) (
	streamRules []StreamRule, total int, ei *ErrorInfo, err error,
) {
	return client.GetStreamRulesContext(context.Background(), streamID)
}

// GetStreamRulesContext returns a list of all stream rules with a context.
func (client *Client) GetStreamRulesContext(
	ctx context.Context, streamID string,
) (streamRules []StreamRule, total int, ei *ErrorInfo, err error) {
	// GET /streams/{streamid}/rules Get a list of all stream rules
	ei, err = client.callReq(
		ctx, http.MethodGet, client.Endpoints.StreamRules(streamID), nil, true)
	if err != nil {
		return nil, 0, ei, err
	}

	body := &StreamRulesBody{}
	if err := json.Unmarshal(ei.ResponseBody, body); err != nil {
		return nil, 0, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as StreamRulesBody: %s",
				string(ei.ResponseBody)))
	}
	return body.StreamRules, body.Total, ei, nil
}

type streamRuleIDBody struct {
	StreamRuleID string `json:"streamrule_id"`
}

// CreateStreamRule creates a stream
func (client *Client) CreateStreamRule(rule *StreamRule) (*ErrorInfo, error) {
	return client.CreateStreamRuleContext(context.Background(), rule)
}

// CreateStreamRuleContext creates a stream with a context
func (client *Client) CreateStreamRuleContext(
	ctx context.Context, rule *StreamRule,
) (*ErrorInfo, error) {
	// POST /streams/{streamid}/rules Create a stream rule
	if rule == nil {
		return nil, errors.New("rule is required")
	}
	cr := *rule
	cr.StreamID = ""
	b, err := json.Marshal(cr)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(stream)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPost, client.Endpoints.StreamRules(rule.StreamID), b, true)
	if err != nil {
		return ei, err
	}

	body := &streamRuleIDBody{}
	if err := json.Unmarshal(ei.ResponseBody, body); err != nil {
		return ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body: %s", string(ei.ResponseBody)))
	}
	rule.ID = body.StreamRuleID
	return ei, nil
}

// UpdateStreamRule updates a stream rule
func (client *Client) UpdateStreamRule(rule *StreamRule) (*ErrorInfo, error) {
	return client.UpdateStreamRuleContext(context.Background(), rule)
}

// UpdateStreamRuleContext updates a stream rule
func (client *Client) UpdateStreamRuleContext(
	ctx context.Context, rule *StreamRule,
) (*ErrorInfo, error) {
	// PUT /streams/{streamid}/rules/{streamRuleID} Update a stream rule
	if rule == nil {
		return nil, errors.New("rule is required")
	}
	streamID := rule.StreamID
	if streamID == "" {
		return nil, errors.New("streamID is empty")
	}
	ruleID := rule.ID
	if ruleID == "" {
		return nil, errors.New("streamRuleID is empty")
	}
	rule.StreamID = ""
	rule.ID = ""
	b, err := json.Marshal(rule)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(rule)")
	}

	return client.callReq(
		ctx, http.MethodPut, client.Endpoints.StreamRule(
			streamID, ruleID), b, false)
}

// DeleteStreamRule deletes a stream rule
func (client *Client) DeleteStreamRule(streamID, ruleID string) (*ErrorInfo, error) {
	return client.DeleteStreamRuleContext(context.Background(), streamID, ruleID)
}

// DeleteStreamRuleContext deletes a stream rule with a context
func (client *Client) DeleteStreamRuleContext(
	ctx context.Context, streamID, ruleID string,
) (*ErrorInfo, error) {
	// DELETE /streams/{streamid}/rules/{streamRuleID} Delete a stream rule
	if streamID == "" {
		return nil, errors.New("stream id is required")
	}
	if ruleID == "" {
		return nil, errors.New("stream rule id is required")
	}
	return client.callReq(
		ctx, http.MethodDelete, client.Endpoints.StreamRule(
			streamID, ruleID), nil, false)
}

// GetStreamRule returns a stream rule
func (client *Client) GetStreamRule(streamID, ruleID string) (*StreamRule, *ErrorInfo, error) {
	return client.GetStreamRuleContext(context.Background(), streamID, ruleID)
}

// GetStreamRuleContext returns a stream rule with a context
func (client *Client) GetStreamRuleContext(
	ctx context.Context, streamID, ruleID string,
) (*StreamRule, *ErrorInfo, error) {
	// GET /streams/{streamid}/rules/{streamRuleID} Get a single stream rules
	if streamID == "" {
		return nil, nil, errors.New("stream id is required")
	}
	if ruleID == "" {
		return nil, nil, errors.New("stream rule id is required")
	}

	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.StreamRule(
			streamID, ruleID), nil, true)
	if err != nil {
		return nil, ei, err
	}

	rule := &StreamRule{}
	if err := json.Unmarshal(ei.ResponseBody, rule); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as StreamRule: %s",
				string(ei.ResponseBody)))
	}
	return rule, ei, nil
}
