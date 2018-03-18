package mockserver

import (
	"fmt"
	"net"
	"net/http/httptest"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
)

var (
	once sync.Once
)

// Server represents a mock of the Graylog API.
type Server struct {
	server   *httptest.Server `json:"-"`
	endpoint string           `json:"-"`

	streamRules map[string]map[string]graylog.StreamRule `json:"stream_rules"`

	store    Store       `json:"-"`
	logger   *log.Logger `json:"-"`
	dataPath string      `json:"-"`
}

func (ms *Server) Logger() *log.Logger {
	return ms.logger
}

// NewServer returns new Server but doesn't start it.
// If addr is an empty string, the free port is assigned automatially.
func NewServer(addr string, store Store) (*Server, error) {
	ms := &Server{
		// indexSetStats: map[string]graylog.IndexSetStats{},
		streamRules: map[string]map[string]graylog.StreamRule{},

		store:  store,
		logger: log.New(),
	}
	// By Default logLevel is error
	// because debug and info logs are often noisy at unit tests.
	ms.logger.SetLevel(log.ErrorLevel)

	server := httptest.NewUnstartedServer(newRouter(ms))
	if addr != "" {
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			return nil, err
		}
		server.Listener = ln
	}
	u := fmt.Sprintf("http://%s/api", server.Listener.Addr().String())
	ms.endpoint = u
	ms.server = server
	if err := ms.InitData(); err != nil {
		return ms, err
	}
	return ms, nil
}

// SetStore sets a store to the mock server.
func (ms *Server) SetStore(store Store) {
	ms.store = store
}

// Start starts a server from NewUnstartedServer.
func (ms *Server) Start() {
	ms.server.Start()
}

// Close shuts down the server and blocks until all outstanding requests
// on this server have completed.
func (ms *Server) Close() {
	ms.Logger().Info("Close Server")
	ms.server.Close()
}

// Save writes Mock Server's data in a file for persistence.
func (ms *Server) Save() error {
	return ms.store.Save()
}

// Load reads Mock Server's data from a file.
func (ms *Server) Load() error {
	return ms.store.Load()
}

func (ms *Server) safeSave() {
	if err := ms.Save(); err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "data_path": ms.dataPath,
		}).Error("Failed to save data")
	}
}

func (ms *Server) GetEndpoint() string {
	return ms.endpoint
}
