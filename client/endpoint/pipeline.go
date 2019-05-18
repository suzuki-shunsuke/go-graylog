package endpoint

import (
	"net/url"
)

// Pipelines returns a Pipeline API's endpoint url.
func (ep *Endpoints) Pipelines() string {
	return ep.pipelines.String()
}

// Pipeline returns a Pipeline API's endpoint url.
func (ep *Endpoints) Pipeline(id string) (*url.URL, error) {
	return urlJoin(ep.pipelines, id)
}
