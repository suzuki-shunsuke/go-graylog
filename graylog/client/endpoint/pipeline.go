package endpoint

// Pipelines returns a Pipeline API's endpoint url.
func (ep *Endpoints) Pipelines() string {
	return ep.pipelines
}

// Pipeline returns a Pipeline API's endpoint url.
func (ep *Endpoints) Pipeline(id string) string {
	return ep.pipelines + "/" + id
}
