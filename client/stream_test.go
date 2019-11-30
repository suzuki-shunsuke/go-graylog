package client_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/gofrs/uuid"

	"github.com/suzuki-shunsuke/go-graylog/v8/client"
	"github.com/suzuki-shunsuke/go-graylog/v8/testdata"
	"github.com/suzuki-shunsuke/go-graylog/v8/testutil"
)

func TestClient_GetStreams(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/streams.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method: "GET",
								Path:   "/api/streams",
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: bodyStr,
							},
						},
					},
				},
			},
		},
	})

	streams, total, _, err := cl.GetStreams(ctx)
	require.Nil(t, err)
	require.Equal(t, testdata.Streams.Total, total)
	require.Equal(t, testdata.Streams.Streams, streams)
}

func TestClient_CreateStream(t *testing.T) {
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

func TestClient_GetEnabledStreams(t *testing.T) {
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

func TestClient_GetStream(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/stream.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method: "GET",
								Path:   "/api/streams/" + testdata.Stream.ID,
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: bodyStr,
							},
						},
					},
				},
			},
		},
	})

	_, _, err = cl.GetStream(ctx, "")
	require.NotNil(t, err)

	stream, _, err := cl.GetStream(ctx, testdata.Stream.ID)
	require.Nil(t, err)
	require.Equal(t, testdata.Stream, stream)
}

func TestClient_UpdateStream(t *testing.T) {
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

func TestClient_DeleteStream(t *testing.T) {
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

func TestClient_PauseStream(t *testing.T) {
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

func TestClient_ResumeStream(t *testing.T) {
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
