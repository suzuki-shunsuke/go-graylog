package logic_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestAddStreamRule(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet("hoge")
	if _, err := lgc.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := lgc.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	rule := testutil.StreamRule()
	rule.StreamID = stream.ID
	if _, err := lgc.AddStreamRule(nil); err == nil {
		t.Fatal("stream is nil")
	}
	if _, err := lgc.AddStreamRule(rule); err != nil {
		t.Fatal(err)
	}
}

func TestUpdateStreamRule(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet("hoge")
	if _, err := lgc.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := lgc.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	rule := testutil.StreamRule()
	rule.StreamID = stream.ID
	if _, err := lgc.AddStreamRule(rule); err != nil {
		t.Fatal(err)
	}
	if _, err := lgc.UpdateStreamRule(nil); err == nil {
		t.Fatal("stream is nil")
	}
	if _, err := lgc.UpdateStreamRule(rule); err != nil {
		t.Fatal(err)
	}
	rule.ID = ""
	if _, err := lgc.UpdateStreamRule(rule); err == nil {
		t.Fatal("stream.ID is required")
	}
}

func TestDeleteStreamRule(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet("hoge")
	if _, err := lgc.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := lgc.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	rule := testutil.StreamRule()
	rule.StreamID = stream.ID
	if _, err := lgc.AddStreamRule(rule); err != nil {
		t.Fatal(err)
	}
	if _, err := lgc.DeleteStreamRule("", rule.ID); err == nil {
		t.Fatal("stream id is required")
	}
	if _, err := lgc.DeleteStreamRule(rule.StreamID, ""); err == nil {
		t.Fatal("stream rule id is required")
	}
	if _, err := lgc.DeleteStreamRule(rule.StreamID, rule.ID); err != nil {
		t.Fatal(err)
	}
}

func TestGetStreamRules(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet("hoge")
	if _, err := lgc.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := lgc.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	rule := testutil.StreamRule()
	rule.StreamID = stream.ID
	if _, err := lgc.AddStreamRule(rule); err != nil {
		t.Fatal(err)
	}
	if _, _, err := lgc.GetStreamRules(""); err == nil {
		t.Fatal("stream id is required")
	}
	if _, _, err := lgc.GetStreamRules(rule.StreamID); err != nil {
		t.Fatal(err)
	}
}

func TestGetStreamRule(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet("hoge")
	if _, err := lgc.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := lgc.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	rule := testutil.StreamRule()
	rule.StreamID = stream.ID
	if _, err := lgc.AddStreamRule(rule); err != nil {
		t.Fatal(err)
	}
	if _, _, err := lgc.GetStreamRule("", rule.ID); err == nil {
		t.Fatal("stream id is required")
	}
	if _, _, err := lgc.GetStreamRule(rule.StreamID, ""); err == nil {
		t.Fatal("stream rule id is required")
	}
	if _, _, err := lgc.GetStreamRule(rule.StreamID, rule.ID); err != nil {
		t.Fatal(err)
	}
}
