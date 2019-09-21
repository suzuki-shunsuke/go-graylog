package client_test

import (
	"context"
	"fmt"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
)

func TestClient_GetStreamAlertConditions(t *testing.T) {
	ctx := context.Background()
	defer gock.Off()
	client, err := client.NewClient("http://example.com/api", "admin", "password")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		resp       string
		total      int
		conds      []graylog.AlertCondition
		checkErr   func(require.TestingT, interface{}, ...interface{})
	}{{
		statusCode: 200,
		resp: `{
  "total": 2,
  "conditions": [
    {
      "id": "d3ecb503-b767-4d59-bf6a-e2c000000000",
      "type": "field_content_value",
      "creator_user_id": "admin",
      "created_at": "2018-12-18T10:09:49.060+0000",
      "parameters": {
        "backlog": 1,
        "repeat_notifications": false,
        "field": "message",
        "query": "*",
        "grace": 0,
        "value": "transport: http2Server.HandleStreams failed to read frame: read tcp"
      },
      "in_grace": false,
      "title": "hello"
    },
    {
      "id": "6d3aafd0-b277-4b55-bfd9-f4a000000000",
      "type": "message_count",
      "creator_user_id": "admin",
      "created_at": "2018-12-06T05:59:33.215+0000",
      "parameters": {
        "backlog": 0,
        "repeat_notifications": true,
        "query": "*",
        "grace": 5,
        "threshold_type": "MORE",
        "threshold": 400,
        "time": 5
      },
      "in_grace": false,
      "title": "hello: too many log"
    }
  ]
}`,
		total: 2,
		conds: []graylog.AlertCondition{
			{
				ID:            "d3ecb503-b767-4d59-bf6a-e2c000000000",
				CreatorUserID: "admin",
				CreatedAt:     "2018-12-18T10:09:49.060+0000",
				Parameters: graylog.FieldContentAlertConditionParameters{
					Backlog:             1,
					RepeatNotifications: false,
					Field:               "message",
					Query:               "*",
					Grace:               0,
					Value:               "transport: http2Server.HandleStreams failed to read frame: read tcp",
				},
				Title: "hello",
			}, {
				ID:            "6d3aafd0-b277-4b55-bfd9-f4a000000000",
				CreatorUserID: "admin",
				CreatedAt:     "2018-12-06T05:59:33.215+0000",
				Parameters: graylog.MessageCountAlertConditionParameters{
					Backlog:             0,
					RepeatNotifications: true,
					Query:               "*",
					Grace:               5,
					ThresholdType:       "MORE",
					Threshold:           400,
					Time:                5,
				},
				Title: "hello: too many log",
			},
		},
		checkErr: require.Nil,
	}}
	for _, d := range data {
		gock.New("http://example.com").
			Get(fmt.Sprintf("/api/streams/%s/alerts/conditions", "xxxxx")).
			MatchType("json").Reply(d.statusCode).
			BodyString(d.resp)
		conds, total, _, err := client.GetStreamAlertConditions(ctx, "xxxxx")
		d.checkErr(t, err)
		if err != nil {
			require.Equal(t, d.conds, conds)
			require.Equal(t, d.total, total)
		}
	}
}

func TestClient_GetStreamAlertCondition(t *testing.T) {
	ctx := context.Background()
	defer gock.Off()
	client, err := client.NewClient("http://example.com/api", "admin", "password")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		resp       string
		cond       graylog.AlertCondition
		checkErr   func(require.TestingT, interface{}, ...interface{})
	}{{
		statusCode: 200,
		resp: `{
      "id": "d3ecb503-b767-4d59-bf6a-e2c000000000",
      "type": "field_content_value",
      "creator_user_id": "admin",
      "created_at": "2018-12-18T10:09:49.060+0000",
      "parameters": {
        "backlog": 1,
        "repeat_notifications": false,
        "field": "message",
        "query": "*",
        "grace": 0,
        "value": "transport: http2Server.HandleStreams failed to read frame: read tcp"
      },
      "in_grace": false,
      "title": "hello"
    }`,
		cond: graylog.AlertCondition{
			ID:            "d3ecb503-b767-4d59-bf6a-e2c000000000",
			CreatorUserID: "admin",
			CreatedAt:     "2018-12-18T10:09:49.060+0000",
			Parameters: graylog.FieldContentAlertConditionParameters{
				Backlog:             1,
				RepeatNotifications: false,
				Field:               "message",
				Query:               "*",
				Grace:               0,
				Value:               "transport: http2Server.HandleStreams failed to read frame: read tcp",
			},
			Title: "hello",
		},
		checkErr: require.Nil,
	}}
	for _, d := range data {
		gock.New("http://example.com").
			Get(fmt.Sprintf("/api/streams/%s/alerts/conditions/%s", "xxxxx", d.cond.ID)).
			MatchType("json").Reply(d.statusCode).
			BodyString(d.resp)
		cond, _, err := client.GetStreamAlertCondition(ctx, "xxxxx", d.cond.ID)
		d.checkErr(t, err)
		if err != nil {
			require.Equal(t, d.cond, cond)
		}
	}
}

func TestClient_CreateStreamAlertCondition(t *testing.T) {
	ctx := context.Background()
	defer gock.Off()
	client, err := client.NewClient("http://example.com/api", "admin", "password")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		condID     string
		cond       graylog.AlertCondition
		checkErr   func(require.TestingT, interface{}, ...interface{})
	}{{
		statusCode: 201,
		condID:     "d3ecb503-b767-4d59-bf6a-e2c000000000",
		cond: graylog.AlertCondition{
			Parameters: graylog.FieldContentAlertConditionParameters{
				Backlog:             1,
				RepeatNotifications: false,
				Field:               "message",
				Query:               "*",
				Grace:               0,
				Value:               "transport: http2Server.HandleStreams failed to read frame: read tcp",
			},
			Title: "hello",
		},
		checkErr: require.Nil,
	}}
	streamID := "xxxxx"
	for _, d := range data {
		gock.New("http://example.com").
			Post(fmt.Sprintf("/api/streams/%s/alerts/conditions", streamID)).
			MatchType("json").Reply(d.statusCode).
			BodyString(fmt.Sprintf(`{
  "alert_condition_id": "%s"
}`, d.condID))
		_, err := client.CreateStreamAlertCondition(ctx, streamID, &d.cond)
		d.checkErr(t, err)
		if err != nil {
			require.Equal(t, d.cond.ID, d.condID)
		}
	}
}

func TestClient_UpdateStreamAlertCondition(t *testing.T) {
	ctx := context.Background()
	defer gock.Off()
	client, err := client.NewClient("http://example.com/api", "admin", "password")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		cond       graylog.AlertCondition
		checkErr   func(require.TestingT, interface{}, ...interface{})
	}{{
		statusCode: 204,
		cond: graylog.AlertCondition{
			ID: "d3ecb503-b767-4d59-bf6a-e2c000000000",
			Parameters: graylog.FieldContentAlertConditionParameters{
				Backlog:             1,
				RepeatNotifications: false,
				Field:               "message",
				Query:               "*",
				Grace:               0,
				Value:               "transport: http2Server.HandleStreams failed to read frame: read tcp",
			},
			Title: "hello",
		},
		checkErr: require.Nil,
	}}
	streamID := "xxxxx"
	for _, d := range data {
		gock.New("http://example.com").
			Put(fmt.Sprintf("/api/streams/%s/alerts/conditions/%s", streamID, d.cond.ID)).
			MatchType("json").Reply(d.statusCode)
		_, err := client.UpdateStreamAlertCondition(ctx, streamID, &d.cond)
		d.checkErr(t, err)
	}
}

func TestClient_DeleteStreamAlertCondition(t *testing.T) {
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
	condID := "d3ecb503-b767-4d59-bf6a-e2c000000000"
	for _, d := range data {
		gock.New("http://example.com").
			Delete(fmt.Sprintf("/api/streams/%s/alerts/conditions/%s", streamID, condID)).
			MatchType("json").Reply(d.statusCode)
		_, err := client.DeleteStreamAlertCondition(ctx, streamID, condID)
		d.checkErr(t, err)
	}
}
