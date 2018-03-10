package graylog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

var (
	once sync.Once
)

// MockServer represents a mock of the Graylog API.
type MockServer struct {
	server   *httptest.Server `json:"-"`
	endpoint string           `json:"-"`

	// users         map[string]User                  `json:"users"`
	// roles         map[string]Role                  `json:"roles"`
	inputs        map[string]Input                 `json:"inputs"`
	indexSets     map[string]IndexSet              `json:"index_sets"`
	indexSetStats map[string]IndexSetStats         `json:"index_set_stats"`
	streams       map[string]Stream                `json:"streams"`
	streamRules   map[string]map[string]StreamRule `json:"stream_rules"`

	store    Store       `json:"-"`
	logger   *log.Logger `json:"-"`
	dataPath string      `json:"-"`
}

func (ms *MockServer) Logger() *log.Logger {
	return ms.logger
}

// NewMockServer returns new MockServer but doesn't start it.
// If addr is an empty string, the free port is assigned automatially.
func NewMockServer(addr string) (*MockServer, error) {
	store := &InMemoryStore{
		roles:         map[string]Role{},
		users:         map[string]User{},
		inputs:        map[string]Input{},
		indexSets:     map[string]IndexSet{},
		indexSetStats: map[string]IndexSetStats{},
		streams:       map[string]Stream{},
		streamRules:   map[string]map[string]StreamRule{},
	}
	ms := &MockServer{
		indexSetStats: map[string]IndexSetStats{},
		streams:       map[string]Stream{},
		streamRules:   map[string]map[string]StreamRule{},

		store:  store,
		logger: log.New(),
	}
	// By Default logLevel is error
	// because debug and info logs are often noisy at unit tests.
	ms.logger.SetLevel(log.ErrorLevel)

	router := httprouter.New()

	router.GET("/api/roles/:rolename", ms.handleGetRole)
	router.PUT("/api/roles/:rolename", ms.handleUpdateRole)
	router.DELETE("/api/roles/:rolename", ms.handleDeleteRole)
	router.GET("/api/roles", ms.handleGetRoles)
	router.POST("/api/roles", ms.handleCreateRole)

	router.GET("/api/users/:username", ms.handleGetUser)
	router.PUT("/api/users/:username", ms.handleUpdateUser)
	router.DELETE("/api/users/:username", ms.handleDeleteUser)
	router.GET("/api/users", ms.handleGetUsers)
	router.POST("/api/users", ms.handleCreateUser)

	router.GET("/api/roles/:rolename/members", ms.handleRoleMembers)
	router.PUT("/api/roles/:rolename/members/:username", ms.handleAddUserToRole)
	router.DELETE(
		"/api/roles/:rolename/members/:username", ms.handleRemoveUserFromRole)

	router.GET("/api/system/inputs", ms.handleGetInputs)
	router.GET("/api/system/inputs/:inputId", ms.handleGetInput)
	router.POST("/api/system/inputs", ms.handleCreateInput)
	router.PUT("/api/system/inputs/:inputId", ms.handleUpdateInput)
	router.DELETE("/api/system/inputs/:inputId", ms.handleDeleteInput)

	router.GET("/api/system/indices/index_sets", ms.handleGetIndexSets)
	router.GET(
		"/api/system/indices/index_sets/:indexSetId", ms.handleGetIndexSet)
	router.POST("/api/system/indices/index_sets", ms.handleCreateIndexSet)
	router.PUT(
		"/api/system/indices/index_sets/:indexSetId", ms.handleUpdateIndexSet)
	router.DELETE(
		"/api/system/indices/index_sets/:indexSetId", ms.handleDeleteIndexSet)
	router.GET(
		"/api/system/indices/index_sets/:indexSetId/stats",
		ms.handleGetIndexSetStats)
	router.PUT(
		"/api/system/indices/index_sets/:indexSetId/default",
		ms.handleSetDefaultIndexSet)

	router.GET("/api/streams", ms.handleGetStreams)
	router.POST("/api/streams", ms.handleCreateStream)
	router.GET("/api/streams/:streamId", ms.handleGetStream)
	router.PUT("/api/streams/:streamId", ms.handleUpdateStream)
	router.DELETE("/api/streams/:streamId", ms.handleDeleteStream)
	router.POST("/api/streams/:streamId/pause", ms.handlePauseStream)
	router.POST("/api/streams/:streamId/resume", ms.handleResumeStream)

	router.GET("/api/streams/:streamId/rules", ms.handleGetStreamRules)
	router.POST("/api/streams/:streamId/rules", ms.handleCreateStreamRule)

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
	if ms.dataPath == "" {
		return nil
	}
	b, err := json.Marshal(ms)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(ms.dataPath, b, 0600)
}

// Load reads Mock Server's data from a file.
func (ms *MockServer) Load() error {
	if ms.dataPath == "" {
		return nil
	}
	if _, err := os.Stat(ms.dataPath); err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err,
			"path":  ms.dataPath}).Info("data file is not found")
		return nil
	}
	b, err := ioutil.ReadFile(ms.dataPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, ms)
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
