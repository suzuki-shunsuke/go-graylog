package handler

import (
	"net/http"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

// HandleGetAlarmCallbacks is the handler of GET AlarmCallbacks API.
func HandleGetAlarmCallbacks(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// GET /alerts/callbacks Get a list of all alarm callbacks
	arr, total, sc, err := lgc.GetAlarmCallbacks()
	if err != nil {
		return arr, sc, err
	}
	return &graylog.AlarmCallbacksBody{AlarmCallbacks: arr, Total: total}, sc, nil
}
