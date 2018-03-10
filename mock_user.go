package graylog

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// HasUser
func (ms *MockServer) HasUser(username string) bool {
	_, ok := ms.users[username]
	return ok
}

// GetUser
func (ms *MockServer) GetUser(username string) (User, bool) {
	s, ok := ms.users[username]
	return s, ok
}

func (ms *MockServer) checkUserRoles(roles []string) (int, error) {
	if roles != nil && len(roles) != 0 {
		if ms.roles == nil || len(ms.roles) == 0 {
			// unfortunately, graylog 2.4.3-1 returns 500 error
			return 500, fmt.Errorf("No role found with name %s", roles[0])
		}
		for _, roleName := range roles {
			if !ms.HasRole(roleName) {
				// unfortunately, graylog 2.4.3-1 returns 500 error
				return 500, fmt.Errorf("No role found with name %s", roleName)
			}
		}
	}
	return 200, nil
}

// AddUser adds a user to the MockServer.
func (ms *MockServer) AddUser(user *User) (*User, int, error) {
	if err := CreateValidator.Struct(user); err != nil {
		return nil, 400, err
	}
	if ms.HasUser(user.Username) {
		return nil, 400, fmt.Errorf(
			"The user %s has already existed.", user.Username)
	}
	s := *user

	// check role exists
	if sc, err := ms.checkUserRoles(user.Roles); err != nil {
		return nil, sc, err
	}
	s.Id = randStringBytesMaskImprSrc(24)
	ms.users[s.Username] = s
	return &s, 200, nil
}

// UpdateUser updates a user of the MockServer.
// "email", "permissions", "full_name", "password"
func (ms *MockServer) UpdateUser(user *User) (int, error) {
	u, ok := ms.GetUser(user.Username)
	if !ok {
		return 404, fmt.Errorf("The user is not found")
	}
	if err := UpdateValidator.Struct(user); err != nil {
		return 400, err
	}
	// check role exists
	if sc, err := ms.checkUserRoles(user.Roles); err != nil {
		return sc, err
	}
	if user.Email != "" {
		u.Email = user.Email
	}
	if user.Permissions != nil {
		u.Permissions = user.Permissions
	}
	if user.FullName != "" {
		u.FullName = user.FullName
	}
	if user.Password != "" {
		u.Password = user.Password
	}
	ms.users[u.Username] = u
	return 200, nil
}

// DeleteUser removes a user from the MockServer.
func (ms *MockServer) DeleteUser(name string) (int, error) {
	if !ms.HasUser(name) {
		return 404, fmt.Errorf("The user is not found")
	}
	delete(ms.users, name)
	return 200, nil
}

// UserList returns a list of all users.
func (ms *MockServer) UserList() []User {
	size := len(ms.users)
	arr := make([]User, size)
	i := 0
	for _, user := range ms.users {
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
		"startpage", "timezone", "session_timeout_ms", "roles"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}

	user := &User{}
	if err := msDecode(body, user); err != nil {
		// if err := decoder.Decode(body); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as User")
		writeApiError(w, 400, "400 Bad Request")
		return
	}

	if _, sc, err := ms.AddUser(user); err != nil {
		writeApiError(w, sc, err.Error())
	}
	ms.safeSave()
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
	user, ok := ms.GetUser(name)
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
	user, ok := ms.GetUser(name)
	if !ok {
		writeApiError(w, 404, "No user found with name %s", name)
		return
	}

	// required fields is nil
	acceptedFields := []string{
		"email", "permissions", "full_name", "password"}
	sc, msg, body := validateRequestBody(b, nil, nil, acceptedFields)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}
	if err := msDecode(body, &user); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as User")
		writeApiError(w, 400, "400 Bad Request")
		return
	}

	if sc, err := ms.UpdateUser(&user); err != nil {
		writeApiError(w, sc, err.Error())
		return
	}
	ms.safeSave()
}

// DELETE /users/{username} Removes a user account
func (ms *MockServer) handleDeleteUser(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	name := ps.ByName("username")
	if sc, err := ms.DeleteUser(name); err != nil {
		writeApiError(w, sc, err.Error())
		return
	}
	ms.safeSave()
}
