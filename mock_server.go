package graylog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	Users map[string]User
	Roles map[string]Role
}

func (ms *MockServer) handleRoles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ms.handleGetRoles(w, r)
	case http.MethodPost:
		ms.handleCreateRole(w, r)
	}
}

func (ms *MockServer) handleCreateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	role := Role{}
	err = json.Unmarshal(b, &role)
	if err != nil {
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	b, err = json.Marshal(&role)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func (ms *MockServer) handleGetRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	size := len(ms.Roles)
	arr := make([]Role, size)
	i := 0
	for _, role := range ms.Roles {
		arr[i] = role
		i++
	}
	roles := rolesBody{
		Roles: arr,
		Total: size}
	b, err := json.Marshal(&roles)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func GetMockServer() (*MockServer, error) {
	m := http.NewServeMux()
	ms := &MockServer{
		Users: map[string]User{},
		Roles: map[string]Role{},
	}

	m.Handle("/api/roles", http.HandlerFunc(ms.handleRoles))
	m.Handle("/api/roles/", http.HandlerFunc(handleRole))
	m.Handle("/api/users", http.HandlerFunc(handleUsers))
	m.Handle("/api/users/", http.HandlerFunc(handleUser))

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
