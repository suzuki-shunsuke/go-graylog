package plain

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

// HasRole
func (store *PlainStore) HasRole(name string) (bool, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	_, ok := store.roles[name]
	return ok, nil
}

// GetRole returns a Role.
// If no role with given name is found, returns nil and not returns an error.
func (store *PlainStore) GetRole(name string) (*graylog.Role, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	s, ok := store.roles[name]
	if ok {
		return &s, nil
	}
	return nil, nil
}

// GetRoles returns Roles.
func (store *PlainStore) GetRoles() ([]graylog.Role, int, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	size := len(store.roles)
	arr := make([]graylog.Role, size)
	i := 0
	for _, role := range store.roles {
		arr[i] = role
		i++
	}
	return arr, size, nil
}

// AddRole adds a new role to the store.
func (store *PlainStore) AddRole(role *graylog.Role) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	store.roles[role.Name] = *role
	return nil
}

// UpdateRole updates a role at the store.
func (store *PlainStore) UpdateRole(name string, role *graylog.Role) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	s, ok := store.roles[name]
	if !ok {
		return fmt.Errorf(`no role with name "%s"`, name)
	}
	if role.Description == "" {
		role.Description = s.Description
	}
	delete(store.roles, name)
	store.roles[role.Name] = *role
	return nil
}

// DeleteRole deletes a role from store.
func (store *PlainStore) DeleteRole(name string) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	delete(store.roles, name)
	return nil
}
