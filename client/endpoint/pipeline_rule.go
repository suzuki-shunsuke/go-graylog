package endpoint

// PipelineRules returns a Pipeline Rules API's endpoint url.
func (ep *Endpoints) PipelineRules() string {
	return ep.pipelineRules
}

// PipelineRule returns a Pipeline Rule API's endpoint url.
func (ep *Endpoints) PipelineRule(id string) string {
	return ep.pipelineRules + "/" + id
}
