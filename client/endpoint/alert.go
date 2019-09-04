package endpoint

// Alert returns an Alert API's endpoint url.
func (ep *Endpoints) Alert(id string) string {
	return ep.alerts + "/" + id
}

// Alerts returns Alerts API's endpoint url.
func (ep *Endpoints) Alerts() string {
	return ep.alerts
}
