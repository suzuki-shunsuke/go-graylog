package endpoint

import (
	"net/url"
	"path"
)

// IndexSet returns an IndexSet API's endpoint url.
func (ep *Endpoints) IndexSet(id string) (*url.URL, error) {
	return urlJoin(ep.indexSets, id)
}

// IndexSets returns an IndexSet API's endpoint url.
func (ep *Endpoints) IndexSets() string {
	return ep.indexSets.String()
}

// SetDefaultIndexSet returns SetDefaultIndexSet API's endpoint url.
func (ep *Endpoints) SetDefaultIndexSet(id string) (*url.URL, error) {
	return urlJoin(ep.indexSets, path.Join(id, "default"))
}

// IndexSetsStats returns all IndexSets stats API's endpoint url.
func (ep *Endpoints) IndexSetsStats() string {
	return ep.indexSetStats.String()
}

// IndexSetStats returns an IndexSet stats API's endpoint url.
func (ep *Endpoints) IndexSetStats(id string) (*url.URL, error) {
	return urlJoin(ep.indexSets, path.Join(id, "stats"))
}
