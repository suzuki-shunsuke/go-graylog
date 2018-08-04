package logic

import (
	"crypto/md5"
	"fmt"

	"github.com/suzuki-shunsuke/go-ptr"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

func encryptPassword(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

// HasUser returns whether the user exists.
func (lgc *Logic) HasUser(username string) (bool, error) {
	return lgc.store.HasUser(username)
}

// GetUser returns a user.
func (lgc *Logic) GetUser(username string) (*graylog.User, int, error) {
	user, err := lgc.store.GetUser(username)
	if err != nil {
		return user, 500, err
	}
	if user == nil {
		return user, 404, fmt.Errorf(`no user "%s" is found`, username)
	}
	return user, 200, nil
}

// GetUsers returns a list of users.
func (lgc *Logic) GetUsers() ([]graylog.User, int, error) {
	users, err := lgc.store.GetUsers()
	if err != nil {
		return users, 500, err
	}
	return users, 200, nil
}

func (lgc *Logic) checkUserRoles(roles []string) (int, error) {
	if len(roles) != 0 {
		for _, roleName := range roles {
			ok, err := lgc.HasRole(roleName)
			if err != nil {
				return 500, err
			}
			if !ok {
				// unfortunately, graylog 2.4.3-1 returns 500 error
				// https://github.com/Graylog2/graylog2-server/issues/4665
				return 500, fmt.Errorf(`no role found with name "%s"`, roleName)
			}
		}
	}
	return 200, nil
}

// AddUser adds a user to the Server.
func (lgc *Logic) AddUser(user *graylog.User) (int, error) {
	// client side validation
	if err := validator.CreateValidator.Struct(user); err != nil {
		return 400, err
	}

	// Check a given username has already used.
	ok, err := lgc.HasUser(user.Username)
	if err != nil {
		return 500, err
	}
	if ok {
		return 400, fmt.Errorf(
			`the user "%s" has already existed`, user.Username)
	}

	// check role exists
	if user.Roles != nil {
		if sc, err := lgc.checkUserRoles(user.Roles.ToList()); err != nil {
			return sc, err
		}
	}

	user.SetDefaultValues()
	user.Password = encryptPassword(user.Password)
	// Add a user
	if err := lgc.store.AddUser(user); err != nil {
		return 500, err
	}
	return 201, nil
}

// UpdateUser updates a user of the Server.
// "email", "permissions", "full_name", "password"
func (lgc *Logic) UpdateUser(prms *graylog.UserUpdateParams) (int, error) {
	if prms == nil {
		return 400, fmt.Errorf("user is nil")
	}
	// Check updated user exists
	ok, err := lgc.HasUser(prms.Username)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf(`the user "%s" is not found`, prms.Username)
	}

	// client side validation
	if err := validator.UpdateValidator.Struct(prms); err != nil {
		return 400, err
	}

	// check role exists
	if prms.Roles != nil {
		if sc, err := lgc.checkUserRoles(prms.Roles.ToList()); err != nil {
			return sc, err
		}
	}
	if prms.Password != nil {
		prms.Password = ptr.PStr(encryptPassword(*prms.Password))
	}

	// update
	if err := lgc.store.UpdateUser(prms); err != nil {
		return 500, err
	}
	return 200, nil
}

// DeleteUser removes a user from the Server.
func (lgc *Logic) DeleteUser(name string) (int, error) {
	// Check deleted user exists
	ok, err := lgc.HasUser(name)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf(`the user "%s" is not found`, name)
	}
	if name == "admin" {
		// graylog spec
		return 404, fmt.Errorf(`the user "%s" is not found`, name)
	}
	// Delete a user
	if err := lgc.store.DeleteUser(name); err != nil {
		return 500, err
	}
	return 204, nil
}

// UserList returns a list of all users.
func (lgc *Logic) UserList() ([]graylog.User, error) {
	return lgc.store.GetUsers()
}
