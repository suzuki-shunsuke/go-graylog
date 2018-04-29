package logic

import (
	"github.com/suzuki-shunsuke/go-graylog"
)

// GetIndexSetStats returns an index set stats.
func (lgc *Logic) GetIndexSetStats(id string) (*graylog.IndexSetStats, int, error) {
	ok, err := lgc.HasIndexSet(id)
	if err != nil {
		return nil, 500, err
	}
	if !ok {
		return nil, 404, nil
	}
	stats, err := lgc.store.GetIndexSetStats(id)
	if err != nil {
		return nil, 500, err
	}
	if stats == nil {
		return &graylog.IndexSetStats{}, 200, err
	}
	return stats, 200, nil
}

// GetTotalIndexSetStats returns all index set's statistics.
func (lgc *Logic) GetTotalIndexSetStats() (*graylog.IndexSetStats, int, error) {
	stats, err := lgc.store.GetTotalIndexSetStats()
	if err != nil {
		return stats, 500, err
	}
	return stats, 200, nil
}

// GetIndexSetStatsMap returns a each Index Set's statistics.
func (lgc *Logic) GetIndexSetStatsMap() (map[string]graylog.IndexSetStats, int, error) {
	m, err := lgc.store.GetIndexSetStatsMap()
	if err != nil {
		return m, 500, err
	}
	return m, 200, err
}

// SetIndexSetStats sets an index set stats to a index set.
// func (lgc *Logic) SetIndexSetStats(id string, stats *graylog.IndexSetStats) (int, error) {
// 	ok, err := lgc.HasIndexSet(id)
// 	if err != nil {
// 		return 500, err
// 	}
// 	if !ok {
// 		return 404, fmt.Errorf("no index set with id <%s> is found", id)
// 	}
//
// 	if err := lgc.store.SetIndexSetStats(id, stats); err != nil {
// 		return 500, err
// 	}
// 	return 200, nil
// }
