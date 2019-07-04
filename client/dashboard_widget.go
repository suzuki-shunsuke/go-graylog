package client

import (
	"context"
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

// CreateDashboardWidget creates a new dashboard widget.
func (client *Client) CreateDashboardWidget(
	ctx context.Context, dashboardID string, widget graylog.Widget,
) (graylog.Widget, *ErrorInfo, error) {
	if dashboardID == "" {
		return widget, nil, fmt.Errorf("dashboard id is required")
	}

	ret := map[string]string{}
	u, err := client.Endpoints().DashboardWidgets(dashboardID)
	if err != nil {
		return widget, nil, err
	}
	ei, err := client.callPost(ctx, u.String(), &widget, &ret)
	if err != nil {
		return widget, ei, err
	}
	if id, ok := ret["widget_id"]; ok {
		widget.ID = id
		return widget, ei, nil
	}
	return widget, ei, fmt.Errorf(`response doesn't have the field "widget_id"`)
}

// UpdateDashboardWidget creates an existing dashboard widget.
func (client *Client) UpdateDashboardWidget(
	ctx context.Context, dashboardID string, widget graylog.Widget,
) (*ErrorInfo, error) {
	if dashboardID == "" {
		return nil, fmt.Errorf("dashboard id is required")
	}
	if widget.ID == "" {
		return nil, fmt.Errorf("dashboard widget id is required")
	}

	u, err := client.Endpoints().DashboardWidget(dashboardID, widget.ID)
	if err != nil {
		return nil, err
	}
	return client.callPut(ctx, u.String(), map[string]interface{}{
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
		return nil, fmt.Errorf("dashboard id is required")
	}
	if widgetID == "" {
		return nil, fmt.Errorf("widget id is required")
	}
	u, err := client.Endpoints().DashboardWidget(dashboardID, widgetID)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}

// GetDashboardWidget gets a dashboard widget.
func (client *Client) GetDashboardWidget(
	ctx context.Context, dashboardID, widgetID string,
) (graylog.Widget, *ErrorInfo, error) {
	widget := graylog.Widget{}
	if dashboardID == "" {
		return widget, nil, fmt.Errorf("dashboard id is required")
	}
	if widgetID == "" {
		return widget, nil, fmt.Errorf("widget id is required")
	}
	u, err := client.Endpoints().DashboardWidget(dashboardID, widgetID)
	if err != nil {
		return widget, nil, err
	}
	ei, err := client.callGet(ctx, u.String(), nil, &widget)
	return widget, ei, err
}

// UpdateDashboardWidgetCacheTime updates an existing dashboard widget cache time.
func (client *Client) UpdateDashboardWidgetCacheTime(
	ctx context.Context, dashboardID, widgetID string, cacheTime int,
) (*ErrorInfo, error) {
	if dashboardID == "" {
		return nil, fmt.Errorf("dashboard id is required")
	}
	if widgetID == "" {
		return nil, fmt.Errorf("dashboard widget id is required")
	}

	u, err := client.Endpoints().DashboardWidgetCacheTime(dashboardID, widgetID)
	if err != nil {
		return nil, err
	}
	return client.callPut(ctx, u.String(), map[string]interface{}{
		"cache_time": cacheTime,
	}, nil)
}

// UpdateDashboardWidgetDescription updates an existing dashboard widget description.
func (client *Client) UpdateDashboardWidgetDescription(
	ctx context.Context, dashboardID, widgetID, description string,
) (*ErrorInfo, error) {
	if dashboardID == "" {
		return nil, fmt.Errorf("dashboard id is required")
	}
	if widgetID == "" {
		return nil, fmt.Errorf("dashboard widget id is required")
	}

	u, err := client.Endpoints().DashboardWidgetDescription(dashboardID, widgetID)
	if err != nil {
		return nil, err
	}
	return client.callPut(ctx, u.String(), map[string]interface{}{
		"description": description,
	}, nil)
}
