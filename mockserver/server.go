package mockserver

import (
	"fmt"
	"net"
	"net/http/httptest"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/inmemory"
)

var (
	once sync.Once
)

// Server represents a mock of the Graylog API.
type Server struct {
	server        *httptest.Server `json:"-"`
	*logic.Server `json:"-"`
	endpoint      string `json:"-"`

	streamRules map[string]map[string]graylog.StreamRule `json:"stream_rules"`
}

// NewServer returns new Server but doesn't start it.
// If addr is an empty string, the free port is assigned automatially.
func NewServer(addr string, store store.Store) (*Server, error) {
	if store == nil {
		store = inmemory.NewStore("")
	}
	logger := log.New()
	// By Default logLevel is error
	// because debug and info logs are often noisy at unit tests.
	logger.SetLevel(log.ErrorLevel)
	srv, err := logic.NewServer(store)
	if err != nil {
		return nil, err
	}
	ms := &Server{
		Server: srv,
	}

	server := httptest.NewUnstartedServer(newRouter(srv))
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
	return ms, nil
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

func (ms *Server) GetEndpoint() string {
	return ms.endpoint
}
