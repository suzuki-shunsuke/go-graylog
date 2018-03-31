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

// GET /roles/{rolename}/members Retrieve the role's members
func HandleRoleMembers(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	name := ps.ByName("rolename")
	ok, sc, err := ms.HasRole(name)
	if err != nil {
		return sc, nil, err
	}
	if !ok {
		return 404, nil, fmt.Errorf("no role found with name %s", name)
	}
	arr, sc, err := ms.RoleMembers(name)
	if err != nil {
		return sc, nil, err
	}
	users := &membersBody{Users: arr, Role: name}
	return sc, users, nil
}

// PUT /roles/{rolename}/members/{username} Add a user to a role
func HandleAddUserToRole(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	sc, err := ms.AddUserToRole(ps.ByName("username"), ps.ByName("rolename"))
	return sc, nil, err
}

// DELETE /roles/{rolename}/members/{username} Remove a user from a role
func HandleRemoveUserFromRole(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	sc, err := ms.RemoveUserFromRole(ps.ByName("username"), ps.ByName("rolename"))
	return sc, nil, err
}
