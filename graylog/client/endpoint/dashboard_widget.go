package endpoint

// DashboardWidgets returns a Dashboard Widget API's endpoint url.
func (ep *Endpoints) DashboardWidgets(dashboardID string) string {
	return ep.dashboards + "/" + dashboardID + "/widgets"
}

// DashboardWidget returns a Dashboard Widget API's endpoint url.
func (ep *Endpoints) DashboardWidget(dashboardID, widgetID string) string {
	return ep.dashboards + "/" + dashboardID + "/widgets/" + widgetID
}

// DashboardWidgetCacheTime returns a Dashboard Widget cache time API's endpoint url.
func (ep *Endpoints) DashboardWidgetCacheTime(dashboardID, widgetID string) string {
	return ep.dashboards + "/" + dashboardID + "/widgets/" + widgetID + "/cachetime"
}

// DashboardWidgetDescription returns a Dashboard Widget description API's endpoint url.
func (ep *Endpoints) DashboardWidgetDescription(dashboardID, widgetID string) string {
	return ep.dashboards + "/" + dashboardID + "/widgets/" + widgetID + "/description"
}
