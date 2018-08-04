package handler

import (
	"net/http"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

// HandleGetIndexSetStats is the handler of Get Index Set Statistics API.
func HandleGetIndexSetStats(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// GET /system/indices/index_sets/{id}/stats Get index set statistics
	// TODO authorization
	id := ps.PathParam("indexSetID")
	return lgc.GetIndexSetStats(id)
}

// HandleGetTotalIndexSetStats is the handler of Get Index Set Statistics of all Index Sets API.
func HandleGetTotalIndexSetStats(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// GET /system/indices/index_sets/stats Get stats of all index sets
	// TODO authorization
	return lgc.GetTotalIndexSetStats()
}
