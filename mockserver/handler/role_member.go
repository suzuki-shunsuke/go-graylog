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

// HandleRoleMembers is the handler of Get the role's members API.
func HandleRoleMembers(
	user *graylog.User, lgc *logic.Logic,
	r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// GET /roles/{rolename}/members Retrieve the role's members
	name := ps.ByName("rolename")
	ok, err := lgc.HasRole(name)
	if err != nil {
		return nil, 500, err
	}
	if !ok {
		return nil, 404, fmt.Errorf("no role found with name %s", name)
	}
	arr, sc, err := lgc.RoleMembers(name)
	if err != nil {
		return nil, sc, err
	}
	users := &membersBody{Users: arr, Role: name}
	return users, sc, nil
}

// HandleAddUserToRole is the handler of Add a user to a role API.
func HandleAddUserToRole(
	user *graylog.User, lgc *logic.Logic,
	r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// PUT /roles/{rolename}/members/{username} Add a user to a role
	sc, err := lgc.AddUserToRole(ps.ByName("username"), ps.ByName("rolename"))
	return nil, sc, err
}

// HandleRemoveUserFromRole is the handler of Remove a user from a role API.
func HandleRemoveUserFromRole(
	user *graylog.User, lgc *logic.Logic,
	r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// DELETE /roles/{rolename}/members/{username} Remove a user from a role
	sc, err := lgc.RemoveUserFromRole(ps.ByName("username"), ps.ByName("rolename"))
	return nil, sc, err
}
