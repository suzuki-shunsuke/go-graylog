package logic

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasRole returns whether the role with given name exists.
func (ms *Logic) HasRole(name string) (bool, error) {
	return ms.store.HasRole(name)
}

// GetRole returns a Role.
// If a role is not found, an error is returns.
func (ms *Logic) GetRole(name string) (*graylog.Role, int, error) {
	role, err := ms.store.GetRole(name)
	if err != nil {
		return role, 500, err
	}
	if role == nil {
		return nil, 404, fmt.Errorf(`no role with name "%s"`, name)
	}
	return role, 200, nil
}

// AddRole adds a new role.
func (ms *Logic) AddRole(role *graylog.Role) (int, error) {
	if err := validator.CreateValidator.Struct(role); err != nil {
		return 400, err
	}
	ok, err := ms.HasRole(role.Name)
	if err != nil {
		return 500, err
	}
	if ok {
		return 400, fmt.Errorf("Role %s already exists.", role.Name)
	}
	if role.Name != "Admin" && role.Name != "Reader" {
		role.ReadOnly = false
	}
	if err := ms.store.AddRole(role); err != nil {
		return 500, err
	}
	return 200, nil
}

// UpdateRole updates a role.
func (ms *Logic) UpdateRole(name string, role *graylog.Role) (int, error) {
	if err := validator.UpdateValidator.Struct(role); err != nil {
		return 400, err
	}
	r, sc, err := ms.GetRole(name)
	if err != nil {
		return sc, err
	}
	if name != role.Name {
		ok, err := ms.HasRole(role.Name)
		if err != nil {
			return 500, err
		}
		if ok {
			return 400, fmt.Errorf("The role %s has already existed.", role.Name)
		}
	}
	if r.ReadOnly {
		return 400, fmt.Errorf("cannot update read only role %s", role.Name)
	}
	if role.Name != "Admin" && role.Name != "Reader" {
		role.ReadOnly = false
	}
	if err := ms.store.UpdateRole(name, role); err != nil {
		return 500, err
	}
	return 204, nil
}

// DeleteRole deletes a role.
func (ms *Logic) DeleteRole(name string) (int, error) {
	role, sc, err := ms.GetRole(name)
	if err != nil {
		return sc, err
	}
	if role.ReadOnly {
		return 400, fmt.Errorf("cannot delete read only role %s", name)
	}
	if err := ms.store.DeleteRole(name); err != nil {
		return 500, err
	}
	return 200, nil
}

// GetRoles returns a list of roles.
func (ms *Logic) GetRoles() ([]graylog.Role, int, error) {
	roles, err := ms.store.GetRoles()
	if err != nil {
		return roles, 500, err
	}
	return roles, 200, nil
}
