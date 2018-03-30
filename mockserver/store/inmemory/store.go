package inmemory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

type InMemoryStore struct {
	users             map[string]graylog.User                  `json:"users"`
	roles             map[string]graylog.Role                  `json:"roles"`
	inputs            map[string]graylog.Input                 `json:"inputs"`
	indexSets         map[string]graylog.IndexSet              `json:"index_sets"`
	defaultIndexSetID string                                   `json:"default_index_set_id"`
	indexSetStats     map[string]graylog.IndexSetStats         `json:"index_set_stats"`
	streams           map[string]graylog.Stream                `json:"streams"`
	streamRules       map[string]map[string]graylog.StreamRule `json:"stream_rules"`
	dataPath          string                                   `json:"-"`
	tokens            map[string]string                        `json:"tokens"`
}

func NewStore(dataPath string) store.Store {
	return &InMemoryStore{
		roles:         map[string]graylog.Role{},
		users:         map[string]graylog.User{},
		inputs:        map[string]graylog.Input{},
		indexSets:     map[string]graylog.IndexSet{},
		indexSetStats: map[string]graylog.IndexSetStats{},
		streams:       map[string]graylog.Stream{},
		streamRules:   map[string]map[string]graylog.StreamRule{},
		tokens:        map[string]string{},
		dataPath:      dataPath,
	}
}

// Save writes Mock Server's data in a file for persistence.
func (store *InMemoryStore) Save() error {
	if store.dataPath == "" {
		return nil
	}
	b, err := json.Marshal(store)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(store.dataPath, b, 0600)
}

// Load reads Mock Server's data from a file.
func (store *InMemoryStore) Load() error {
	if store.dataPath == "" {
		return nil
	}
	if _, err := os.Stat(store.dataPath); err != nil {
		return nil
	}
	b, err := ioutil.ReadFile(store.dataPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, store)
}

func (store *InMemoryStore) Auth(user *graylog.User, scope string, args ...string) (bool, error) {
	perm := fmt.Sprintf("%s:%s", scope, strings.Join(args, ":"))
	// check user permissions
	if user.Permissions != nil {
		if user.Permissions.HasAny("*", scope, perm) {
			return true, nil
		}
	}
	// check user roles
	if user.Roles == nil {
		return false, nil
	}
	for k := range user.Roles.ToMap(false) {
		// get role
		role, err := store.GetRole(k)
		if err != nil {
			return false, err
		}
		// check role permissions
		if role.Permissions == nil {
			continue
		}
		if role.Permissions.HasAny("*", scope, perm) {
			return true, nil
		}
	}
	return false, nil
}
