package logic

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog"
)

const (
	adminName string = "admin"
)

// RoleMembers returns members of a given role.
func (lgc *Logic) RoleMembers(name string) ([]graylog.User, int, error) {
	ok, err := lgc.HasRole(name)
	if err != nil {
		return nil, 500, err
	}
	if !ok {
		return nil, 404, fmt.Errorf("no role found with name %s", name)
	}
	users := []graylog.User{}
	us, sc, err := lgc.GetUsers()
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
func (lgc *Logic) AddUserToRole(userName, roleName string) (int, error) {
	ok, err := lgc.HasRole(roleName)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no role found with name %s", roleName)
	}

	if userName == adminName {
		return 500, fmt.Errorf("cannot modify local root user, this is a bug")
	}
	user, sc, err := lgc.GetUser(userName)
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
	return lgc.UpdateUser(user.NewUpdateParams())
}

// RemoveUserFromRole removes a user from a role.
func (lgc *Logic) RemoveUserFromRole(
	userName, roleName string,
) (int, error) {
	ok, err := lgc.HasRole(roleName)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf(`no role found with name "%s"`, roleName)
	}

	if userName == adminName {
		return 500, fmt.Errorf("cannot modify local root user, this is a bug")
	}
	user, sc, err := lgc.GetUser(userName)
	if err != nil {
		return sc, err
	}
	if user == nil {
		return 404, fmt.Errorf("no user found with name %s", userName)
	}
	if user.Roles != nil {
		user.Roles.Remove(roleName)
	}
	return lgc.UpdateUser(user.NewUpdateParams())
}
