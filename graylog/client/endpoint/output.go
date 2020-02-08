package endpoint

// Outputs returns a Output API's endpoint url.
func (ep *Endpoints) Outputs() string {
	return ep.outputs
}

// AvailableOutputs returns a Output API's endpoint url.
func (ep *Endpoints) AvailableOutputs() string {
	return ep.availableOutputs
}

// Output returns a Output API's endpoint url.
func (ep *Endpoints) Output(id string) string {
	return ep.outputs + "/" + id
}
