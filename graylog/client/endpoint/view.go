package endpoint

// Views returns a View API's endpoint url.
func (ep *Endpoints) Views() string {
	return ep.views
}

// View returns a View API's endpoint url.
func (ep *Endpoints) View(id string) string {
	return ep.views + "/" + id
}

// ViewDefault returns a View API's endpoint url.
func (ep *Endpoints) ViewDefault(id string) string {
	return ep.views + "/" + id + "/default"
}
