package graylog

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// AddUser adds a user to the MockServer.
// If ms.DataPath != "", the data is written in a file for persistence.
func (ms *MockServer) AddUser(user *User) {
	ms.Users[user.Username] = *user
	ms.safeSave()
}

// UpdateUser updates a user of the MockServer.
// If ms.DataPath != "", the data is written in a file for persistence.
func (ms *MockServer) UpdateUser(name string, user *User) {
	delete(ms.Users, name)
	ms.AddUser(user)
}

// DeleteUser removes a user from the MockServer.
// If ms.DataPath != "", the data is written in a file for persistence.
func (ms *MockServer) DeleteUser(name string) {
	delete(ms.Users, name)
	ms.safeSave()
}

// UserList returns a list of all users.
func (ms *MockServer) UserList() []User {
	if ms.Users == nil {
		return []User{}
	}
	size := len(ms.Users)
	arr := make([]User, size)
	i := 0
	for _, user := range ms.Users {
		arr[i] = user
		i++
	}
	return arr
}

// POST /users Create a new user account.
func (ms *MockServer) handleCreateUser(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		write500Error(w)
		return
	}

	requiredFields := []string{
		"username", "email", "permissions", "full_name", "password"}
	allowedFields := []string{
		"startpage", "permissions", "username", "timezone", "password", "email",
		"session_timeout_ms", "full_name", "roles"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}

	user := &User{}
	if err := msDecode(body, user); err != nil {
		// if err := decoder.Decode(body); err != nil {
		ms.Logger.WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as User")
		writeApiError(w, 400, "400 Bad Request")
		return
	}

	if err := UpdateValidator.Struct(user); err != nil {
		writeApiError(w, 400, err.Error())
		return
	}

	if _, ok := ms.Users[user.Username]; ok {
		writeApiError(w, 400, "User %s already exists.", user.Username)
		return
	}
	ms.AddUser(user)
}

// GET /users List all users
func (ms *MockServer) handleGetUsers(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	ms.handleInit(w, r, false)
	arr := ms.UserList()
	users := &usersBody{Users: arr}
	writeOr500Error(w, users)
}

// GET /users/{username} Get user details
func (ms *MockServer) handleGetUser(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	name := ps.ByName("username")
	user, ok := ms.Users[name]
	if !ok {
		writeApiError(w, 404, "No user found with name %s", name)
		return
	}
	writeOr500Error(w, &user)
}

// PUT /users/{username} Modify user details.
func (ms *MockServer) handleUpdateUser(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		write500Error(w)
		return
	}
	name := ps.ByName("username")
	user, ok := ms.Users[name]
	if !ok {
		writeApiError(w, 404, "No user found with name %s", name)
		return
	}

	acceptedFields := []string{
		"email", "permissions", "full_name", "password"}
	sc, msg, body := validateRequestBody(b, nil, nil, acceptedFields)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}
	if err := msDecode(body, &user); err != nil {
		ms.Logger.WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as User")
		writeApiError(w, 400, "400 Bad Request")
		return
	}
	if err := UpdateValidator.Struct(&user); err != nil {
		writeApiError(w, 400, err.Error())
		return
	}

	ms.UpdateUser(name, &user)
}

// DELETE /users/{username} Removes a user account
func (ms *MockServer) handleDeleteUser(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	name := ps.ByName("username")
	_, ok := ms.Users[name]
	if !ok {
		writeApiError(w, 404, "No user found with name %s", name)
		return
	}
	ms.DeleteUser(name)
}
