package graylog

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func testUpdateStream400(t *testing.T, endpoint string, body io.Reader) {
	req, err := http.NewRequest(
		http.MethodPut, endpoint, body)
	if err != nil {
		t.Fatal(err)
	}
	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Fatalf("resp.StatusCode == %d, wanted 400", resp.StatusCode)
	}
}

func TestMockServerHandleUpdateStream(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	stream := dummyStream()
	server.Streams[stream.Id] = *stream
	endpoint := client.endpoints.Stream(stream.Id)

	body := bytes.NewBuffer([]byte("hoge"))
	testUpdateStream400(t, endpoint, body)

	body = bytes.NewBuffer([]byte(`{"title": 0}`))
	testUpdateStream400(t, endpoint, body)

	body = bytes.NewBuffer([]byte(`{"description": 0}`))
	testUpdateStream400(t, endpoint, body)

	body = bytes.NewBuffer([]byte(`{"matching_type": 0}`))
	testUpdateStream400(t, endpoint, body)

	body = bytes.NewBuffer([]byte(`{"remove_matches_from_default_stream": 0}`))
	testUpdateStream400(t, endpoint, body)

	body = bytes.NewBuffer([]byte(`{"index_set_id": 0}`))
	testUpdateStream400(t, endpoint, body)
}

func TestMockServerHandleCreateStream(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPost, client.endpoints.Streams, body)
	if err != nil {
		t.Fatal(err)
	}
	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Fatalf("resp.StatusCode == %d, wanted 400", resp.StatusCode)
	}
}
