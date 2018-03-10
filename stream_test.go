package graylog

import (
	"testing"
)

func addDummyStream(server *MockServer) (*IndexSet, *Stream, error) {
	indexSet, _, err := server.AddIndexSet(dummyNewIndexSet())
	if err != nil {
		return nil, nil, err
	}
	stream := dummyNewStream()
	stream.IndexSetId = indexSet.Id
	stream, _, err = server.AddStream(stream)
	return indexSet, stream, err
}

func dummyNewStream() *Stream {
	return &Stream{
		MatchingType: "AND",
		Description:  "Stream containing all messages",
		Rules:        []StreamRule{},
		Title:        "All messages",
	}
}

func dummyStream() *Stream {
	return &Stream{
		Id:              "000000000000000000000001",
		CreatorUserId:   "local:admin",
		Outputs:         []Output{},
		MatchingType:    "AND",
		Description:     "Stream containing all messages",
		CreatedAt:       "2018-02-20T11:37:19.371Z",
		Rules:           []StreamRule{},
		AlertConditions: []AlertCondition{},
		AlertReceivers: &AlertReceivers{
			Emails: []string{},
			Users:  []string{},
		},
		Title:      "All messages",
		IndexSetId: "5a8c086fc006c600013ca6f5",
		// "content_pack": null,
	}
}

func TestGetStreams(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	_, stream, err := addDummyStream(server)
	if err != nil {
		t.Fatal(err)
	}
	streams, total, _, err := client.GetStreams()
	if err != nil {
		t.Fatal("Failed to GetStreams", err)
	}
	if total != 1 {
		t.Fatalf("total == %d, wanted %d", total, 1)
	}
	if streams[0].Id != stream.Id {
		t.Fatalf("streams[0].Id == %s, wanted %s", streams[0].Id, stream.Id)
	}
}

func TestCreateStream(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	stream := dummyStream()
	if _, _, err := client.CreateStream(stream); err == nil {
		t.Fatal("CreateStream() must be failed")
	}
	indexSet, _, err := server.AddIndexSet(dummyNewIndexSet())
	if err != nil {
		t.Fatal(err)
	}
	stream = dummyNewStream()
	stream.IndexSetId = indexSet.Id
	id, _, err := client.CreateStream(stream)
	if err != nil {
		t.Fatal("Failed to CreateStream", err)
	}
	if id == "" {
		t.Fatal(`client.CreateStream() == ""`)
	}
	stream.Id = "h"
	if _, _, err := client.CreateStream(stream); err == nil {
		t.Fatal("id must be empty")
	}
	stream.Id = ""
	stream.CreatorUserId = "h"
	if _, _, err := client.CreateStream(stream); err == nil {
		t.Fatal("creator_user_id must be empty")
	}
	stream.CreatorUserId = ""
	stream.CreatedAt = "h"
	if _, _, err := client.CreateStream(stream); err == nil {
		t.Fatal("created_at must be empty")
	}
	stream.CreatedAt = ""
	stream.Disabled = true
	if _, _, err := client.CreateStream(stream); err == nil {
		t.Fatal("disabled must be false")
	}
	stream.Disabled = false
	stream.IsDefault = true
	if _, _, err := client.CreateStream(stream); err == nil {
		t.Fatal("is_default must be false")
	}

	copiedStream := *stream
	stream.IsDefault = false
	stream.Title = ""
	if _, _, err := client.CreateStream(stream); err == nil {
		t.Fatal("title is required")
	}
	stream.Title = copiedStream.Title
	stream.IndexSetId = ""
	if _, _, err := client.CreateStream(stream); err == nil {
		t.Fatal("index_set_id is required")
	}
	stream.IndexSetId = copiedStream.IndexSetId
	stream.AlertReceivers = &AlertReceivers{}
	if _, _, err := client.CreateStream(stream); err == nil {
		t.Fatal("alert_receiver is required")
	}
}

func TestGetEnabledStreams(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	_, stream, err := addDummyStream(server)
	if err != nil {
		t.Fatal(err)
	}
	exp := []Stream{*stream}
	streams, total, _, err := client.GetEnabledStreams()
	if err != nil {
		t.Fatal("Failed to GetStreams", err)
	}
	if total != 1 {
		t.Fatalf("total == %d, wanted %d", total, 1)
	}
	if streams[0].Id != exp[0].Id {
		t.Fatalf("streams[0].Id == %s, wanted %s", streams[0].Id, exp[0].Id)
	}

	stream.Disabled = true
	if _, err := server.UpdateStream(stream); err != nil {
		t.Fatal(err)
	}
	streams, total, _, err = client.GetEnabledStreams()
	if err != nil {
		t.Fatal("Failed to GetStreams", err)
	}
	if total != 0 {
		t.Fatalf("total == %d, wanted %d", total, 0)
	}
}

func TestGetStream(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	_, stream, err := addDummyStream(server)
	if err != nil {
		t.Fatal(err)
	}

	act, _, err := client.GetStream(stream.Id)
	if err != nil {
		t.Fatal("Failed to GetStream", err)
	}
	if act.Title != stream.Title {
		t.Fatalf("act.Title == %s, wanted %s", act.Title, stream.Title)
	}
	if _, _, err := client.GetStream(""); err == nil {
		t.Fatal("id is required")
	}
	if _, _, err := client.GetStream("h"); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}
}

func TestUpdateStream(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	_, stream, err := addDummyStream(server)
	if err != nil {
		t.Fatal(err)
	}

	stream.Description = "changed!"
	updatedStream, _, err := client.UpdateStream(stream.Id, stream)
	if err != nil {
		t.Fatal("Failed to UpdateStream", err)
	}
	if updatedStream == nil {
		t.Fatal("UpdateStream() == nil, nil")
	}
	if updatedStream.Title != stream.Title {
		t.Fatalf(
			"updatedStream.Title == %s, wanted %s",
			updatedStream.Title, stream.Title)
	}
	if _, _, err := client.UpdateStream("", stream); err == nil {
		t.Fatal("id is required")
	}
	if _, _, err := client.UpdateStream("h", stream); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}
}

func TestDeleteStream(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	_, stream, err := addDummyStream(server)
	if err != nil {
		t.Fatal(err)
	}

	if _, err = client.DeleteStream(""); err == nil {
		t.Fatal("id is required")
	}
	if _, err := client.DeleteStream("h"); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}
	if _, err := client.DeleteStream(stream.Id); err != nil {
		t.Fatal("Failed to DeleteStream", err)
	}
	s := len(server.streams)
	if s != 0 {
		t.Fatalf("len(server.streams) == %d, wanted 0", s)
	}
}

func TestPauseStream(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	_, stream, err := addDummyStream(server)
	if err != nil {
		t.Fatal(err)
	}

	if _, err = client.PauseStream(stream.Id); err != nil {
		t.Fatal("Failed to PauseStream", err)
	}
	if _, err := client.PauseStream(""); err == nil {
		t.Fatal("id is required")
	}
	if _, err := client.PauseStream("h"); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}
	// TODO test pause
}

func TestResumeStream(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	_, stream, err := addDummyStream(server)
	if err != nil {
		t.Fatal(err)
	}

	if _, err = client.ResumeStream(stream.Id); err != nil {
		t.Fatal("Failed to ResumeStream", err)
	}

	if _, err = client.ResumeStream(""); err == nil {
		t.Fatal("id is required")
	}

	if _, err = client.ResumeStream("h"); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}
	// TODO test resume
}
