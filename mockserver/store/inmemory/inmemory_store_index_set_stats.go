package inmemory

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetIndexSetStats returns an index set stats.
func (store *InMemoryStore) GetIndexSetStats(id string) (*graylog.IndexSetStats, error) {
	s, ok := store.indexSetStats[id]
	if ok {
		return &s, nil
	}
	return nil, nil
}

// GetTotalIndexSetsStats returns all index set's statistics.
func (store *InMemoryStore) GetTotalIndexSetsStats() (*graylog.IndexSetStats, error) {
	indexSetStats := &graylog.IndexSetStats{}
	for _, stats := range store.indexSetStats {
		indexSetStats.Indices += stats.Indices
		indexSetStats.Documents += stats.Documents
		indexSetStats.Size += stats.Size
	}
	return indexSetStats, nil
}

func (store *InMemoryStore) SetIndexSetStats(id string, stats *graylog.IndexSetStats) error {
	ok, err := store.HasIndexSet(id)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("no index set with id <%s> is found", id)
	}
	store.indexSetStats[id] = *stats
	return nil
}
