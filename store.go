package graylog

// Store manage data.
// Basically Store doesn't have responsibility to validate a request from user.
type Store interface {
	HasRole(name string) (bool, error)
	// GetRole returns a role.
	// If no role with given name is found, returns nil and not returns an error.
	GetRole(name string) (*Role, error)
	GetRoles() ([]Role, error)
	AddRole(role *Role) error
	UpdateRole(name string, role *Role) error
	DeleteRole(name string) error

	HasUser(username string) (bool, error)
	GetUser(username string) (*User, error)
	GetUsers() ([]User, error)
	AddUser(user *User) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(name string) error

	HasInput(id string) (bool, error)
	GetInput(id string) (*Input, error)
	GetInputs() ([]Input, error)
	AddInput(input *Input) (*Input, error)
	UpdateInput(input *Input) error
	DeleteInput(id string) error

	HasIndexSet(id string) (bool, error)
	GetIndexSet(id string) (*IndexSet, error)
	GetIndexSets() ([]IndexSet, error)
	AddIndexSet(indexSet *IndexSet) (*IndexSet, error)
	UpdateIndexSet(indexSet *IndexSet) error
	DeleteIndexSet(id string) error
	SetDefaultIndexSetID(id string) error
	GetDefaultIndexSetID() (string, error)
	IsConflictIndexPrefix(id, indexPrefix string) (bool, error)

	HasStream(id string) (bool, error)
	GetStream(id string) (Stream, bool, error)
	GetStreams() ([]Stream, error)
	AddStream(stream *Stream) (*Stream, int, error)
	UpdateStream(stream *Stream) (int, error)
	DeleteStream(id string) (int, error)
	GetEnabledStreams() ([]Stream, error)
}
