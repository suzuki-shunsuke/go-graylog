package endpoint

import (
	"net/url"
	"path"
)

// DashboardWidgets returns a Dashboard Widget API's endpoint url.
func (ep *Endpoints) DashboardWidgets(dashboardID string) (*url.URL, error) {
	return urlJoin(ep.dashboards, path.Join(dashboardID, "widgets"))
}

// DashboardWidget returns a Dashboard Widget API's endpoint url.
func (ep *Endpoints) DashboardWidget(dashboardID, widgetID string) (*url.URL, error) {
	return urlJoin(ep.dashboards, path.Join(dashboardID, "widgets", widgetID))
}

// DashboardWidgetCacheTime returns a Dashboard Widget cache time API's endpoint url.
func (ep *Endpoints) DashboardWidgetCacheTime(dashboardID, widgetID string) (*url.URL, error) {
	return urlJoin(ep.dashboards, path.Join(dashboardID, "widgets", widgetID, "cachetime"))
}

// DashboardWidgetDescription returns a Dashboard Widget description API's endpoint url.
func (ep *Endpoints) DashboardWidgetDescription(dashboardID, widgetID string) (*url.URL, error) {
	return urlJoin(ep.dashboards, path.Join(dashboardID, "widgets", widgetID, "description"))
}
