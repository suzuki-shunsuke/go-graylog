package graylog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

// /roles
func (ms *MockServer) handleRoles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ms.handleGetRoles(w, r)
	case http.MethodPost:
		ms.handleCreateRole(w, r)
	}
}

// /roles/{rolename}
func (ms *MockServer) handleRole(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ms.handleGetRole(w, r)
	case http.MethodPut:
		ms.handleUpdateRole(w, r)
	case http.MethodDelete:
		ms.handleDeleteRole(w, r)
	}
}

func (ms *MockServer) handleGetRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := path.Base(r.URL.Path)
	role, ok := ms.Roles[name]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(`{"type": "ApiError", "message": "No role found with name %s"}`, name)))
		return
	}
	b, err := json.Marshal(&role)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func (ms *MockServer) handleUpdateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	name := path.Base(r.URL.Path)
	if _, ok := ms.Roles[name]; !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(`{"type": "ApiError", "message": "No role found with name %s"}`, name)))
		return
	}
	role := Role{}
	err = json.Unmarshal(b, &role)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	sc, msg := validateRole(&role)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write(msg)
		return
	}
	delete(ms.Roles, name)
	ms.Roles[role.Name] = role
	b, err = json.Marshal(&role)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func (ms *MockServer) handleDeleteRole(w http.ResponseWriter, r *http.Request) {}

func validateRole(role *Role) (int, []byte) {
	if role.Name == "" {
		return 400, []byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.models.roles.responses.RoleResponse, problem: Null name\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@472db3c8; line: 1, column: 31]"}`)
	}
	if role.Permissions == nil || len(role.Permissions) == 0 {
		return 400, []byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.models.roles.responses.RoleResponse, problem: Null permissions\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@7e64a22d; line: 1, column: 16]"}`)
	}
	return 200, []byte("")
}

func (ms *MockServer) handleCreateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	role := Role{}
	err = json.Unmarshal(b, &role)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	sc, msg := validateRole(&role)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write(msg)
		return
	}
	if _, ok := ms.Roles[role.Name]; ok {
		w.WriteHeader(400)
		w.Write([]byte(`{"type": "ApiError", "message": "Role Admin already exists."}`))
		return
	}
	ms.Roles[role.Name] = role
	b, err = json.Marshal(&role)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func (ms *MockServer) handleGetRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	arr := ms.RoleList()
	roles := rolesBody{Roles: arr, Total: len(arr)}
	b, err := json.Marshal(&roles)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}
