package endpoint

// Extractors returns Stream Rules API's endpoint url.
func (ep *Endpoints) Extractors(inputID string) string {
	// /system/inputs/{inputID}/extractors
	return ep.inputs + "/" + inputID + "/extractors"
}

// Extractor returns a Stream Rule API's endpoint url.
func (ep *Endpoints) Extractor(inputID, extractorID string) string {
	// /system/inputs/{inputID}/extractors/{extractorID}
	return ep.inputs + "/" + inputID + "/extractors/" + extractorID
}
