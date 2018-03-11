package graylog

// HasRole
func (store *InMemoryStore) HasRole(name string) (bool, error) {
	_, ok := store.roles[name]
	return ok, nil
}

// GetRole returns a Role.
// If no role with given name is found, returns nil and not returns an error.
func (store *InMemoryStore) GetRole(name string) (*Role, error) {
	s, ok := store.roles[name]
	if ok {
		return &s, nil
	}
	return nil, nil
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
func (store *InMemoryStore) AddRole(role *Role) error {
	store.roles[role.Name] = *role
	return nil
}

// UpdateRole updates a role at the store.
func (store *InMemoryStore) UpdateRole(name string, role *Role) error {
	delete(store.roles, name)
	store.roles[role.Name] = *role
	return nil
}

// DeleteRole deletes a role from store.
func (store *InMemoryStore) DeleteRole(name string) error {
	delete(store.roles, name)
	return nil
}
