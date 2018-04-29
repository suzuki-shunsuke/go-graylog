package logic

import (
	"github.com/suzuki-shunsuke/go-graylog/mockserver/seed"
)

// InitData sets an initial data.
func (lgc *Logic) InitData() error {
	role := seed.Role()
	if _, err := lgc.AddRole(role); err != nil {
		return err
	}
	if _, err := lgc.AddUser(seed.User()); err != nil {
		return err
	}
	if _, err := lgc.AddUser(seed.Nobody()); err != nil {
		return err
	}
	lgc.AddInput(seed.Input())
	is := seed.IndexSet()
	if _, err := lgc.AddIndexSet(is); err != nil {
		return err
	}
	is, _, err := lgc.SetDefaultIndexSet(is.ID)
	if err != nil {
		return err
	}
	stream := seed.Stream()
	stream.IndexSetID = is.ID
	if _, err := lgc.AddStream(stream); err != nil {
		return err
	}
	rule := seed.StreamRule()
	rule.StreamID = stream.ID
	_, err = lgc.AddStreamRule(rule)
	return err
}
