package test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetStreamRules(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	_, stream, err := addDummyStream(server)
	if err != nil {
		t.Fatal(err)
	}
	streamRule := testutil.DummyNewStreamRule()
	streamRule.StreamID = stream.ID
	if _, err := server.AddStreamRule(streamRule); err != nil {
		t.Fatal(err)
	}
	rules, total, _, err := client.GetStreamRules(stream.ID)
	if err != nil {
		t.Fatal("Failed to GetStreamRules", err)
	}
	if total != 1 {
		t.Fatalf("total == %d, wanted %d", total, 1)
	}
	if len(rules) != 1 {
		t.Fatalf("len(rules) == %d, wanted %d", len(rules), 1)
	}
	if _, _, _, err := client.GetStreamRules("h"); err == nil {
		t.Fatal(`no stream with id "h" is found`)
	}
}

func TestCreateStreamRule(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	is := testutil.IndexSet("hoge")
	if _, err := server.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err = server.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	rule := testutil.DummyNewStreamRule()
	rule.StreamID = stream.ID
	if _, err := client.CreateStreamRule(rule); err != nil {
		t.Fatal(err)
	}
	if _, err := client.CreateStreamRule(rule); err == nil {
		t.Fatal("stream rule id should be empty")
	}
	rule.ID = ""
	rule.StreamID = ""
	if _, err := client.CreateStreamRule(rule); err == nil {
		t.Fatal("stream id is required")
	}
	rule.StreamID = "h"
	if _, err := client.CreateStreamRule(rule); err == nil {
		t.Fatal(`no stream with id "h" is not found`)
	}

	if _, err := client.CreateStreamRule(nil); err == nil {
		t.Fatal("stream rule is nil")
	}
}

func TestUpdateStreamRule(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
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

	rule.Description += " changed!"
	if _, err := client.UpdateStreamRule(&rule); err != nil {
		t.Fatal(err)
	}
	streamID := rule.StreamID
	rule.StreamID = ""
	if _, err := client.UpdateStreamRule(&rule); err == nil {
		t.Fatal("stream id is required")
	}
	rule.StreamID = streamID
	// ruleID = rule.ID
	rule.ID = ""
	if _, err := client.UpdateStreamRule(&rule); err == nil {
		t.Fatal("stream rule id is required")
	}
	rule.ID = "h"
	if _, err := client.UpdateStreamRule(&rule); err == nil {
		t.Fatal(`no stream rule with id "h" is not found`)
	}

	if _, err := client.UpdateStreamRule(nil); err == nil {
		t.Fatal("stream rule is nil")
	}
}

func TestDeleteStreamRule(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
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
	if _, err := client.DeleteStreamRule("", rule.ID); err == nil {
		t.Fatal("stream id is required")
	}
	if _, err := client.DeleteStreamRule(rule.StreamID, ""); err == nil {
		t.Fatal("stream rule id is required")
	}
	if _, err := client.DeleteStreamRule(rule.StreamID, rule.ID); err != nil {
		t.Fatal(err)
	}
	if _, _, err := server.GetStreamRule(rule.StreamID, rule.ID); err == nil {
		t.Fatal("stream rule should be deleted")
	}
}

func TestGetStreamRule(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
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
	if _, _, err := client.GetStreamRule("", rule.ID); err == nil {
		t.Fatal("stream id is required")
	}
	if _, _, err := client.GetStreamRule(rule.StreamID, ""); err == nil {
		t.Fatal("stream rule id is required")
	}
	if _, _, err := client.GetStreamRule("h", rule.ID); err == nil {
		t.Fatal(`no stream with id "h" is found`)
	}
	if _, _, err := client.GetStreamRule(rule.StreamID, "h"); err == nil {
		t.Fatal(`no stream rule with id "h" is found`)
	}
	r, _, err := client.GetStreamRule(rule.StreamID, rule.ID)
	if err != nil {
		t.Fatal(err)
	}
	if r.ID != rule.ID {
		t.Fatalf("rule.ID = %s, wanted %s", r.ID, rule.ID)
	}
}
