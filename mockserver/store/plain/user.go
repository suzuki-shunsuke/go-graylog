package plain

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
	st "github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

// HasUser
func (store *PlainStore) HasUser(username string) (bool, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	_, ok := store.users[username]
	return ok, nil
}

// GetUser returns a user.
// If the user is not found, this method returns nil and doesn't raise an error.
func (store *PlainStore) GetUser(username string) (*graylog.User, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	s, ok := store.users[username]
	if ok {
		return &s, nil
	}
	return nil, nil
}

// GetUsers returns users
func (store *PlainStore) GetUsers() ([]graylog.User, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	arr := make([]graylog.User, len(store.users))
	i := 0
	for _, user := range store.users {
		arr[i] = user
		i++
	}
	return arr, nil
}

// AddUser adds a user to the PlainStore.
func (store *PlainStore) AddUser(user *graylog.User) error {
	if user == nil {
		return fmt.Errorf("user is nil")
	}
	if user.ID == "" {
		user.ID = st.NewObjectID()
	}
	store.imutex.Lock()
	defer store.imutex.Unlock()
	store.users[user.Username] = *user
	return nil
}

// UpdateUser updates a user of the PlainStore.
// "email", "permissions", "full_name", "password"
func (store *PlainStore) UpdateUser(user *graylog.User) error {
	if user == nil {
		return fmt.Errorf("user is nil")
	}
	store.imutex.Lock()
	defer store.imutex.Unlock()
	u, ok := store.users[user.Username]
	if !ok {
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
	store.users[u.Username] = u
	return nil
}

// DeleteUser removes a user from the PlainStore.
func (store *PlainStore) DeleteUser(name string) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	delete(store.users, name)
	return nil
}

// GetUserByAccessToken returns a user name.
// If the user is not found, this method returns nil and doesn't raise an error.
func (store *PlainStore) GetUserByAccessToken(token string) (*graylog.User, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	username, ok := store.tokens[token]
	if !ok {
		return nil, nil
	}
	s, ok := store.users[username]
	if ok {
		return &s, nil
	}
	return nil, nil
}
