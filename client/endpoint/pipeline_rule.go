package endpoint

// PipelineRules returns a Pipeline Rules API's endpoint url.
func (ep *Endpoints) PipelineRules() string {
	return ep.pipelineRules.String()
}
