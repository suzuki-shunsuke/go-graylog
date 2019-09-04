package endpoint

// IndexSet returns an IndexSet API's endpoint url.
func (ep *Endpoints) IndexSet(id string) string {
	return ep.indexSets + "/" + id
}

// IndexSets returns an IndexSet API's endpoint url.
func (ep *Endpoints) IndexSets() string {
	return ep.indexSets
}

// SetDefaultIndexSet returns SetDefaultIndexSet API's endpoint url.
func (ep *Endpoints) SetDefaultIndexSet(id string) string {
	return ep.indexSets + "/" + id + "/default"
}

// IndexSetsStats returns all IndexSets stats API's endpoint url.
func (ep *Endpoints) IndexSetsStats() string {
	return ep.indexSetStats
}

// IndexSetStats returns an IndexSet stats API's endpoint url.
func (ep *Endpoints) IndexSetStats(id string) string {
	return ep.indexSets + "/" + id + "/stats"
}
