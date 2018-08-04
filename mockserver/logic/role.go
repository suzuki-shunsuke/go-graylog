package logic

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasRole returns whether the role with given name exists.
func (lgc *Logic) HasRole(name string) (bool, error) {
	return lgc.store.HasRole(name)
}

// GetRole returns a Role.
// If a role is not found, an error is returns.
func (lgc *Logic) GetRole(name string) (*graylog.Role, int, error) {
	role, err := lgc.store.GetRole(name)
	if err != nil {
		return role, 500, err
	}
	if role == nil {
		return nil, 404, fmt.Errorf(`no role with name "%s"`, name)
	}
	return role, 200, nil
}

// AddRole adds a new role.
func (lgc *Logic) AddRole(role *graylog.Role) (int, error) {
	if err := validator.CreateValidator.Struct(role); err != nil {
		return 400, err
	}
	ok, err := lgc.HasRole(role.Name)
	if err != nil {
		return 500, err
	}
	if ok {
		return 400, fmt.Errorf("role %s already exists", role.Name)
	}
	if role.Name != "Admin" && role.Name != "Reader" {
		role.ReadOnly = false
	}
	if err := lgc.store.AddRole(role); err != nil {
		return 500, err
	}
	return 200, nil
}

// UpdateRole updates a role.
func (lgc *Logic) UpdateRole(name string, prms *graylog.RoleUpdateParams) (*graylog.Role, int, error) {
	if err := validator.UpdateValidator.Struct(prms); err != nil {
		return nil, 400, err
	}
	role, sc, err := lgc.GetRole(name)
	if err != nil {
		return nil, sc, err
	}
	if name != prms.Name {
		ok, err := lgc.HasRole(prms.Name)
		if err != nil {
			return nil, 500, err
		}
		if ok {
			return nil, 400, fmt.Errorf("the role %s has already existed", prms.Name)
		}
	}
	if role.ReadOnly {
		return nil, 400, fmt.Errorf("cannot update read only role %s", role.Name)
	}
	role, err = lgc.store.UpdateRole(name, prms)
	if err != nil {
		return nil, 500, err
	}
	return role, 200, nil
}

// DeleteRole deletes a role.
func (lgc *Logic) DeleteRole(name string) (int, error) {
	role, sc, err := lgc.GetRole(name)
	if err != nil {
		return sc, err
	}
	if role.ReadOnly {
		return 400, fmt.Errorf("cannot delete read only role %s", name)
	}
	if err := lgc.store.DeleteRole(name); err != nil {
		return 500, err
	}
	return 200, nil
}

// GetRoles returns a list of roles.
func (lgc *Logic) GetRoles() ([]graylog.Role, int, int, error) {
	roles, total, err := lgc.store.GetRoles()
	if err != nil {
		return nil, 0, 500, err
	}
	return roles, total, 200, nil
}
