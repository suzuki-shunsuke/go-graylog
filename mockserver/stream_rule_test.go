package mockserver_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetStreamRules(t *testing.T) {
	test.TestGetStreamRules(t)
}

func TestCreateStreamRule(t *testing.T) {
	test.TestCreateStreamRule(t)
}

func TestUpdateStreamRule(t *testing.T) {
	test.TestUpdateStreamRule(t)
}

func TestDeleteStreamRule(t *testing.T) {
	test.TestDeleteStreamRule(t)
}

func TestGetStreamRule(t *testing.T) {
	test.TestGetStreamRule(t)
}

func TestServerAddStreamRule(t *testing.T) {
	server, _, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	streams, err := server.GetStreams()
	if err != nil {
		t.Fatal(err)
	}
	stream := streams[0]
	rules, _, err := server.GetStreamRules(stream.ID)
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

func TestServerUpdateStreamRule(t *testing.T) {
	server, _, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	streams, err := server.GetStreams()
	if err != nil {
		t.Fatal(err)
	}
	stream := streams[0]
	rules, _, err := server.GetStreamRules(stream.ID)
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

func TestServerDeleteStreamRule(t *testing.T) {
	server, _, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	streams, err := server.GetStreams()
	if err != nil {
		t.Fatal(err)
	}
	stream := streams[0]
	rules, _, err := server.GetStreamRules(stream.ID)
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
