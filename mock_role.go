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
	admin := Role{
		Name:        "Admin",
		Description: "Grants all permissions for Graylog administrators (built-in)",
		Permissions: []string{"*"},
		ReadOnly:    true,
	}
	name := path.Base(r.URL.Path)
	if name == "Admin" {
		b, err := json.Marshal(&admin)
		if err != nil {
			w.Write([]byte(`{"message":"500 Internal Server Error"}`))
			return
		}
		w.Write(b)
		return
	}
	t := Error{
		Message: fmt.Sprintf("No role found with name %s", name),
		Type:    "ApiError"}
	b, err := json.Marshal(&t)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func (ms *MockServer) handleUpdateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := path.Base(r.URL.Path)
	if name != "Admin" {
		t := Error{
			Message: fmt.Sprintf("No role found with name %s", name),
			Type:    "ApiError"}
		b, err := json.Marshal(&t)
		if err != nil {
			w.Write([]byte(`{"message":"500 Internal Server Error"}`))
			return
		}
		w.Write(b)
		return
	}
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

func (ms *MockServer) handleDeleteRole(w http.ResponseWriter, r *http.Request) {}

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
	if role.Name == "" {
		w.WriteHeader(400)
		w.Write([]byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.models.roles.responses.RoleResponse, problem: Null name\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@472db3c8; line: 1, column: 31]"}`))
		return
	}
	if role.Permissions == nil || len(role.Permissions) == 0 {
		w.WriteHeader(400)
		w.Write([]byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.models.roles.responses.RoleResponse, problem: Null permissions\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@7e64a22d; line: 1, column: 16]"}`))
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
