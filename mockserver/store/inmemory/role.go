package inmemory

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

// HasRole
func (store *InMemoryStore) HasRole(name string) (bool, error) {
	_, ok := store.roles[name]
	return ok, nil
}

// GetRole returns a Role.
// If no role with given name is found, returns nil and not returns an error.
func (store *InMemoryStore) GetRole(name string) (*graylog.Role, error) {
	s, ok := store.roles[name]
	if ok {
		return &s, nil
	}
	return nil, nil
}

// GetRoles returns Roles.
func (store *InMemoryStore) GetRoles() ([]graylog.Role, error) {
	size := len(store.roles)
	arr := make([]graylog.Role, size)
	i := 0
	for _, role := range store.roles {
		arr[i] = role
		i++
	}
	return arr, nil
}

// AddRole adds a new role to the store.
func (store *InMemoryStore) AddRole(role *graylog.Role) error {
	store.roles[role.Name] = *role
	return nil
}

// UpdateRole updates a role at the store.
func (store *InMemoryStore) UpdateRole(name string, role *graylog.Role) error {
	delete(store.roles, name)
	store.roles[role.Name] = *role
	return nil
}

// DeleteRole deletes a role from store.
func (store *InMemoryStore) DeleteRole(name string) error {
	delete(store.roles, name)
	return nil
}

// AuthRolesRead
func (store *InMemoryStore) AuthRolesRead(user *graylog.User, roleName string) (bool, error) {
	perm := fmt.Sprintf("users:read:%s", roleName)
	// check user permissions
	if user.Permissions != nil {
		if user.Permissions.HasAny("*", "users:read", perm) {
			return true, nil
		}
	}
	// check user roles
	if user.Roles == nil {
		return false, nil
	}
	for k, _ := range user.Roles.ToMap(false) {
		role, err := store.GetRole(k)
		if err != nil {
			return false, err
		}
		if role.Permissions == nil {
			continue
		}
		if role.Permissions.HasAny("*", "users:read", perm) {
			return true, nil
		}
	}
	return false, nil
}
