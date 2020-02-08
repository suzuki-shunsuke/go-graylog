package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog/v10"
)

// GetStreamAlertConditions gets all alert conditions of this stream.
func (client *Client) GetStreamAlertConditions(
	ctx context.Context, streamID string,
) (conds []graylog.AlertCondition, total int, ei *ErrorInfo, err error) {
	conditions := &graylog.AlertConditionsBody{}
	ei, err = client.callGet(
		ctx, client.Endpoints().StreamAlertConditions(streamID), nil, conditions)
	return conditions.AlertConditions, conditions.Total, ei, err
}

// GetStreamAlertCondition gets an alert condition.
func (client *Client) GetStreamAlertCondition(
	ctx context.Context, streamID, id string,
) (graylog.AlertCondition, *ErrorInfo, error) {
	cond := graylog.AlertCondition{}
	if streamID == "" {
		return cond, nil, errStreamIDRequired
	}
	if id == "" {
		return cond, nil, errIDRequired
	}
	ei, err := client.callGet(ctx, client.Endpoints().StreamAlertCondition(streamID, id), nil, &cond)
	return cond, ei, err
}

// CreateStreamAlertCondition creates an alert condition.
func (client *Client) CreateStreamAlertCondition(
	ctx context.Context, streamID string, cond *graylog.AlertCondition,
) (*ErrorInfo, error) {
	errMsg := "failed to create an alert condition"
	if cond == nil {
		return nil, errors.New(errMsg + ": alert condition is nil")
	}
	ret := map[string]string{}
	ei, err := client.callPost(ctx, client.Endpoints().StreamAlertConditions(streamID), cond, &ret)
	if err != nil {
		return ei, fmt.Errorf(errMsg+": %w", err)
	}
	if id, ok := ret["alert_condition_id"]; ok {
		cond.ID = id
		return ei, nil
	}
	return ei, errors.New(errMsg + `: response doesn't have the field "alert_condition_id"`)
}

// UpdateStreamAlertCondition modifies an alert condition.
func (client *Client) UpdateStreamAlertCondition(
	ctx context.Context, streamID string, cond *graylog.AlertCondition,
) (*ErrorInfo, error) {
	errMsg := "failed to update an alert condition"
	if streamID == "" {
		return nil, errors.New(errMsg + ": stream id is empty")
	}
	if cond == nil {
		return nil, errors.New(errMsg + ": alert condition is nil")
	}
	condID := cond.ID
	if condID == "" {
		return nil, errors.New(errMsg + ": alert condition id is empty")
	}
	cond.ID = ""
	ei, err := client.callPut(ctx, client.Endpoints().StreamAlertCondition(streamID, condID), cond, nil)
	cond.ID = condID
	if err != nil {
		return ei, fmt.Errorf(errMsg+": %w", err)
	}
	return ei, nil
}

// DeleteStreamAlertCondition deletes an alert condition.
func (client *Client) DeleteStreamAlertCondition(
	ctx context.Context, streamID, id string,
) (*ErrorInfo, error) {
	errMsg := "failed to delete an alert condition"
	if streamID == "" {
		return nil, errors.New(errMsg + ": stream id is empty")
	}
	if id == "" {
		return nil, errors.New(errMsg + ": alert condition id is empty")
	}
	return client.callDelete(ctx, client.Endpoints().StreamAlertCondition(streamID, id), nil, nil)
}
