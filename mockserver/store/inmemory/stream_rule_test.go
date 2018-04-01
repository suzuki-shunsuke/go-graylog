package inmemory_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/inmemory"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestHasStreamRule(t *testing.T) {
	store := inmemory.NewStore("")
	_, err := store.HasStreamRule("", "")
	if err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	rule := testutil.StreamRule()
	if err := store.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	rule.StreamID = stream.ID
	if err := store.AddStreamRule(rule); err != nil {
		t.Fatal(err)
	}
	_, err = store.HasStreamRule(stream.ID, rule.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetStreamRule(t *testing.T) {
	store := inmemory.NewStore("")
	_, err := store.GetStreamRule("", "")
	if err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	rule := testutil.StreamRule()
	if err := store.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	rule.StreamID = stream.ID
	if err := store.AddStreamRule(rule); err != nil {
		t.Fatal(err)
	}
	_, err = store.GetStreamRule(stream.ID, rule.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetStreamRule(stream.ID, "")
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
	stream := testutil.Stream()
	rule := testutil.StreamRule()
	if err := store.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	rule.StreamID = stream.ID
	if err := store.AddStreamRule(rule); err != nil {
		t.Fatal(err)
	}
	_, err = store.GetStreamRules(stream.ID)
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
	store := inmemory.NewStore("")
	stream := testutil.Stream()
	if err := store.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	rule := testutil.StreamRule()
	rule.StreamID = stream.ID
	if err := store.AddStreamRule(rule); err != nil {
		t.Fatal(err)
	}
	if err := store.UpdateStreamRule(rule); err != nil {
		t.Fatal(err)
	}
	rule.StreamID = ""
	if err := store.UpdateStreamRule(rule); err == nil {
		t.Fatal("stream id is empty")
	}
	if err := store.UpdateStreamRule(nil); err == nil {
		t.Fatal("rule is nil")
	}
}

func TestDeleteStreamRule(t *testing.T) {
	store := inmemory.NewStore("")
	if err := store.DeleteStreamRule("", ""); err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	rule := testutil.StreamRule()
	if err := store.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	rule.StreamID = stream.ID
	if err := store.AddStreamRule(rule); err != nil {
		t.Fatal(err)
	}
	if err := store.DeleteStreamRule(stream.ID, rule.ID); err != nil {
		t.Fatal(err)
	}
}
