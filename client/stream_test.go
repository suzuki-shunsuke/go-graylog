package client_test

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetStreams(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, _, _, err := client.GetStreams(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestCreateStream(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	// nil check
	if _, err := client.CreateStream(ctx, nil); err == nil {
		t.Fatal("stream is nil")
	}
	// success
	u, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet(u.String())
	if _, err := client.CreateIndexSet(ctx, is); err != nil {
		t.Fatal(err)
	}
	testutil.WaitAfterCreateIndexSet(server)
	// clean
	defer func(id string) {
		if _, err := client.DeleteIndexSet(ctx, id); err != nil {
			t.Fatal(err)
		}
		testutil.WaitAfterDeleteIndexSet(server)
	}(is.ID)

	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := client.CreateStream(ctx, stream); err != nil {
		t.Fatal(err)
	}
	// clean
	defer client.DeleteStream(ctx, stream.ID)
}

func TestGetEnabledStreams(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	_, total, _, err := client.GetEnabledStreams(ctx)
	if err != nil {
		t.Fatal("Failed to GetStreams", err)
	}
	if total != 1 {
		t.Fatalf("total == %d, wanted %d", total, 1)
	}
}

func TestGetStream(t *testing.T) {
	ctx := context.Background()
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
	if _, err := client.CreateIndexSet(ctx, is); err != nil {
		t.Fatal(err)
	}
	testutil.WaitAfterCreateIndexSet(server)
	defer func(id string) {
		client.DeleteIndexSet(ctx, id)
		testutil.WaitAfterDeleteIndexSet(server)
	}(is.ID)
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := client.CreateStream(ctx, stream); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteStream(ctx, stream.ID)

	r, _, err := client.GetStream(ctx, stream.ID)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("stream is nil")
	}
	if r.ID != stream.ID {
		t.Fatalf(`stream.ID = "%s", wanted "%s"`, r.ID, stream.ID)
	}
	if _, _, err := client.GetStream(ctx, ""); err == nil {
		t.Fatal("id is required")
	}
	if _, _, err := client.GetStream(ctx, "h"); err == nil {
		t.Fatal("stream should not be found")
	}
}

func TestUpdateStream(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	stream, f, err := testutil.GetStream(ctx, client, server, 2)
	if err != nil {
		t.Fatal(err)
	}
	if f != nil {
		defer f(stream.ID)
	}

	stream.Description = "changed!"
	if _, err := client.UpdateStream(ctx, stream); err != nil {
		t.Fatal(err)
	}
	stream.ID = ""
	if _, err := client.UpdateStream(ctx, stream); err == nil {
		t.Fatal("id is required")
	}
	stream.ID = "h"
	if _, err := client.UpdateStream(ctx, stream); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}
	if _, err := client.UpdateStream(ctx, nil); err == nil {
		t.Fatal("stream is nil")
	}
}

func TestDeleteStream(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	// id required
	if _, err := client.DeleteStream(ctx, ""); err == nil {
		t.Fatal("id is required")
	}
	// invalid id
	if _, err := client.DeleteStream(ctx, "h"); err == nil {
		t.Fatal(`no stream with id "h" is found`)
	}
}

func TestPauseStream(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, err := client.PauseStream(ctx, ""); err == nil {
		t.Fatal("id is required")
	}
	if _, err := client.PauseStream(ctx, "h"); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}

	u, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet(u.String())
	if _, err := client.CreateIndexSet(ctx, is); err != nil {
		t.Fatal(err)
	}
	testutil.WaitAfterCreateIndexSet(server)
	defer func(id string) {
		client.DeleteIndexSet(ctx, id)
		testutil.WaitAfterDeleteIndexSet(server)
	}(is.ID)
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := client.CreateStream(ctx, stream); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteStream(ctx, stream.ID)

	if _, err = client.PauseStream(ctx, stream.ID); err != nil {
		t.Fatal("Failed to PauseStream", err)
	}
	// TODO test pause
}

func TestResumeStream(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, err = client.ResumeStream(ctx, ""); err == nil {
		t.Fatal("id is required")
	}

	if _, err = client.ResumeStream(ctx, "h"); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}

	u, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet(u.String())
	if _, err := client.CreateIndexSet(ctx, is); err != nil {
		t.Fatal(err)
	}
	testutil.WaitAfterCreateIndexSet(server)
	defer func(id string) {
		client.DeleteIndexSet(ctx, id)
		testutil.WaitAfterDeleteIndexSet(server)
	}(is.ID)
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := client.CreateStream(ctx, stream); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteStream(ctx, stream.ID)

	if _, err = client.ResumeStream(ctx, stream.ID); err != nil {
		t.Fatal("Failed to ResumeStream", err)
	}
	// TODO test resume
}
