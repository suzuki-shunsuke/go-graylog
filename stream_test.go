package graylog

import (
	"testing"
)

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
	stream := dummyStream()
	exp := []Stream{*stream}
	server.Streams[stream.Id] = *stream
	streams, total, _, err := client.GetStreams()
	if err != nil {
		t.Fatal("Failed to GetStreams", err)
	}
	if total != 1 {
		t.Fatalf("total == %d, wanted %d", total, 1)
	}
	if streams[0].Id != exp[0].Id {
		t.Fatalf("streams[0].Id == %s, wanted %s", streams[0].Id, exp[0].Id)
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
	stream = &Stream{
		MatchingType: "AND",
		Description:  "Stream containing all messages",
		Rules:        []StreamRule{},
		Title:        "All messages",
		IndexSetId:   "5a8c086fc006c600013ca6f5",
	}
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
	stream := dummyStream()
	exp := []Stream{*stream}
	server.Streams[stream.Id] = *stream
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
	server.Streams[stream.Id] = *stream
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
	exp := dummyStream()
	server.Streams[exp.Id] = *exp
	act, _, err := client.GetStream(exp.Id)
	if err != nil {
		t.Fatal("Failed to GetStream", err)
	}
	if act.Title != exp.Title {
		t.Fatalf("act.Title == %s, wanted %s", act.Title, exp.Title)
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
	stream := dummyStream()
	server.Streams[stream.Id] = *stream
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
	stream := dummyStream()
	server.Streams[stream.Id] = *stream
	if _, err = client.DeleteStream(""); err == nil {
		t.Fatal("id is required")
	}
	if _, err := client.DeleteStream("h"); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}
	if _, err := client.DeleteStream(stream.Id); err != nil {
		t.Fatal("Failed to DeleteStream", err)
	}
	s := len(server.Streams)
	if s != 0 {
		t.Fatalf("len(server.Streams) == %d, wanted 0", s)
	}
}

func TestPauseStream(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	stream := dummyStream()
	server.Streams[stream.Id] = *stream

	if _, err = client.PauseStream(stream.Id); err != nil {
		t.Fatal("Failed to PauseStream", err)
	}
	s := len(server.Streams)
	if s != 1 {
		t.Fatalf("len(server.Streams) == %d, wanted 1", s)
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
	stream := dummyStream()
	server.Streams[stream.Id] = *stream

	if _, err = client.ResumeStream(stream.Id); err != nil {
		t.Fatal("Failed to ResumeStream", err)
	}
	s := len(server.Streams)
	if s != 1 {
		t.Fatalf("len(server.Streams) == %d, wanted 1", s)
	}

	if _, err = client.ResumeStream(""); err == nil {
		t.Fatal("id is required")
	}

	if _, err = client.ResumeStream("h"); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}
	// TODO test resume
}
