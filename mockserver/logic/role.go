package logic

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// authRolesRead
func (ms *Server) authRolesRead(user *graylog.User, name string) (int, error) {
	if user == nil {
		return 200, nil
	}
	ok, err := ms.store.Auth(user, "roles:read", name)
	if err != nil {
		return 500, err
	}
	if ok {
		return 200, nil
	}
	return 403, fmt.Errorf("authorization failure")
}

// HasRole returns whether the role with given name exists.
func (ms *Server) HasRole(user *graylog.User, name string) (bool, int, error) {
	sc, err := ms.authRolesRead(user, name)
	if err != nil {
		return false, sc, err
	}
	ok, err := ms.store.HasRole(name)
	if err != nil {
		return false, 500, err
	}
	return ok, 200, nil
}

// GetRole returns a Role.
// If a role is not found, an error is returns.
func (ms *Server) GetRole(name string) (*graylog.Role, int, error) {
	role, err := ms.store.GetRole(name)
	if err != nil {
		return role, 500, err
	}
	if role == nil {
		return nil, 404, fmt.Errorf(`no role with name "%s"`, name)
	}
	return role, 200, nil
}

// AddRole adds a new role to the mock server.
func (ms *Server) AddRole(role *graylog.Role) (int, error) {
	if err := validator.CreateValidator.Struct(role); err != nil {
		return 400, err
	}
	ok, sc, err := ms.HasRole(nil, role.Name)
	if err != nil {
		return sc, err
	}
	if ok {
		return 400, fmt.Errorf("Role %s already exists.", role.Name)
	}
	if err := ms.store.AddRole(role); err != nil {
		return 500, err
	}
	return 200, nil
}

// UpdateRole updates a role.
func (ms *Server) UpdateRole(name string, role *graylog.Role) (int, error) {
	if err := validator.UpdateValidator.Struct(role); err != nil {
		return 400, err
	}
	ok, sc, err := ms.HasRole(nil, name)
	if err != nil {
		return sc, err
	}
	if !ok {
		return 404, fmt.Errorf("No role found with name %s", name)
	}
	if name != role.Name {
		ok, sc, err := ms.HasRole(nil, role.Name)
		if err != nil {
			return sc, err
		}
		if ok {
			return 400, fmt.Errorf("The role %s has already existed.", role.Name)
		}
	}
	if err := ms.store.UpdateRole(name, role); err != nil {
		return 500, err
	}
	return 204, nil
}

// DeleteRole deletes a role.
func (ms *Server) DeleteRole(name string) (int, error) {
	ok, sc, err := ms.HasRole(nil, name)
	if err != nil {
		return sc, err
	}
	if !ok {
		return 404, fmt.Errorf("No role found with name %s", name)
	}
	if err := ms.store.DeleteRole(name); err != nil {
		return 500, err
	}
	return 200, nil
}

// GetRoles returns a list of roles.
func (ms *Server) GetRoles() ([]graylog.Role, int, error) {
	roles, err := ms.store.GetRoles()
	if err != nil {
		return roles, 500, err
	}
	return roles, 200, nil
}
