package graylog

import (
	"fmt"
	"net"
	"net/http/httptest"
	"sync"

	"github.com/julienschmidt/httprouter"
)

var (
	once sync.Once
)

type MockServer struct {
	Server   *httptest.Server
	Endpoint string

	Users     map[string]User
	Roles     map[string]Role
	Inputs    map[string]Input
	IndexSets map[string]IndexSet
}

// NewMockServer returns new MockServer but doesn't start it.
// If addr is an empty string, the free port is assigned automatially.
func NewMockServer(addr string) (*MockServer, error) {
	ms := &MockServer{
		Users:     map[string]User{},
		Roles:     map[string]Role{},
		Inputs:    map[string]Input{},
		IndexSets: map[string]IndexSet{},
	}

	router := httprouter.New()

	router.GET("/api/roles/:rolename", ms.handleGetRole)
	router.PUT("/api/roles/:rolename", ms.handleUpdateRole)
	router.DELETE("/api/roles/:rolename", ms.handleDeleteRole)
	router.GET("/api/roles", ms.handleGetRoles)
	router.POST("/api/roles", ms.handleCreateRole)

	router.GET("/api/users/:username", ms.handleGetUser)
	router.PUT("/api/users/:username", ms.handleUpdateUser)
	router.DELETE("/api/users/:username", ms.handleDeleteUser)
	router.GET("/api/users", ms.handleGetUsers)
	router.POST("/api/users", ms.handleCreateUser)

	router.GET("/api/roles/:rolename/members", ms.handleRoleMembers)
	router.PUT("/api/roles/:rolename/members/:username", ms.handleAddUserToRole)
	router.DELETE("/api/roles/:rolename/members/:username", ms.handleRemoveUserFromRole)

	router.GET("/api/system/inputs", ms.handleGetInputs)
	router.GET("/api/system/inputs/:inputId", ms.handleGetInput)
	router.POST("/api/system/inputs", ms.handleCreateInput)
	router.PUT("/api/system/inputs/:inputId", ms.handleUpdateInput)
	router.DELETE("/api/system/inputs/:inputId", ms.handleDeleteInput)

	router.GET("/api/system/indices/index_sets", ms.handleGetIndexSets)
	router.GET("/api/system/indices/index_sets/:indexSetId", ms.handleGetIndexSet)
	router.POST("/api/system/indices/index_sets", ms.handleCreateIndexSet)
	router.PUT("/api/system/indices/index_sets/:indexSetId", ms.handleUpdateIndexSet)
	router.DELETE("/api/system/indices/index_sets/:indexSetId", ms.handleDeleteIndexSet)

	server := httptest.NewUnstartedServer(router)
	if addr != "" {
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			return nil, err
		}
		server.Listener = ln
	}
	u := fmt.Sprintf("http://%s/api", server.Listener.Addr().String())
	ms.Endpoint = u
	ms.Server = server
	return ms, nil
}

// Start starts a server from NewUnstartedServer.
func (ms *MockServer) Start() {
	ms.Server.Start()
}

// Close shuts down the server and blocks until all outstanding requests on this server have completed.
func (ms *MockServer) Close() {
	ms.Server.Close()
}
