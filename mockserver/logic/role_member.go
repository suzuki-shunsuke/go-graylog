package logic

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-set"
)

// RoleMembers returns members of a given role.
func (ms *Server) RoleMembers(name string) ([]graylog.User, int, error) {
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
func (ms *Server) AddUserToRole(userName, roleName string) (int, error) {
	ok, sc, err := ms.HasRole(roleName)
	if err != nil {
		return sc, err
	}
	if !ok {
		return 404, fmt.Errorf("no role found with name %s", roleName)
	}

	user, sc, err := ms.GetUser(userName)
	if err != nil {
		return sc, err
	}
	if user.Roles == nil {
		user.Roles = set.NewStrSet(roleName)
	} else {
		user.Roles.Add(roleName)
	}
	return ms.UpdateUser(user)
}

// RemoveUserFromRole removes a user from a role.
func (ms *Server) RemoveUserFromRole(
	userName, roleName string,
) (int, error) {
	ok, sc, err := ms.HasRole(roleName)
	if err != nil {
		return sc, err
	}
	if !ok {
		return 404, fmt.Errorf(`no role found with name "%s"`, roleName)
	}

	user, sc, err := ms.GetUser(userName)
	if err != nil {
		return sc, err
	}
	if user.Roles != nil {
		user.Roles.Remove(roleName)
	}
	return ms.UpdateUser(user)
}
