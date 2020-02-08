package client_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog/v10/client"
	"github.com/suzuki-shunsuke/go-graylog/v10/testdata"
)

func TestClient_GetInputs(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/input/inputs.json")
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
								Method:       "GET",
								Path:         "/api/system/inputs",
								PartOfHeader: getTestHeader(),
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

	inputs, total, _, err := cl.GetInputs(ctx)
	require.Nil(t, err)
	require.Equal(t, testdata.Inputs().Inputs, inputs)
	require.Equal(t, testdata.Inputs().Total, total)
}

func TestClient_GetInput(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/input/input.json")
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
								Method:       "GET",
								Path:         "/api/system/inputs/" + testdata.Input().ID,
								PartOfHeader: getTestHeader(),
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

	input, _, err := cl.GetInput(ctx, testdata.Input().ID)
	require.Nil(t, err)
	require.Equal(t, testdata.Input(), input)
}

func TestClient_CreateInput(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/input/create_input.json")
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
								Method:         "POST",
								Path:           "/api/system/inputs",
								PartOfHeader:   getTestHeader(),
								BodyJSONString: bodyStr,
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 201,
								},
								BodyString: `{
  "id": "5e1460d1a1de18000d890cd5"
}`,
							},
						},
					},
				},
			},
		},
	})

	// nil check
	if _, err := cl.CreateInput(ctx, nil); err == nil {
		t.Fatal("input is nil")
	}
	input := testdata.CreateInput()
	if _, err := cl.CreateInput(ctx, &input); err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdateInput(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/input/create_input.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	inputID := "5e1460d1a1de18000d890cd5"

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method:         "PUT",
								Path:           "/api/system/inputs/" + inputID,
								PartOfHeader:   getTestHeader(),
								BodyJSONString: bodyStr,
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 201,
								},
								BodyString: `{
  "id": "5e1460d1a1de18000d890cd5"
}`,
							},
						},
					},
				},
			},
		},
	})

	input := testdata.UpdateInput()
	if _, _, err := cl.UpdateInput(ctx, &input); err == nil {
		t.Fatal("id is required")
	}
	input = testdata.UpdateInput()
	input.ID = inputID
	if _, _, err := cl.UpdateInput(ctx, &input); err != nil {
		t.Fatal(err)
	}
	if _, _, err := cl.UpdateInput(ctx, nil); err == nil {
		t.Fatal("input is required")
	}
}

func TestClient_DeleteInput(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	inputID := "5e1460d1a1de18000d890cd5"

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Matcher: &flute.Matcher{
								Path: "/api/system/inputs/" + inputID,
							},
							Tester: &flute.Tester{
								Method:       "DELETE",
								PartOfHeader: getTestHeader(),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 204,
								},
							},
						},
						{
							Matcher: &flute.Matcher{
								Path: "/api/system/inputs/h",
							},
							Tester: &flute.Tester{
								Method:       "DELETE",
								PartOfHeader: getTestHeader(),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 404,
								},
								BodyString: `{
  "type": "ApiError",
  "message": "Input id <h> is invalid!"
}`,
							},
						},
					},
				},
			},
		},
	})

	if _, err := cl.DeleteInput(ctx, ""); err == nil {
		t.Fatal("input id is required")
	}
	if _, err := cl.DeleteInput(ctx, "h"); err == nil {
		t.Fatal(`no input with id "h" is found`)
	}
	if _, err := cl.DeleteInput(ctx, inputID); err != nil {
		t.Fatal(err)
	}
}
