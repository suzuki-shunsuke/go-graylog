package client_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog/v8"
	"github.com/suzuki-shunsuke/go-graylog/v8/client"
	"github.com/suzuki-shunsuke/go-graylog/v8/testdata"
)

func TestClient_GetStreamRules(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/stream_rules.json")
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
								Path:   "/api/streams/" + id + "/rules",
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

	_, _, _, err = cl.GetStreamRules(ctx, "")
	require.NotNil(t, err)

	rules, total, _, err := cl.GetStreamRules(ctx, id)
	require.Nil(t, err)
	require.Equal(t, testdata.StreamRules.Total, total)
	require.Equal(t, testdata.StreamRules.StreamRules, rules)
}

func TestClient_CreateStreamRule(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/create_stream_rule.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	rule := testdata.CreateStreamRule

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method: "POST",
								Path:   "/api/streams/" + rule.StreamID + "/rules",
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
								},
								BodyString: `{
  "streamrule_id": "5e1539d0a1de18000d89d7fe"
}`,
							},
						},
					},
				},
			},
		},
	})
	ruleID := "5e1539d0a1de18000d89d7fe"

	if _, err := cl.CreateStreamRule(ctx, &rule); err != nil {
		t.Fatal(err)
	}
	require.Equal(t, ruleID, rule.ID)
	rule.ID = ""
	rule.StreamID = ""
	if _, err := cl.CreateStreamRule(ctx, &rule); err == nil {
		t.Fatal("stream id is required")
	}

	if _, err := cl.CreateStreamRule(ctx, nil); err == nil {
		t.Fatal("stream rule is nil")
	}
}

func TestClient_UpdateStreamRule(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/create_stream_rule.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	rule := testdata.CreateStreamRule
	rule.ID = "5e1539d0a1de18000d89d7fe"

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method: "PUT",
								Path:   "/api/streams/" + rule.StreamID + "/rules/" + rule.ID,
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
								},
								BodyString: `{
  "streamrule_id": "5e1539d0a1de18000d89d7fe"
}`,
							},
						},
					},
				},
			},
		},
	})

	if _, err := cl.UpdateStreamRule(ctx, &rule); err != nil {
		t.Fatal(err)
	}
	streamID := rule.StreamID
	rule.StreamID = ""
	if _, err := cl.UpdateStreamRule(ctx, &rule); err == nil {
		t.Fatal("stream id is required")
	}
	rule.StreamID = streamID
	// ruleID = rule.ID
	rule.ID = ""
	if _, err := cl.UpdateStreamRule(ctx, &rule); err == nil {
		t.Fatal("stream rule id is required")
	}

	if _, err := cl.UpdateStreamRule(ctx, nil); err == nil {
		t.Fatal("stream rule is nil")
	}
}

func TestClient_DeleteStreamRule(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	rule := testdata.CreateStreamRule
	rule.ID = "5e1539d0a1de18000d89d7fe"

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method: "DELETE",
								Path:   "/api/streams/" + rule.StreamID + "/rules/" + rule.ID,
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 204,
								},
							},
						},
					},
				},
			},
		},
	})

	if _, err := cl.DeleteStreamRule(ctx, "", rule.ID); err == nil {
		t.Fatal("stream id is required")
	}
	if _, err := cl.DeleteStreamRule(ctx, rule.StreamID, ""); err == nil {
		t.Fatal("stream rule id is required")
	}
	if _, err := cl.DeleteStreamRule(ctx, rule.StreamID, rule.ID); err != nil {
		t.Fatal(err)
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
