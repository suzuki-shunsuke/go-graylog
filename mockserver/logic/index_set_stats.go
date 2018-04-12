package logic

import (
	"github.com/suzuki-shunsuke/go-graylog"
)

// GetIndexSetStats returns an index set stats.
func (ms *Logic) GetIndexSetStats(id string) (*graylog.IndexSetStats, int, error) {
	ok, err := ms.HasIndexSet(id)
	if err != nil {
		return nil, 500, err
	}
	if !ok {
		return nil, 404, nil
	}
	stats, err := ms.store.GetIndexSetStats(id)
	if err != nil {
		return nil, 500, err
	}
	if stats == nil {
		return &graylog.IndexSetStats{}, 200, err
	}
	return stats, 200, nil
}

// GetTotalIndexSetStats returns all index set's statistics.
func (ms *Logic) GetTotalIndexSetStats() (*graylog.IndexSetStats, int, error) {
	stats, err := ms.store.GetTotalIndexSetStats()
	if err != nil {
		return stats, 500, err
	}
	return stats, 200, nil
}

func (ms *Logic) GetIndexSetStatsMap() (map[string]graylog.IndexSetStats, int, error) {
	m, err := ms.store.GetIndexSetStatsMap()
	if err != nil {
		return m, 500, err
	}
	return m, 200, err
}

// SetIndexSetStats sets an index set stats to a index set.
// func (ms *Logic) SetIndexSetStats(id string, stats *graylog.IndexSetStats) (int, error) {
// 	ok, err := ms.HasIndexSet(id)
// 	if err != nil {
// 		return 500, err
// 	}
// 	if !ok {
// 		return 404, fmt.Errorf("no index set with id <%s> is found", id)
// 	}
//
// 	if err := ms.store.SetIndexSetStats(id, stats); err != nil {
// 		return 500, err
// 	}
// 	return 200, nil
// }
