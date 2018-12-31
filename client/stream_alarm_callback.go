package client

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetStreamAlarmCallbacks gets all alarm callbacks of this stream.
func (client *Client) GetStreamAlarmCallbacks(streamID string) (
	acs []graylog.AlarmCallback, total int, ei *ErrorInfo, err error,
) {
	return client.GetStreamAlarmCallbacksContext(context.Background(), streamID)
}

// GetStreamAlarmCallbacksContext gets all alarm callbacks of this stream with a context.
func (client *Client) GetStreamAlarmCallbacksContext(
	ctx context.Context, streamID string,
) (acs []graylog.AlarmCallback, total int, ei *ErrorInfo, err error) {
	errMsg := "failed to get stream's alarm callbacks"
	callbacks := &graylog.AlarmCallbacksBody{}
	u, err := client.Endpoints().StreamAlarmCallbacks(streamID)
	if err != nil {
		return nil, 0, nil, errors.Wrap(err, errMsg)
	}
	ei, err = client.callGet(ctx, u.String(), nil, callbacks)
	return callbacks.AlarmCallbacks, callbacks.Total, ei, err
}

// GetStreamAlarmCallback gets an alarm callback.
func (client *Client) GetStreamAlarmCallback(
	streamID, id string,
) (graylog.AlarmCallback, *ErrorInfo, error) {
	return client.GetStreamAlarmCallbackContext(context.Background(), streamID, id)
}

// GetStreamAlarmCallbackContext gets an alarm callback with a context.
func (client *Client) GetStreamAlarmCallbackContext(
	ctx context.Context, streamID, id string,
) (graylog.AlarmCallback, *ErrorInfo, error) {
	ac := graylog.AlarmCallback{}
	if streamID == "" {
		return ac, nil, errors.New("stream id is empty")
	}
	if id == "" {
		return ac, nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().StreamAlarmCallback(streamID, id)
	if err != nil {
		return ac, nil, err
	}
	ei, err := client.callGet(ctx, u.String(), nil, &ac)
	return ac, ei, err
}

// CreateStreamAlarmCallback creates an alarm callback.
func (client *Client) CreateStreamAlarmCallback(ac *graylog.AlarmCallback) (*ErrorInfo, error) {
	return client.CreateStreamAlarmCallbackContext(context.Background(), ac)
}

// CreateStreamAlarmCallbackContext creates an alarm callback with a context.
func (client *Client) CreateStreamAlarmCallbackContext(
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
	u, err := client.Endpoints().StreamAlarmCallbacks(streamID)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	ret := map[string]string{}
	ac.StreamID = ""
	defer func() {
		ac.StreamID = streamID
	}()
	ei, err := client.callPost(ctx, u.String(), ac, &ret)
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
func (client *Client) UpdateStreamAlarmCallback(ac *graylog.AlarmCallback) (*ErrorInfo, error) {
	return client.UpdateStreamAlarmCallbackContext(context.Background(), ac)
}

// UpdateStreamAlarmCallbackContext modifies an alarm callback with a context.
func (client *Client) UpdateStreamAlarmCallbackContext(
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
	u, err := client.Endpoints().StreamAlarmCallback(streamID, acID)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	ac.ID = ""
	ac.StreamID = ""
	defer func() {
		ac.ID = acID
		ac.StreamID = streamID
	}()
	ei, err := client.callPut(ctx, u.String(), ac, nil)
	if err != nil {
		return ei, errors.Wrap(err, errMsg)
	}
	return ei, nil
}

// DeleteStreamAlarmCallback deletes an alarm callback.
func (client *Client) DeleteStreamAlarmCallback(streamID, id string) (*ErrorInfo, error) {
	return client.DeleteStreamAlarmCallbackContext(context.Background(), streamID, id)
}

// DeleteStreamAlarmCallbackContext deletes an alarm callback with a context.
func (client *Client) DeleteStreamAlarmCallbackContext(
	ctx context.Context, streamID, id string,
) (*ErrorInfo, error) {
	errMsg := "failed to delete an alarm callback"
	if streamID == "" {
		return nil, fmt.Errorf("%s: stream id is empty", errMsg)
	}
	if id == "" {
		return nil, fmt.Errorf("%s: alarm callback id is empty", errMsg)
	}
	u, err := client.Endpoints().StreamAlarmCallback(streamID, id)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
