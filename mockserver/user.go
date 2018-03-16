package mockserver

import (
	"fmt"

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
func (ms *MockServer) AddUser(user *graylog.User) (int, error) {
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
			"The user %s has already existed.", user.Username)
	}

	// check role exists
	if sc, err := ms.checkUserRoles(user.Roles); err != nil {
		return sc, err
	}

	// generate ID
	user.ID = randStringBytesMaskImprSrc(24)

	// Add a user
	if err := ms.store.AddUser(user); err != nil {
		return 500, err
	}
	return 200, nil
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
