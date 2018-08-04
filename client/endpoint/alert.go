package endpoint

import (
	"net/url"
)

// Alert returns an Alert API's endpoint url.
func (ep *Endpoints) Alert(id string) (*url.URL, error) {
	return urlJoin(ep.alerts, id)
}

// Alerts returns Alerts API's endpoint url.
func (ep *Endpoints) Alerts() string {
	return ep.alerts.String()
}
