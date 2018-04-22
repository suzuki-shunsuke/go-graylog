package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/util"
	"github.com/suzuki-shunsuke/go-set"
)

// HandleGetUsers is the handler of GET Users API.
func HandleGetUsers(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// GET /users List all users
	users, sc, err := ms.GetUsers()
	for i, u := range users {
		u.Password = ""
		users[i] = u
	}
	if err != nil {
		return nil, sc, err
	}
	// TODO authorization
	return &graylog.UsersBody{Users: users}, sc, nil
}

// HandleGetUser is the handler of GET User API.
func HandleGetUser(
	u *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// GET /users/{username} Get user details
	name := ps.ByName("username")
	// TODO authorization
	user, sc, err := ms.GetUser(name)
	if user != nil {
		user.Password = ""
	}
	return user, sc, err
}

// HandleCreateUser is the handler of Create User API.
func HandleCreateUser(
	u *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// POST /users Create a new user account.
	if sc, err := ms.Authorize(u, "users:create"); err != nil {
		return nil, sc, err
	}
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required: set.NewStrSet("username", "email", "permissions", "full_name", "password"),
			Optional: set.NewStrSet("startpage", "timezone", "session_timeout_ms", "roles"),
		})
	if err != nil {
		return nil, sc, err
	}

	user := &graylog.User{}
	if err := util.MSDecode(body, user); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as User")
		return nil, 400, err
	}

	if sc, err := ms.AddUser(user); err != nil {
		return nil, sc, err
	}
	if err := ms.Save(); err != nil {
		return nil, 500, err
	}
	return nil, 201, nil
}

// HandleUpdateUser is the handler of Update User API.
func HandleUpdateUser(
	u *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// PUT /users/{username} Modify user details.
	userName := ps.ByName("username")
	if sc, err := ms.Authorize(u, "users:edit", userName); err != nil {
		return nil, sc, err
	}
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Optional:     set.NewStrSet("email", "permissions", "full_name", "password", "timezone", "session_timeout_ms", "start_page", "roles"),
			Ignored:      set.NewStrSet("id", "preferences", "external", "read_only", "session_active", "last_activity", "client_address"),
			ExtForbidden: false,
		})
	if err != nil {
		return nil, sc, err
	}

	prms := &graylog.UserUpdateParams{Username: userName}
	if err := util.MSDecode(body, prms); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as UserUpdateParams")
		return nil, 400, err
	}
	if sc, err := ms.UpdateUser(prms); err != nil {
		return nil, sc, err
	}
	if err := ms.Save(); err != nil {
		return nil, 500, err
	}
	return nil, 200, nil
}

// HandleDeleteUser is the handler of Delete User API.
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
