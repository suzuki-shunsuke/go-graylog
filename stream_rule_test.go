package graylog

import (
	"testing"
)

func dummyStreamRule() *StreamRule {
	return &StreamRule{
		Id:       "5a9b53c7c006c6000127f965",
		Type:     1,
		Value:    "test",
		StreamId: "5a94abdac006c60001f04fc1",
		Field:    "tag",
	}
}

func TestGetStreamRules(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	stream := dummyStream()
	streamRule := dummyStreamRule()
	streamRule.StreamId = stream.Id
	server.Streams[stream.Id] = *stream
	if err := server.AddStreamRule(streamRule); err != nil {
		t.Fatal(err)
	}
	rules, total, _, err := client.GetStreamRules(stream.Id)
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
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	stream := dummyStream()
	rule := dummyStreamRule()
	rule.Id = ""
	rule.StreamId = ""
	server.AddStream(stream)
	if _, _, err := client.CreateStreamRule(stream.Id, rule); err != nil {
		t.Fatal(err)
	}
}
