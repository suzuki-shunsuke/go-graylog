package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

// HandleGetIndexSetStats
func HandleGetIndexSetStats(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// GET /system/indices/index_sets/{id}/stats Get index set statistics
	// TODO authorization
	id := ps.ByName("indexSetID")
	return ms.GetIndexSetStats(id)
}

// HandleGetTotalIndexSetStats
func HandleGetTotalIndexSetStats(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// GET /system/indices/index_sets/stats Get stats of all index sets
	// TODO authorization
	return ms.GetTotalIndexSetStats()
}
