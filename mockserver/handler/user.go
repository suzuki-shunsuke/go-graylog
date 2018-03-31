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

// POST /users Create a new user account.
func HandleCreateUser(
	u *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	if sc, err := ms.Authorize(u, "users:create"); err != nil {
		return sc, nil, err
	}
	requiredFields := set.NewStrSet(
		"username", "email", "permissions", "full_name", "password")
	allowedFields := set.NewStrSet(
		"startpage", "timezone", "session_timeout_ms", "roles")
	body, sc, err := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if err != nil {
		return sc, nil, err
	}

	user := &graylog.User{}
	if err := msDecode(body, user); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as User")
		return 400, nil, err
	}

	if sc, err := ms.AddUser(user); err != nil {
		fmt.Println(sc, err)
		return sc, nil, err
	}
	if err := ms.Save(); err != nil {
		return 500, nil, err
	}
	return 201, nil, nil
}

// GET /users List all users
func HandleGetUsers(
	user *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	users, sc, err := ms.GetUsers()
	if err != nil {
		return sc, users, err
	}
	return sc, &graylog.UsersBody{Users: users}, nil
}

// GET /users/{username} Get user details
func HandleGetUser(
	u *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	name := ps.ByName("username")
	user, sc, err := ms.GetUser(name)
	return sc, user, err
}

// PUT /users/{username} Modify user details.
func HandleUpdateUser(
	u *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	userName := ps.ByName("username")
	if sc, err := ms.Authorize(u, "users:edit", userName); err != nil {
		return sc, nil, err
	}
	// required fields is nil
	acceptedFields := set.NewStrSet(
		"email", "permissions", "full_name", "password")
	body, sc, err := validateRequestBody(r.Body, nil, nil, acceptedFields)
	if err != nil {
		return sc, nil, err
	}

	user := &graylog.User{Username: userName}
	if err := msDecode(body, user); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as User")
		return 400, nil, err
	}

	if sc, err := ms.UpdateUser(user); err != nil {
		return sc, nil, err
	}
	if err := ms.Save(); err != nil {
		return 500, nil, err
	}
	return 200, nil, nil
}

// DELETE /users/{username} Removes a user account
func HandleDeleteUser(
	user *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	name := ps.ByName("username")
	if sc, err := ms.DeleteUser(name); err != nil {
		return sc, nil, err
	}
	if err := ms.Save(); err != nil {
		return 500, nil, err
	}
	return 204, nil, nil
}
