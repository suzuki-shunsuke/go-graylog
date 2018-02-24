package graylog

import (
	"fmt"
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

	Users  map[string]User
	Roles  map[string]Role
	Inputs map[string]Input
}

func GetMockServer() (*MockServer, error) {
	ms := &MockServer{
		Users:  map[string]User{},
		Roles:  map[string]Role{},
		Inputs: map[string]Input{},
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

	server := httptest.NewServer(router)
	u := fmt.Sprintf("http://%s/api", server.Listener.Addr().String())
	ms.Server = server
	ms.Endpoint = u
	return ms, nil
}
