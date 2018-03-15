package mockserver_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func testUpdateStreamStatusCode(
	t *testing.T, endpoint string, body io.Reader, statusCode int,
) error {
	req, err := http.NewRequest(
		http.MethodPut, endpoint, body)
	if err != nil {
		return err
	}
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

func TestMockServerHandleUpdateStream(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSet := testutil.DummyNewIndexSet("hoge")
	is, _, err := server.AddIndexSet(indexSet)
	if err != nil {
		t.Fatal(err)
	}
	stream := testutil.DummyNewStream()
	stream.IndexSetID = is.ID
	s, _, err := server.AddStream(stream)
	if err != nil {
		t.Fatal(err)
	}
	endpoint := client.Endpoints.Stream(s.ID)

	body := bytes.NewBuffer([]byte("hoge"))
	if err := testUpdateStreamStatusCode(t, endpoint, body, 400); err != nil {
		t.Fatal(err)
	}

	body = bytes.NewBuffer([]byte(`{"title": 0}`))
	if err := testUpdateStreamStatusCode(t, endpoint, body, 400); err != nil {
		t.Fatal(err)
	}

	body = bytes.NewBuffer([]byte(`{"description": 0}`))
	if err := testUpdateStreamStatusCode(t, endpoint, body, 400); err != nil {
		t.Fatal(err)
	}

	body = bytes.NewBuffer([]byte(`{"matching_type": 0}`))
	if err := testUpdateStreamStatusCode(t, endpoint, body, 400); err != nil {
		t.Fatal(err)
	}

	body = bytes.NewBuffer([]byte(`{"remove_matches_from_default_stream": 0}`))
	if err := testUpdateStreamStatusCode(t, endpoint, body, 400); err != nil {
		t.Fatal(err)
	}

	body = bytes.NewBuffer([]byte(`{"index_set_id": 0}`))
	if err := testUpdateStreamStatusCode(t, endpoint, body, 400); err != nil {
		t.Fatal(err)
	}
}

func TestMockServerHandleCreateStream(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPost, client.Endpoints.Streams, body)
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

func TestGetStreams(t *testing.T) {
	test.TestGetStreams(t)
}

func TestCreateStream(t *testing.T) {
	test.TestCreateStream(t)
}

func TestGetEnabledStreams(t *testing.T) {
	test.TestGetEnabledStreams(t)
}

func TestGetStream(t *testing.T) {
	test.TestGetStream(t)
}

func TestUpdateStream(t *testing.T) {
	test.TestUpdateStream(t)
}

func TestDeleteStream(t *testing.T) {
	test.TestDeleteStream(t)
}

func TestPauseStream(t *testing.T) {
	test.TestPauseStream(t)
}

func TestResumeStream(t *testing.T) {
	test.TestResumeStream(t)
}
