package endpoint

// EventNotifications returns a EventNotification API's endpoint url.
func (ep *Endpoints) EventNotifications() string {
	return ep.eventNotifications
}

// EventNotification returns a EventNotification API's endpoint url.
func (ep *Endpoints) EventNotification(id string) string {
	return ep.eventNotifications + "/" + id
}
