package logic

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/inmemory"
)

// Logic represents a mock of the Graylog API.
type Logic struct {
	authEnabled bool                                     `json:"-"`
	streamRules map[string]map[string]graylog.StreamRule `json:"stream_rules"`

	store  store.Store `json:"-"`
	logger *log.Logger `json:"-"`
}

func (ms *Logic) Logger() *log.Logger {
	return ms.logger
}

// NewServer returns new Server.
func NewServer(store store.Store) (*Logic, error) {
	if store == nil {
		store = inmemory.NewStore("")
	}
	ms := &Logic{
		// indexSetStats: map[string]graylog.IndexSetStats{},
		streamRules: map[string]map[string]graylog.StreamRule{},

		store:  store,
		logger: log.New(),
	}
	// By Default logLevel is error
	// because debug and info logs are often noisy at unit tests.
	ms.logger.SetLevel(log.WarnLevel)
	if err := ms.InitData(); err != nil {
		return ms, err
	}
	return ms, nil
}

// SetStore sets a store to the mock server.
func (ms *Logic) SetStore(store store.Store) {
	ms.store = store
}

// Save writes Mock Server's data in a file for persistence.
func (ms *Logic) Save() error {
	return ms.store.Save()
}

// Load reads Mock Server's data from a file.
func (ms *Logic) Load() error {
	return ms.store.Load()
}

// SetAuth sets whether the authentication and authentication are enabled.
func (ms *Logic) SetAuth(authEnabled bool) {
	ms.authEnabled = authEnabled
}

// GetAuth reruns whether the authentication and authentication are enabled.
func (ms *Logic) GetAuth() bool {
	return ms.authEnabled
}

// Authorize
func (ms *Logic) Authorize(user *graylog.User, scope string, args ...string) (int, error) {
	if user == nil {
		return 200, nil
	}
	ok, err := ms.store.Authorize(user, scope, args...)
	if err != nil {
		return 500, err
	}
	if ok {
		return 200, nil
	}
	return 403, fmt.Errorf("authorization failure")
}
