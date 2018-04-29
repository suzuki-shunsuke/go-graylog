package client_test

import (
	"os"
	"testing"

	"github.com/satori/go.uuid"
	"github.com/suzuki-shunsuke/go-graylog/test"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetStreams(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, _, _, err := client.GetStreams(); err != nil {
		t.Fatal(err)
	}
}

func TestCreateStream(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	// nil check
	if _, err := client.CreateStream(nil); err == nil {
		t.Fatal("stream is nil")
	}
	// success
	u, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet(u.String())
	if _, err := client.CreateIndexSet(is); err != nil {
		t.Fatal(err)
	}
	testutil.WaitAfterCreateIndexSet(server)
	// clean
	defer func(id string) {
		if _, err := client.DeleteIndexSet(id); err != nil {
			t.Fatal(err)
		}
		testutil.WaitAfterDeleteIndexSet(server)
	}(is.ID)

	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := client.CreateStream(stream); err != nil {
		t.Fatal(err)
	}
	// clean
	defer client.DeleteStream(stream.ID)
}

func TestGetEnabledStreams(t *testing.T) {
	test.TestGetEnabledStreams(t)
}

func TestGetStream(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	u, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet(u.String())
	if _, err := client.CreateIndexSet(is); err != nil {
		t.Fatal(err)
	}
	testutil.WaitAfterCreateIndexSet(server)
	defer func(id string) {
		client.DeleteIndexSet(id)
		testutil.WaitAfterDeleteIndexSet(server)
	}(is.ID)
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := client.CreateStream(stream); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteStream(stream.ID)

	r, _, err := client.GetStream(stream.ID)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("stream is nil")
	}
	if r.ID != stream.ID {
		t.Fatalf(`stream.ID = "%s", wanted "%s"`, r.ID, stream.ID)
	}
	if _, _, err := client.GetStream(""); err == nil {
		t.Fatal("id is required")
	}
	if _, _, err := client.GetStream("h"); err == nil {
		t.Fatal("stream should not be found")
	}
}

func TestUpdateStream(t *testing.T) {
	test.TestUpdateStream(t)
}

func TestDeleteStream(t *testing.T) {
	if err := os.Setenv("GRAYLOG_WEB_ENDPOINT_URI", "http://localhost:9000/api"); err != nil {
		t.Fatal(err)
	}
	defer os.Unsetenv("GRAYLOG_WEB_ENDPOINT_URI")

	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	// id required
	if _, err := client.DeleteStream(""); err == nil {
		t.Fatal("id is required")
	}
	// invalid id
	if _, err := client.DeleteStream("h"); err == nil {
		t.Fatal(`no stream with id "h" is found`)
	}
}

func TestPauseStream(t *testing.T) {
	test.TestPauseStream(t)
}

func TestResumeStream(t *testing.T) {
	test.TestResumeStream(t)
}
