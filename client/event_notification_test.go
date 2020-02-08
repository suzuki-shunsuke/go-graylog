package client_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog/v10"
	"github.com/suzuki-shunsuke/go-graylog/v10/client"
	"github.com/suzuki-shunsuke/go-graylog/v10/testdata"
)

func TestClient_CreateEventNotification(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	reqBuf, err := ioutil.ReadFile("../testdata/event_notification/request_create_event_notification.json")
	require.Nil(t, err)
	reqBody := string(reqBuf)

	respBuf, err := ioutil.ReadFile("../testdata/event_notification/response_create_event_notification.json")
	require.Nil(t, err)
	respBody := string(respBuf)

	_, err = cl.CreateEventNotification(ctx, nil)
	require.NotNil(t, err, "event notification should not be nil")

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
								Path:   "/api/events/notifications",
							},
							Tester: &flute.Tester{
								PartOfHeader:   getTestHeader(),
								BodyJSONString: reqBody,
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 201,
									Header: http.Header{
										"Content-Type": []string{"application/json"},
									},
								},
								BodyString: respBody,
							},
						},
					},
				},
			},
		},
	})
	notification := &graylog.EventNotification{
		Title:       "http",
		Description: "",
		Config: map[string]interface{}{
			"type": "http-notification-v1",
			"url":  "http://example.com",
		},
	}
	_, err = cl.CreateEventNotification(ctx, notification)
	require.Nil(t, err)
	require.Equal(t, testdata.EventNotification().ID, notification.ID)
}

func TestClient_DeleteEventNotification(t *testing.T) {
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
								Path:   "/api/events/notifications/" + testdata.EventNotification().ID,
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
	if _, err := cl.DeleteEventNotification(ctx, ""); err == nil {
		t.Fatal("id is required")
	}
	// invalid id
	_, err = cl.DeleteEventNotification(ctx, testdata.EventNotification().ID)
	require.Nil(t, err)
}

func TestClient_GetEventNotification(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/event_notification/response_create_event_notification.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	_, _, err = cl.GetEventNotification(ctx, "")
	require.NotNil(t, err, "event notification id should not be empty")

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
								Path:   "/api/events/notifications/" + testdata.EventNotification().ID,
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
	notification, _, err := cl.GetEventNotification(ctx, testdata.EventNotification().ID)
	require.Nil(t, err)
	require.Equal(t, testdata.EventNotification(), notification)
}

func TestClient_GetEventNotifications(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/event_notification/event_notifications.json")
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
								Path:   "/api/events/notifications",
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

	_, _, err = cl.GetEventNotifications(ctx)
	require.Nil(t, err)
}

func TestClient_UpdateEventNotification(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/event_notification/response_create_event_notification.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	_, err = cl.UpdateEventNotification(ctx, nil)
	require.NotNil(t, err, "event notification should not be nil")

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
								Path:   "/api/events/notifications/" + testdata.EventNotification().ID,
							},
							Tester: &flute.Tester{
								PartOfHeader:   getTestHeader(),
								BodyJSONString: ``,
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
	_, err = cl.UpdateEventNotification(ctx, testdata.EventNotification())
	require.Nil(t, err)
}
