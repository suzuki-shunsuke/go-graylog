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

func GetMockServer() (*httptest.Server, string, error) {
	once.Do(handlerFuncs)
	server := httptest.NewServer(nil)
	u := fmt.Sprintf("http://%s/api", server.Listener.Addr().String())
	return server, u, nil
}
