package endpoint

import (
	"fmt"
	"net/url"
)

// Extractors returns an Extractor API's endpoint url.
func (ep *Endpoints) Extractors(inputID string) (*url.URL, error) {
	// /system/inputs/{inputID}/extractors
	return urlJoin(ep.extractors, fmt.Sprintf("%s/extractors", inputID))
}

// Extractor returns an Extractor API's endpoint url.
func (ep *Endpoints) Extractor(inputID string, extractorID string) (*url.URL, error) {
	// /system/inputs/{inputID}/extractors/{extractorinputID}
	return urlJoin(ep.extractors, fmt.Sprintf("%s/extractors/%s", inputID, extractorID))
}
