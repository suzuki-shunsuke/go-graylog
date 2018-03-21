package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

// GET /roles/{rolename} Retrieve permissions for a single role
func HandleGetRole(
	ms *logic.Server,
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
		return 404, nil, fmt.Errorf("no role found with name %s", name)
	}
	return 200, role, nil
}

// PUT /roles/{rolename} Update an existing role
func HandleUpdateRole(
	ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	name := ps.ByName("rolename")
	requiredFields := []string{"name", "permissions"}
	allowedFields := []string{"description", "read_only"}
	body, sc, err := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if err != nil {
		return sc, nil, err
	}

	role := &graylog.Role{}
	if err := msDecode(body, role); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as Role")
		return 400, nil, err
	}

	if sc, err := ms.UpdateRole(name, role); err != nil {
		return sc, nil, err
	}
	ms.SafeSave()
	return 204, role, nil
}

// DELETE /roles/{rolename} Remove the named role and dissociate any users from it
func HandleDeleteRole(
	ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	name := ps.ByName("rolename")
	sc, err := ms.DeleteRole(name)
	if err != nil {
		return sc, nil, err
	}
	ms.SafeSave()
	return 204, nil, nil
}

// POST /roles Create a new role
func HandleCreateRole(
	ms *logic.Server,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	requiredFields := []string{"name", "permissions"}
	allowedFields := []string{"description", "read_only"}
	body, sc, err := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if err != nil {
		return sc, nil, err
	}

	role := &graylog.Role{}
	if err := msDecode(body, &role); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as Role")
		return 400, nil, err
	}

	if sc, err := ms.AddRole(role); err != nil {
		return sc, nil, err
	}
	ms.SafeSave()
	return 201, role, nil
}

// GET /roles List all roles
func HandleGetRoles(
	ms *logic.Server,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	arr, err := ms.GetRoles()
	if err != nil {
		return 500, nil, err
	}
	roles := &graylog.RolesBody{Roles: arr, Total: len(arr)}
	return 200, roles, nil
}
