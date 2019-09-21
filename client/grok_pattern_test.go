package client_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
)

const (
	grokPatternID = "5d6a29e86df4af0000000000"
)

func TestCreateGrokPattern(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	_, err = cl.CreateGrokPattern(ctx, nil)
	require.NotNil(t, err, "grok pattern should not be nil")

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
								Path:   "/api/system/grok",
							},
							Tester: &flute.Tester{
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
								BodyJSONString: `{
								  "name": "grok pattern name",
								  "pattern": "grok pattern"
								}`,
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 201,
									Header: http.Header{
										"Content-Type": []string{"application/json"},
									},
								},
								BodyString: `{
								  "name": "grok pattern name",
								  "pattern": "grok pattern",
								  "content_pack": null,
								  "id": "5d6a29e86df4af0000000000"
								}`,
							},
						},
					},
				},
			},
		},
	})
	pattern := &graylog.GrokPattern{
		Name:    "grok pattern name",
		Pattern: "grok pattern",
	}
	_, err = cl.CreateGrokPattern(ctx, pattern)
	require.Nil(t, err)
	require.Equal(t, grokPatternID, pattern.ID)
}

func TestDeleteGrokPattern(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	_, err = cl.CreateGrokPattern(ctx, nil)
	require.NotNil(t, err, "grok pattern should not be nil")

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
								Path:   "/api/system/grok/" + grokPatternID,
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
	if _, err := cl.DeleteGrokPattern(ctx, ""); err == nil {
		t.Fatal("id is required")
	}
	// invalid id
	_, err = cl.DeleteGrokPattern(ctx, grokPatternID)
	require.Nil(t, err)
}

func TestGetGrokPattern(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	_, _, err = cl.GetGrokPattern(ctx, "")
	require.NotNil(t, err, "grok pattern id should not be empty")

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
								Path:   "/api/system/grok/" + grokPatternID,
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
								BodyString: `{
								  "name": "grok pattern name",
								  "pattern": "grok pattern",
								  "content_pack": null,
								  "id": "5d6a29e86df4af0000000000"
								}`,
							},
						},
					},
				},
			},
		},
	})
	pattern, _, err := cl.GetGrokPattern(ctx, grokPatternID)
	require.Nil(t, err)
	require.Equal(t, grokPatternID, pattern.ID)
	require.Equal(t, "grok pattern name", pattern.Name)
	require.Equal(t, "grok pattern", pattern.Pattern)
}

func TestGetGrokPatterns(t *testing.T) {
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
								Method: "GET",
								Path:   "/api/system/grok",
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
								BodyString: `{"patterns":[{
								  "name": "grok pattern name",
								  "pattern": "grok pattern",
								  "content_pack": null,
								  "id": "5d6a29e86df4af0000000000"
								}]}`,
							},
						},
					},
				},
			},
		},
	})

	patterns, _, err := cl.GetGrokPatterns(ctx)
	require.Nil(t, err)
	require.Equal(t, []graylog.GrokPattern{
		{
			Name:    "grok pattern name",
			Pattern: "grok pattern",
			ID:      grokPatternID,
		},
	}, patterns)
}

func TestUpdateGrokPattern(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	_, err = cl.UpdateGrokPattern(ctx, nil)
	require.NotNil(t, err, "grok pattern should not be nil")

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
								Path:   "/api/system/grok/" + grokPatternID,
							},
							Tester: &flute.Tester{
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
								BodyJSONString: `{
								  "id": "5d6a29e86df4af0000000000",
								  "name": "grok pattern name",
								  "pattern": "grok pattern"
								}`,
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: `{
								  "name": "grok pattern name",
								  "pattern": "grok pattern",
								  "content_pack": null,
								  "id": "5d6a29e86df4af0000000000"
								}`,
							},
						},
					},
				},
			},
		},
	})
	pattern := &graylog.GrokPattern{
		ID:      grokPatternID,
		Name:    "grok pattern name",
		Pattern: "grok pattern",
	}
	_, err = cl.UpdateGrokPattern(ctx, pattern)
	require.Nil(t, err)
	require.Equal(t, grokPatternID, pattern.ID)
	require.Equal(t, "grok pattern name", pattern.Name)
	require.Equal(t, "grok pattern", pattern.Pattern)
}
