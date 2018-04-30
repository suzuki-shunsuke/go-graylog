package client_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetStreamRules(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	stream, f, err := testutil.GetStream(client, server, 2)
	if err != nil {
		t.Fatal(err)
	}
	if f != nil {
		defer f(stream.ID)
	}

	if _, _, _, err := client.GetStreamRules(stream.ID); err != nil {
		t.Fatal("Failed to GetStreamRules", err)
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
	if server != nil {
		defer server.Close()
	}
	stream, f, err := testutil.GetStream(client, server, 2)
	if err != nil {
		t.Fatal(err)
	}
	if f != nil {
		defer f(stream.ID)
	}
	rule := testutil.StreamRule()
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
	if server != nil {
		defer server.Close()
	}
	stream, f, err := testutil.GetStream(client, server, 2)
	if err != nil {
		t.Fatal(err)
	}
	if f != nil {
		defer f(stream.ID)
	}
	rules, _, _, err := client.GetStreamRules(stream.ID)
	if err != nil {
		t.Fatal(err)
	}
	rule := testutil.StreamRule()
	if len(rules) == 0 {
		rule.StreamID = stream.ID
		if _, err := client.CreateStreamRule(rule); err != nil {
			t.Fatal(err)
		}
	} else {
		rule = &(rules[0])
	}

	rule.Description += " changed!"
	if _, err := client.UpdateStreamRule(rule); err != nil {
		t.Fatal(err)
	}
	streamID := rule.StreamID
	rule.StreamID = ""
	if _, err := client.UpdateStreamRule(rule); err == nil {
		t.Fatal("stream id is required")
	}
	rule.StreamID = streamID
	// ruleID = rule.ID
	rule.ID = ""
	if _, err := client.UpdateStreamRule(rule); err == nil {
		t.Fatal("stream rule id is required")
	}
	rule.ID = "h"
	if _, err := client.UpdateStreamRule(rule); err == nil {
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
	if server != nil {
		defer server.Close()
	}
	stream, f, err := testutil.GetStream(client, server, 2)
	if err != nil {
		t.Fatal(err)
	}
	if f != nil {
		defer f(stream.ID)
	}
	rules, _, _, err := client.GetStreamRules(stream.ID)
	if err != nil {
		t.Fatal(err)
	}
	rule := testutil.StreamRule()
	if len(rules) == 0 {
		rule.StreamID = stream.ID
		if _, err := client.CreateStreamRule(rule); err != nil {
			t.Fatal(err)
		}
	} else {
		rule = &(rules[0])
	}

	if _, err := client.DeleteStreamRule("", rule.ID); err == nil {
		t.Fatal("stream id is required")
	}
	if _, err := client.DeleteStreamRule(rule.StreamID, ""); err == nil {
		t.Fatal("stream rule id is required")
	}
	if _, err := client.DeleteStreamRule(rule.StreamID, rule.ID); err != nil {
		t.Fatal(err)
	}
	if _, _, err := client.GetStreamRule(rule.StreamID, rule.ID); err == nil {
		t.Fatal("stream rule should be deleted")
	}
}
