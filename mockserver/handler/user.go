package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-set"
)

// HandleCreateUser
func HandleCreateUser(
	u *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// POST /users Create a new user account.
	if sc, err := ms.Authorize(u, "users:create"); err != nil {
		return nil, sc, err
	}
	requiredFields := set.NewStrSet(
		"username", "email", "permissions", "full_name", "password")
	allowedFields := set.NewStrSet(
		"startpage", "timezone", "session_timeout_ms", "roles")
	body, sc, err := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if err != nil {
		return nil, sc, err
	}

	user := &graylog.User{}
	if err := msDecode(body, user); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as User")
		return nil, 400, err
	}

	if sc, err := ms.AddUser(user); err != nil {
		fmt.Println(sc, err)
		return nil, sc, err
	}
	if err := ms.Save(); err != nil {
		return nil, 500, err
	}
	return nil, 201, nil
}

// HandleGetUsers
func HandleGetUsers(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// GET /users List all users
	users, sc, err := ms.GetUsers()
	if err != nil {
		return nil, sc, err
	}
	// TODO authorization
	return &graylog.UsersBody{Users: users}, sc, nil
}

// HandleGetUser
func HandleGetUser(
	u *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// GET /users/{username} Get user details
	name := ps.ByName("username")
	// TODO authorization
	return ms.GetUser(name)
}

// HandleUpdateUser
func HandleUpdateUser(
	u *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// PUT /users/{username} Modify user details.
	userName := ps.ByName("username")
	if sc, err := ms.Authorize(u, "users:edit", userName); err != nil {
		return nil, sc, err
	}
	// required fields is nil
	acceptedFields := set.NewStrSet(
		"email", "permissions", "full_name", "password")
	body, sc, err := validateRequestBody(r.Body, nil, nil, acceptedFields)
	if err != nil {
		return nil, sc, err
	}

	user := &graylog.User{Username: userName}
	if err := msDecode(body, user); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as User")
		return nil, 400, err
	}

	if sc, err := ms.UpdateUser(user); err != nil {
		return nil, sc, err
	}
	if err := ms.Save(); err != nil {
		return nil, 500, err
	}
	return nil, 200, nil
}

// HandleDeleteUser
func HandleDeleteUser(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// DELETE /users/{username} Removes a user account
	name := ps.ByName("username")
	// TODO authorization
	if sc, err := ms.DeleteUser(name); err != nil {
		return nil, sc, err
	}
	if err := ms.Save(); err != nil {
		return nil, 500, err
	}
	return nil, 204, nil
}
