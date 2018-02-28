package graylog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func (ms *MockServer) AddRole(role *Role) {
	ms.Roles[role.Name] = *role
	ms.safeSave()
}

func (ms *MockServer) UpdateRole(name string, role *Role) {
	delete(ms.Roles, name)
	ms.AddRole(role)
}

func (ms *MockServer) DeleteRole(name string) {
	delete(ms.Roles, name)
	ms.safeSave()
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

// GET /roles/{rolename} Retrieve permissions for a single role
func (ms *MockServer) handleGetRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	name := ps.ByName("rolename")
	ms.Logger.WithFields(log.Fields{
		"handler": "handleGetRole", "rolename": name}).Info("request start")
	role, ok := ms.Roles[name]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "No role found with name %s"}`, name)))
		return
	}
	b, err := json.Marshal(&role)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

// PUT /roles/{rolename} Update an existing role
func (ms *MockServer) handleUpdateRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	name := ps.ByName("rolename")
	if _, ok := ms.Roles[name]; !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "No role found with name %s"}`, name)))
		return
	}
	role := &Role{}
	err = json.Unmarshal(b, role)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	sc, msg := validateRole(role)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write(msg)
		return
	}
	ms.UpdateRole(name, role)
	b, err = json.Marshal(role)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

// DELETE /roles/{rolename} Remove the named role and dissociate any users from it
func (ms *MockServer) handleDeleteRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	name := ps.ByName("rolename")
	_, ok := ms.Roles[name]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "No role found with name %s"}`, name)))
		return
	}
	ms.DeleteRole(name)
}

func validateRole(role *Role) (int, []byte) {
	if role.Name == "" {
		return 400, []byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.models.roles.responses.RoleResponse, problem: Null name\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@472db3c8; line: 1, column: 31]"}`)
	}
	if role.Permissions == nil || len(role.Permissions) == 0 {
		return 400, []byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.models.roles.responses.RoleResponse, problem: Null permissions\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@7e64a22d; line: 1, column: 16]"}`)
	}
	return 200, []byte("")
}

// POST /roles Create a new role
func (ms *MockServer) handleCreateRole(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	role := &Role{}
	err = json.Unmarshal(b, role)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	sc, msg := validateRole(role)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write(msg)
		return
	}
	if _, ok := ms.Roles[role.Name]; ok {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "Role %s already exists."}`,
			role.Name)))
		return
	}
	ms.AddRole(role)
	b, err = json.Marshal(role)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

// GET /roles List all roles
func (ms *MockServer) handleGetRoles(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
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
