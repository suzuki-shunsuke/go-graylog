package client

import (
	"context"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetAlertConditions returns all alert conditions.
func (client *Client) GetAlertConditions() ([]graylog.AlertCondition, int, *ErrorInfo, error) {
	return client.GetAlertConditionsContext(context.Background())
}

// GetAlertConditionsContext returns all alert conditions with a context.
func (client *Client) GetAlertConditionsContext(ctx context.Context) (
	[]graylog.AlertCondition, int, *ErrorInfo, error,
) {
	conditions := &graylog.AlertConditionsBody{}
	ei, err := client.callGet(
		ctx, client.Endpoints().AlertConditions(), nil, conditions)
	return conditions.AlertConditions, conditions.Total, ei, err
}
