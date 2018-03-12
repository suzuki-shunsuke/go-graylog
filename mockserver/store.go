package mockserver

import (
	"github.com/suzuki-shunsuke/go-graylog"
)

// Store manage data.
// Basically Store doesn't have responsibility to validate a request from user.
type Store interface {
	Save() error
	Load() error

	HasRole(name string) (bool, error)
	// GetRole returns a role.
	// If no role with given name is found, returns nil and not returns an error.
	GetRole(name string) (*graylog.Role, error)
	GetRoles() ([]graylog.Role, error)
	AddRole(role *graylog.Role) error
	UpdateRole(name string, role *graylog.Role) error
	DeleteRole(name string) error

	HasUser(username string) (bool, error)
	GetUser(username string) (*graylog.User, error)
	GetUsers() ([]graylog.User, error)
	AddUser(user *graylog.User) (*graylog.User, error)
	UpdateUser(user *graylog.User) error
	DeleteUser(name string) error

	HasInput(id string) (bool, error)
	GetInput(id string) (*graylog.Input, error)
	GetInputs() ([]graylog.Input, error)
	AddInput(input *graylog.Input) (*graylog.Input, error)
	UpdateInput(input *graylog.Input) error
	DeleteInput(id string) error

	HasIndexSet(id string) (bool, error)
	GetIndexSet(id string) (*graylog.IndexSet, error)
	GetIndexSets() ([]graylog.IndexSet, error)
	AddIndexSet(indexSet *graylog.IndexSet) (*graylog.IndexSet, error)
	UpdateIndexSet(indexSet *graylog.IndexSet) error
	DeleteIndexSet(id string) error
	SetDefaultIndexSetID(id string) error
	GetDefaultIndexSetID() (string, error)
	IsConflictIndexPrefix(id, indexPrefix string) (bool, error)

	GetIndexSetStats(id string) (*graylog.IndexSetStats, error)
	GetIndexSetsStats() ([]graylog.IndexSetStats, error)
	GetTotalIndexSetsStats() (*graylog.IndexSetStats, error)
	SetIndexSetStats(id string, stats *graylog.IndexSetStats) error

	HasStream(id string) (bool, error)
	GetStream(id string) (*graylog.Stream, error)
	GetStreams() ([]graylog.Stream, error)
	AddStream(stream *graylog.Stream) (*graylog.Stream, error)
	UpdateStream(stream *graylog.Stream) error
	DeleteStream(id string) error
	GetEnabledStreams() ([]graylog.Stream, error)
}
