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

	Users map[string]User
	Roles map[string]Role
}

func (ms *MockServer) UserList() []User {
	if ms.Users == nil {
		return []User{}
	}
	size := len(ms.Users)
	arr := make([]User, size)
	i := 0
	for _, user := range ms.Users {
		arr[i] = user
		i++
	}
	return arr
}

func (ms *MockServer) RoleList() []Role {
	if ms.Roles == nil {
		return []Role{}
	}
	size := len(ms.Roles)
	arr := make([]Role, size)
	i := 0
	for _, role := range ms.Roles {
		arr[i] = role
		i++
	}
	return arr
}

func GetMockServer() (*MockServer, error) {
	ms := &MockServer{
		Users: map[string]User{},
		Roles: map[string]Role{},
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

	server := httptest.NewServer(router)
	u := fmt.Sprintf("http://%s/api", server.Listener.Addr().String())
	ms.Server = server
	ms.Endpoint = u
	return ms, nil
}
