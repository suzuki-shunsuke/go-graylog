package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

// NewRouter returns a new HTTP router.
func NewRouter(lgc *logic.Logic) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/roles/:rolename", wrapHandle(lgc, HandleGetRole))
	router.PUT("/api/roles/:rolename", wrapHandle(lgc, HandleUpdateRole))
	router.DELETE("/api/roles/:rolename", wrapHandle(lgc, HandleDeleteRole))
	router.GET("/api/roles", wrapHandle(lgc, HandleGetRoles))
	router.POST("/api/roles", wrapHandle(lgc, HandleCreateRole))

	router.GET("/api/users/:username", wrapHandle(lgc, HandleGetUser))
	router.PUT("/api/users/:username", wrapHandle(lgc, HandleUpdateUser))
	router.DELETE("/api/users/:username", wrapHandle(lgc, HandleDeleteUser))
	router.GET("/api/users", wrapHandle(lgc, HandleGetUsers))
	router.POST("/api/users", wrapHandle(lgc, HandleCreateUser))

	router.GET("/api/roles/:rolename/members", wrapHandle(lgc, HandleRoleMembers))
	router.PUT("/api/roles/:rolename/members/:username", wrapHandle(lgc, HandleAddUserToRole))
	router.DELETE(
		"/api/roles/:rolename/members/:username", wrapHandle(lgc, HandleRemoveUserFromRole))

	router.GET("/api/system/inputs", wrapHandle(lgc, HandleGetInputs))
	router.GET("/api/system/inputs/:inputID", wrapHandle(lgc, HandleGetInput))
	router.POST("/api/system/inputs", wrapHandle(lgc, HandleCreateInput))
	router.PUT("/api/system/inputs/:inputID", wrapHandle(lgc, HandleUpdateInput))
	router.DELETE("/api/system/inputs/:inputID", wrapHandle(lgc, HandleDeleteInput))

	router.GET("/api/system/indices/index_sets", wrapHandle(lgc, HandleGetIndexSets))
	router.GET(
		"/api/system/indices/index_sets/:indexSetID", wrapHandle(lgc, HandleGetIndexSet))
	router.POST("/api/system/indices/index_sets", wrapHandle(lgc, HandleCreateIndexSet))
	router.PUT(
		"/api/system/indices/index_sets/:indexSetID", wrapHandle(lgc, HandleUpdateIndexSet))
	router.DELETE(
		"/api/system/indices/index_sets/:indexSetID", wrapHandle(lgc, HandleDeleteIndexSet))
	router.PUT(
		"/api/system/indices/index_sets/:indexSetID/default",
		wrapHandle(lgc, HandleSetDefaultIndexSet))

	router.GET(
		"/api/system/indices/index_sets/:indexSetID/stats",
		wrapHandle(lgc, HandleGetIndexSetStats))

	router.GET("/api/streams", wrapHandle(lgc, HandleGetStreams))
	router.POST("/api/streams", wrapHandle(lgc, HandleCreateStream))
	router.GET("/api/streams/:streamID", wrapHandle(lgc, HandleGetStream))
	router.PUT("/api/streams/:streamID", wrapHandle(lgc, HandleUpdateStream))
	router.DELETE("/api/streams/:streamID", wrapHandle(lgc, HandleDeleteStream))
	router.POST("/api/streams/:streamID/pause", wrapHandle(lgc, HandlePauseStream))
	router.POST("/api/streams/:streamID/resume", wrapHandle(lgc, HandleResumeStream))

	router.GET("/api/streams/:streamID/rules", wrapHandle(lgc, HandleGetStreamRules))
	router.POST("/api/streams/:streamID/rules", wrapHandle(lgc, HandleCreateStreamRule))
	router.PUT("/api/streams/:streamID/rules/:streamRuleID", wrapHandle(lgc, HandleUpdateStreamRule))
	router.DELETE("/api/streams/:streamID/rules/:streamRuleID", wrapHandle(lgc, HandleDeleteStreamRule))
	router.GET("/api/streams/:streamID/rules/:streamRuleID", wrapHandle(lgc, HandleGetStreamRule))

	router.GET("/api/alerts/conditions", wrapHandle(lgc, HandleGetAlertConditions))

	router.NotFound = HandleNotFound(lgc)
	return router
}
