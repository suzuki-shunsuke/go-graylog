package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-set"
)

// GET /roles/{rolename} Retrieve permissions for a single role
func HandleGetRole(
	user *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	name := ps.ByName("rolename")
	ms.Logger().WithFields(log.Fields{
		"handler": "handleGetRole", "rolename": name}).Info("request start")
	if sc, err := ms.Authorize(user, "roles:read", name); err != nil {
		return sc, nil, err
	}
	role, sc, err := ms.GetRole(name)
	return sc, role, err
}

// PUT /roles/{rolename} Update an existing role
func HandleUpdateRole(
	user *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	name := ps.ByName("rolename")
	if sc, err := ms.Authorize(user, "roles:edit", name); err != nil {
		return sc, nil, err
	}
	requiredFields := set.NewStrSet("name", "permissions")
	allowedFields := set.NewStrSet("description", "read_only")
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
	user *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	name := ps.ByName("rolename")
	if sc, err := ms.Authorize(user, "roles:delete", name); err != nil {
		return sc, nil, err
	}
	sc, err := ms.DeleteRole(name)
	if err != nil {
		return sc, nil, err
	}
	ms.SafeSave()
	return 204, nil, nil
}

// POST /roles Create a new role
func HandleCreateRole(
	user *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	if sc, err := ms.Authorize(user, "roles:create"); err != nil {
		return sc, nil, err
	}
	requiredFields := set.NewStrSet("name", "permissions")
	allowedFields := set.NewStrSet("description", "read_only")
	body, sc, err := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if err != nil {
		return sc, nil, err
	}

	role := &graylog.Role{}
	if err := msDecode(body, &role); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Warn("Failed to parse request body as Role")
		return 400, nil, err
	}

	if sc, err := ms.AddRole(role); err != nil {
		return sc, nil, err
	}
	ms.SafeSave()
	return sc, role, nil
}

// GET /roles List all roles
func HandleGetRoles(
	user *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	arr, sc, err := ms.GetRoles()
	if err != nil {
		return sc, nil, err
	}
	return sc, &graylog.RolesBody{Roles: arr, Total: len(arr)}, nil
}
