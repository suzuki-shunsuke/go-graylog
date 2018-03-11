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
	GetInput(id string) (Input, bool, error)
	GetInputs() ([]Input, error)
	AddInput(input *Input) (*Input, int, error)
	UpdateInput(input *Input) (int, error)
	DeleteInput(id string) (int, error)

	HasIndexSet(id string) (bool, error)
	GetIndexSet(id string) (IndexSet, bool, error)
	GetIndexSets() ([]IndexSet, error)
	AddIndexSet(indexSet *IndexSet) (*IndexSet, int, error)
	UpdateIndexSet(indexSet *IndexSet) (int, error)
	DeleteIndexSet(id string) (int, error)

	HasStream(id string) (bool, error)
	GetStream(id string) (Stream, bool, error)
	GetStreams() ([]Stream, error)
	AddStream(stream *Stream) (*Stream, int, error)
	UpdateStream(stream *Stream) (int, error)
	DeleteStream(id string) (int, error)
	GetEnabledStreams() ([]Stream, error)
}
