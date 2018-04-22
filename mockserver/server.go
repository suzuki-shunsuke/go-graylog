package mockserver

import (
	"fmt"
	"net"
	"net/http/httptest"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/handler"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/plain"
)

// Server represents a mock of the Graylog API.
type Server struct {
	*logic.Logic `json:"-"`
	server       *httptest.Server
	endpoint     string
}

// NewServer returns new Server but doesn't start it.
// If addr is an empty string, the free port is assigned automatially.
func NewServer(addr string, store store.Store) (*Server, error) {
	if store == nil {
		store = plain.NewStore("")
	}
	srv, err := logic.NewLogic(store)
	if err != nil {
		return nil, err
	}
	ms := &Server{
		Logic:  srv,
		server: httptest.NewUnstartedServer(handler.NewRouter(srv)),
	}
	if addr != "" {
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			return nil, err
		}
		ms.server.Listener = ln
	}
	ms.endpoint = fmt.Sprintf("http://%s/api", ms.server.Listener.Addr().String())
	return ms, nil
}

// Start starts a server from NewUnstartedServer.
func (ms *Server) Start() {
	ms.server.Start()
}

// Close shuts down the server and blocks until all outstanding requests on this server have completed.
func (ms *Server) Close() {
	ms.Logger().Info("Close Server")
	ms.server.Close()
}

// Endpoint returns the endpoint url.
func (ms *Server) Endpoint() string {
	return ms.endpoint
}
