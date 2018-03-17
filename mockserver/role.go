package mockserver

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasRole
func (ms *MockServer) HasRole(name string) (bool, error) {
	return ms.store.HasRole(name)
}

// GetRole returns a Role.
func (ms *MockServer) GetRole(name string) (*graylog.Role, error) {
	return ms.store.GetRole(name)
}

// AddRole adds a new role to the mock server.
func (ms *MockServer) AddRole(role *graylog.Role) (int, error) {
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
	if err := ms.store.AddRole(role); err != nil {
		return 500, err
	}
	return 200, nil
}

// UpdateRole updates a role.
func (ms *MockServer) UpdateRole(name string, role *graylog.Role) (int, error) {
	if err := validator.UpdateValidator.Struct(role); err != nil {
		return 400, err
	}
	ok, err := ms.HasRole(name)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("No role found with name %s", name)
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
	if err := ms.store.UpdateRole(name, role); err != nil {
		return 500, err
	}
	return 204, nil
}

// DeleteRole
func (ms *MockServer) DeleteRole(name string) (int, error) {
	ok, err := ms.HasRole(name)
	if err != nil {
		return 500, err
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
func (ms *MockServer) GetRoles() ([]graylog.Role, error) {
	return ms.store.GetRoles()
}
