package graylog

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type membersBody struct {
	Role  string `json:"role"`
	Users []User `json:"users"`
}

// RoleMembers returns members of a given role.
func (ms *MockServer) RoleMembers(name string) ([]User, error) {
	users := []User{}
	us, err := ms.UsersList()
	if err != nil {
		return nil, err
	}
	for _, user := range us {
		if user.Roles == nil {
			continue
		}
		for _, roleName := range user.Roles {
			if roleName == name {
				users = append(users, user)
				break
			}
		}
	}
	return users, nil
}

// GET /roles/{rolename}/members Retrieve the role's members
func (ms *MockServer) handleRoleMembers(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	name := ps.ByName("rolename")
	ok, err := ms.HasRole(name)
	if err != nil {
		writeApiError(w, 500, err.Error())
		return
	}
	if !ok {
		writeApiError(w, 404, "No role found with name %s", name)
		return
	}
	arr, err := ms.RoleMembers(name)
	if err != nil {
		writeApiError(w, 500, err.Error())
		return
	}
	users := &membersBody{Users: arr, Role: name}
	writeOr500Error(w, users)
}

// PUT /roles/{rolename}/members/{username} Add a user to a role
func (ms *MockServer) handleAddUserToRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	roleName := ps.ByName("rolename")
	userName := ps.ByName("username")
	ok, err := ms.HasRole(roleName)
	if err != nil {
		writeApiError(w, 500, err.Error())
		return
	}
	if !ok {
		writeApiError(w, 404, "No role found with name %s", roleName)
		return
	}

	user, err := ms.GetUser(userName)
	if err != nil {
		writeApiError(w, 500, err.Error())
		return
	}
	if user == nil {
		writeApiError(w, 404, "User %s has not been found.", userName)
		return
	}
	user.Roles = addToStringArray(user.Roles, roleName)
	ms.AddUser(user)
}

// DELETE /roles/{rolename}/members/{username} Remove a user from a role
func (ms *MockServer) handleRemoveUserFromRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	roleName := ps.ByName("rolename")
	userName := ps.ByName("username")
	ok, err := ms.HasRole(roleName)
	if err != nil {
		writeApiError(w, 500, err.Error())
		return
	}
	if !ok {
		writeApiError(w, 404, "No role found with name %s", roleName)
		return
	}

	user, err := ms.GetUser(userName)
	if err != nil {
		writeApiError(w, 500, err.Error())
		return
	}
	if user == nil {
		writeApiError(w, 404, "User %s has not been found.", userName)
		return
	}
	user.Roles = removeFromStringArray(user.Roles, roleName)
	if sc, err := ms.UpdateUser(user); err != nil {
		writeApiError(w, sc, err.Error())
		return
	}
}
