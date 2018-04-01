package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

type membersBody struct {
	Role  string         `json:"role"`
	Users []graylog.User `json:"users"`
}

// HandleRoleMembers
func HandleRoleMembers(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// GET /roles/{rolename}/members Retrieve the role's members
	name := ps.ByName("rolename")
	ok, err := ms.HasRole(name)
	if err != nil {
		return nil, 500, err
	}
	if !ok {
		return nil, 404, fmt.Errorf("no role found with name %s", name)
	}
	arr, sc, err := ms.RoleMembers(name)
	if err != nil {
		return nil, sc, err
	}
	users := &membersBody{Users: arr, Role: name}
	return users, sc, nil
}

// HandleAddUserToRole
func HandleAddUserToRole(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// PUT /roles/{rolename}/members/{username} Add a user to a role
	sc, err := ms.AddUserToRole(ps.ByName("username"), ps.ByName("rolename"))
	return nil, sc, err
}

// HandleRemoveUserFromRole
func HandleRemoveUserFromRole(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// DELETE /roles/{rolename}/members/{username} Remove a user from a role
	sc, err := ms.RemoveUserFromRole(ps.ByName("username"), ps.ByName("rolename"))
	return nil, sc, err
}
