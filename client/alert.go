package client

import (
	"context"
	"errors"
	"net/url"
	"strconv"

	"github.com/suzuki-shunsuke/go-graylog/v9"
)

// GetAlert returns an alert.
func (client *Client) GetAlert(ctx context.Context, id string) (
	*graylog.Alert, *ErrorInfo, error,
) {
	// GET /streams/alerts/{alertId} Get an alert by ID
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	alert := &graylog.Alert{}
	ei, err := client.callGet(ctx, client.Endpoints().Alert(id), nil, alert)
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
	ei, err := client.callGet(
		ctx, client.Endpoints().Alerts()+"?"+v.Encode(), nil, body)
	return body.Alerts, body.Total, ei, err
}
