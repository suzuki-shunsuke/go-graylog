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
	Server   *httptest.Server `json:"-"`
	Endpoint string           `json:"-"`

	Users         map[string]User          `json:"users"`
	Roles         map[string]Role          `json:"roles"`
	Inputs        map[string]Input         `json:"inputs"`
	IndexSets     map[string]IndexSet      `json:"index_sets"`
	IndexSetStats map[string]IndexSetStats `json:"index_set_stats"`
	Streams       map[string]Stream        `json:"streams"`

	Logger   *log.Logger `json:"-"`
	DataPath string      `json:"-"`
}

// NewMockServer returns new MockServer but doesn't start it.
// If addr is an empty string, the free port is assigned automatially.
func NewMockServer(addr string) (*MockServer, error) {
	ms := &MockServer{
		Users:         map[string]User{},
		Roles:         map[string]Role{},
		Inputs:        map[string]Input{},
		IndexSets:     map[string]IndexSet{},
		IndexSetStats: map[string]IndexSetStats{},
		Streams:       map[string]Stream{},
		Logger:        log.New(),
	}
	ms.Logger.SetLevel(log.PanicLevel)

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
	ms.Endpoint = u
	ms.Server = server
	return ms, nil
}

// Start starts a server from NewUnstartedServer.
func (ms *MockServer) Start() {
	ms.Server.Start()
}

// Close shuts down the server and blocks until all outstanding requests
// on this server have completed.
func (ms *MockServer) Close() {
	ms.Server.Close()
}

// Save writes Mock Server's data in a file for persistence.
func (ms *MockServer) Save() error {
	if ms.DataPath == "" {
		return nil
	}
	b, err := json.Marshal(ms)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(ms.DataPath, b, 0600)
}

// Load reads Mock Server's data from a file.
func (ms *MockServer) Load() error {
	if ms.DataPath == "" {
		return nil
	}
	if _, err := os.Stat(ms.DataPath); err != nil {
		ms.Logger.WithFields(log.Fields{
			"error": err,
			"path":  ms.DataPath}).Info("data file is not found")
		return nil
	}
	b, err := ioutil.ReadFile(ms.DataPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, ms)
}

func (ms *MockServer) safeSave() {
	if err := ms.Save(); err != nil {
		ms.Logger.WithFields(log.Fields{
			"error": err, "data_path": ms.DataPath,
		}).Error("Failed to save data")
	}
}

func (ms *MockServer) handleNotFound(w http.ResponseWriter, r *http.Request) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
		"message": "404 Page Not Found",
	}).Info("request start")
	w.WriteHeader(404)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(
		`{"message":"Page Not Found %s %s"}`, r.Method, r.URL.Path)))
}
