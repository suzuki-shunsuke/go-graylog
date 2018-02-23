package graylog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
)

func (ms *MockServer) handleRoleMember(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		ms.handleAddUserToRole(w, r)
	case http.MethodDelete:
		ms.handleRemoveUserFromRole(w, r)
	}
}

type membersBody struct {
	Role  string `json:"role"`
	Users []User `json:"users"`
}

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
func (ms *MockServer) handleRoleMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := r.URL.Path
	name := path.Base(p[:len(p)-len("/members")])
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
func (ms *MockServer) handleAddUserToRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ps := strings.Split(r.URL.Path, "/")
	s := len(ps)
	roleName := ps[s-3]
	userName := ps[s-1]
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
	for _, rn := range user.Roles {
		if rn == roleName {
			return
		}
	}
	user.Roles = append(user.Roles, roleName)
}

// DELETE /roles/{rolename}/members/{username} Remove a user from a role
func (ms *MockServer) handleRemoveUserFromRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ps := strings.Split(r.URL.Path, "/")
	s := len(ps)
	roleName := ps[s-3]
	userName := ps[s-1]
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
	roles := []string{}
	for _, rn := range user.Roles {
		if rn != roleName {
			roles = append(roles, rn)
		}
	}
	user.Roles = roles
	ms.Users[userName] = user
}
