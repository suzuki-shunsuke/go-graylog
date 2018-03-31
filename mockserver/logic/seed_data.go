package logic

import (
	"github.com/suzuki-shunsuke/go-graylog/mockserver/seed"
)

// InitData sets an initial data.
func (ms *Logic) InitData() error {
	role := seed.Role()
	if _, err := ms.AddRole(role); err != nil {
		return err
	}
	if _, err := ms.AddUser(seed.User()); err != nil {
		return err
	}
	if _, err := ms.AddUser(seed.Nobody()); err != nil {
		return err
	}
	ms.AddInput(seed.Input())
	is := seed.IndexSet()
	if _, err := ms.AddIndexSet(is); err != nil {
		return err
	}
	is, _, err := ms.SetDefaultIndexSet(is.ID)
	if err != nil {
		return err
	}
	stream := seed.Stream()
	stream.IndexSetID = is.ID
	if _, err := ms.AddStream(stream); err != nil {
		return err
	}
	rule := seed.StreamRule()
	rule.StreamID = stream.ID
	if _, err := ms.AddStreamRule(rule); err != nil {
		return err
	}
	return nil
}
