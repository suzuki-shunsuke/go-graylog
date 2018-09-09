package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

// NewRouter returns a new HTTP router.
func NewRouter(lgc *logic.Logic) http.Handler {
	e := echo.New()

	// Alarm Callback
	e.GET("/api/alerts/callbacks", wrapEchoHandle(lgc, HandleGetAlarmCallbacks))

	// Alert Condition
	e.GET("/api/alerts/conditions", wrapEchoHandle(lgc, HandleGetAlertConditions))

	// Dashboard
	e.GET("/api/dashboards/:dashboardID", wrapEchoHandle(lgc, HandleGetDashboard))
	e.GET("/api/dashboards", wrapEchoHandle(lgc, HandleGetDashboards))
	e.POST("/api/dashboards", wrapEchoHandle(lgc, HandleCreateDashboard))
	e.PUT("/api/dashboards/:dashboardID", wrapEchoHandle(lgc, HandleUpdateDashboard))
	e.DELETE("/api/dashboards/:dashboardID", wrapEchoHandle(lgc, HandleDeleteDashboard))

	// Collector Configuration
	e.GET(
		"/api/plugins/org.graylog.plugins.collector/configurations",
		wrapEchoHandle(lgc, HandleGetCollectorConfigurations))
	e.POST(
		"/api/plugins/org.graylog.plugins.collector/configurations",
		wrapEchoHandle(lgc, HandleCreateCollectorConfiguration))
	e.PUT(
		"/api/plugins/org.graylog.plugins.collector/configurations/:collectorConfigurationID/name",
		wrapEchoHandle(lgc, HandleUpdateCollectorConfigurationName))
	e.GET(
		"/api/plugins/org.graylog.plugins.collector/configurations/:collectorConfigurationID",
		wrapEchoHandle(lgc, HandleGetCollectorConfiguration))
	e.DELETE(
		"/api/plugins/org.graylog.plugins.collector/configurations/:collectorConfigurationID",
		wrapEchoHandle(lgc, HandleDeleteCollectorConfiguration))

	// Collector Configuration Input
	e.POST(
		"/api/plugins/org.graylog.plugins.collector/configurations/:collectorConfigurationID/inputs",
		wrapEchoHandle(lgc, HandleCreateCollectorConfigurationInput))
	e.PUT(
		"/api/plugins/org.graylog.plugins.collector/configurations/:collectorConfigurationID/inputs/:collectorConfigurationInputID",
		wrapEchoHandle(lgc, HandleUpdateCollectorConfigurationInput))
	e.DELETE(
		"/api/plugins/org.graylog.plugins.collector/configurations/:collectorConfigurationID/inputs/:collectorConfigurationInputID",
		wrapEchoHandle(lgc, HandleDeleteCollectorConfigurationInput))

	// Collector Configuration Output
	e.POST(
		"/api/plugins/org.graylog.plugins.collector/configurations/:collectorConfigurationID/outputs",
		wrapEchoHandle(lgc, HandleCreateCollectorConfigurationOutput))
	e.PUT(
		"/api/plugins/org.graylog.plugins.collector/configurations/:collectorConfigurationID/outputs/:collectorConfigurationOutputID",
		wrapEchoHandle(lgc, HandleUpdateCollectorConfigurationOutput))
	e.DELETE(
		"/api/plugins/org.graylog.plugins.collector/configurations/:collectorConfigurationID/outputs/:collectorConfigurationOutputID",
		wrapEchoHandle(lgc, HandleDeleteCollectorConfigurationOutput))

	// Role
	e.GET("/api/roles/:rolename", wrapEchoHandle(lgc, HandleGetRole))
	e.GET("/api/roles", wrapEchoHandle(lgc, HandleGetRoles))
	e.PUT("/api/roles/:rolename", wrapEchoHandle(lgc, HandleUpdateRole))
	e.DELETE("/api/roles/:rolename", wrapEchoHandle(lgc, HandleDeleteRole))
	e.POST("/api/roles", wrapEchoHandle(lgc, HandleCreateRole))

	// Role member
	e.GET("/api/roles/:rolename/members", wrapEchoHandle(lgc, HandleRoleMembers))
	e.PUT(
		"/api/roles/:rolename/members/:username",
		wrapEchoHandle(lgc, HandleAddUserToRole))
	e.DELETE(
		"/api/roles/:rolename/members/:username",
		wrapEchoHandle(lgc, HandleRemoveUserFromRole))

	// Alert
	e.GET("/api/streams/alerts/:alertID", wrapEchoHandle(lgc, HandleGetAlert))
	e.GET("/api/streams/alerts", wrapEchoHandle(lgc, HandleGetAlerts))

	// Stream
	e.GET("/api/streams/enabled", wrapEchoHandle(lgc, HandleGetEnabledStreams))
	e.GET("/api/streams/:streamID", wrapEchoHandle(lgc, HandleGetStream))
	e.GET("/api/streams", wrapEchoHandle(lgc, HandleGetStreams))
	e.POST("/api/streams", wrapEchoHandle(lgc, HandleCreateStream))
	e.PUT("/api/streams/:streamID", wrapEchoHandle(lgc, HandleUpdateStream))
	e.DELETE("/api/streams/:streamID", wrapEchoHandle(lgc, HandleDeleteStream))
	e.POST(
		"/api/streams/:streamID/pause", wrapEchoHandle(lgc, HandlePauseStream))
	e.POST(
		"/api/streams/:streamID/resume", wrapEchoHandle(lgc, HandleResumeStream))

	// Stream Rule
	e.GET("/api/streams/:streamID/rules/:streamRuleID", wrapEchoHandle(lgc, HandleGetStreamRule))
	e.GET("/api/streams/:streamID/rules", wrapEchoHandle(lgc, HandleGetStreamRules))
	e.POST("/api/streams/:streamID/rules", wrapEchoHandle(lgc, HandleCreateStreamRule))
	e.PUT("/api/streams/:streamID/rules/:streamRuleID", wrapEchoHandle(lgc, HandleUpdateStreamRule))
	e.DELETE("/api/streams/:streamID/rules/:streamRuleID", wrapEchoHandle(lgc, HandleDeleteStreamRule))

	// Input
	e.GET("/api/system/inputs/:inputID", wrapEchoHandle(lgc, HandleGetInput))
	e.GET("/api/system/inputs", wrapEchoHandle(lgc, HandleGetInputs))
	e.PUT("/api/system/inputs/:inputID", wrapEchoHandle(lgc, HandleUpdateInput))
	e.DELETE(
		"/api/system/inputs/:inputID", wrapEchoHandle(lgc, HandleDeleteInput))
	e.POST("/api/system/inputs", wrapEchoHandle(lgc, HandleCreateInput))

	// IndexSet
	e.GET(
		"/api/system/indices/index_sets/stats",
		wrapEchoHandle(lgc, HandleGetTotalIndexSetStats))
	e.GET(
		"/api/system/indices/index_sets/:indexSetID/stats",
		wrapEchoHandle(lgc, HandleGetIndexSetStats))
	e.GET(
		"/api/system/indices/index_sets/:indexSetID",
		wrapEchoHandle(lgc, HandleGetIndexSet))
	e.GET(
		"/api/system/indices/index_sets", wrapEchoHandle(lgc, HandleGetIndexSets))
	e.PUT(
		"/api/system/indices/index_sets/:indexSetID/default",
		wrapEchoHandle(lgc, HandleSetDefaultIndexSet))
	e.PUT(
		"/api/system/indices/index_sets/:indexSetID",
		wrapEchoHandle(lgc, HandleUpdateIndexSet))
	e.DELETE(
		"/api/system/indices/index_sets/:indexSetID",
		wrapEchoHandle(lgc, HandleDeleteIndexSet))
	e.POST(
		"/api/system/indices/index_sets",
		wrapEchoHandle(lgc, HandleCreateIndexSet))

	// LDAP Setting
	e.GET(
		"/api/system/ldap/settings", wrapEchoHandle(lgc, HandleGetLDAPSetting))
	e.PUT(
		"/api/system/ldap/settings", wrapEchoHandle(lgc, HandleUpdateLDAPSetting))
	e.DELETE(
		"/api/system/ldap/settings", wrapEchoHandle(lgc, HandleDeleteLDAPSetting))

	// User
	e.GET("/api/users/:username", wrapEchoHandle(lgc, HandleGetUser))
	e.GET("/api/users", wrapEchoHandle(lgc, HandleGetUsers))
	e.PUT("/api/users/:username", wrapEchoHandle(lgc, HandleUpdateUser))
	e.DELETE("/api/users/:username", wrapEchoHandle(lgc, HandleDeleteUser))
	e.POST("/api/users", wrapEchoHandle(lgc, HandleCreateUser))

	echo.NotFoundHandler = HandleNotFound(lgc)
	return e
}
