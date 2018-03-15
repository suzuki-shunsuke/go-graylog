package mockserver

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/suzuki-shunsuke/go-graylog"
)

// GET /system/indices/index_sets/{id}/stats Get index set statistics
func (ms *MockServer) handleGetIndexSetStats(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	ms.handleInit(w, r, false)
	id := ps.ByName("indexSetID")
	indexSetStats, err := ms.GetIndexSetStats(id)
	if err != nil {
		return 500, nil, err
	}
	if indexSetStats == nil {
		return 404, nil, fmt.Errorf("No indexSet found with id %s", id)
	}
	return 200, indexSetStats, nil
}

// GET /system/indices/index_sets/stats Get stats of all index sets
func (ms *MockServer) handleGetAllIndexSetsStats(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	ms.handleInit(w, r, false)
	s, err := ms.GetTotalIndexSetsStats()
	if err != nil {
		return 500, nil, err
	}
	return 200, s, nil
}

func (ms *MockServer) GetIndexSetStats(id string) (*graylog.IndexSetStats, error) {
	ok, err := ms.HasIndexSet(id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	return ms.store.GetIndexSetStats(id)
}

func (ms *MockServer) GetIndexSetsStats() ([]graylog.IndexSetStats, error) {
	return ms.store.GetIndexSetsStats()
}

// GetTotalIndexSetsStats returns all index set's statistics.
func (ms *MockServer) GetTotalIndexSetsStats() (*graylog.IndexSetStats, error) {
	return ms.store.GetTotalIndexSetsStats()
}

func (ms *MockServer) SetIndexSetStats(id string, stats *graylog.IndexSetStats) (int, error) {
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
