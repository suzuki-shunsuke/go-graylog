package endpoint

// Outputs returns a Output API's endpoint url.
func (ep *Endpoints) StreamOutputs(streamID string) string {
	return ep.streams + "/" + streamID + "/outputs"
}

// Output returns a Output API's endpoint url.
func (ep *Endpoints) StreamOutput(streamID, outputID string) string {
	return ep.streams + "/" + streamID + "/outputs/" + outputID
}
