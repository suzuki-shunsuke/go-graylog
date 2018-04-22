package logic

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-set"
)

// RoleMembers returns members of a given role.
func (ms *Logic) RoleMembers(name string) ([]graylog.User, int, error) {
	ok, err := ms.HasRole(name)
	if err != nil {
		return nil, 500, err
	}
	if !ok {
		return nil, 404, fmt.Errorf("no role found with name %s", name)
	}
	users := []graylog.User{}
	us, sc, err := ms.GetUsers()
	if err != nil {
		return us, sc, err
	}
	for _, user := range us {
		if user.Roles == nil {
			continue
		}
		for roleName := range user.Roles.ToMap(false) {
			if roleName == name {
				users = append(users, user)
				break
			}
		}
	}
	return users, 200, nil
}

// AddUserToRole adds a user to a role.
func (ms *Logic) AddUserToRole(userName, roleName string) (int, error) {
	ok, err := ms.HasRole(roleName)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no role found with name %s", roleName)
	}

	if userName == "admin" {
		return 500, fmt.Errorf("cannot modify local root user, this is a bug")
	}
	user, sc, err := ms.GetUser(userName)
	if err != nil {
		return sc, err
	}
	if user == nil {
		return 404, fmt.Errorf("no user found with name %s", userName)
	}
	if user.Roles == nil {
		user.Roles = set.NewStrSet(roleName)
	} else {
		user.Roles.Add(roleName)
	}
	return ms.UpdateUser(user.NewUpdateParams())
}

// RemoveUserFromRole removes a user from a role.
func (ms *Logic) RemoveUserFromRole(
	userName, roleName string,
) (int, error) {
	ok, err := ms.HasRole(roleName)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf(`no role found with name "%s"`, roleName)
	}

	if userName == "admin" {
		return 500, fmt.Errorf("cannot modify local root user, this is a bug")
	}
	user, sc, err := ms.GetUser(userName)
	if err != nil {
		return sc, err
	}
	if user == nil {
		return 404, fmt.Errorf("no user found with name %s", userName)
	}
	if user.Roles != nil {
		user.Roles.Remove(roleName)
	}
	return ms.UpdateUser(user.NewUpdateParams())
}
