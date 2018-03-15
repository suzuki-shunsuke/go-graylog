package mockserver

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasUser
func (ms *MockServer) HasUser(username string) (bool, error) {
	return ms.store.HasUser(username)
}

// GetUser returns a user.
func (ms *MockServer) GetUser(username string) (*graylog.User, error) {
	return ms.store.GetUser(username)
}

// UsersList returns a user.
func (ms *MockServer) UsersList() ([]graylog.User, error) {
	return ms.store.GetUsers()
}

func (ms *MockServer) checkUserRoles(roles []string) (int, error) {
	if roles != nil && len(roles) != 0 {
		for _, roleName := range roles {
			ok, err := ms.HasRole(roleName)
			if err != nil {
				return 500, err
			}
			if !ok {
				// unfortunately, graylog 2.4.3-1 returns 500 error
				return 500, fmt.Errorf("No role found with name %s", roleName)
			}
		}
	}
	return 200, nil
}

// AddUser adds a user to the MockServer.
func (ms *MockServer) AddUser(user *graylog.User) (*graylog.User, int, error) {
	// client side validation
	if err := validator.CreateValidator.Struct(user); err != nil {
		return nil, 400, err
	}

	// Check a given username has already used.
	ok, err := ms.HasUser(user.Username)
	if err != nil {
		return nil, 500, err
	}
	if ok {
		return nil, 400, fmt.Errorf(
			"The user %s has already existed.", user.Username)
	}
	s := *user

	// check role exists
	if sc, err := ms.checkUserRoles(user.Roles); err != nil {
		return nil, sc, err
	}

	// generate ID
	s.ID = randStringBytesMaskImprSrc(24)

	// Add a user
	u, err := ms.store.AddUser(&s)
	if err != nil {
		return u, 500, err
	}
	return u, 200, nil
}

// UpdateUser updates a user of the MockServer.
// "email", "permissions", "full_name", "password"
func (ms *MockServer) UpdateUser(user *graylog.User) (int, error) {
	// Check updated user exists
	ok, err := ms.HasUser(user.Username)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("The user is not found")
	}

	// client side validation
	if err := validator.UpdateValidator.Struct(user); err != nil {
		return 400, err
	}

	// check role exists
	if sc, err := ms.checkUserRoles(user.Roles); err != nil {
		return sc, err
	}

	// update
	if err := ms.store.UpdateUser(user); err != nil {
		return 500, err
	}
	return 200, nil
}

// DeleteUser removes a user from the MockServer.
func (ms *MockServer) DeleteUser(name string) (int, error) {
	// Check deleted user exists
	ok, err := ms.HasUser(name)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("The user is not found")
	}

	// Delete a user
	if err := ms.store.DeleteUser(name); err != nil {
		return 500, err
	}
	return 200, nil
}

// UserList returns a list of all users.
func (ms *MockServer) UserList() ([]graylog.User, error) {
	return ms.store.GetUsers()
}

// POST /users Create a new user account.
func (ms *MockServer) handleCreateUser(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		return 500, nil, err
	}

	requiredFields := []string{
		"username", "email", "permissions", "full_name", "password"}
	allowedFields := []string{
		"startpage", "timezone", "session_timeout_ms", "roles"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		return sc, nil, fmt.Errorf(msg)
	}

	user := &graylog.User{}
	if err := msDecode(body, user); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as User")
		return 400, nil, err
	}

	if _, sc, err := ms.AddUser(user); err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	return 200, nil, nil
}

// GET /users List all users
func (ms *MockServer) handleGetUsers(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	ms.handleInit(w, r, false)
	arr, err := ms.UserList()
	if err != nil {
		return 500, nil, err
	}
	users := &graylog.UsersBody{Users: arr}
	return 200, users, nil
}

// GET /users/{username} Get user details
func (ms *MockServer) handleGetUser(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	ms.handleInit(w, r, false)
	name := ps.ByName("username")
	user, err := ms.GetUser(name)
	if err != nil {
		return 500, nil, err
	}
	if user == nil {
		return 404, nil, fmt.Errorf("No user found with name %s", name)
	}
	return 200, user, nil
}

// PUT /users/{username} Modify user details.
func (ms *MockServer) handleUpdateUser(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		return 500, nil, err
	}
	name := ps.ByName("username")

	user, err := ms.GetUser(name)
	if err != nil {
		return 500, nil, err
	}
	if user == nil {
		return 404, nil, fmt.Errorf("No user found with name %s", name)
	}

	// required fields is nil
	acceptedFields := []string{
		"email", "permissions", "full_name", "password"}
	sc, msg, body := validateRequestBody(b, nil, nil, acceptedFields)
	if sc != 200 {
		return sc, nil, fmt.Errorf(msg)
	}
	if err := msDecode(body, &user); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as User")
		return 400, nil, err
	}

	if sc, err := ms.UpdateUser(user); err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	return 200, nil, nil
}

// DELETE /users/{username} Removes a user account
func (ms *MockServer) handleDeleteUser(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	ms.handleInit(w, r, false)
	name := ps.ByName("username")
	if sc, err := ms.DeleteUser(name); err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	return 200, nil, nil
}
