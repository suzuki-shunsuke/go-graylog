package inmemory_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/inmemory"
)

func TestGetStreamRule(t *testing.T) {
	store := inmemory.NewStore("")
	_, err := store.GetStreamRule("", "")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetStreamRules(t *testing.T) {
	store := inmemory.NewStore("")
	_, err := store.GetStreamRules("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddStreamRule(t *testing.T) {
	store := inmemory.NewStore("")
	if err := store.AddStreamRule(nil); err == nil {
		t.Fatal("rule is nil")
	}
	if err := store.AddStreamRule(&graylog.StreamRule{}); err != nil {
		t.Fatal(err)
	}
}

func TestUpdateStreamRule(t *testing.T) {
}

func TestDeleteStreamRule(t *testing.T) {
	store := inmemory.NewStore("")
	if err := store.DeleteStreamRule("", ""); err != nil {
		t.Fatal(err)
	}
}
