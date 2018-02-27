package graylog

import (
	// "reflect"
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
		t.Error(err)
		return
	}
	defer server.Close()
	stream := dummyStream()
	exp := []Stream{*stream}
	server.Streams[stream.Id] = *stream
	streams, total, err := client.GetStreams()
	if err != nil {
		t.Error("Failed to GetStreams", err)
		return
	}
	if total != 1 {
		t.Errorf("total == %d, wanted %d", total, 1)
	}
	if streams[0].Id != exp[0].Id {
		t.Errorf("streams[0].Id == %s, wanted %s", streams[0].Id, exp[0].Id)
	}
	// if !reflect.DeepEqual(streams, exp) {
	// 	t.Errorf("client.GetStreams() %v != %v", streams, exp)
	// }
}

func TestCreateStream(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Close()
	stream := dummyStream()
	_, err = client.CreateStream(stream)
	if err == nil {
		t.Error("CreateStream() must be failed")
		return
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
		t.Error("Failed to CreateStream", err)
		return
	}
	if id == "" {
		t.Error(`client.CreateStream() == ""`)
		return
	}
}

func TestGetEnabledStreams(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Close()
	stream := dummyStream()
	exp := []Stream{*stream}
	server.Streams[stream.Id] = *stream
	streams, total, err := client.GetEnabledStreams()
	if err != nil {
		t.Error("Failed to GetStreams", err)
		return
	}
	if total != 1 {
		t.Errorf("total == %d, wanted %d", total, 1)
	}
	if streams[0].Id != exp[0].Id {
		t.Errorf("streams[0].Id == %s, wanted %s", streams[0].Id, exp[0].Id)
	}

	stream.Disabled = true
	server.Streams[stream.Id] = *stream
	streams, total, err = client.GetEnabledStreams()
	if err != nil {
		t.Error("Failed to GetStreams", err)
		return
	}
	if total != 0 {
		t.Errorf("total == %d, wanted %d", total, 0)
	}
}

func TestGetStream(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Close()
	exp := dummyStream()
	server.Streams[exp.Id] = *exp
	act, err := client.GetStream(exp.Id)
	if err != nil {
		t.Error("Failed to GetStream", err)
		return
	}
	if act.Title != exp.Title {
		t.Errorf("act.Title == %s, wanted %s", act.Title, exp.Title)
	}
}

func TestUpdateStream(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Close()
	stream := dummyStream()
	server.Streams[stream.Id] = *stream
	stream.Description = "changed!"
	updatedStream, err := client.UpdateStream(stream.Id, stream)
	if err != nil {
		t.Error("Failed to UpdateStream", err)
		return
	}
	if updatedStream.Title != stream.Title {
		t.Errorf(
			"updatedStream.Title == %s, wanted %s",
			updatedStream.Title, stream.Title)
	}
}

func TestDeleteStream(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Close()
	stream := dummyStream()
	server.Streams[stream.Id] = *stream
	err = client.DeleteStream(stream.Id)
	if err != nil {
		t.Error("Failed to DeleteStream", err)
		return
	}
	s := len(server.Streams)
	if s != 0 {
		t.Errorf("len(server.Streams) == %d, wanted 0", s)
		return
	}
}

func TestPauseStream(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Close()
	stream := dummyStream()
	server.Streams[stream.Id] = *stream
	err = client.PauseStream(stream.Id)
	if err != nil {
		t.Error("Failed to PauseStream", err)
		return
	}
	s := len(server.Streams)
	if s != 1 {
		t.Errorf("len(server.Streams) == %d, wanted 1", s)
		return
	}
	// TODO test pause
}

func TestResumeStream(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Close()
	stream := dummyStream()
	server.Streams[stream.Id] = *stream
	err = client.ResumeStream(stream.Id)
	if err != nil {
		t.Error("Failed to ResumeStream", err)
		return
	}
	s := len(server.Streams)
	if s != 1 {
		t.Errorf("len(server.Streams) == %d, wanted 1", s)
		return
	}
	// TODO test resume
}
