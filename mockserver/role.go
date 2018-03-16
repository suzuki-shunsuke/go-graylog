package mockserver

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasRole
func (ms *MockServer) HasRole(name string) (bool, error) {
	return ms.store.HasRole(name)
}

// GetRole returns a Role.
func (ms *MockServer) GetRole(name string) (*graylog.Role, error) {
	return ms.store.GetRole(name)
}

// AddRole adds a new role to the mock server.
func (ms *MockServer) AddRole(role *graylog.Role) (int, error) {
	if err := validator.CreateValidator.Struct(role); err != nil {
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
func (ms *MockServer) UpdateRole(name string, role *graylog.Role) (int, error) {
	if err := validator.UpdateValidator.Struct(role); err != nil {
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
	return 204, nil
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

func (ms *MockServer) GetRoles() ([]graylog.Role, error) {
	return ms.store.GetRoles()
}

// GET /roles/{rolename} Retrieve permissions for a single role
func (ms *MockServer) handleGetRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	name := ps.ByName("rolename")
	ms.Logger().WithFields(log.Fields{
		"handler": "handleGetRole", "rolename": name}).Info("request start")
	role, err := ms.GetRole(name)
	if err != nil {
		return 500, nil, err
	}
	if role == nil {
		return 404, nil, fmt.Errorf("No role found with name %s", name)
	}
	return 200, role, nil
}

// PUT /roles/{rolename} Update an existing role
func (ms *MockServer) handleUpdateRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	name := ps.ByName("rolename")
	requiredFields := []string{"name", "permissions"}
	allowedFields := []string{"description", "read_only"}
	sc, msg, body := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if sc != 200 {
		return sc, nil, fmt.Errorf(msg)
	}

	role := &graylog.Role{}
	if err := msDecode(body, role); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as Role")
		return 400, nil, fmt.Errorf("400 Bad Request")
	}

	if sc, err := ms.UpdateRole(name, role); err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	return 204, role, nil
}

// DELETE /roles/{rolename} Remove the named role and dissociate any users from it
func (ms *MockServer) handleDeleteRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	name := ps.ByName("rolename")
	sc, err := ms.DeleteRole(name)
	if err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	return 204, nil, nil
}

// POST /roles Create a new role
func (ms *MockServer) handleCreateRole(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	requiredFields := []string{"name", "permissions"}
	allowedFields := []string{"description", "read_only"}
	sc, msg, body := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if sc != 200 {
		return sc, nil, fmt.Errorf(msg)
	}

	role := &graylog.Role{}
	if err := msDecode(body, &role); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as Role")
		return 400, nil, fmt.Errorf("400 Bad Request")
	}

	if sc, err := ms.AddRole(role); err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	return 201, role, nil
}

// GET /roles List all roles
func (ms *MockServer) handleGetRoles(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	arr, err := ms.GetRoles()
	if err != nil {
		return 500, nil, err
	}
	roles := &graylog.RolesBody{Roles: arr, Total: len(arr)}
	return 200, roles, nil
}
