package client_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog/v8"
	"github.com/suzuki-shunsuke/go-graylog/client/v8"
	"github.com/suzuki-shunsuke/go-graylog/testdata"
	"github.com/suzuki-shunsuke/go-graylog/testutil/v8"
)

func TestClient_GetInputs(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/inputs.json")
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
								Path:   "/api/system/inputs",
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

	inputs, total, _, err := cl.GetInputs(ctx)
	require.Nil(t, err)
	require.Equal(t, testdata.Inputs.Inputs, inputs)
	require.Equal(t, testdata.Inputs.Total, total)
}

func TestClient_GetInput(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/input.json")
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
								Path:   "/api/system/inputs/" + testdata.Input.ID,
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

	input, _, err := cl.GetInput(ctx, testdata.Input.ID)
	require.Nil(t, err)
	require.Equal(t, testdata.Input, input)
}

func TestClient_CreateInput(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	// nil check
	if _, err := client.CreateInput(ctx, nil); err == nil {
		t.Fatal("input is nil")
	}
	input := testutil.Input()
	if _, err := client.CreateInput(ctx, input); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteInput(ctx, input.ID)
	attrs := input.Attrs.(*graylog.InputBeatsAttrs)
	if attrs.BindAddress == "" {
		t.Fatal(`attrs.BindAddress == ""`)
	}
	// error check
	if _, err := client.CreateInput(ctx, input); err == nil {
		t.Fatal("input id should be empty")
	}
}

func TestClient_UpdateInput(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	input := testutil.Input()
	if _, err := client.CreateInput(ctx, input); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteInput(ctx, input.ID)
	attrs := input.Attrs.(*graylog.InputBeatsAttrs)
	if attrs.BindAddress == "" {
		t.Fatal(`attrs.BindAddress == ""`)
	}
	if _, _, err := client.UpdateInput(ctx, input.NewUpdateParams()); err != nil {
		t.Fatal(err)
	}
	input.ID = ""
	if _, _, err := client.UpdateInput(ctx, input.NewUpdateParams()); err == nil {
		t.Fatal("input id is required")
	}
	if _, _, err := client.UpdateInput(ctx, nil); err == nil {
		t.Fatal("input is required")
	}
	input.ID = "h"
	if _, _, err := client.UpdateInput(ctx, input.NewUpdateParams()); err == nil {
		t.Fatal("input should no be found")
	}
}

func TestClient_DeleteInput(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, err := client.DeleteInput(ctx, ""); err == nil {
		t.Fatal("input id is required")
	}
	if _, err := client.DeleteInput(ctx, "h"); err == nil {
		t.Fatal(`no input with id "h" is found`)
	}
}
