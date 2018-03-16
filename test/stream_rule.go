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
}

func TestCreateStreamRule(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	is := testutil.DummyNewIndexSet("hoge")
	if _, err := server.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream := testutil.DummyNewStream()
	stream.IndexSetID = is.ID
	if _, err = server.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	rule := testutil.DummyNewStreamRule()
	rule.StreamID = stream.ID
	if _, _, err := client.CreateStreamRule(rule); err != nil {
		t.Fatal(err)
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
		t.Fatal("Failed to UpdateStream", err)
	}
}
