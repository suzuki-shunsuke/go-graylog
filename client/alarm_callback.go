package client

import (
	"context"

	"github.com/suzuki-shunsuke/go-graylog/v8"
)

// GetAlarmCallbacksContext returns all alarm callbacks.
func (client *Client) GetAlarmCallbacks(ctx context.Context) (
	[]graylog.AlarmCallback, int, *ErrorInfo, error,
) {
	body := &graylog.AlarmCallbacksBody{}
	ei, err := client.callGet(
		ctx, client.Endpoints().AlarmCallbacks(), nil, body)
	return body.AlarmCallbacks, body.Total, ei, err
}
