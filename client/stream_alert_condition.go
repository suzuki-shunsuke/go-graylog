package client

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetStreamAlertConditions gets all alert conditions of this stream.
func (client *Client) GetStreamAlertConditions(
	ctx context.Context, streamID string,
) (conds []graylog.AlertCondition, total int, ei *ErrorInfo, err error) {
	errMsg := "failed to get stream's alert conditions"
	conditions := &graylog.AlertConditionsBody{}
	u, err := client.Endpoints().StreamAlertConditions(streamID)
	if err != nil {
		return nil, 0, nil, errors.Wrap(err, errMsg)
	}
	ei, err = client.callGet(
		ctx, u.String(), nil, conditions)
	return conditions.AlertConditions, conditions.Total, ei, err
}

// GetStreamAlertCondition gets an alert condition.
func (client *Client) GetStreamAlertCondition(
	ctx context.Context, streamID, id string,
) (graylog.AlertCondition, *ErrorInfo, error) {
	cond := graylog.AlertCondition{}
	if streamID == "" {
		return cond, nil, errors.New("stream id is empty")
	}
	if id == "" {
		return cond, nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().StreamAlertCondition(streamID, id)
	if err != nil {
		return cond, nil, err
	}
	ei, err := client.callGet(ctx, u.String(), nil, &cond)
	return cond, ei, err
}

// CreateStreamAlertCondition creates an alert condition.
func (client *Client) CreateStreamAlertCondition(
	ctx context.Context, streamID string, cond *graylog.AlertCondition,
) (*ErrorInfo, error) {
	errMsg := "failed to create an alert condition"
	if cond == nil {
		return nil, fmt.Errorf("%s: alert condition is nil", errMsg)
	}
	u, err := client.Endpoints().StreamAlertConditions(streamID)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	ret := map[string]string{}
	ei, err := client.callPost(ctx, u.String(), cond, &ret)
	if err != nil {
		return ei, errors.Wrap(err, errMsg)
	}
	if id, ok := ret["alert_condition_id"]; ok {
		cond.ID = id
		return ei, nil
	}
	return ei, fmt.Errorf(`%s: response doesn't have the field "alert_condition_id"`, errMsg)
}

// UpdateStreamAlertCondition modifies an alert condition.
func (client *Client) UpdateStreamAlertCondition(
	ctx context.Context, streamID string, cond *graylog.AlertCondition,
) (*ErrorInfo, error) {
	errMsg := "failed to update an alert condition"
	if streamID == "" {
		return nil, fmt.Errorf("%s: stream id is empty", errMsg)
	}
	if cond == nil {
		return nil, fmt.Errorf("%s: alert condition is nil", errMsg)
	}
	condID := cond.ID
	if condID == "" {
		return nil, fmt.Errorf("%s: alert condition id is empty", errMsg)
	}
	u, err := client.Endpoints().StreamAlertCondition(streamID, condID)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	cond.ID = ""
	ei, err := client.callPut(ctx, u.String(), cond, nil)
	cond.ID = condID
	if err != nil {
		return ei, errors.Wrap(err, errMsg)
	}
	return ei, nil
}

// DeleteStreamAlertCondition deletes an alert condition.
func (client *Client) DeleteStreamAlertCondition(
	ctx context.Context, streamID, id string,
) (*ErrorInfo, error) {
	errMsg := "failed to delete an alert condition"
	if streamID == "" {
		return nil, fmt.Errorf("%s: stream id is empty", errMsg)
	}
	if id == "" {
		return nil, fmt.Errorf("%s: alert condition id is empty", errMsg)
	}
	u, err := client.Endpoints().StreamAlertCondition(streamID, id)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
