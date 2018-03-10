package graylog

import (
	"testing"
)

func dummyNewStreamRule() *StreamRule {
	return &StreamRule{
		Type:  1,
		Value: "test",
		Field: "tag",
	}
}

func dummyStreamRule() *StreamRule {
	return &StreamRule{
		ID:       "5a9b53c7c006c6000127f965",
		Type:     1,
		Value:    "test",
		StreamID: "5a94abdac006c60001f04fc1",
		Field:    "tag",
	}
}

func TestGetStreamRules(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	_, stream, err := addDummyStream(server)
	if err != nil {
		t.Fatal(err)
	}
	streamRule := dummyStreamRule()
	streamRule.StreamID = stream.ID
	if err := server.AddStreamRule(streamRule); err != nil {
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
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSet := dummyNewIndexSet()
	is, _, err := server.AddIndexSet(indexSet)
	if err != nil {
		t.Fatal(err)
	}
	stream := dummyNewStream()
	stream.IndexSetID = is.ID
	stream, _, err = server.AddStream(stream)
	if err != nil {
		t.Fatal(err)
	}
	rule := dummyNewStreamRule()
	rule.StreamID = stream.ID
	if _, _, err := client.CreateStreamRule(rule); err != nil {
		t.Fatal(err)
	}
}

// func TestUpdateStreamRule(t *testing.T) {
// 	server, client, err := getServerAndClient()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer server.Close()
// 	indexSet := dummyIndexSet()
// 	is, _, err := client.CreateIndexSet(indexSet)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	stream := dummyStream()
// 	stream.IndexSetID = is.ID
// 	is, _, err = client.CreateStream(stream)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	server.streams[stream.ID] = *stream
// 	stream.Description = "changed!"
// 	updatedStream, _, err := client.UpdateStream(stream.ID, stream)
// 	if err != nil {
// 		t.Fatal("Failed to UpdateStream", err)
// 	}
// 	if updatedStream == nil {
// 		t.Fatal("UpdateStream() == nil, nil")
// 	}
// 	if updatedStream.Title != stream.Title {
// 		t.Fatalf(
// 			"updatedStream.Title == %s, wanted %s",
// 			updatedStream.Title, stream.Title)
// 	}
// 	if _, _, err := client.UpdateStream("", stream); err == nil {
// 		t.Fatal("id is required")
// 	}
// 	if _, _, err := client.UpdateStream("h", stream); err == nil {
// 		t.Fatal(`no stream whose id is "h"`)
// 	}
// }
