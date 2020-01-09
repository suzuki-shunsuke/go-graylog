package client_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog/v9/client"
	"github.com/suzuki-shunsuke/go-graylog/v9/testdata/event_definition/create"
	"github.com/suzuki-shunsuke/go-graylog/v9/testdata/event_definition/get"
	"github.com/suzuki-shunsuke/go-graylog/v9/testdata/event_definition/update"
)

func TestClient_CreateEventDefinition(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	req, err := ioutil.ReadFile("../testdata/event_definition/create/request.json")
	require.Nil(t, err)
	reqStr := string(req)

	resp, err := ioutil.ReadFile("../testdata/event_definition/create/response.json")
	require.Nil(t, err)
	respStr := string(resp)

	_, err = cl.CreateEventDefinition(ctx, nil)
	require.NotNil(t, err, "definition should not be nil")

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
								Path:   "/api/events/definitions",
							},
							Tester: &flute.Tester{
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
								BodyJSONString: reqStr,
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 201,
									Header: http.Header{
										"Content-Type": []string{"application/json"},
									},
								},
								BodyString: respStr,
							},
						},
					},
				},
			},
		},
	})
	ed := create.Request()
	_, err = cl.CreateEventDefinition(ctx, ed)
	require.Nil(t, err)
	require.Equal(t, create.Response().ID, ed.ID)
}

func TestClient_DeleteEventDefinition(t *testing.T) {
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
								Path:   "/api/events/definitions/" + get.Response().ID,
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
	if _, err := cl.DeleteEventDefinition(ctx, ""); err == nil {
		t.Fatal("id is required")
	}
	// invalid id
	_, err = cl.DeleteEventDefinition(ctx, get.Response().ID)
	require.Nil(t, err)
}

func TestClient_GetEventDefinition(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/event_definition/get/response.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	_, _, err = cl.GetEventDefinition(ctx, "")
	require.NotNil(t, err, "definition id should not be empty")

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
								Path:   "/api/events/definitions/" + get.Response().ID,
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
	definition, _, err := cl.GetEventDefinition(ctx, get.Response().ID)
	require.Nil(t, err)
	require.Equal(t, get.Response().ID, definition.ID)
}

func TestClient_GetEventDefinitions(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/event_definition/gets/response.json")
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
								Path:   "/api/events/definitions",
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

	_, _, err = cl.GetEventDefinitions(ctx)
	require.Nil(t, err)
}

func TestClient_UpdateEventDefinition(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	req, err := ioutil.ReadFile("../testdata/event_definition/update/request.json")
	require.Nil(t, err)
	reqStr := string(req)

	resp, err := ioutil.ReadFile("../testdata/event_definition/update/response.json")
	require.Nil(t, err)
	respStr := string(resp)

	_, err = cl.UpdateEventDefinition(ctx, nil)
	require.NotNil(t, err, "definition should not be nil")

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
								Path:   "/api/events/definitions/" + update.Request().ID,
							},
							Tester: &flute.Tester{
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
								BodyJSONString: reqStr,
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: respStr,
							},
						},
					},
				},
			},
		},
	})
	_, err = cl.UpdateEventDefinition(ctx, update.Request())
	require.Nil(t, err)
}
