package graylog

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// HasRole
func (ms *MockServer) HasRole(name string) (bool, error) {
	return ms.store.HasRole(name)
}

// GetRole returns a Role.
func (ms *MockServer) GetRole(name string) (*Role, error) {
	return ms.store.GetRole(name)
}

// AddRole adds a new role to the mock server.
func (ms *MockServer) AddRole(role *Role) (int, error) {
	if err := CreateValidator.Struct(role); err != nil {
		return 400, err
	}
	ok, err := ms.HasRole(role.Name)
	if err != nil {
		return 500, err
	}
	if ok {
		return 400, fmt.Errorf("Role %s already exists.", role.Name)
	}
	if err := ms.store.AddRole(role); err != nil {
		return 500, err
	}
	return 200, nil
}

// UpdateRole updates a role.
func (ms *MockServer) UpdateRole(name string, role *Role) (int, error) {
	if err := UpdateValidator.Struct(role); err != nil {
		return 400, err
	}
	ok, err := ms.HasRole(name)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("No role found with name %s", name)
	}
	if name != role.Name {
		ok, err := ms.HasRole(role.Name)
		if err != nil {
			return 500, err
		}
		if ok {
			return 400, fmt.Errorf("The role %s has already existed.", role.Name)
		}
	}
	if err := ms.store.UpdateRole(name, role); err != nil {
		return 500, err
	}
	return 200, nil
}

// DeleteRole
func (ms *MockServer) DeleteRole(name string) (int, error) {
	ok, err := ms.HasRole(name)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("No role found with name %s", name)
	}
	if err := ms.store.DeleteRole(name); err != nil {
		return 500, err
	}
	return 200, nil
}

func (ms *MockServer) RoleList() ([]Role, error) {
	return ms.store.GetRoles()
}

// GET /roles/{rolename} Retrieve permissions for a single role
func (ms *MockServer) handleGetRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	name := ps.ByName("rolename")
	ms.Logger().WithFields(log.Fields{
		"handler": "handleGetRole", "rolename": name}).Info("request start")
	role, err := ms.GetRole(name)
	if err != nil {
		writeApiError(w, 500, err.Error())
		return
	}
	if role == nil {
		writeApiError(w, 404, "No role found with name %s", name)
		return
	}
	writeOr500Error(w, role)
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
	sc, err := ms.DeleteRole(name)
	if err != nil {
		writeApiError(w, sc, err.Error())
		return
	}
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
	arr, err := ms.RoleList()
	if err != nil {
		writeApiError(w, 500, err.Error())
		return
	}
	roles := &rolesBody{Roles: arr, Total: len(arr)}
	writeOr500Error(w, roles)
}
