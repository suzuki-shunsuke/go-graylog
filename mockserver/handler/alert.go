package handler

import (
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

// HandleGetAlert is the handler of Get an Alert API.
func HandleGetAlert(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// GET /streams/alerts/{id} Get an alarm
	id := ps.PathParam("alertID")
	// TODO authorization
	// if sc, err := lgc.Authorize(user, "inputs:read", id); err != nil {
	// 	return nil, sc, err
	// }
	return lgc.GetAlert(id)
}

// HandleGetAlerts is the handler of GET Alerts API.
func HandleGetAlerts(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// GET /streams/alerts Get the most recent alarms of all streams
	since := 0
	limit := 0
	query := r.URL.Query()
	s, ok := query["since"]
	var err error
	if ok && len(s) > 0 {
		since, err = strconv.Atoi(s[0])
		if err != nil {
			lgc.Logger().WithFields(log.Fields{
				"error": err, "param_name": "since", "value": s[0],
			}).Warn("failed to convert string to integer")
			// Unfortunately, graylog 2.4.3 returns 404
			return nil, 404, fmt.Errorf("HTTP 404 Not Found")
		}
		if since < 0 {
			lgc.Logger().WithFields(log.Fields{
				"error": err, "param_name": "since", "value": since,
			}).Warn("must be greater than or equal to 0")
			return nil, 400, fmt.Errorf("must be greater than or equal to 0")
		}
	}
	l, ok := query["limit"]
	if ok && len(l) > 0 {
		limit, err = strconv.Atoi(l[0])
		if err != nil {
			lgc.Logger().WithFields(log.Fields{
				"error": err, "param_name": "limit", "value": l[0],
			}).Warn("failed to convert string to integer")
			// Unfortunately, graylog 2.4.3 returns 404
			return nil, 404, fmt.Errorf("HTTP 404 Not Found")
		}
		if limit < 1 {
			lgc.Logger().WithFields(log.Fields{
				"error": err, "param_name": "limit", "value": limit,
			}).Warn("must be greater than or equal to 1")
			return nil, 400, fmt.Errorf("must be greater than or equal to 1")
		}
	}
	arr, total, sc, err := lgc.GetAlerts(since, limit)
	if err != nil {
		return arr, sc, err
	}
	return &graylog.AlertsBody{Alerts: arr, Total: total}, sc, nil
}
