package mockserver

import (
	"github.com/julienschmidt/httprouter"
)

func newRouter(ms *Server) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/roles/:rolename", wrapHandle(ms, ms.handleGetRole))
	router.PUT("/api/roles/:rolename", wrapHandle(ms, ms.handleUpdateRole))
	router.DELETE("/api/roles/:rolename", wrapHandle(ms, ms.handleDeleteRole))
	router.GET("/api/roles", wrapHandle(ms, ms.handleGetRoles))
	router.POST("/api/roles", wrapHandle(ms, ms.handleCreateRole))

	router.GET("/api/users/:username", wrapHandle(ms, ms.handleGetUser))
	router.PUT("/api/users/:username", wrapHandle(ms, ms.handleUpdateUser))
	router.DELETE("/api/users/:username", wrapHandle(ms, ms.handleDeleteUser))
	router.GET("/api/users", wrapHandle(ms, ms.handleGetUsers))
	router.POST("/api/users", wrapHandle(ms, ms.handleCreateUser))

	router.GET("/api/roles/:rolename/members", wrapHandle(ms, ms.handleRoleMembers))
	router.PUT("/api/roles/:rolename/members/:username", wrapHandle(ms, ms.handleAddUserToRole))
	router.DELETE(
		"/api/roles/:rolename/members/:username", wrapHandle(ms, ms.handleRemoveUserFromRole))

	router.GET("/api/system/inputs", wrapHandle(ms, ms.handleGetInputs))
	router.GET("/api/system/inputs/:inputID", wrapHandle(ms, ms.handleGetInput))
	router.POST("/api/system/inputs", wrapHandle(ms, ms.handleCreateInput))
	router.PUT("/api/system/inputs/:inputID", wrapHandle(ms, ms.handleUpdateInput))
	router.DELETE("/api/system/inputs/:inputID", wrapHandle(ms, ms.handleDeleteInput))

	router.GET("/api/system/indices/index_sets", wrapHandle(ms, ms.handleGetIndexSets))
	router.GET(
		"/api/system/indices/index_sets/:indexSetID", wrapHandle(ms, ms.handleGetIndexSet))
	router.POST("/api/system/indices/index_sets", wrapHandle(ms, ms.handleCreateIndexSet))
	router.PUT(
		"/api/system/indices/index_sets/:indexSetID", wrapHandle(ms, ms.handleUpdateIndexSet))
	router.DELETE(
		"/api/system/indices/index_sets/:indexSetID", wrapHandle(ms, ms.handleDeleteIndexSet))
	router.PUT(
		"/api/system/indices/index_sets/:indexSetID/default",
		wrapHandle(ms, ms.handleSetDefaultIndexSet))

	router.GET(
		"/api/system/indices/index_sets/:indexSetID/stats",
		wrapHandle(ms, ms.handleGetIndexSetStats))

	router.GET("/api/streams", wrapHandle(ms, ms.handleGetStreams))
	router.POST("/api/streams", wrapHandle(ms, ms.handleCreateStream))
	router.GET("/api/streams/:streamID", wrapHandle(ms, ms.handleGetStream))
	router.PUT("/api/streams/:streamID", wrapHandle(ms, ms.handleUpdateStream))
	router.DELETE("/api/streams/:streamID", wrapHandle(ms, ms.handleDeleteStream))
	router.POST("/api/streams/:streamID/pause", wrapHandle(ms, ms.handlePauseStream))
	router.POST("/api/streams/:streamID/resume", wrapHandle(ms, ms.handleResumeStream))

	router.GET("/api/streams/:streamID/rules", wrapHandle(ms, ms.handleGetStreamRules))
	router.POST("/api/streams/:streamID/rules", wrapHandle(ms, ms.handleCreateStreamRule))
	router.PUT("/api/streams/:streamID/rules/:streamRuleID", wrapHandle(ms, ms.handleUpdateStreamRule))
	router.DELETE("/api/streams/:streamID/rules/:streamRuleID", wrapHandle(ms, ms.handleDeleteStreamRule))
	router.GET("/api/streams/:streamID/rules/:streamRuleID", wrapHandle(ms, ms.handleGetStreamRule))

	router.NotFound = ms.handleNotFound
	return router
}
