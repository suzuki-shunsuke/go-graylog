package endpoint

// AlarmCallbacks returns AlarmCallbacks API's endpoint url.
func (ep *Endpoints) AlarmCallbacks() string {
	return ep.alarmCallbacks.String()
}
