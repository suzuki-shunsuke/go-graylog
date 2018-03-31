package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

func NewRouter(ms *logic.Server) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/roles/:rolename", wrapHandle(ms, HandleGetRole))
	router.PUT("/api/roles/:rolename", wrapHandle(ms, HandleUpdateRole))
	router.DELETE("/api/roles/:rolename", wrapHandle(ms, HandleDeleteRole))
	router.GET("/api/roles", wrapHandle(ms, HandleGetRoles))
	router.POST("/api/roles", wrapHandle(ms, HandleCreateRole))

	router.GET("/api/users/:username", wrapHandle(ms, HandleGetUser))
	router.PUT("/api/users/:username", wrapHandle(ms, HandleUpdateUser))
	router.DELETE("/api/users/:username", wrapHandle(ms, HandleDeleteUser))
	router.GET("/api/users", wrapHandle(ms, HandleGetUsers))
	router.POST("/api/users", wrapHandle(ms, HandleCreateUser))

	router.GET("/api/roles/:rolename/members", wrapHandle(ms, HandleRoleMembers))
	router.PUT("/api/roles/:rolename/members/:username", wrapHandle(ms, HandleAddUserToRole))
	router.DELETE(
		"/api/roles/:rolename/members/:username", wrapHandle(ms, HandleRemoveUserFromRole))

	router.GET("/api/system/inputs", wrapHandle(ms, HandleGetInputs))
	router.GET("/api/system/inputs/:inputID", wrapHandle(ms, HandleGetInput))
	router.POST("/api/system/inputs", wrapHandle(ms, HandleCreateInput))
	router.PUT("/api/system/inputs/:inputID", wrapHandle(ms, HandleUpdateInput))
	router.DELETE("/api/system/inputs/:inputID", wrapHandle(ms, HandleDeleteInput))

	router.GET("/api/system/indices/index_sets", wrapHandle(ms, HandleGetIndexSets))
	router.GET(
		"/api/system/indices/index_sets/:indexSetID", wrapHandle(ms, HandleGetIndexSet))
	router.POST("/api/system/indices/index_sets", wrapHandle(ms, HandleCreateIndexSet))
	router.PUT(
		"/api/system/indices/index_sets/:indexSetID", wrapHandle(ms, HandleUpdateIndexSet))
	router.DELETE(
		"/api/system/indices/index_sets/:indexSetID", wrapHandle(ms, HandleDeleteIndexSet))
	router.PUT(
		"/api/system/indices/index_sets/:indexSetID/default",
		wrapHandle(ms, HandleSetDefaultIndexSet))

	router.GET(
		"/api/system/indices/index_sets/:indexSetID/stats",
		wrapHandle(ms, HandleGetIndexSetStats))

	router.GET("/api/streams", wrapHandle(ms, HandleGetStreams))
	router.POST("/api/streams", wrapHandle(ms, HandleCreateStream))
	router.GET("/api/streams/:streamID", wrapHandle(ms, HandleGetStream))
	router.PUT("/api/streams/:streamID", wrapHandle(ms, HandleUpdateStream))
	router.DELETE("/api/streams/:streamID", wrapHandle(ms, HandleDeleteStream))
	router.POST("/api/streams/:streamID/pause", wrapHandle(ms, HandlePauseStream))
	router.POST("/api/streams/:streamID/resume", wrapHandle(ms, HandleResumeStream))

	router.GET("/api/streams/:streamID/rules", wrapHandle(ms, HandleGetStreamRules))
	router.POST("/api/streams/:streamID/rules", wrapHandle(ms, HandleCreateStreamRule))
	router.PUT("/api/streams/:streamID/rules/:streamRuleID", wrapHandle(ms, HandleUpdateStreamRule))
	router.DELETE("/api/streams/:streamID/rules/:streamRuleID", wrapHandle(ms, HandleDeleteStreamRule))
	router.GET("/api/streams/:streamID/rules/:streamRuleID", wrapHandle(ms, HandleGetStreamRule))

	router.NotFound = HandleNotFound(ms)
	return router
}
