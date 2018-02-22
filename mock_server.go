package graylog

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
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
	m := http.NewServeMux()
	ms := &MockServer{
		Users: map[string]User{},
		Roles: map[string]Role{},
	}

	m.Handle("/api/roles", http.HandlerFunc(ms.handleRoles))
	m.Handle("/api/roles/", http.HandlerFunc(ms.handleRole))
	m.Handle("/api/users", http.HandlerFunc(ms.handleUsers))
	m.Handle("/api/users/", http.HandlerFunc(ms.handleUser))

	server := httptest.NewServer(m)
	u := fmt.Sprintf("http://%s/api", server.Listener.Addr().String())
	ms.Server = server
	ms.Endpoint = u

	ms.Roles = map[string]Role{
		"Admin": {
			Name:        "Admin",
			Description: "Grants all permissions for Graylog administrators (built-in)",
			Permissions: []string{"*"},
			ReadOnly:    true},
	}

	return ms, nil
}
