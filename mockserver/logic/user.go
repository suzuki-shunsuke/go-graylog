package logic

import (
	"crypto/md5"
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
	"github.com/suzuki-shunsuke/go-ptr"
)

func encryptPassword(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

// HasUser returns whether the user exists.
func (ms *Logic) HasUser(username string) (bool, error) {
	return ms.store.HasUser(username)
}

// GetUser returns a user.
func (ms *Logic) GetUser(username string) (*graylog.User, int, error) {
	user, err := ms.store.GetUser(username)
	if err != nil {
		return user, 500, err
	}
	if user == nil {
		return user, 404, fmt.Errorf(`no user "%s" is found`, username)
	}
	return user, 200, nil
}

// GetUsers returns a list of users.
func (ms *Logic) GetUsers() ([]graylog.User, int, error) {
	users, err := ms.store.GetUsers()
	if err != nil {
		return users, 500, err
	}
	return users, 200, nil
}

func (ms *Logic) checkUserRoles(roles []string) (int, error) {
	if len(roles) != 0 {
		for _, roleName := range roles {
			ok, err := ms.HasRole(roleName)
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
func (ms *Logic) AddUser(user *graylog.User) (int, error) {
	// client side validation
	if err := validator.CreateValidator.Struct(user); err != nil {
		return 400, err
	}

	// Check a given username has already used.
	ok, err := ms.HasUser(user.Username)
	if err != nil {
		return 500, err
	}
	if ok {
		return 400, fmt.Errorf(
			`the user "%s" has already existed`, user.Username)
	}

	// check role exists
	if user.Roles != nil {
		if sc, err := ms.checkUserRoles(user.Roles.ToList()); err != nil {
			return sc, err
		}
	}

	user.SetDefaultValues()
	user.Password = encryptPassword(user.Password)
	// Add a user
	if err := ms.store.AddUser(user); err != nil {
		return 500, err
	}
	return 200, nil
}

// UpdateUser updates a user of the Server.
// "email", "permissions", "full_name", "password"
func (ms *Logic) UpdateUser(prms *graylog.UserUpdateParams) (int, error) {
	if prms == nil {
		return 400, fmt.Errorf("user is nil")
	}
	// Check updated user exists
	ok, err := ms.HasUser(prms.Username)
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
		if sc, err := ms.checkUserRoles(prms.Roles.ToList()); err != nil {
			return sc, err
		}
	}
	if prms.Password != nil {
		prms.Password = ptr.PStr(encryptPassword(*prms.Password))
	}

	// update
	if err := ms.store.UpdateUser(prms); err != nil {
		return 500, err
	}
	return 200, nil
}

// DeleteUser removes a user from the Server.
func (ms *Logic) DeleteUser(name string) (int, error) {
	// Check deleted user exists
	ok, err := ms.HasUser(name)
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
	if err := ms.store.DeleteUser(name); err != nil {
		return 500, err
	}
	return 200, nil
}

// UserList returns a list of all users.
func (ms *Logic) UserList() ([]graylog.User, error) {
	return ms.store.GetUsers()
}
