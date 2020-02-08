package endpoint

// EventDefinitions returns a EventDefinition API's endpoint url.
func (ep *Endpoints) EventDefinitions() string {
	return ep.eventDefinitions
}

// EventDefinition returns a EventDefinition API's endpoint url.
func (ep *Endpoints) EventDefinition(id string) string {
	return ep.eventDefinitions + "/" + id
}
