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

func handlerFuncs() {
	http.HandleFunc("/api/roles", handleRoles)
	http.HandleFunc("/api/roles/", handleRole)

	http.HandleFunc("/api/users", handleUsers)
	http.HandleFunc("/api/users/", handleUser)
}

type MockServer struct {
	Server   *httptest.Server
	Endpoint string

	Users []User
	Roles []Role
}

func GetMockServer() (*MockServer, error) {
	m := http.NewServeMux()

	m.Handle("/api/roles", http.HandlerFunc(handleRoles))
	m.Handle("/api/roles/", http.HandlerFunc(handleRole))
	m.Handle("/api/users", http.HandlerFunc(handleUsers))
	m.Handle("/api/users/", http.HandlerFunc(handleUser))

	server := httptest.NewServer(m)
	u := fmt.Sprintf("http://%s/api", server.Listener.Addr().String())
	ms := &MockServer{
		Server:   server,
		Endpoint: u,
		Users:    []User{},
		Roles:    []Role{},
	}
	return ms, nil
}
