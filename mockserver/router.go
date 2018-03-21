package mockserver

import (
	"github.com/julienschmidt/httprouter"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/handler"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

func newRouter(ms *logic.Server) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/roles/:rolename", wrapHandle(ms, handler.HandleGetRole))
	router.PUT("/api/roles/:rolename", wrapHandle(ms, handler.HandleUpdateRole))
	router.DELETE("/api/roles/:rolename", wrapHandle(ms, handler.HandleDeleteRole))
	router.GET("/api/roles", wrapHandle(ms, handler.HandleGetRoles))
	router.POST("/api/roles", wrapHandle(ms, handler.HandleCreateRole))

	router.GET("/api/users/:username", wrapHandle(ms, handler.HandleGetUser))
	router.PUT("/api/users/:username", wrapHandle(ms, handler.HandleUpdateUser))
	router.DELETE("/api/users/:username", wrapHandle(ms, handler.HandleDeleteUser))
	router.GET("/api/users", wrapHandle(ms, handler.HandleGetUsers))
	router.POST("/api/users", wrapHandle(ms, handler.HandleCreateUser))

	router.GET("/api/roles/:rolename/members", wrapHandle(ms, handler.HandleRoleMembers))
	router.PUT("/api/roles/:rolename/members/:username", wrapHandle(ms, handler.HandleAddUserToRole))
	router.DELETE(
		"/api/roles/:rolename/members/:username", wrapHandle(ms, handler.HandleRemoveUserFromRole))

	router.GET("/api/system/inputs", wrapHandle(ms, handler.HandleGetInputs))
	router.GET("/api/system/inputs/:inputID", wrapHandle(ms, handler.HandleGetInput))
	router.POST("/api/system/inputs", wrapHandle(ms, handler.HandleCreateInput))
	router.PUT("/api/system/inputs/:inputID", wrapHandle(ms, handler.HandleUpdateInput))
	router.DELETE("/api/system/inputs/:inputID", wrapHandle(ms, handler.HandleDeleteInput))

	router.GET("/api/system/indices/index_sets", wrapHandle(ms, handler.HandleGetIndexSets))
	router.GET(
		"/api/system/indices/index_sets/:indexSetID", wrapHandle(ms, handler.HandleGetIndexSet))
	router.POST("/api/system/indices/index_sets", wrapHandle(ms, handler.HandleCreateIndexSet))
	router.PUT(
		"/api/system/indices/index_sets/:indexSetID", wrapHandle(ms, handler.HandleUpdateIndexSet))
	router.DELETE(
		"/api/system/indices/index_sets/:indexSetID", wrapHandle(ms, handler.HandleDeleteIndexSet))
	router.PUT(
		"/api/system/indices/index_sets/:indexSetID/default",
		wrapHandle(ms, handler.HandleSetDefaultIndexSet))

	router.GET(
		"/api/system/indices/index_sets/:indexSetID/stats",
		wrapHandle(ms, handler.HandleGetIndexSetStats))

	router.GET("/api/streams", wrapHandle(ms, handler.HandleGetStreams))
	router.POST("/api/streams", wrapHandle(ms, handler.HandleCreateStream))
	router.GET("/api/streams/:streamID", wrapHandle(ms, handler.HandleGetStream))
	router.PUT("/api/streams/:streamID", wrapHandle(ms, handler.HandleUpdateStream))
	router.DELETE("/api/streams/:streamID", wrapHandle(ms, handler.HandleDeleteStream))
	router.POST("/api/streams/:streamID/pause", wrapHandle(ms, handler.HandlePauseStream))
	router.POST("/api/streams/:streamID/resume", wrapHandle(ms, handler.HandleResumeStream))

	router.GET("/api/streams/:streamID/rules", wrapHandle(ms, handler.HandleGetStreamRules))
	router.POST("/api/streams/:streamID/rules", wrapHandle(ms, handler.HandleCreateStreamRule))
	router.PUT("/api/streams/:streamID/rules/:streamRuleID", wrapHandle(ms, handler.HandleUpdateStreamRule))
	router.DELETE("/api/streams/:streamID/rules/:streamRuleID", wrapHandle(ms, handler.HandleDeleteStreamRule))
	router.GET("/api/streams/:streamID/rules/:streamRuleID", wrapHandle(ms, handler.HandleGetStreamRule))

	router.NotFound = handler.HandleNotFound(ms)
	return router
}
