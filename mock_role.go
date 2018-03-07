package graylog

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
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
	ms.handleInit(w, r, false)
	name := ps.ByName("rolename")
	ms.Logger.WithFields(log.Fields{
		"handler": "handleGetRole", "rolename": name}).Info("request start")
	role, ok := ms.Roles[name]
	if !ok {
		writeApiError(w, 404, "No role found with name %s", name)
		return
	}
	writeOr500Error(w, &role)
}

// PUT /roles/{rolename} Update an existing role
func (ms *MockServer) handleUpdateRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		write500Error(w)
		return
	}

	name := ps.ByName("rolename")
	if _, ok := ms.Roles[name]; !ok {
		writeApiError(w, 404, "No role found with name %s", name)
		return
	}

	requiredFields := []string{"name", "permissions"}
	allowedFields := []string{
		"name", "description", "read_only", "permissions"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}

	role := &Role{}
	if err := mapstructure.Decode(body, role); err != nil {
		ms.Logger.WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as Role")
		writeApiError(w, 400, "400 Bad Request")
		return
	}

	if err := UpdateValidator.Struct(role); err != nil {
		writeApiError(w, 400, err.Error())
		return
	}
	ms.UpdateRole(name, role)
	writeOr500Error(w, role)
}

// DELETE /roles/{rolename} Remove the named role and dissociate any users from it
func (ms *MockServer) handleDeleteRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	name := ps.ByName("rolename")
	_, ok := ms.Roles[name]
	if !ok {
		writeApiError(w, 404, "No role found with name %s", name)
		return
	}
	ms.DeleteRole(name)
}

// POST /roles Create a new role
func (ms *MockServer) handleCreateRole(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		write500Error(w)
		return
	}

	requiredFields := []string{"name", "permissions"}
	allowedFields := []string{
		"name", "description", "read_only", "permissions"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}

	role := &Role{}
	if err := mapstructure.Decode(body, &role); err != nil {
		ms.Logger.WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as Role")
		writeApiError(w, 400, "400 Bad Request")
		return
	}

	if err := CreateValidator.Struct(role); err != nil {
		writeApiError(w, 400, err.Error())
		return
	}
	if _, ok := ms.Roles[role.Name]; ok {
		writeApiError(w, 400, "Role %s already exists.", role.Name)
		return
	}
	ms.AddRole(role)
	writeOr500Error(w, role)
}

// GET /roles List all roles
func (ms *MockServer) handleGetRoles(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	ms.handleInit(w, r, false)
	arr := ms.RoleList()
	roles := &rolesBody{Roles: arr, Total: len(arr)}
	writeOr500Error(w, roles)
}
