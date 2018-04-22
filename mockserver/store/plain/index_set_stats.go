package plain

import (
	"github.com/suzuki-shunsuke/go-graylog"
)

// GetIndexSetStats returns an index set stats.
func (store *Store) GetIndexSetStats(id string) (*graylog.IndexSetStats, error) {
	ok, err := store.HasIndexSet(id)
	if err != nil {
		return nil, err
	}
	if ok {
		// TODO returns correct index set stats
		return &graylog.IndexSetStats{}, nil
	}
	return nil, nil
}

// GetIndexSetStatsMap returns all of index set stats.
func (store *Store) GetIndexSetStatsMap() (map[string]graylog.IndexSetStats, error) {
	m := map[string]graylog.IndexSetStats{}
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	for _, is := range store.indexSets {
		// TODO returns correct index set stats
		m[is.ID] = graylog.IndexSetStats{}
	}
	return m, nil
}

// GetTotalIndexSetStats returns all index set's statistics.
func (store *Store) GetTotalIndexSetStats() (*graylog.IndexSetStats, error) {
	// TODO returns correct index set stats
	indexSetStats := &graylog.IndexSetStats{}
	return indexSetStats, nil
}
