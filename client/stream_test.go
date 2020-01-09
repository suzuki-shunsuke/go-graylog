package client_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog/v9/client"
	"github.com/suzuki-shunsuke/go-graylog/v9/testdata"
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
	require.Equal(t, testdata.Streams().Total, total)
	require.Equal(t, testdata.Streams().Streams, streams)
}

func TestClient_CreateStream(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/create_stream.json")
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
								Method: "POST",
								Path:   "/api/streams",
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
								BodyJSONString: bodyStr,
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: `{
  "stream_id": "5e151c31a1de18000d89a83f"
}`,
							},
						},
					},
				},
			},
		},
	})

	// nil check
	if _, err := cl.CreateStream(ctx, nil); err == nil {
		t.Fatal("stream is nil")
	}

	stream := testdata.CreateStream()
	if _, err := cl.CreateStream(ctx, &stream); err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "5e151c31a1de18000d89a83f", stream.ID)
}

func TestClient_GetEnabledStreams(t *testing.T) {
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
								Path:   "/api/streams/enabled",
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

	_, total, _, err := cl.GetEnabledStreams(ctx)
	require.Nil(t, err)
	require.Equal(t, testdata.Streams().Total, total)
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
								Path:   "/api/streams/" + testdata.Stream().ID,
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

	stream, _, err := cl.GetStream(ctx, testdata.Stream().ID)
	require.Nil(t, err)
	require.Equal(t, testdata.Stream(), stream)
}

func TestClient_UpdateStream(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	reqBuf, err := ioutil.ReadFile("../testdata/create_stream.json")
	require.Nil(t, err)
	reqBodyStr := string(reqBuf)

	respBuf, err := ioutil.ReadFile("../testdata/stream.json")
	require.Nil(t, err)
	respBodyStr := string(respBuf)

	streamID := "5e151c31a1de18000d89a83f"

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Matcher: &flute.Matcher{
								Path: "/api/streams/" + streamID,
							},
							Tester: &flute.Tester{
								Method: "PUT",
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
								BodyJSONString: reqBodyStr,
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: respBodyStr,
							},
						},
						{
							Matcher: &flute.Matcher{
								Path: "/api/streams/h",
							},
							Tester: &flute.Tester{
								Method: "PUT",
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
								BodyJSONString: reqBodyStr,
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 404,
								},
								BodyString: `{
  "type": "ApiError",
  "message": "Stream <h> not found!"
}`,
							},
						},
					},
				},
			},
		},
	})

	stream := testdata.CreateStream()

	if _, err := cl.UpdateStream(ctx, &stream); err == nil {
		t.Fatal("id is required")
	}
	stream = testdata.CreateStream()
	stream.ID = streamID
	if _, err := cl.UpdateStream(ctx, &stream); err != nil {
		t.Fatal(err)
	}
	stream = testdata.CreateStream()
	stream.ID = "h"
	if _, err := cl.UpdateStream(ctx, &stream); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}
	if _, err := cl.UpdateStream(ctx, nil); err == nil {
		t.Fatal("stream is nil")
	}
}

func TestClient_DeleteStream(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	streamID := "5e151c31a1de18000d89a83f"

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method: "DELETE",
								Path:   "/api/streams/" + streamID,
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 204,
								},
							},
						},
						{
							Tester: &flute.Tester{
								Method: "DELETE",
								Path:   "/api/streams/h",
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 404,
								},
								BodyString: `"message"
"Stream <h> not found!"`,
							},
						},
					},
				},
			},
		},
	})

	// id required
	if _, err := cl.DeleteStream(ctx, ""); err == nil {
		t.Fatal("id is required")
	}
	// invalid id
	if _, err := cl.DeleteStream(ctx, streamID); err != nil {
		t.Fatal(err)
	}
}

func TestClient_PauseStream(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	streamID := "5e151c31a1de18000d89a83f"

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method: "POST",
								Path:   "/api/streams/" + streamID + "/pause",
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 204,
								},
							},
						},
					},
				},
			},
		},
	})

	if _, err := cl.PauseStream(ctx, ""); err == nil {
		t.Fatal("id is required")
	}
	if _, err := cl.PauseStream(ctx, streamID); err != nil {
		t.Fatal(err)
	}
}

func TestClient_ResumeStream(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	streamID := "5e151c31a1de18000d89a83f"

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method: "POST",
								Path:   "/api/streams/" + streamID + "/resume",
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 204,
								},
							},
						},
					},
				},
			},
		},
	})

	if _, err := cl.ResumeStream(ctx, ""); err == nil {
		t.Fatal("id is required")
	}
	if _, err := cl.ResumeStream(ctx, streamID); err != nil {
		t.Fatal(err)
	}
}
