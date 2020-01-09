package client_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/suzuki-shunsuke/go-set/v6"

	"github.com/suzuki-shunsuke/go-graylog/v9"
	"github.com/suzuki-shunsuke/go-graylog/v9/client"
	"github.com/suzuki-shunsuke/go-graylog/v9/testdata"
)

func TestClient_GetStreamAlarmCallbacks(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/stream_alarm_callbacks.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	id := "5d84c1a92ab79c000d35d6ca"

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
								Path:   "/api/streams/" + id + "/alarmcallbacks",
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

	_, _, _, err = cl.GetStreamAlarmCallbacks(ctx, "")
	require.NotNil(t, err)
	acs, total, _, err := cl.GetStreamAlarmCallbacks(ctx, id)
	require.Nil(t, err)
	require.Equal(t, testdata.StreamAlarmCallbacks().Total, total)
	require.Equal(t, testdata.StreamAlarmCallbacks().AlarmCallbacks, acs)
}

func TestClient_GetStreamAlarmCallback(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/slack_stream_alarm_callback.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	id := "5d84c1a92ab79c000d35d6ca"
	callbackID := "5d84c1a92ab79c000d35d6d5"

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
								Path:   "/api/streams/" + id + "/alarmcallbacks/" + callbackID,
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

	_, _, err = cl.GetStreamAlarmCallback(ctx, "", callbackID)
	require.NotNil(t, err)

	_, _, err = cl.GetStreamAlarmCallback(ctx, id, "")
	require.NotNil(t, err)

	ac, _, err := cl.GetStreamAlarmCallback(ctx, id, callbackID)
	require.Nil(t, err)
	require.Equal(t, testdata.SlackStreamAlarmCallback(), ac)
}

func TestClient_CreateStreamAlarmCallback(t *testing.T) {
	ctx := context.Background()
	defer gock.Off()
	client, err := client.NewClient("http://example.com/api", "admin", "password")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		acID       string
		req        interface{}
		ac         graylog.AlarmCallback
		checkErr   func(require.TestingT, interface{}, ...interface{})
	}{{
		statusCode: 201,
		acID:       "d3ecb503-b767-4d59-bf6a-e2c000000000",
		req: map[string]interface{}{
			"type":  "org.graylog2.alarmcallbacks.hipchat.HipChatAlarmCallback",
			"title": "hipchat alarm callback",
			"configuration": map[string]interface{}{
				"color":            "yellow",
				"api_url":          "https://api.hipchat.com",
				"message_template": "test template",
				"api_token":        "test",
				"graylog_base_url": "http://localhost:9000",
				"room":             "test",
				"notify":           true,
			},
		},
		ac: graylog.AlarmCallback{
			Configuration: &graylog.GeneralAlarmCallbackConfiguration{
				Type: "org.graylog2.alarmcallbacks.hipchat.HipChatAlarmCallback",
				Configuration: map[string]interface{}{
					"color":            "yellow",
					"api_url":          "https://api.hipchat.com",
					"message_template": "test template",
					"api_token":        "test",
					"graylog_base_url": "http://localhost:9000",
					"room":             "test",
					"notify":           true,
				},
			},
			StreamID: "5b93b425c9e0000000000000",
			Title:    "hipchat alarm callback",
		},
		checkErr: require.Nil,
	}, {
		statusCode: 201,
		acID:       "d3ecb503-b767-4d59-bf6a-e2c000000000",
		req: map[string]interface{}{
			"type":  graylog.SlackAlarmCallbackType,
			"title": "slack alarm callback",
			"configuration": map[string]interface{}{
				"graylog2_url":   "https://graylog.example.com",
				"link_names":     true,
				"color":          "#FF0000",
				"webhook_url":    "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX",
				"user_name":      "Graylog",
				"backlog_items":  5,
				"channel":        "#general",
				"custom_message": "${alert_acition.title}\n\n${foreach backlog message}\n<https://graylog.example.com/streams/${stream.id}/search?rangetype=absolute&from=${message.timestamp}&to=${message.timestamp} | link> ${message.message}\n${end}",
				"notify_channel": false,
			},
		},
		ac: graylog.AlarmCallback{
			StreamID: "5b93b425c9e0000000000000",
			Title:    "slack alarm callback",
			Configuration: &graylog.SlackAlarmCallbackConfiguration{
				Graylog2URL:   "https://graylog.example.com",
				LinkNames:     true,
				Color:         "#FF0000",
				WebhookURL:    "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX",
				UserName:      "Graylog",
				BacklogItems:  5,
				Channel:       "#general",
				CustomMessage: "${alert_acition.title}\n\n${foreach backlog message}\n<https://graylog.example.com/streams/${stream.id}/search?rangetype=absolute&from=${message.timestamp}&to=${message.timestamp} | link> ${message.message}\n${end}",
				NotifyChannel: false,
			},
		},
		checkErr: require.Nil,
	}, {
		statusCode: 201,
		acID:       "d3ecb503-b767-4d59-bf6a-e2c000000000",
		req: map[string]interface{}{
			"type":  graylog.EmailAlarmCallbackType,
			"title": "email alarm callback",
			"configuration": map[string]interface{}{
				"user_receivers":  []string{"example"},
				"body":            "##########\nAlert Description: ${check_result.resultDescription}\nDate: ${check_result.triggeredAt}\nStream ID: ${stream.id}\nStream title: ${stream.title}\nStream description: ${stream.description}\nAlert Condition Title: ${alertCondition.title}\n${if stream_url}Stream URL: ${stream_url}${end}\n\nTriggered acition: ${check_result.triggeredCondition}\n##########\n\n${if backlog}Last messages accounting for this alert:\n${foreach backlog message}${message}\n\n${end}${else}<No backlog>\n${end}\n",
				"sender":          "graylog@example.org",
				"subject":         "Graylog alert for stream: ${stream.title}: ${check_result.resultDescription}",
				"email_receivers": []string{"graylog@example.com"},
			},
		},
		ac: graylog.AlarmCallback{
			StreamID: "5b93b425c9e0000000000000",
			Title:    "email alarm callback",
			Configuration: &graylog.EmailAlarmCallbackConfiguration{
				UserReceivers:  set.NewStrSet("example"),
				Body:           "##########\nAlert Description: ${check_result.resultDescription}\nDate: ${check_result.triggeredAt}\nStream ID: ${stream.id}\nStream title: ${stream.title}\nStream description: ${stream.description}\nAlert Condition Title: ${alertCondition.title}\n${if stream_url}Stream URL: ${stream_url}${end}\n\nTriggered acition: ${check_result.triggeredCondition}\n##########\n\n${if backlog}Last messages accounting for this alert:\n${foreach backlog message}${message}\n\n${end}${else}<No backlog>\n${end}\n",
				Sender:         "graylog@example.org",
				Subject:        "Graylog alert for stream: ${stream.title}: ${check_result.resultDescription}",
				EmailReceivers: set.NewStrSet("graylog@example.com"),
			},
		},
		checkErr: require.Nil,
	}, {
		statusCode: 201,
		acID:       "d3ecb503-b767-4d59-bf6a-e2c000000000",
		req: map[string]interface{}{
			"type":  graylog.HTTPAlarmCallbackType,
			"title": "http alarm callback",
			"configuration": map[string]interface{}{
				"url": "https://example.com",
			},
		},
		ac: graylog.AlarmCallback{
			StreamID: "5b93b425c9e0000000000000",
			Title:    "http alarm callback",
			Configuration: &graylog.HTTPAlarmCallbackConfiguration{
				URL: "https://example.com",
			},
		},
		checkErr: require.Nil,
	}}
	for _, d := range data {
		gock.New("http://example.com").
			Post(fmt.Sprintf("/api/streams/%s/alarmcallbacks", d.ac.StreamID)).
			MatchType("json").JSON(d.req).Reply(d.statusCode).
			JSON(map[string]string{"alarmcallback_id": d.acID})
		_, err := client.CreateStreamAlarmCallback(ctx, &d.ac)
		d.checkErr(t, err)
		if err != nil {
			require.Equal(t, d.ac.ID, d.acID)
		}
	}
}

func TestClient_UpdateStreamAlarmCallback(t *testing.T) {
	ctx := context.Background()
	defer gock.Off()
	client, err := client.NewClient("http://example.com/api", "admin", "password")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		ac         graylog.AlarmCallback
		checkErr   func(require.TestingT, interface{}, ...interface{})
	}{{
		statusCode: 204,
		ac: graylog.AlarmCallback{
			ID: "d3ecb503-b767-4d59-bf6a-e2c000000000",
			Configuration: &graylog.HTTPAlarmCallbackConfiguration{
				URL: "https://example.com",
			},
			StreamID:      "5b93b425c9e0000000000000",
			Title:         "http alarm callback",
			CreatedAt:     "2018-12-30T08:47:32.865+0000",
			CreatorUserID: "admin",
		},
		checkErr: require.Nil,
	}}
	for _, d := range data {
		gock.New("http://example.com").
			Put(fmt.Sprintf("/api/streams/%s/alarmcallbacks/%s", d.ac.StreamID, d.ac.ID)).
			MatchType("json").Reply(d.statusCode)
		_, err := client.UpdateStreamAlarmCallback(ctx, &d.ac)
		d.checkErr(t, err)
	}
}

func TestClient_DeleteStreamAlarmCallback(t *testing.T) {
	ctx := context.Background()
	defer gock.Off()
	client, err := client.NewClient("http://example.com/api", "admin", "password")
	require.Nil(t, err)
	data := []struct {
		statusCode int
		checkErr   func(require.TestingT, interface{}, ...interface{})
	}{{
		statusCode: 204,
		checkErr:   require.Nil,
	}}
	streamID := "xxxxx"
	acID := "d3ecb503-b767-4d59-bf6a-e2c000000000"
	for _, d := range data {
		gock.New("http://example.com").
			Delete(fmt.Sprintf("/api/streams/%s/alarmcallbacks/%s", streamID, acID)).
			MatchType("json").Reply(d.statusCode)
		_, err := client.DeleteStreamAlarmCallback(ctx, streamID, acID)
		d.checkErr(t, err)
	}
}
