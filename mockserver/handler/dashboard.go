package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/util"
)

// HandleCreateDashboard is the handler of Create an Dashboard API.
func HandleCreateDashboard(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// POST /dashboards Create a dashboard
	if sc, err := lgc.Authorize(user, "dashboards:create"); err != nil {
		return nil, sc, err
	}
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("title", "description"),
			Forbidden:    set.NewStrSet("created_at", "id", "widgets"),
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}

	dashboard := &graylog.Dashboard{}
	if err := util.MSDecode(body, dashboard); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as DashboardData")
		return nil, 400, err
	}

	sc, err = lgc.AddDashboard(dashboard)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return &map[string]string{"dashboard_id": dashboard.ID}, sc, nil
}

// HandleDeleteDashboard is the handler of Delete an Dashboard API.
func HandleDeleteDashboard(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// DELETE /dashboards/{dashboardId} Delete a dashboard and all its widgets
	id := ps.PathParam("dashboardID")
	if sc, err := lgc.Authorize(user, "dashboards:edit", id); err != nil {
		return nil, sc, err
	}
	sc, err := lgc.DeleteDashboard(id)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return nil, sc, nil
}

// HandleGetDashboard is the handler of Get an Dashboard API.
func HandleGetDashboard(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// GET /dashboards/{dashboardId} Get a single dashboards and all configurations of its widgets.
	id := ps.PathParam("dashboardID")
	if sc, err := lgc.Authorize(user, "dashboards:read", id); err != nil {
		return nil, sc, err
	}
	return lgc.GetDashboard(id)
}

// HandleGetDashboards is the handler of Get Dashboards API.
func HandleGetDashboards(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// GET /dashboards Get a list of all dashboards and all configurations of their widgets.
	arr, total, sc, err := lgc.GetDashboards()
	if err != nil {
		return arr, sc, err
	}
	dashboards := &graylog.DashboardsBody{Dashboards: arr, Total: total}
	return dashboards, sc, nil
}

// HandleUpdateDashboard is the handler of Update an Dashboard API.
func HandleUpdateDashboard(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// PUT /dashboards/{dashboardId} Update the settings of a dashboard.
	// 204 No Content
	id := ps.PathParam("dashboardID")
	if sc, err := lgc.Authorize(user, "dashboards:edit", id); err != nil {
		return nil, sc, err
	}
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Optional:     set.NewStrSet("title", "description"),
			Forbidden:    set.NewStrSet("created_at", "id", "widgets"),
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}
	dashboard := &graylog.Dashboard{}
	if err := util.MSDecode(body, dashboard); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as Dashboard")
		return nil, 400, err
	}

	lgc.Logger().WithFields(log.Fields{
		"body": body, "dashboard": dashboard, "id": id,
	}).Debug("request body")

	dashboard.ID = id
	sc, err = lgc.UpdateDashboard(dashboard)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return nil, sc, nil
}
