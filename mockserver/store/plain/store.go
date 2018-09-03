package plain

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

// Store is the implementation of the Store interface with pure golang.
type Store struct {
	alarmCallbacks  map[string]graylog.AlarmCallback
	alerts          map[string]graylog.Alert
	alertConditions map[string]graylog.AlertCondition
	dashboards      map[string]graylog.Dashboard
	indexSets       []graylog.IndexSet
	inputs          map[string]graylog.Input
	roles           map[string]graylog.Role
	streams         map[string]graylog.Stream
	users           map[string]graylog.User
	ldapSetting     *graylog.LDAPSetting

	tokens map[string]string

	dataPath          string
	defaultIndexSetID string

	imutex sync.RWMutex
}

type plainStore struct {
	AlarmCallbacks  map[string]graylog.AlarmCallback  `json:"alarm_callbacks"`
	Alerts          map[string]graylog.Alert          `json:"alerts"`
	AlertConditions map[string]graylog.AlertCondition `json:"alert_conditions"`
	Dashboards      map[string]graylog.Dashboard      `json:"dashboards"`
	Inputs          map[string]graylog.Input          `json:"inputs"`
	IndexSets       []graylog.IndexSet                `json:"index_sets"`
	Roles           map[string]graylog.Role           `json:"roles"`
	Streams         map[string]graylog.Stream         `json:"streams"`
	Users           map[string]graylog.User           `json:"users"`
	LDAPSetting     *graylog.LDAPSetting              `json:"ldap_setting"`

	Tokens map[string]string `json:"tokens"`

	DefaultIndexSetID string `json:"default_index_set_id"`
}

// MarshalJSON is the implementation of the json.Marshaler interface.
func (store *Store) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"alarm_callbacks":  store.alarmCallbacks,
		"alerts":           store.alerts,
		"alert_conditions": store.alertConditions,
		"inputs":           store.inputs,
		"index_sets":       store.indexSets,
		"roles":            store.roles,
		"streams":          store.streams,
		"users":            store.users,
		"ldap_setting":     store.ldapSetting,

		"tokens": store.tokens,

		"default_index_set_id": store.defaultIndexSetID,
	}
	return json.Marshal(data)
}

// UnmarshalJSON is the implementation of the json.Unmarshaler interface.
func (store *Store) UnmarshalJSON(b []byte) error {
	s := &plainStore{}
	if err := json.Unmarshal(b, s); err != nil {
		return err
	}
	store.alarmCallbacks = s.AlarmCallbacks
	store.alerts = s.Alerts
	store.alertConditions = s.AlertConditions
	store.dashboards = s.Dashboards
	store.inputs = s.Inputs
	store.indexSets = s.IndexSets
	store.roles = s.Roles
	store.users = s.Users
	store.streams = s.Streams
	store.ldapSetting = s.LDAPSetting

	store.tokens = s.Tokens

	store.defaultIndexSetID = s.DefaultIndexSetID
	return nil
}

// NewStore returns a new Store.
// the argument `dataPath` is the file path where write the data.
// If `dataPath` is empty, the data aren't written to the file.
func NewStore(dataPath string) store.Store {
	return &Store{
		alarmCallbacks:  map[string]graylog.AlarmCallback{},
		alerts:          map[string]graylog.Alert{},
		alertConditions: map[string]graylog.AlertCondition{},
		dashboards:      map[string]graylog.Dashboard{},
		inputs:          map[string]graylog.Input{},
		indexSets:       []graylog.IndexSet{},
		roles:           map[string]graylog.Role{},
		streams:         map[string]graylog.Stream{},
		users:           map[string]graylog.User{},
		ldapSetting:     defaultLDAPSetting(),

		tokens: map[string]string{},

		dataPath: dataPath,
	}
}

// Save writes Mock Server's data in a file for persistence.
func (store *Store) Save() error {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
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
func (store *Store) Load() error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
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

// Authorize authorizes the user.
func (store *Store) Authorize(user *graylog.User, scope string, args ...string) (bool, error) {
	if user == nil {
		return true, nil
	}
	perm := scope
	if len(args) != 0 {
		perm += ":" + strings.Join(args, ":")
	}
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
