package logic

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

func (ms *Server) GetIndexSetStats(id string) (*graylog.IndexSetStats, error) {
	ok, err := ms.HasIndexSet(id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	return ms.store.GetIndexSetStats(id)
}

// GetTotalIndexSetsStats returns all index set's statistics.
func (ms *Server) GetTotalIndexSetsStats() (*graylog.IndexSetStats, error) {
	return ms.store.GetTotalIndexSetsStats()
}

func (ms *Server) SetIndexSetStats(id string, stats *graylog.IndexSetStats) (int, error) {
	ok, err := ms.HasIndexSet(id)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no index set with id <%s> is found", id)
	}

	if err := ms.store.SetIndexSetStats(id, stats); err != nil {
		return 500, err
	}
	return 200, nil
}
