package graylog

type Store interface {
	HasRole(name string) bool
	GetRole(name string) (Role, bool)
	GetRoles() ([]Role, error)
	AddRole(role *Role) (int, error)
	UpdateRole(name string, role *Role) (int, error)
	DeleteRole(name string) (int, error)

	HasUser(username string) bool
	GetUser(username string) (User, bool)
	GetUsers() ([]User, error)
	AddUser(user *User) (*User, int, error)
	UpdateUser(user *User) (int, error)
	DeleteUser(name string) (int, error)

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
