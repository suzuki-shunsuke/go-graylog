package endpoint

import (
	"net/url"
)

// Dashboards returns an Dashboard API's endpoint url.
func (ep *Endpoints) Dashboards() string {
	return ep.dashboards.String()
}

// Dashboard returns an Dashboard API's endpoint url.
func (ep *Endpoints) Dashboard(id string) (*url.URL, error) {
	return urlJoin(ep.dashboards, id)
}
