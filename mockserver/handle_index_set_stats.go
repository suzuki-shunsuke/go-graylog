package mockserver

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GET /system/indices/index_sets/{id}/stats Get index set statistics
func (ms *Server) handleGetIndexSetStats(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	id := ps.ByName("indexSetID")
	indexSetStats, err := ms.GetIndexSetStats(id)
	if err != nil {
		return 500, nil, err
	}
	if indexSetStats == nil {
		return 404, nil, fmt.Errorf("no indexSet found with id %s", id)
	}
	return 200, indexSetStats, nil
}

// GET /system/indices/index_sets/stats Get stats of all index sets
func (ms *Server) handleGetAllIndexSetsStats(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	s, err := ms.GetTotalIndexSetsStats()
	if err != nil {
		return 500, nil, err
	}
	return 200, s, nil
}
