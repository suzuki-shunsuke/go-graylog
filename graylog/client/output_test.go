package client_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/client"
	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/testdata"
)

func TestClient_CreateOutput(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/output/stdout_output.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	_, err = cl.CreateOutput(ctx, nil)
	require.NotNil(t, err, "output should not be nil")

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
								Path:   "/api/system/outputs",
							},
							Tester: &flute.Tester{
								PartOfHeader: getTestHeader(),
								BodyJSONString: `{
								  "title": "test-stdout",
								  "type": "org.graylog2.outputs.LoggingOutput",
								  "configuration": {
								    "prefix": "Writing message: "
								  } 	
								}`,
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
	output := &graylog.Output{
		Title: "test-stdout",
		Type:  "org.graylog2.outputs.LoggingOutput",
		Configuration: map[string]interface{}{
			"prefix": "Writing message: ",
		},
	}
	_, err = cl.CreateOutput(ctx, output)
	require.Nil(t, err)
	require.Equal(t, testdata.StdoutOutput().ID, output.ID)
}

func TestClient_DeleteOutput(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

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
								Path:   "/api/system/outputs/" + testdata.StdoutOutput().ID,
							},
							Tester: &flute.Tester{
								PartOfHeader: getTestHeader(),
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
	if _, err := cl.DeleteOutput(ctx, ""); err == nil {
		t.Fatal("id is required")
	}
	// invalid id
	_, err = cl.DeleteOutput(ctx, testdata.StdoutOutput().ID)
	require.Nil(t, err)
}

func TestClient_GetOutput(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/output/stdout_output.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	_, _, err = cl.GetOutput(ctx, "")
	require.NotNil(t, err, "output id should not be empty")

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
								Path:   "/api/system/outputs/" + testdata.StdoutOutput().ID,
							},
							Tester: &flute.Tester{
								PartOfHeader: getTestHeader(),
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
	output, _, err := cl.GetOutput(ctx, testdata.StdoutOutput().ID)
	require.Nil(t, err)
	require.Equal(t, testdata.StdoutOutput(), output)
}

func TestClient_GetOutputs(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/output/outputs.json")
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
							Matcher: &flute.Matcher{
								Method: "GET",
								Path:   "/api/system/outputs",
							},
							Tester: &flute.Tester{
								PartOfHeader: getTestHeader(),
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

	_, total, _, err := cl.GetOutputs(ctx)
	require.Nil(t, err)
	require.Equal(t, testdata.Outputs().Total, total)
}

func TestClient_UpdateOutput(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/output/stdout_output.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	_, err = cl.UpdateOutput(ctx, nil)
	require.NotNil(t, err, "output should not be nil")

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Matcher: &flute.Matcher{
								Method: "PUT",
								Path:   "/api/system/outputs/" + testdata.StdoutOutput().ID,
							},
							Tester: &flute.Tester{
								PartOfHeader: getTestHeader(),
								BodyJSONString: `{
  "title": "test",
  "type": "org.graylog2.outputs.LoggingOutput",
  "configuration": {
    "prefix": "Writing message: "
  }
}`,
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
	_, err = cl.UpdateOutput(ctx, testdata.StdoutOutput())
	require.Nil(t, err)
}
