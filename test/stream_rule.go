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

// func TestUpdateStreamRule(t *testing.T) {
// 	server, client, err := testutil.GetServerAndClient()
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
