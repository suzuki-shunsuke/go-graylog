package logic

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

func (ms *Server) Authenticate(name, password string) (*graylog.User, int, error) {
	if name == "" || password == "" {
		return nil, 401, fmt.Errorf("authentication failure")
	}
	if password == "session" {
		// session token is not supported
		return nil, 400, fmt.Errorf(`{"message":"mock server doesn't support session token"}`)
	}
	if password == "token" {
		// access token
		user, err := ms.store.GetUserByAccessToken(name)
		if err != nil {
			return nil, 500, err
		}
		if user == nil {
			return nil, 401, fmt.Errorf("authentication failure")
		}
		return user, 200, nil
	}
	// TODO authentication
	user, err := ms.store.GetUser(name)
	if err != nil {
		return user, 500, err
	}
	return user, 200, nil
}
