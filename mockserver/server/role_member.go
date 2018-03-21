package server

import (
	"github.com/suzuki-shunsuke/go-graylog"
)

// RoleMembers returns members of a given role.
func (ms *Server) RoleMembers(name string) ([]graylog.User, error) {
	users := []graylog.User{}
	us, err := ms.UsersList()
	if err != nil {
		return nil, err
	}
	for _, user := range us {
		if user.Roles == nil {
			continue
		}
		for _, roleName := range user.Roles {
			if roleName == name {
				users = append(users, user)
				break
			}
		}
	}
	return users, nil
}
