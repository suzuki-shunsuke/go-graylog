package client

import (
	"context"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetAlarmCallbacks returns all alarm callbacks.
func (client *Client) GetAlarmCallbacks() (
	[]graylog.AlarmCallback, int, *ErrorInfo, error,
) {
	return client.GetAlarmCallbacksContext(context.Background())
}

// GetAlarmCallbacksContext returns all alarm callbacks with a context.
func (client *Client) GetAlarmCallbacksContext(ctx context.Context) (
	[]graylog.AlarmCallback, int, *ErrorInfo, error,
) {
	body := &graylog.AlarmCallbacksBody{}
	ei, err := client.callGet(
		ctx, client.Endpoints().AlarmCallbacks(), nil, body)
	return body.AlarmCallbacks, body.Total, ei, err
}
