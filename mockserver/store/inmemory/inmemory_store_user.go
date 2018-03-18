package inmemory

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

// HasUser
func (store *InMemoryStore) HasUser(username string) (bool, error) {
	_, ok := store.users[username]
	return ok, nil
}

// GetUser returns a user.
func (store *InMemoryStore) GetUser(username string) (*graylog.User, error) {
	s, ok := store.users[username]
	if ok {
		return &s, nil
	}
	return nil, nil
}

// GetUsers returns users
func (store *InMemoryStore) GetUsers() ([]graylog.User, error) {
	arr := make([]graylog.User, len(store.users))
	i := 0
	for _, user := range store.users {
		arr[i] = user
		i++
	}
	return arr, nil
}

// AddUser adds a user to the InMemoryStore.
func (store *InMemoryStore) AddUser(user *graylog.User) error {
	store.users[user.Username] = *user
	return nil
}

// UpdateUser updates a user of the InMemoryStore.
// "email", "permissions", "full_name", "password"
func (store *InMemoryStore) UpdateUser(user *graylog.User) error {
	u, err := store.GetUser(user.Username)
	if err != nil {
		return err
	}
	if u == nil {
		return fmt.Errorf("The user is not found")
	}
	if user.Email != "" {
		u.Email = user.Email
	}
	if user.Permissions != nil {
		u.Permissions = user.Permissions
	}
	if user.FullName != "" {
		u.FullName = user.FullName
	}
	if user.Password != "" {
		u.Password = user.Password
	}
	store.users[u.Username] = *u
	return nil
}

// DeleteUser removes a user from the InMemoryStore.
func (store *InMemoryStore) DeleteUser(name string) error {
	delete(store.users, name)
	return nil
}
