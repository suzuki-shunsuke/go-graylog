package graylog

// HasRole
func (store *InMemoryStore) HasRole(name string) bool {
	_, ok := store.roles[name]
	return ok
}

// GetRole returns a Role.
func (store *InMemoryStore) GetRole(name string) (Role, bool) {
	s, ok := store.roles[name]
	return s, ok
}

// GetRoles returns Roles.
func (store *InMemoryStore) GetRoles() ([]Role, error) {
	size := len(store.roles)
	arr := make([]Role, size)
	i := 0
	for _, role := range store.roles {
		arr[i] = role
		i++
	}
	return arr, nil
}

// AddRole adds a new role to the store.
func (store *InMemoryStore) AddRole(role *Role) (int, error) {
	store.roles[role.Name] = *role
	return 200, nil
}

// UpdateRole updates a role at the store.
func (store *InMemoryStore) UpdateRole(name string, role *Role) (int, error) {
	delete(store.roles, name)
	store.roles[role.Name] = *role
	return 200, nil
}

// DeleteRole deletes a role from store.
func (store *InMemoryStore) DeleteRole(name string) (int, error) {
	delete(store.roles, name)
	return 200, nil
}
