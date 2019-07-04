package client

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetAlert returns an alert.
func (client *Client) GetAlert(ctx context.Context, id string) (
	*graylog.Alert, *ErrorInfo, error,
) {
	// GET /streams/alerts/{alertId} Get an alert by ID
	if id == "" {
		return nil, nil, fmt.Errorf("id is empty")
	}
	u, err := client.Endpoints().Alert(id)
	if err != nil {
		return nil, nil, err
	}
	alert := &graylog.Alert{}
	ei, err := client.callGet(
		ctx, u.String(), nil, alert)
	return alert, ei, err
}

// GetAlerts returns all alerts.
func (client *Client) GetAlerts(ctx context.Context, skip, limit int) (
	[]graylog.Alert, int, *ErrorInfo, error,
) {
	body := &graylog.AlertsBody{}
	v := url.Values{
		"skip":  []string{strconv.Itoa(skip)},
		"limit": []string{strconv.Itoa(limit)},
	}
	u := fmt.Sprintf("%s?%s", client.Endpoints().Alerts(), v.Encode())
	ei, err := client.callGet(ctx, u, nil, body)
	return body.Alerts, body.Total, ei, err
}
