package endpoint

import (
	"net/url"
)

// PipelineRules returns a Pipeline Rules API's endpoint url.
func (ep *Endpoints) PipelineRules() string {
	return ep.pipelineRules.String()
}

// PipelineRule returns a Pipeline Rule API's endpoint url.
func (ep *Endpoints) PipelineRule(id string) (*url.URL, error) {
	return urlJoin(ep.pipelineRules, id)
}
