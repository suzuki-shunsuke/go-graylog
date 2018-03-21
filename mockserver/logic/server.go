package logic

import (
	"net/http/httptest"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/inmemory"
)

var (
	once sync.Once
)

// Server represents a mock of the Graylog API.
type Server struct {
	server   *httptest.Server `json:"-"`
	endpoint string           `json:"-"`

	streamRules map[string]map[string]graylog.StreamRule `json:"stream_rules"`

	store    store.Store `json:"-"`
	logger   *log.Logger `json:"-"`
	dataPath string      `json:"-"`
}

func (ms *Server) Logger() *log.Logger {
	return ms.logger
}

// NewServer returns new Server.
func NewServer(store store.Store) (*Server, error) {
	if store == nil {
		store = inmemory.NewStore("")
	}
	ms := &Server{
		// indexSetStats: map[string]graylog.IndexSetStats{},
		streamRules: map[string]map[string]graylog.StreamRule{},

		store:  store,
		logger: log.New(),
	}
	// By Default logLevel is error
	// because debug and info logs are often noisy at unit tests.
	ms.logger.SetLevel(log.ErrorLevel)
	if err := ms.InitData(); err != nil {
		return ms, err
	}
	return ms, nil
}

// SetStore sets a store to the mock server.
func (ms *Server) SetStore(store store.Store) {
	ms.store = store
}

// Save writes Mock Server's data in a file for persistence.
func (ms *Server) Save() error {
	return ms.store.Save()
}

// Load reads Mock Server's data from a file.
func (ms *Server) Load() error {
	return ms.store.Load()
}

func (ms *Server) SafeSave() {
	if err := ms.Save(); err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "data_path": ms.dataPath,
		}).Error("Failed to save data")
	}
}
