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

const (
	streamID = "5de4fcf7a1de1800127e2fbe"
)

func TestClient_CreateStreamOutputs(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/request_create_stream_outputs.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	_, err = cl.CreateStreamOutputs(ctx, "", nil)
	require.NotNil(t, err, "stream id is required")

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Matcher: &flute.Matcher{
								Method: "POST",
								Path:   "/api/streams/" + streamID + "/outputs",
							},
							Tester: &flute.Tester{
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
								BodyJSONString: bodyStr,
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 201,
									Header: http.Header{
										"Content-Type": []string{"application/json"},
									},
								},
								BodyString: bodyStr,
							},
						},
					},
				},
			},
		},
	})
	_, err = cl.CreateStreamOutputs(ctx, streamID, []string{
		"5de4fcf8a1de1800127e2fc4",
		"5de4fcf7a1de1800127e2fc0",
	})
	require.Nil(t, err)
}

func TestClient_DeleteStreamOutput(t *testing.T) {
	ctx := context.Background()

	outputID := "5de4fcf8a1de1800127e2fc4"

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	_, err = cl.DeleteStreamOutput(ctx, "", outputID)
	require.NotNil(t, err, "stream id is required")
	_, err = cl.DeleteStreamOutput(ctx, streamID, "")
	require.NotNil(t, err, "output id is required")

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Matcher: &flute.Matcher{
								Method: "DELETE",
								Path:   "/api/streams/" + streamID + "/outputs/" + outputID,
							},
							Tester: &flute.Tester{
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 204,
									Header: http.Header{
										"Content-Type": []string{"application/json"},
									},
								},
							},
						},
					},
				},
			},
		},
	})

	// id required
	_, err = cl.DeleteStreamOutput(ctx, streamID, outputID)
	require.Nil(t, err)
}

func TestClient_GetStreamOutputs(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/outputs.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	_, _, _, err = cl.GetStreamOutputs(ctx, "")
	require.NotNil(t, err, "stream id should not be empty")

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Matcher: &flute.Matcher{
								Method: "GET",
								Path:   "/api/streams/" + streamID + "/outputs",
							},
							Tester: &flute.Tester{
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
									Header: http.Header{
										"Content-Type": []string{"application/json"},
									},
								},
								BodyString: bodyStr,
							},
						},
					},
				},
			},
		},
	})
	_, total, _, err := cl.GetStreamOutputs(ctx, streamID)
	require.Nil(t, err)
	require.Equal(t, testdata.Outputs().Total, total)
}
