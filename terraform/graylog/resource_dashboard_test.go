package graylog

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/v9/testdata"
)

func TestAccDashboard(t *testing.T) {
	setEnv()

	ds := testdata.Dashboard()

	tc := &testCase{
		t:          t,
		Name:       "dashboard",
		CreatePath: "/api/dashboards",
		GetPath:    "/api/dashboards/" + ds.ID,

		CreateReqBodyMap: map[string]interface{}{
			"title":       "test",
			"description": "test",
		},
		UpdateReqBodyMap: map[string]interface{}{
			"title":       "updated title",
			"description": "updated description",
		},
		CreatedDataPath:    "dashboard/dashboard.json",
		UpdatedDataPath:    "dashboard/updated_dashboard.json",
		CreateRespBodyPath: "dashboard/create_dashboard_response.json",
		CreateTFPath:       "dashboard/dashboard.tf",
		UpdateTFPath:       "dashboard/update_dashboard.tf",
	}
	tc.Test()
}
