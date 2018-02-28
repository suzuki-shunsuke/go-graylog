package graylog

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type membersBody struct {
	Role  string `json:"role"`
	Users []User `json:"users"`
}

// RoleMembers returns members of a given role.
func (ms *MockServer) RoleMembers(name string) []User {
	users := []User{}
	for _, user := range ms.Users {
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
	return users
}

// GET /roles/{rolename}/members Retrieve the role's members
func (ms *MockServer) handleRoleMembers(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	name := ps.ByName("rolename")
	arr := ms.RoleMembers(name)
	users := membersBody{Users: arr, Role: name}
	b, err := json.Marshal(&users)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

// PUT /roles/{rolename}/members/{username} Add a user to a role
func (ms *MockServer) handleAddUserToRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	roleName := ps.ByName("rolename")
	userName := ps.ByName("username")
	if _, ok := ms.Roles[roleName]; !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "No role found with name %s"}`,
			roleName)))
		return
	}
	user, ok := ms.Users[userName]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "User %s has not been found."}`,
			userName)))
		return
	}
	user.Roles = addToStringArray(user.Roles, roleName)
	ms.AddUser(&user)
}

// DELETE /roles/{rolename}/members/{username} Remove a user from a role
func (ms *MockServer) handleRemoveUserFromRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	roleName := ps.ByName("rolename")
	userName := ps.ByName("username")
	if _, ok := ms.Roles[roleName]; !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "No role found with name %s"}`,
			roleName)))
		return
	}
	user, ok := ms.Users[userName]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "User %s has not been found."}`,
			userName)))
		return
	}
	user.Roles = removeFromStringArray(user.Roles, roleName)
	ms.AddUser(&user)
}
