package mockserver

import (
	"github.com/suzuki-shunsuke/go-graylog/mockserver/seed"
)

// InitData sets an initial data.
func (ms *MockServer) InitData() error {
	role := seed.Role()
	if _, err := ms.AddRole(role); err != nil {
		return err
	}
	_, _, err := ms.AddUser(seed.User())
	if err != nil {
		return err
	}
	ms.AddInput(seed.Input())
	indexSet, _, err := ms.AddIndexSet(seed.IndexSet())
	if err != nil {
		return err
	}
	stream := seed.Stream()
	stream.IndexSetID = indexSet.ID
	stream, _, err = ms.AddStream(stream)
	if err != nil {
		return err
	}
	rule := seed.StreamRule()
	rule.StreamID = stream.ID
	if _, _, err := ms.AddStreamRule(rule); err != nil {
		return err
	}
	return nil
}
