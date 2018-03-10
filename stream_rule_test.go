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
	server.streams[stream.Id] = *stream
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
	indexSet := dummyNewIndexSet()
	is, _, err := server.AddIndexSet(indexSet)
	if err != nil {
		t.Fatal(err)
	}
	stream := dummyNewStream()
	stream.IndexSetId = is.Id
	stream, _, err = server.AddStream(stream)
	if err != nil {
		t.Fatal(err)
	}
	rule := dummyNewStreamRule()
	rule.StreamId = stream.Id
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
// 	stream.IndexSetId = is.Id
// 	is, _, err = client.CreateStream(stream)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	server.streams[stream.Id] = *stream
// 	stream.Description = "changed!"
// 	updatedStream, _, err := client.UpdateStream(stream.Id, stream)
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
