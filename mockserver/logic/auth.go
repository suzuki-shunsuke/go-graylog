package logic

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

// Authenticate authenticates a user.
func (lgc *Logic) Authenticate(name, password string) (*graylog.User, int, error) {
	if name == "" || password == "" {
		return nil, 401, fmt.Errorf("authentication failure")
	}
	if password == "session" {
		// session token is not supported
		return nil, 400, fmt.Errorf(`{"message":"mock server doesn't support session token"}`)
	}
	if password == "token" {
		// access token
		user, err := lgc.store.GetUserByAccessToken(name)
		if err != nil {
			return nil, 500, err
		}
		if user == nil {
			return nil, 401, fmt.Errorf("authentication failure")
		}
		return user, 200, nil
	}
	user, err := lgc.store.GetUser(name)
	if err != nil {
		return nil, 500, err
	}
	if user == nil {
		return nil, 401, fmt.Errorf("authentication failure")
	}
	if user.Password != encryptPassword(password) {
		return nil, 401, fmt.Errorf("authentication failure")
	}
	return user, 200, nil
}
