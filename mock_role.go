package graylog

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// HasRole
func (ms *MockServer) HasRole(name string) bool {
	return ms.store.HasRole(name)
	// _, ok := ms.roles[name]
	// return ok
}

// GetRole returns a Role.
func (ms *MockServer) GetRole(name string) (Role, bool) {
	return ms.store.GetRole(name)
	// s, ok := ms.roles[name]
	// return s, ok
}

// AddRole adds a new role to the mock server.
func (ms *MockServer) AddRole(role *Role) (int, error) {
	if err := CreateValidator.Struct(role); err != nil {
		return 400, err
	}
	if ms.HasRole(role.Name) {
		return 400, fmt.Errorf("Role %s already exists.", role.Name)
	}
	return ms.store.AddRole(role)
}

// UpdateRole updates a role.
func (ms *MockServer) UpdateRole(name string, role *Role) (int, error) {
	if err := UpdateValidator.Struct(role); err != nil {
		return 400, err
	}
	if !ms.HasRole(name) {
		return 404, fmt.Errorf("No role found with name %s", name)
	}
	if name != role.Name && ms.HasRole(role.Name) {
		return 400, fmt.Errorf("The role %s has already existed.", name)
	}
	return ms.store.UpdateRole(name, role)
}

// DeleteRole
func (ms *MockServer) DeleteRole(name string) (int, error) {
	if !ms.HasRole(name) {
		return 404, fmt.Errorf("No role found with name %s", name)
	}
	return ms.store.DeleteRole(name)
}

func (ms *MockServer) RoleList() []Role {
	arr, _ := ms.store.GetRoles()
	return arr
}

// GET /roles/{rolename} Retrieve permissions for a single role
func (ms *MockServer) handleGetRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	name := ps.ByName("rolename")
	ms.Logger().WithFields(log.Fields{
		"handler": "handleGetRole", "rolename": name}).Info("request start")
	role, ok := ms.GetRole(name)
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

	requiredFields := []string{"name", "permissions"}
	allowedFields := []string{"description", "read_only"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}

	role := &Role{}
	if err := msDecode(body, role); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as Role")
		writeApiError(w, 400, "400 Bad Request")
		return
	}

	if sc, err := ms.UpdateRole(name, role); err != nil {
		writeApiError(w, sc, err.Error())
		return
	}
	ms.safeSave()
	writeOr500Error(w, role)
}

// DELETE /roles/{rolename} Remove the named role and dissociate any users from it
func (ms *MockServer) handleDeleteRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	name := ps.ByName("rolename")
	_, ok := ms.GetRole(name)
	if !ok {
		writeApiError(w, 404, "No role found with name %s", name)
		return
	}
	ms.DeleteRole(name)
	ms.safeSave()
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
	allowedFields := []string{"description", "read_only"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}

	role := &Role{}
	if err := msDecode(body, &role); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as Role")
		writeApiError(w, 400, "400 Bad Request")
		return
	}

	if sc, err := ms.AddRole(role); err != nil {
		writeApiError(w, sc, err.Error())
		return
	}
	ms.safeSave()
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
