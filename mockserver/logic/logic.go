package logic

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/plain"
)

// Logic represents a mock of the Graylog API.
// This is embedded to mockserver.Server.
type Logic struct {
	authEnabled bool
	streamRules map[string]map[string]graylog.StreamRule

	store  store.Store
	logger *log.Logger
}

// Logger returns a logger.
// This logger is logrus.Logger .
// https://github.com/sirupsen/logrus
// You can change the Logic's logger configuration freely.
//
//   lgc := logic.NewLogic(nil)
//   logger := lgc.Logger()
//   logger.SetFormatter(&log.JSONFormatter{})
//   logger.SetLevel(log.WarnLevel)
func (lgc *Logic) Logger() *log.Logger {
	return lgc.logger
}

// NewLogic returns new Server.
// The argument `store` is the store which the server uses.
// If `store` is nil, the default plain store is used and data is not persisted.
func NewLogic(store store.Store) (*Logic, error) {
	if store == nil {
		store = plain.NewStore("")
	}
	lgc := &Logic{
		// indexSetStats: map[string]graylog.IndexSetStats{},
		streamRules: map[string]map[string]graylog.StreamRule{},

		store:  store,
		logger: log.New(),
		// By default the authentication is enabled
		authEnabled: true,
	}
	// By Default logLevel is warn,
	// because debug and info logs are often noisy at unit tests.
	lgc.logger.SetLevel(log.WarnLevel)
	err := lgc.InitData()
	return lgc, err
}

// SetStore sets a store to the mock server.
func (lgc *Logic) SetStore(store store.Store) {
	lgc.store = store
}

// Save writes Mock Server's data in a file for persistence.
func (lgc *Logic) Save() error {
	return lgc.store.Save()
}

// Load reads Mock Server's data from a file.
func (lgc *Logic) Load() error {
	return lgc.store.Load()
}

// SetAuth sets whether the authentication and authentication are enabled.
// Disable the authentication.
//
//   lgc.SetAuth(false)
//
// Enable the authentication.
//
//   lgc.SetAuth(true)
func (lgc *Logic) SetAuth(authEnabled bool) {
	lgc.authEnabled = authEnabled
}

// Auth returns whether the authentication and authentication are enabled.
func (lgc *Logic) Auth() bool {
	return lgc.authEnabled
}

// Authorize authorizes a user.
// If the user doesn't have the permission, an error is returned.
//
//   // whether the user has the permission to read all roles
//   if sc, err := lgc.Authorize(user, "roles:read", ""); err != nil {
//   	fmt.Println(sc, err) // 403, "authorization failure"
//   }
//
//   // whether the user has the permission to read the role "foo"
//   sc, err := lgc.Authorize(admin, "roles:read", "foo")
//   fmt.Println(sc, err) // 200, nil
func (lgc *Logic) Authorize(user *graylog.User, scope string, args ...string) (int, error) {
	if user == nil {
		return 200, nil
	}
	ok, err := lgc.store.Authorize(user, scope, args...)
	if err != nil {
		return 500, err
	}
	if ok {
		return 200, nil
	}
	return 403, fmt.Errorf("authorization failure")
}
