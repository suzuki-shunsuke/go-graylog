package handler_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestHandleGetStreamRules(t *testing.T) {
	test.TestGetStreamRules(t)
}

func TestHandleGetStreamRule(t *testing.T) {
	test.TestGetStreamRule(t)
}

func TestHandleCreateStreamRule(t *testing.T) {
	test.TestCreateStreamRule(t)
	server, _, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	streams, _, _, err := server.GetStreams()
	if err != nil {
		t.Fatal(err)
	}
	stream := streams[0]
	rules, _, _, err := server.GetStreamRules(stream.ID)
	if err != nil {
		t.Fatal(err)
	}
	rule := rules[0]
	rule.ID = ""
	rule.StreamID = ""
	if _, err := server.AddStreamRule(&rule); err == nil {
		t.Fatal("stream id is required")
	}
	rule.StreamID = "h"
	if _, err := server.AddStreamRule(&rule); err == nil {
		t.Fatal(`no stream with id "h" is found`)
	}
}

func TestHandleUpdateStreamRule(t *testing.T) {
	test.TestUpdateStreamRule(t)
	server, _, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	streams, _, _, err := server.GetStreams()
	if err != nil {
		t.Fatal(err)
	}
	stream := streams[0]
	rules, _, _, err := server.GetStreamRules(stream.ID)
	if err != nil {
		t.Fatal(err)
	}
	rule := rules[0]
	rule.StreamID = ""
	if _, err := server.UpdateStreamRule(&rule); err == nil {
		t.Fatal("stream id is required")
	}
	rule.StreamID = "h"
	if _, err := server.UpdateStreamRule(&rule); err == nil {
		t.Fatal(`no stream with id "h" is found`)
	}
}

func TestHandleDeleteStreamRule(t *testing.T) {
	test.TestDeleteStreamRule(t)
	server, _, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	streams, _, _, err := server.GetStreams()
	if err != nil {
		t.Fatal(err)
	}
	stream := streams[0]
	rules, _, _, err := server.GetStreamRules(stream.ID)
	if err != nil {
		t.Fatal(err)
	}
	rule := rules[0]
	if _, err := server.DeleteStreamRule("h", rule.ID); err == nil {
		t.Fatal(`no stream with id "h" is found`)
	}
	if _, err := server.DeleteStreamRule(rule.StreamID, "h"); err == nil {
		t.Fatal(`no stream rule with id "h" is found`)
	}
}
