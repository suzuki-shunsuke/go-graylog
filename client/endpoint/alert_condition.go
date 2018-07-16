package endpoint

// AlertConditions returns a Alert Condition API's endpoint url.
func (ep *Endpoints) AlertConditions() string {
	return ep.alertConditions.String()
}
