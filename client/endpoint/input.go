package endpoint

import (
	"net/url"
)

// Inputs returns an Input API's endpoint url.
func (ep *Endpoints) Inputs() string {
	return ep.inputs.String()
}

// Input returns an Input API's endpoint url.
func (ep *Endpoints) Input(id string) (*url.URL, error) {
	return urlJoin(ep.inputs, id)
}
