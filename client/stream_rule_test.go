package client_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestClient_GetStreamRules(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	stream, f, err := testutil.GetStream(ctx, client, server, 2)
	if err != nil {
		t.Fatal(err)
	}
	if f != nil {
		defer f(stream.ID)
	}

	if _, _, _, err := client.GetStreamRules(ctx, stream.ID); err != nil {
		t.Fatal("Failed to GetStreamRules", err)
	}
	if _, _, _, err := client.GetStreamRules(ctx, "h"); err == nil {
		t.Fatal(`no stream with id "h" is found`)
	}
}

func TestClient_CreateStreamRule(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	stream, f, err := testutil.GetStream(ctx, client, server, 2)
	if err != nil {
		t.Fatal(err)
	}
	if f != nil {
		defer f(stream.ID)
	}
	rule := testutil.StreamRule()
	rule.StreamID = stream.ID
	if _, err := client.CreateStreamRule(ctx, rule); err != nil {
		t.Fatal(err)
	}
	if _, err := client.CreateStreamRule(ctx, rule); err == nil {
		t.Fatal("stream rule id should be empty")
	}
	rule.ID = ""
	rule.StreamID = ""
	if _, err := client.CreateStreamRule(ctx, rule); err == nil {
		t.Fatal("stream id is required")
	}
	rule.StreamID = "h"
	if _, err := client.CreateStreamRule(ctx, rule); err == nil {
		t.Fatal(`no stream with id "h" is not found`)
	}

	if _, err := client.CreateStreamRule(ctx, nil); err == nil {
		t.Fatal("stream rule is nil")
	}
}

func TestClient_UpdateStreamRule(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	stream, f, err := testutil.GetStream(ctx, client, server, 2)
	if err != nil {
		t.Fatal(err)
	}
	if f != nil {
		defer f(stream.ID)
	}
	rules, _, _, err := client.GetStreamRules(ctx, stream.ID)
	if err != nil {
		t.Fatal(err)
	}
	rule := testutil.StreamRule()
	if len(rules) == 0 {
		rule.StreamID = stream.ID
		if _, err := client.CreateStreamRule(ctx, rule); err != nil {
			t.Fatal(err)
		}
	} else {
		rule = &(rules[0])
	}

	rule.Description += " changed!"
	if _, err := client.UpdateStreamRule(ctx, rule); err != nil {
		t.Fatal(err)
	}
	streamID := rule.StreamID
	rule.StreamID = ""
	if _, err := client.UpdateStreamRule(ctx, rule); err == nil {
		t.Fatal("stream id is required")
	}
	rule.StreamID = streamID
	// ruleID = rule.ID
	rule.ID = ""
	if _, err := client.UpdateStreamRule(ctx, rule); err == nil {
		t.Fatal("stream rule id is required")
	}
	rule.ID = "h"
	if _, err := client.UpdateStreamRule(ctx, rule); err == nil {
		t.Fatal(`no stream rule with id "h" is not found`)
	}

	if _, err := client.UpdateStreamRule(ctx, nil); err == nil {
		t.Fatal("stream rule is nil")
	}
}

func TestClient_DeleteStreamRule(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	stream, f, err := testutil.GetStream(ctx, client, server, 2)
	if err != nil {
		t.Fatal(err)
	}
	if f != nil {
		defer f(stream.ID)
	}
	rules, _, _, err := client.GetStreamRules(ctx, stream.ID)
	if err != nil {
		t.Fatal(err)
	}
	rule := testutil.StreamRule()
	if len(rules) == 0 {
		rule.StreamID = stream.ID
		if _, err := client.CreateStreamRule(ctx, rule); err != nil {
			t.Fatal(err)
		}
	} else {
		rule = &(rules[0])
	}

	if _, err := client.DeleteStreamRule(ctx, "", rule.ID); err == nil {
		t.Fatal("stream id is required")
	}
	if _, err := client.DeleteStreamRule(ctx, rule.StreamID, ""); err == nil {
		t.Fatal("stream rule id is required")
	}
	if _, err := client.DeleteStreamRule(ctx, rule.StreamID, rule.ID); err != nil {
		t.Fatal(err)
	}
	if _, _, err := client.GetStreamRule(ctx, rule.StreamID, rule.ID); err == nil {
		t.Fatal("stream rule should be deleted")
	}
}

func TestClient_GetStreamRuleTypes(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	_, _, err = cl.GetStreamRuleTypes(ctx, "")
	require.NotNil(t, err, "stream id is required")

	streamID := "5d6a08796df0000000000000"

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
								Path:   "/api/streams/" + streamID + "/rules/types",
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
								},
								BodyString: `[
  {
    "id": 1,
    "name": "EXACT",
    "short_desc": "match exactly",
    "long_desc": "match exactly"
  },
  {
    "id": 2,
    "name": "REGEX",
    "short_desc": "match regular expression",
    "long_desc": "match regular expression"
  },
  {
    "id": 3,
    "name": "GREATER",
    "short_desc": "greater than",
    "long_desc": "be greater than"
  },
  {
    "id": 4,
    "name": "SMALLER",
    "short_desc": "smaller than",
    "long_desc": "be smaller than"
  },
  {
    "id": 5,
    "name": "PRESENCE",
    "short_desc": "field presence",
    "long_desc": "be present"
  },
  {
    "id": 6,
    "name": "CONTAINS",
    "short_desc": "contain",
    "long_desc": "contain"
  },
  {
    "id": 7,
    "name": "ALWAYS_MATCH",
    "short_desc": "always match",
    "long_desc": "always match"
  }
]`,
							},
						},
					},
				},
			},
		},
	})

	types, _, err := cl.GetStreamRuleTypes(ctx, streamID)
	require.Nil(t, err)
	require.Equal(t, []graylog.StreamRuleType{
		{
			ID:        1,
			Name:      "EXACT",
			ShortDesc: "match exactly",
			LongDesc:  "match exactly",
		},
		{
			ID:        2,
			Name:      "REGEX",
			ShortDesc: "match regular expression",
			LongDesc:  "match regular expression",
		},
		{
			ID:        3,
			Name:      "GREATER",
			ShortDesc: "greater than",
			LongDesc:  "be greater than",
		},
		{
			ID:        4,
			Name:      "SMALLER",
			ShortDesc: "smaller than",
			LongDesc:  "be smaller than",
		},
		{
			ID:        5,
			Name:      "PRESENCE",
			ShortDesc: "field presence",
			LongDesc:  "be present",
		},
		{
			ID:        6,
			Name:      "CONTAINS",
			ShortDesc: "contain",
			LongDesc:  "contain",
		},
		{
			ID:        7,
			Name:      "ALWAYS_MATCH",
			ShortDesc: "always match",
			LongDesc:  "always match",
		},
	}, types)
}
