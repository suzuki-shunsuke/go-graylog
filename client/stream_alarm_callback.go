package client

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetStreamAlarmCallbacks gets all alarm callbacks of this stream.
func (client *Client) GetStreamAlarmCallbacks(
	ctx context.Context, streamID string,
) (acs []graylog.AlarmCallback, total int, ei *ErrorInfo, err error) {
	callbacks := &graylog.AlarmCallbacksBody{}
	ei, err = client.callGet(ctx, client.Endpoints().StreamAlarmCallbacks(streamID), nil, callbacks)
	return callbacks.AlarmCallbacks, callbacks.Total, ei, err
}

// GetStreamAlarmCallback gets an alarm callback.
func (client *Client) GetStreamAlarmCallback(
	ctx context.Context, streamID, id string,
) (graylog.AlarmCallback, *ErrorInfo, error) {
	ac := graylog.AlarmCallback{}
	if streamID == "" {
		return ac, nil, errors.New("stream id is empty")
	}
	if id == "" {
		return ac, nil, errors.New("id is empty")
	}
	ei, err := client.callGet(ctx, client.Endpoints().StreamAlarmCallback(streamID, id), nil, &ac)
	return ac, ei, err
}

// CreateStreamAlarmCallback creates an alarm callback.
func (client *Client) CreateStreamAlarmCallback(
	ctx context.Context, ac *graylog.AlarmCallback,
) (*ErrorInfo, error) {
	errMsg := "failed to create an alarm callback"
	if ac == nil {
		return nil, fmt.Errorf("%s: alarm callback is nil", errMsg)
	}
	streamID := ac.StreamID
	if streamID == "" {
		return nil, fmt.Errorf("%s: stream id is empty", errMsg)
	}
	ret := map[string]string{}
	ac.StreamID = ""
	defer func() {
		ac.StreamID = streamID
	}()
	ei, err := client.callPost(ctx, client.Endpoints().StreamAlarmCallbacks(streamID), ac, &ret)
	if err != nil {
		return ei, errors.Wrap(err, errMsg)
	}
	if id, ok := ret["alarmcallback_id"]; ok {
		ac.ID = id
		return ei, nil
	}
	return ei, fmt.Errorf(`%s: response doesn't have the field "alarmcallback_id"`, errMsg)
}

// UpdateStreamAlarmCallback modifies an alarm callback.
func (client *Client) UpdateStreamAlarmCallback(
	ctx context.Context, ac *graylog.AlarmCallback,
) (*ErrorInfo, error) {
	errMsg := "failed to update an alarm callback"
	if ac == nil {
		return nil, fmt.Errorf("%s: alarm callback is nil", errMsg)
	}
	streamID := ac.StreamID
	if streamID == "" {
		return nil, fmt.Errorf("%s: stream id is empty", errMsg)
	}
	acID := ac.ID
	if acID == "" {
		return nil, fmt.Errorf("%s: alarm callback id is empty", errMsg)
	}
	ac.ID = ""
	ac.StreamID = ""
	defer func() {
		ac.ID = acID
		ac.StreamID = streamID
	}()
	ei, err := client.callPut(ctx, client.Endpoints().StreamAlarmCallback(streamID, acID), ac, nil)
	if err != nil {
		return ei, errors.Wrap(err, errMsg)
	}
	return ei, nil
}

// DeleteStreamAlarmCallback deletes an alarm callback.
func (client *Client) DeleteStreamAlarmCallback(
	ctx context.Context, streamID, id string,
) (*ErrorInfo, error) {
	errMsg := "failed to delete an alarm callback"
	if streamID == "" {
		return nil, fmt.Errorf("%s: stream id is empty", errMsg)
	}
	if id == "" {
		return nil, fmt.Errorf("%s: alarm callback id is empty", errMsg)
	}
	return client.callDelete(ctx, client.Endpoints().StreamAlarmCallback(streamID, id), nil, nil)
}
