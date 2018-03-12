package mockserver

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
)

var (
	once sync.Once
)

// MockServer represents a mock of the Graylog API.
type MockServer struct {
	server   *httptest.Server `json:"-"`
	endpoint string           `json:"-"`

	streamRules map[string]map[string]graylog.StreamRule `json:"stream_rules"`

	store    Store       `json:"-"`
	logger   *log.Logger `json:"-"`
	dataPath string      `json:"-"`
}

func (ms *MockServer) Logger() *log.Logger {
	return ms.logger
}

// NewMockServer returns new MockServer but doesn't start it.
// If addr is an empty string, the free port is assigned automatially.
func NewMockServer(addr string, store Store) (*MockServer, error) {
	ms := &MockServer{
		// indexSetStats: map[string]graylog.IndexSetStats{},
		streamRules: map[string]map[string]graylog.StreamRule{},

		store:  store,
		logger: log.New(),
	}
	// By Default logLevel is error
	// because debug and info logs are often noisy at unit tests.
	ms.logger.SetLevel(log.ErrorLevel)

	router := httprouter.New()

	router.GET("/api/roles/:rolename", wrapHandle(ms.handleGetRole))
	router.PUT("/api/roles/:rolename", wrapHandle(ms.handleUpdateRole))
	router.DELETE("/api/roles/:rolename", wrapHandle(ms.handleDeleteRole))
	router.GET("/api/roles", wrapHandle(ms.handleGetRoles))
	router.POST("/api/roles", wrapHandle(ms.handleCreateRole))

	router.GET("/api/users/:username", wrapHandle(ms.handleGetUser))
	router.PUT("/api/users/:username", wrapHandle(ms.handleUpdateUser))
	router.DELETE("/api/users/:username", wrapHandle(ms.handleDeleteUser))
	router.GET("/api/users", wrapHandle(ms.handleGetUsers))
	router.POST("/api/users", wrapHandle(ms.handleCreateUser))

	router.GET("/api/roles/:rolename/members", wrapHandle(ms.handleRoleMembers))
	router.PUT("/api/roles/:rolename/members/:username", wrapHandle(ms.handleAddUserToRole))
	router.DELETE(
		"/api/roles/:rolename/members/:username", wrapHandle(ms.handleRemoveUserFromRole))

	router.GET("/api/system/inputs", wrapHandle(ms.handleGetInputs))
	router.GET("/api/system/inputs/:inputID", wrapHandle(ms.handleGetInput))
	router.POST("/api/system/inputs", wrapHandle(ms.handleCreateInput))
	router.PUT("/api/system/inputs/:inputID", wrapHandle(ms.handleUpdateInput))
	router.DELETE("/api/system/inputs/:inputID", wrapHandle(ms.handleDeleteInput))

	router.GET("/api/system/indices/index_sets", wrapHandle(ms.handleGetIndexSets))
	router.GET(
		"/api/system/indices/index_sets/:indexSetID", wrapHandle(ms.handleGetIndexSet))
	router.POST("/api/system/indices/index_sets", wrapHandle(ms.handleCreateIndexSet))
	router.PUT(
		"/api/system/indices/index_sets/:indexSetID", wrapHandle(ms.handleUpdateIndexSet))
	router.DELETE(
		"/api/system/indices/index_sets/:indexSetID", wrapHandle(ms.handleDeleteIndexSet))
	router.PUT(
		"/api/system/indices/index_sets/:indexSetID/default",
		wrapHandle(ms.handleSetDefaultIndexSet))

	router.GET(
		"/api/system/indices/index_sets/:indexSetID/stats",
		wrapHandle(ms.handleGetIndexSetStats))

	router.GET("/api/streams", wrapHandle(ms.handleGetStreams))
	router.POST("/api/streams", wrapHandle(ms.handleCreateStream))
	router.GET("/api/streams/:streamID", wrapHandle(ms.handleGetStream))
	router.PUT("/api/streams/:streamID", wrapHandle(ms.handleUpdateStream))
	router.DELETE("/api/streams/:streamID", wrapHandle(ms.handleDeleteStream))
	router.POST("/api/streams/:streamID/pause", wrapHandle(ms.handlePauseStream))
	router.POST("/api/streams/:streamID/resume", wrapHandle(ms.handleResumeStream))

	router.GET("/api/streams/:streamID/rules", wrapHandle(ms.handleGetStreamRules))
	router.POST("/api/streams/:streamID/rules", wrapHandle(ms.handleCreateStreamRule))

	router.NotFound = ms.handleNotFound

	server := httptest.NewUnstartedServer(router)
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

// SetStore sets a store to the mock server.
func (ms *MockServer) SetStore(store Store) {
	ms.store = store
}

// Start starts a server from NewUnstartedServer.
func (ms *MockServer) Start() {
	ms.server.Start()
}

// Close shuts down the server and blocks until all outstanding requests
// on this server have completed.
func (ms *MockServer) Close() {
	ms.Logger().Info("Close Server")
	ms.server.Close()
}

// Save writes Mock Server's data in a file for persistence.
func (ms *MockServer) Save() error {
	return ms.store.Save()
}

// Load reads Mock Server's data from a file.
func (ms *MockServer) Load() error {
	return ms.store.Load()
}

func (ms *MockServer) safeSave() {
	if err := ms.Save(); err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "data_path": ms.dataPath,
		}).Error("Failed to save data")
	}
}

func (ms *MockServer) handleNotFound(w http.ResponseWriter, r *http.Request) {
	ms.Logger().WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
		"message": "404 Page Not Found",
	}).Info("request start")
	w.WriteHeader(404)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(
		`{"message":"Page Not Found %s %s"}`, r.Method, r.URL.Path)))
}

func (ms *MockServer) handleInit(
	w http.ResponseWriter, r *http.Request, isReadBody bool,
) ([]byte, error) {
	ms.Logger().WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	if isReadBody {
		return ioutil.ReadAll(r.Body)
	}
	return nil, nil
}

func (ms *MockServer) GetEndpoint() string {
	return ms.endpoint
}

type Handler func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (int, interface{}, error)

func wrapHandle(handler Handler) httprouter.Handle {
	// ms.Logger().WithFields(log.Fields{
	// 	"path": r.URL.Path, "method": r.Method,
	// }).Info("request start")
	// w.Header().Set("Content-Type", "application/json")
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		sc, body, err := handler(w, r, ps)
		if err != nil {
			writeApiError(w, sc, err.Error())
			return
		}
		if body != nil {
			writeOr500Error(w, body)
		}
	}
}
