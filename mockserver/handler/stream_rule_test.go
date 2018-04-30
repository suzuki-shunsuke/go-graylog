package handler_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestHandleGetStreamRules(t *testing.T) {
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

func TestHandleGetStreamRule(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	streams, _, _, err := client.GetStreams()
	if err != nil {
		t.Fatal(err)
	}
	stream := streams[0]
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

func TestHandleCreateStreamRule(t *testing.T) {
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
	rule.StreamID = "h"
	if _, err := client.CreateStreamRule(rule); err == nil {
		t.Fatal(`no stream with id "h" is not found`)
	}

	pc := &plainClient{Name: client.Name(), Password: client.Password()}
	e, err := client.Endpoints().StreamRules(stream.ID)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := pc.Post(e.String(), `{"value": 0, "field": 0}`)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 400 {
		t.Fatalf("resp.StatusCode = %d, wanted 400", resp.StatusCode)
	}
}

func TestHandleUpdateStreamRule(t *testing.T) {
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

	ruleID := rule.ID
	streamID := rule.StreamID

	rule.Description += " changed!"
	if _, err := client.UpdateStreamRule(rule); err != nil {
		t.Fatal(err)
	}
	rule.StreamID = ""
	if _, err := client.UpdateStreamRule(rule); err == nil {
		t.Fatal("stream id is required")
	}
	rule.ID = "h"
	if _, err := client.UpdateStreamRule(rule); err == nil {
		t.Fatal(`no stream rule with id "h" is not found`)
	}

	pc := &plainClient{Name: client.Name(), Password: client.Password()}
	e, err := client.Endpoints().StreamRule(streamID, ruleID)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := pc.Put(e.String(), `{"value": 0, "field": 0}`)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 400 {
		t.Fatalf("resp.StatusCode = %d, wanted 400", resp.StatusCode)
	}
	resp, err = pc.Put(e.String(), `{"field": 0}`)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 400 {
		t.Fatalf("resp.StatusCode = %d, wanted 400", resp.StatusCode)
	}
}

func TestHandleDeleteStreamRule(t *testing.T) {
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
