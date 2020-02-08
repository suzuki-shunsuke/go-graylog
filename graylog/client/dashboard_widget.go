package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
)

// CreateDashboardWidget creates a new dashboard widget.
func (client *Client) CreateDashboardWidget(
	ctx context.Context, dashboardID string, widget graylog.Widget,
) (graylog.Widget, *ErrorInfo, error) {
	if dashboardID == "" {
		return widget, nil, errors.New("dashboard id is required")
	}

	ret := map[string]string{}
	ei, err := client.callPost(
		ctx, client.Endpoints().DashboardWidgets(dashboardID), &widget, &ret)
	if err != nil {
		return widget, ei, err
	}
	if id, ok := ret["widget_id"]; ok {
		widget.ID = id
		return widget, ei, nil
	}
	return widget, ei, errors.New(`response doesn't have the field "widget_id"`)
}

// UpdateDashboardWidget creates an existing dashboard widget.
func (client *Client) UpdateDashboardWidget(
	ctx context.Context, dashboardID string, widget graylog.Widget,
) (*ErrorInfo, error) {
	if dashboardID == "" {
		return nil, errors.New("dashboard id is required")
	}
	if widget.ID == "" {
		return nil, errors.New("dashboard widget id is required")
	}

	return client.callPut(
		ctx, client.Endpoints().DashboardWidget(dashboardID, widget.ID), map[string]interface{}{
			"description": widget.Description,
			"type":        widget.Type(),
			"config":      widget.Config,
		}, nil)
}

// DeleteDashboardWidget deletes a given dashboard widget.
func (client *Client) DeleteDashboardWidget(
	ctx context.Context, dashboardID, widgetID string,
) (*ErrorInfo, error) {
	if dashboardID == "" {
		return nil, errors.New("dashboard id is required")
	}
	if widgetID == "" {
		return nil, errors.New("widget id is required")
	}
	return client.callDelete(
		ctx, client.Endpoints().DashboardWidget(dashboardID, widgetID), nil, nil)
}

// GetDashboardWidget gets a dashboard widget.
func (client *Client) GetDashboardWidget(
	ctx context.Context, dashboardID, widgetID string,
) (graylog.Widget, *ErrorInfo, error) {
	widget := graylog.Widget{}
	if dashboardID == "" {
		return widget, nil, errors.New("dashboard id is required")
	}
	if widgetID == "" {
		return widget, nil, errors.New("widget id is required")
	}
	ei, err := client.callGet(
		ctx, client.Endpoints().DashboardWidget(dashboardID, widgetID), nil, &widget)
	return widget, ei, err
}

// UpdateDashboardWidgetCacheTime updates an existing dashboard widget cache time.
func (client *Client) UpdateDashboardWidgetCacheTime(
	ctx context.Context, dashboardID, widgetID string, cacheTime int,
) (*ErrorInfo, error) {
	if dashboardID == "" {
		return nil, errors.New("dashboard id is required")
	}
	if widgetID == "" {
		return nil, errors.New("dashboard widget id is required")
	}

	return client.callPut(
		ctx, client.Endpoints().DashboardWidgetCacheTime(dashboardID, widgetID), map[string]interface{}{
			"cache_time": cacheTime,
		}, nil)
}

// UpdateDashboardWidgetDescription updates an existing dashboard widget description.
func (client *Client) UpdateDashboardWidgetDescription(
	ctx context.Context, dashboardID, widgetID, description string,
) (*ErrorInfo, error) {
	if dashboardID == "" {
		return nil, errors.New("dashboard id is required")
	}
	if widgetID == "" {
		return nil, errors.New("dashboard widget id is required")
	}

	return client.callPut(
		ctx, client.Endpoints().DashboardWidgetDescription(dashboardID, widgetID), map[string]interface{}{
			"description": description,
		}, nil)
}
