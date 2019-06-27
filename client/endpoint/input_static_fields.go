package endpoint

import (
	"net/url"
	"path"
)

// InputStaticFields returns the StaticFields API's endpoint url.
func (ep *Endpoints) InputStaticFields(inputID string) (*url.URL, error) {
	// /system/inputs/{inputId}/staticfields
	return urlJoin(ep.inputs, path.Join(inputID, "staticfields"))
}

// InputStaticField returns the StaticFields API's endpoint url.
func (ep *Endpoints) InputStaticField(inputID, key string) (*url.URL, error) {
	// /system/inputs/{inputId}/staticfields/{key}
	return urlJoin(ep.inputs, path.Join(inputID, "staticfields", key))
}
