package client_test

import (
	"context"
	"fmt"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
)

func TestGetStreamAlarmCallbacks(t *testing.T) {
	ctx := context.Background()
	defer gock.Off()
	client, err := client.NewClient("http://example.com/api", "admin", "password")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		resp       string
		total      int
		acs        []graylog.AlarmCallback
		checkErr   func(require.TestingT, interface{}, ...interface{})
	}{{
		statusCode: 200,
		resp: `{
  "total": 4,
  "alarmcallbacks": [
    {
      "id": "5c08bb0dc9e77c0000000000",
      "type": "org.graylog2.plugins.slack.callback.SlackAlarmCallback",
      "configuration": {
        "icon_url": "",
        "graylog2_url": "https://graylog.example.com",
        "link_names": true,
        "color": "#FF0000",
        "webhook_url": "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX",
        "icon_emoji": "",
        "user_name": "Graylog",
        "backlog_items": 5,
        "proxy_address": "",
        "channel": "#general",
        "custom_message": "${alert_acition.title}\n\n${foreach backlog message}\n<https://graylog.example.com/streams/${stream.id}/search?rangetype=absolute&from=${message.timestamp}&to=${message.timestamp} | link> ${message.message}\n${end}",
        "notify_channel": false
      },
      "stream_id": "5b93b425c9e0000000000000",
      "title": "slack alarm callback",
      "created_at": "2018-12-06T06:00:45.717+0000",
      "creator_user_id": "admin"
    },
    {
      "id": "5c28857bc9e77c0000000000",
      "type": "org.graylog2.alarmcallbacks.EmailAlarmCallback",
      "configuration": {
        "user_receivers": [
          "example"
        ],
        "body": "##########\nAlert Description: ${check_result.resultDescription}\nDate: ${check_result.triggeredAt}\nStream ID: ${stream.id}\nStream title: ${stream.title}\nStream description: ${stream.description}\nAlert Condition Title: ${alertCondition.title}\n${if stream_url}Stream URL: ${stream_url}${end}\n\nTriggered acition: ${check_result.triggeredCondition}\n##########\n\n${if backlog}Last messages accounting for this alert:\n${foreach backlog message}${message}\n\n${end}${else}<No backlog>\n${end}\n",
        "sender": "graylog@example.org",
        "subject": "Graylog alert for stream: ${stream.title}: ${check_result.resultDescription}",
        "email_receivers": [
          "graylog@example.com"
        ]
      },
      "stream_id": "5b93b425c9e0000000000000",
      "title": "email alarm callback",
      "created_at": "2018-12-30T08:44:43.088+0000",
      "creator_user_id": "admin"
    },
    {
      "id": "5c288624c9e77c0000000000",
      "type": "org.graylog2.alarmcallbacks.HTTPAlarmCallback",
      "configuration": {
        "url": "https://example.com"
      },
      "stream_id": "5b93b425c9e0000000000000",
      "title": "http alarm callback",
      "created_at": "2018-12-30T08:47:32.865+0000",
      "creator_user_id": "admin"
    },
		{
	    "id": "5c29bb09df46c60001ab3af3",
      "type": "org.graylog2.alarmcallbacks.hipchat.HipChatAlarmCallback",
      "configuration": {
        "color": "green",
        "api_url": "https://api.hipchat.com",
        "message_template": "test template",
        "api_token": "test",
        "graylog_base_url": "http://localhost:9000",
        "notify": true,
        "room": "test"
      },
      "stream_id": "000000000000000000000001",
      "title": "test",
      "created_at": "2018-12-31T06:45:29.907+0000",
      "creator_user_id": "admin"	
		}
  ]
}`,
		total: 4,
		acs: []graylog.AlarmCallback{
			{
				ID:            "5c08bb0dc9e77c0000000000",
				StreamID:      "5b93b425c9e0000000000000",
				Title:         "slack alarm callback",
				CreatorUserID: "admin",
				CreatedAt:     "2018-12-06T06:00:45.717+0000",
				Configuration: &graylog.SlackAlarmCallbackConfiguration{
					IconURL:       "",
					Graylog2URL:   "https://graylog.example.com",
					LinkNames:     true,
					Color:         "#FF0000",
					WebhookURL:    "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX",
					IconEmoji:     "",
					UserName:      "Graylog",
					BacklogItems:  5,
					ProxyAddress:  "",
					Channel:       "#general",
					CustomMessage: "${alert_acition.title}\n\n${foreach backlog message}\n<https://graylog.example.com/streams/${stream.id}/search?rangetype=absolute&from=${message.timestamp}&to=${message.timestamp} | link> ${message.message}\n${end}",
					NotifyChannel: false,
				},
			}, {
				ID: "5c28857bc9e77c0000000000",
				Configuration: &graylog.EmailAlarmCallbackConfiguration{
					UserReceivers:  set.NewStrSet("example"),
					Body:           "##########\nAlert Description: ${check_result.resultDescription}\nDate: ${check_result.triggeredAt}\nStream ID: ${stream.id}\nStream title: ${stream.title}\nStream description: ${stream.description}\nAlert Condition Title: ${alertCondition.title}\n${if stream_url}Stream URL: ${stream_url}${end}\n\nTriggered acition: ${check_result.triggeredCondition}\n##########\n\n${if backlog}Last messages accounting for this alert:\n${foreach backlog message}${message}\n\n${end}${else}<No backlog>\n${end}\n",
					Sender:         "graylog@example.org",
					Subject:        "Graylog alert for stream: ${stream.title}: ${check_result.resultDescription}",
					EmailReceivers: set.NewStrSet("graylog@example.com"),
				},
				StreamID:      "5b93b425c9e0000000000000",
				Title:         "email alarm callback",
				CreatedAt:     "2018-12-30T08:44:43.088+0000",
				CreatorUserID: "admin",
			}, {
				ID: "5c288624c9e77c0000000000",
				Configuration: &graylog.HTTPAlarmCallbackConfiguration{
					URL: "https://example.com",
				},
				StreamID:      "5b93b425c9e0000000000000",
				Title:         "http alarm callback",
				CreatedAt:     "2018-12-30T08:47:32.865+0000",
				CreatorUserID: "admin",
			}, {
				ID: "5c29bb09df46c60001ab3af3",
				Configuration: &graylog.GeneralAlarmCallbackConfiguration{
					Type: "org.graylog2.alarmcallbacks.hipchat.HipChatAlarmCallback",
					Configuration: map[string]interface{}{
						"color":            "green",
						"api_url":          "https://api.hipchat.com",
						"message_template": "test template",
						"api_token":        "test",
						"graylog_base_url": "http://localhost:9000",
						"notify":           true,
						"room":             "test",
					},
				},
				StreamID:      "000000000000000000000001",
				Title:         "test",
				CreatedAt:     "2018-12-31T06:45:29.907+0000",
				CreatorUserID: "admin",
			},
		},
		checkErr: require.Nil,
	}}
	for _, d := range data {
		gock.New("http://example.com").
			Get(fmt.Sprintf("/api/streams/%s/alarmcallbacks", "xxxxx")).
			MatchType("json").Reply(d.statusCode).
			BodyString(d.resp)
		acs, total, _, err := client.GetStreamAlarmCallbacks(ctx, "xxxxx")
		d.checkErr(t, err)
		if err != nil {
			require.Equal(t, d.acs, acs)
			require.Equal(t, d.total, total)
		}
	}
}

func TestGetStreamAlarmCallback(t *testing.T) {
	ctx := context.Background()
	defer gock.Off()
	client, err := client.NewClient("http://example.com/api", "admin", "password")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		resp       string
		ac         graylog.AlarmCallback
		checkErr   func(require.TestingT, interface{}, ...interface{})
	}{{
		statusCode: 200,
		resp: `{
      "id": "5c288624c9e77c0000000000",
      "type": "org.graylog2.alarmcallbacks.HTTPAlarmCallback",
      "configuration": {
        "url": "https://example.com"
      },
      "stream_id": "5b93b425c9e0000000000000",
      "title": "http alarm callback",
      "created_at": "2018-12-30T08:47:32.865+0000",
      "creator_user_id": "admin"
    }`,
		ac: graylog.AlarmCallback{
			ID: "5c288624c9e77c0000000000",
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
			Get(fmt.Sprintf("/api/streams/%s/alarmcallbacks/%s", "xxxxx", d.ac.ID)).
			MatchType("json").Reply(d.statusCode).
			BodyString(d.resp)
		ac, _, err := client.GetStreamAlarmCallback(ctx, "xxxxx", d.ac.ID)
		d.checkErr(t, err)
		if err != nil {
			require.Equal(t, d.ac, ac)
		}
	}
}

func TestCreateStreamAlarmCallback(t *testing.T) {
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

func TestUpdateStreamAlarmCallback(t *testing.T) {
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

func TestDeleteStreamAlarmCallback(t *testing.T) {
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
