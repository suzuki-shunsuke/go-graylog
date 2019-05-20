package endpoint

import (
	"net/url"
	"path"
)

// Extractors returns Stream Rules API's endpoint url.
func (ep *Endpoints) Extractors(inputID string) (*url.URL, error) {
	// /system/inputs/{inputID}/extractors
	return urlJoin(ep.inputs, path.Join(inputID, "extractors"))
}

// Extractor returns a Stream Rule API's endpoint url.
func (ep *Endpoints) Extractor(inputID, extractorID string) (*url.URL, error) {
	// /system/inputs/{inputID}/extractors/{extractorID}
	return urlJoin(ep.inputs, path.Join(inputID, "extractors", extractorID))
}
