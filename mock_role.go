package graylog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
