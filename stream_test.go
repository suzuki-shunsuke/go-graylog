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
	streams, total, err := client.GetStreams()
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
	_, err = client.CreateStream(stream)
	if err == nil {
		t.Fatal("CreateStream() must be failed")
	}
	stream = &Stream{
		MatchingType: "AND",
		Description:  "Stream containing all messages",
		Rules:        []StreamRule{},
		Title:        "All messages",
		IndexSetId:   "5a8c086fc006c600013ca6f5",
	}
	id, err := client.CreateStream(stream)
	if err != nil {
		t.Fatal("Failed to CreateStream", err)
	}
	if id == "" {
		t.Fatal(`client.CreateStream() == ""`)
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
	streams, total, err := client.GetEnabledStreams()
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
	streams, total, err = client.GetEnabledStreams()
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
	act, err := client.GetStream(exp.Id)
	if err != nil {
		t.Fatal("Failed to GetStream", err)
	}
	if act.Title != exp.Title {
		t.Fatalf("act.Title == %s, wanted %s", act.Title, exp.Title)
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
	updatedStream, err := client.UpdateStream(stream.Id, stream)
	if err != nil {
		t.Fatal("Failed to UpdateStream", err)
	}
	if updatedStream.Title != stream.Title {
		t.Fatalf(
			"updatedStream.Title == %s, wanted %s",
			updatedStream.Title, stream.Title)
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
	err = client.DeleteStream(stream.Id)
	if err != nil {
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
	err = client.PauseStream(stream.Id)
	if err != nil {
		t.Fatal("Failed to PauseStream", err)
	}
	s := len(server.Streams)
	if s != 1 {
		t.Fatalf("len(server.Streams) == %d, wanted 1", s)
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
	err = client.ResumeStream(stream.Id)
	if err != nil {
		t.Fatal("Failed to ResumeStream", err)
	}
	s := len(server.Streams)
	if s != 1 {
		t.Fatalf("len(server.Streams) == %d, wanted 1", s)
	}
	// TODO test resume
}
