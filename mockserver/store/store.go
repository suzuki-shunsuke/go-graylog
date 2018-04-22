package store

import (
	"github.com/suzuki-shunsuke/go-graylog"
)

// Store manage data.
// Basically Store doesn't have responsibility to validate a request from user.
type Store interface {
	Save() error
	Load() error
	Authorize(user *graylog.User, scope string, args ...string) (bool, error)

	AddRole(*graylog.Role) error
	// GetRole returns a role.
	// If no role with given name is found, returns nil and not returns an error.
	GetRole(name string) (*graylog.Role, error)
	GetRoles() ([]graylog.Role, int, error)
	UpdateRole(name string, role *graylog.RoleUpdateParams) (*graylog.Role, error)
	DeleteRole(name string) error
	HasRole(name string) (bool, error)

	AddUser(user *graylog.User) error
	GetUser(username string) (*graylog.User, error)
	GetUsers() ([]graylog.User, error)
	UpdateUser(*graylog.UserUpdateParams) error
	DeleteUser(name string) error
	HasUser(username string) (bool, error)
	GetUserByAccessToken(token string) (*graylog.User, error)

	AddInput(*graylog.Input) error
	GetInput(id string) (*graylog.Input, error)
	GetInputs() ([]graylog.Input, int, error)
	UpdateInput(*graylog.InputUpdateParams) (*graylog.Input, error)
	DeleteInput(id string) error
	HasInput(id string) (bool, error)

	AddIndexSet(*graylog.IndexSet) error
	GetIndexSet(id string) (*graylog.IndexSet, error)
	GetIndexSets(skip, limit int) ([]graylog.IndexSet, int, error)
	UpdateIndexSet(*graylog.IndexSetUpdateParams) (*graylog.IndexSet, error)
	DeleteIndexSet(id string) error
	HasIndexSet(id string) (bool, error)
	IsConflictIndexPrefix(id, indexPrefix string) (bool, error)
	SetDefaultIndexSetID(id string) error
	GetDefaultIndexSetID() (string, error)

	// SetIndexSetStats(id string, stats *graylog.IndexSetStats) error
	GetIndexSetStats(id string) (*graylog.IndexSetStats, error)
	GetTotalIndexSetStats() (*graylog.IndexSetStats, error)
	GetIndexSetStatsMap() (map[string]graylog.IndexSetStats, error)

	AddStream(*graylog.Stream) error
	GetStream(id string) (*graylog.Stream, error)
	GetStreams() ([]graylog.Stream, int, error)
	GetEnabledStreams() ([]graylog.Stream, int, error)
	UpdateStream(*graylog.Stream) error
	DeleteStream(id string) error
	HasStream(id string) (bool, error)

	AddStreamRule(*graylog.StreamRule) error
	GetStreamRules(id string) ([]graylog.StreamRule, int, error)
	GetStreamRule(streamID, streamRuleID string) (*graylog.StreamRule, error)
	UpdateStreamRule(*graylog.StreamRule) error
	DeleteStreamRule(streamID, streamRuleID string) error
	HasStreamRule(streamID, streamRuleID string) (bool, error)
}
