package handler

import (
	"net/http"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

// HandleGetAlertConditions is the handler of GET Alert Conditions API.
func HandleGetAlertConditions(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// GET /alerts/conditions Get a list of all alert conditions
	arr, total, sc, err := lgc.GetAlertConditions()
	if err != nil {
		return arr, sc, err
	}
	return &graylog.AlertConditionsBody{
		AlertConditions: arr, Total: total}, sc, nil
}
