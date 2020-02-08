package endpoint

// Inputs returns an Input API's endpoint url.
func (ep *Endpoints) Inputs() string {
	return ep.inputs
}

// Input returns an Input API's endpoint url.
func (ep *Endpoints) Input(id string) string {
	return ep.inputs + "/" + id
}
