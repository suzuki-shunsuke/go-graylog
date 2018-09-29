package endpoint

import (
	"net/url"
	"path"
)

// DashboardWidgets returns an Dashboard Widget API's endpoint url.
func (ep *Endpoints) DashboardWidgets(dashboardID string) (*url.URL, error) {
	return urlJoin(ep.dashboards, path.Join(dashboardID, "widgets"))
}

// DashboardWidget returns an Dashboard Widget API's endpoint url.
func (ep *Endpoints) DashboardWidget(dashboardID, widgetID string) (*url.URL, error) {
	return urlJoin(ep.dashboards, path.Join(dashboardID, "widgets", widgetID))
}
