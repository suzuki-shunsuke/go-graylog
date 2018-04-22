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
func (store *PlainStore) UpdateUser(prms *graylog.UserUpdateParams) error {
	if prms == nil {
		return fmt.Errorf("user is nil")
	}
	store.imutex.Lock()
	defer store.imutex.Unlock()
	user, ok := store.users[prms.Username]
	if !ok {
		return fmt.Errorf(`the user "%s" is not found`, prms.Username)
	}

	if prms.Email != nil {
		user.Email = *prms.Email
	}
	if prms.FullName != nil {
		user.FullName = *prms.FullName
	}
	if prms.Password != nil {
		user.Password = *prms.Password
	}
	if prms.Timezone != nil {
		user.Timezone = *prms.Timezone
	}
	if prms.SessionTimeoutMs != nil {
		user.SessionTimeoutMs = *prms.SessionTimeoutMs
	}
	if prms.Permissions != nil {
		user.Permissions = prms.Permissions
	}
	if prms.Startpage != nil {
		user.Startpage = prms.Startpage
	}
	if prms.Roles != nil {
		user.Roles = prms.Roles
	}
	store.users[user.Username] = user
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
