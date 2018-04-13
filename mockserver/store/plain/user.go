package plain

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
	st "github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

// HasUser
func (store *InMemoryStore) HasUser(username string) (bool, error) {
	_, ok := store.users[username]
	return ok, nil
}

// GetUser returns a user.
// If the user is not found, this method returns nil and doesn't raise an error.
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
	if user == nil {
		return fmt.Errorf("user is nil")
	}
	if user.ID == "" {
		user.ID = st.NewObjectID()
	}
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
		return fmt.Errorf(`the user "%s" is not found`, user.Username)
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

// GetUserByAccessToken returns a user name.
// If the user is not found, this method returns nil and doesn't raise an error.
func (store *InMemoryStore) GetUserByAccessToken(token string) (*graylog.User, error) {
	userName, ok := store.tokens[token]
	if !ok {
		return nil, nil
	}
	return store.GetUser(userName)
}
