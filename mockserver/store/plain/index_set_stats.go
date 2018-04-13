package plain

import (
	"github.com/suzuki-shunsuke/go-graylog"
)

// GetIndexSetStats returns an index set stats.
func (store *InMemoryStore) GetIndexSetStats(id string) (*graylog.IndexSetStats, error) {
	// TODO
	ok, err := store.HasIndexSet(id)
	if err != nil {
		return nil, err
	}
	if ok {
		return &graylog.IndexSetStats{}, nil
	}
	return nil, nil
}

// GetIndexSetStatsMap returns all of index set stats.
func (store *InMemoryStore) GetIndexSetStatsMap() (map[string]graylog.IndexSetStats, error) {
	// TODO
	m := map[string]graylog.IndexSetStats{}
	for _, is := range store.indexSets {
		// TODO
		m[is.ID] = graylog.IndexSetStats{}
	}
	return m, nil
}

// GetTotalIndexSetStats returns all index set's statistics.
func (store *InMemoryStore) GetTotalIndexSetStats() (*graylog.IndexSetStats, error) {
	// TODO
	indexSetStats := &graylog.IndexSetStats{}
	return indexSetStats, nil
}
