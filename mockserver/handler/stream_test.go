package handler_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog"
	. "github.com/suzuki-shunsuke/go-graylog/mockserver"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func testUpdateStreamStatusCode(
	endpoint, name, password string, body io.Reader, statusCode int,
) error {
	req, err := http.NewRequest(
		http.MethodPut, endpoint, body)
	if err != nil {
		return err
	}
	req.SetBasicAuth(name, password)
	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != statusCode {
		return fmt.Errorf("resp.StatusCode == %d, wanted 400", resp.StatusCode)
	}
	return nil
}

func TestGetStream(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	_, stream, err := addDummyStream(server)
	if err != nil {
		t.Fatal(err)
	}

	act, _, err := client.GetStream(stream.ID)
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

func TestGetStreams(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	streams, total, _, err := client.GetStreams()
	if err != nil {
		t.Fatal("Failed to GetStreams", err)
	}
	if total != 1 {
		t.Fatalf("total == %d, wanted %d", total, 1)
	}
	if len(streams) != 1 {
		t.Fatalf("len(stream) == %d, wanted %d", len(streams), 1)
	}
}

func TestHandleCreateStream(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	stream := testutil.DummyStream()
	if _, err := client.CreateStream(stream); err == nil {
		t.Fatal("CreateStream() must be failed")
	}
	indexSet := testutil.IndexSet("hoge")
	if _, err := server.AddIndexSet(indexSet); err != nil {
		t.Fatal(err)
	}
	stream = testutil.Stream()
	stream.IndexSetID = indexSet.ID

	if _, err := client.CreateStream(stream); err != nil {
		t.Fatal("Failed to CreateStream", err)
	}
	if stream.ID == "" {
		t.Fatal(`stream id is empty`)
	}
	if _, err := client.CreateStream(stream); err == nil {
		t.Fatal("id must be empty")
	}
	stream.ID = ""
	stream.CreatorUserID = "h"
	if _, err := client.CreateStream(stream); err == nil {
		t.Fatal("creator_user_id must be empty")
	}
	stream.CreatorUserID = ""
	stream.CreatedAt = "h"
	if _, err := client.CreateStream(stream); err == nil {
		t.Fatal("created_at must be empty")
	}
	stream.CreatedAt = ""
	stream.Disabled = true
	if _, err := client.CreateStream(stream); err == nil {
		t.Fatal("disabled must be false")
	}
	stream.Disabled = false
	stream.IsDefault = true
	if _, err := client.CreateStream(stream); err == nil {
		t.Fatal("is_default must be false")
	}

	copiedStream := *stream
	stream.IsDefault = false
	stream.Title = ""
	if _, err := client.CreateStream(stream); err == nil {
		t.Fatal("title is required")
	}
	stream.Title = copiedStream.Title
	stream.IndexSetID = ""
	if _, err := client.CreateStream(stream); err == nil {
		t.Fatal("index_set_id is required")
	}
	stream.IndexSetID = copiedStream.IndexSetID
	stream.AlertReceivers = &graylog.AlertReceivers{}
	if _, err := client.CreateStream(stream); err == nil {
		t.Fatal("alert_receiver is required")
	}
	if _, err := client.CreateStream(nil); err == nil {
		t.Fatal("stream is nil")
	}

	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPost, client.Endpoints().Streams(), body)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth(client.Name(), client.Password())
	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Fatalf("resp.StatusCode == %d, wanted 400", resp.StatusCode)
	}
}

func TestServerHandleUpdateStream(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	is := testutil.IndexSet("hoge")
	if _, err := server.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := server.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	u, err := client.Endpoints().Stream(stream.ID)
	if err != nil {
		t.Fatal(err)
	}
	endpoint := u.String()

	body := bytes.NewBuffer([]byte("hoge"))
	if err := testUpdateStreamStatusCode(endpoint, client.Name(), client.Password(), body, 400); err != nil {
		t.Fatal(err)
	}

	body = bytes.NewBuffer([]byte(`{"title": 0}`))
	if err := testUpdateStreamStatusCode(endpoint, client.Name(), client.Password(), body, 400); err != nil {
		t.Fatal(err)
	}

	body = bytes.NewBuffer([]byte(`{"description": 0}`))
	if err := testUpdateStreamStatusCode(endpoint, client.Name(), client.Password(), body, 400); err != nil {
		t.Fatal(err)
	}

	body = bytes.NewBuffer([]byte(`{"matching_type": 0}`))
	if err := testUpdateStreamStatusCode(endpoint, client.Name(), client.Password(), body, 400); err != nil {
		t.Fatal(err)
	}

	body = bytes.NewBuffer([]byte(`{"remove_matches_from_default_stream": 0}`))
	if err := testUpdateStreamStatusCode(endpoint, client.Name(), client.Password(), body, 400); err != nil {
		t.Fatal(err)
	}

	body = bytes.NewBuffer([]byte(`{"index_set_id": 0}`))
	if err := testUpdateStreamStatusCode(endpoint, client.Name(), client.Password(), body, 400); err != nil {
		t.Fatal(err)
	}

	// nil check
	if _, _, err := server.UpdateStream(nil); err == nil {
		t.Fatal("stream is nil")
	}

	stream, f, err := testutil.GetStream(client, server, 2)
	if err != nil {
		t.Fatal(err)
	}
	if f != nil {
		defer f(stream.ID)
	}

	stream.Description = "changed!"
	if _, err := client.UpdateStream(stream); err != nil {
		t.Fatal(err)
	}
	stream.ID = "h"
	if _, err := client.UpdateStream(stream); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}
	if _, err := client.UpdateStream(nil); err == nil {
		t.Fatal("stream is nil")
	}
}

func addDummyStream(server *Server) (*graylog.IndexSet, *graylog.Stream, error) {
	indexSet := testutil.IndexSet("hoge")
	if _, err := server.AddIndexSet(indexSet); err != nil {
		return nil, nil, err
	}
	stream := testutil.Stream()
	stream.IndexSetID = indexSet.ID
	_, err := server.AddStream(stream)
	return indexSet, stream, err
}

func TestDeleteStream(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
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
	if _, err := client.DeleteStream(stream.ID); err != nil {
		t.Fatal("Failed to DeleteStream", err)
	}
}

func TestGetEnabledStreams(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	_, total, _, err := client.GetEnabledStreams()
	if err != nil {
		t.Fatal("Failed to GetStreams", err)
	}
	if total != 1 {
		t.Fatalf("total == %d, wanted %d", total, 1)
	}
}

func TestPauseStream(t *testing.T) {
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
	if _, err = client.PauseStream(stream.ID); err != nil {
		t.Fatal("Failed to PauseStream", err)
	}
	if _, err := client.PauseStream("h"); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}
	// TODO test pause
}

func TestResumeStream(t *testing.T) {
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

	if _, err = client.ResumeStream(stream.ID); err != nil {
		t.Fatal("Failed to ResumeStream", err)
	}

	if _, err = client.ResumeStream("h"); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}
	// TODO test resume
}
