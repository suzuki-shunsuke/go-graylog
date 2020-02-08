package client_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog/v10"
	"github.com/suzuki-shunsuke/go-graylog/v10/client"
)

func TestClient_GetPipelines(t *testing.T) {
	ctx := context.Background()
	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		resp       string
		exp        []graylog.Pipeline
		isErr      bool
	}{{
		statusCode: 200,
		resp: `[
  {
    "id": "000000000000000000000000",
    "title": "test",
    "description": null,
    "source": "pipeline \"test\"\nstage 0 match either\nend",
    "created_at": "2019-05-17T19:42:09.643Z",
    "modified_at": "2019-05-17T19:42:09.643Z",
    "stages": [
      {
        "stage": 0,
        "match_all": false,
        "rules": []
      }
    ],
    "errors": null
  }]`,
		exp: []graylog.Pipeline{
			{
				ID:     "000000000000000000000000",
				Title:  "test",
				Source: "pipeline \"test\"\nstage 0 match either\nend",
				Stages: []graylog.PipelineStage{{Rules: []string{}}},
			},
		},
		isErr: false,
	}}
	for _, d := range data {
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
									Path:         "/api/plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/pipeline",
									PartOfHeader: getTestHeader(),
								},
								Response: &flute.Response{
									Base: http.Response{
										StatusCode: d.statusCode,
									},
									BodyString: d.resp,
								},
							},
						},
					},
				},
			},
		})
		pipes, _, err := cl.GetPipelines(ctx)
		if d.isErr {
			require.NotNil(t, err)
		} else {
			require.Nil(t, err)
			require.Equal(t, d.exp, pipes)
		}
	}
}

func TestClient_GetPipeline(t *testing.T) {
	ctx := context.Background()
	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		resp       string
		exp        *graylog.Pipeline
		isErr      bool
	}{{
		statusCode: 200,
		resp: `{
  "id": "000000000000000000000000",
  "title": "test",
  "description": null,
  "source": "pipeline \"test\"\nstage 0 match either\nend",
  "created_at": "2019-05-17T19:42:09.643Z",
  "modified_at": "2019-05-17T19:42:09.643Z",
  "stages": [
    {
      "stage": 0,
      "match_all": false,
      "rules": []
    }
  ],
  "errors": null
}`,
		exp: &graylog.Pipeline{
			ID:     "000000000000000000000000",
			Title:  "test",
			Source: "pipeline \"test\"\nstage 0 match either\nend",
			Stages: []graylog.PipelineStage{{Rules: []string{}}},
		},
		isErr: false,
	}}
	for _, d := range data {
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
									Path:         "/api/plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/pipeline/" + d.exp.ID,
									PartOfHeader: getTestHeader(),
								},
								Response: &flute.Response{
									Base: http.Response{
										StatusCode: d.statusCode,
									},
									BodyString: d.resp,
								},
							},
						},
					},
				},
			},
		})

		pipe, _, err := cl.GetPipeline(ctx, d.exp.ID)
		if d.isErr {
			require.NotNil(t, err)
		} else {
			require.Nil(t, err)
			require.Equal(t, d.exp, pipe)
		}
	}
}

func TestClient_CreatePipeline(t *testing.T) {
	ctx := context.Background()
	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		resp       string
		pipeline   *graylog.Pipeline
		exp        *graylog.Pipeline
		isErr      bool
	}{{
		statusCode: 200,
		resp: `{
  "id": "000000000000000000000000",
  "title": "test",
  "description": null,
  "source": "pipeline \"test\"\nstage 0 match either\nend",
  "created_at": "2019-05-17T19:42:09.643Z",
  "modified_at": "2019-05-17T19:42:09.643Z",
  "stages": [
    {
      "stage": 0,
      "match_all": false,
      "rules": []
    }
  ],
  "errors": null
}`,
		exp: &graylog.Pipeline{
			ID:     "000000000000000000000000",
			Title:  "test",
			Source: "pipeline \"test\"\nstage 0 match either\nend",
			Stages: []graylog.PipelineStage{{Rules: []string{}}},
		},
		pipeline: &graylog.Pipeline{
			Title:  "test",
			Source: "pipeline \"test\"\nstage 0 match either\nend",
		},
		isErr: false,
	}}
	for _, d := range data {
		cl.SetHTTPClient(&http.Client{
			Transport: &flute.Transport{
				T: t,
				Services: []flute.Service{
					{
						Endpoint: "http://example.com",
						Routes: []flute.Route{
							{
								Tester: &flute.Tester{
									Method:       "POST",
									Path:         "/api/plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/pipeline",
									PartOfHeader: getTestHeader(),
									BodyJSON:     d.pipeline,
								},
								Response: &flute.Response{
									Base: http.Response{
										StatusCode: d.statusCode,
									},
									BodyString: d.resp,
								},
							},
						},
					},
				},
			},
		})

		_, err = cl.CreatePipeline(ctx, d.pipeline)
		if d.isErr {
			require.NotNil(t, err)
		} else {
			require.Nil(t, err)
			require.Equal(t, d.exp, d.pipeline)
		}
	}
}

func TestClient_UpdatePipeline(t *testing.T) {
	ctx := context.Background()
	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		resp       string
		exp        *graylog.Pipeline
		pipeline   *graylog.Pipeline
		isErr      bool
	}{{
		statusCode: 200,
		resp: `{
  "id": "000000000000000000000000",
  "title": "test",
  "description": null,
  "source": "pipeline \"test\"\nstage 0 match either\nend",
  "created_at": "2019-05-17T19:42:09.643Z",
  "modified_at": "2019-05-17T19:42:09.643Z",
  "stages": [
    {
      "stage": 0,
      "match_all": false,
      "rules": []
    }
  ],
  "errors": null
}`,
		exp: &graylog.Pipeline{
			ID:     "000000000000000000000000",
			Title:  "test",
			Source: "pipeline \"test\"\nstage 0 match either\nend",
			Stages: []graylog.PipelineStage{{Rules: []string{}}},
		},
		pipeline: &graylog.Pipeline{
			ID:     "000000000000000000000000",
			Title:  "test",
			Source: "pipeline \"test\"\nstage 0 match either\nend",
		},
		isErr: false,
	}}
	for _, d := range data {
		cl.SetHTTPClient(&http.Client{
			Transport: &flute.Transport{
				T: t,
				Services: []flute.Service{
					{
						Endpoint: "http://example.com",
						Routes: []flute.Route{
							{
								Tester: &flute.Tester{
									Method:       "PUT",
									Path:         "/api/plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/pipeline/" + d.exp.ID,
									PartOfHeader: getTestHeader(),
								},
								Response: &flute.Response{
									Base: http.Response{
										StatusCode: d.statusCode,
									},
									BodyString: d.resp,
								},
							},
						},
					},
				},
			},
		})

		_, err = cl.UpdatePipeline(ctx, d.pipeline)
		if d.isErr {
			require.NotNil(t, err)
		} else {
			require.Nil(t, err)
			require.Equal(t, d.exp, d.pipeline)
		}
	}
}

func TestClient_DeletePipeline(t *testing.T) {
	ctx := context.Background()
	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		id         string
		isErr      bool
	}{{
		statusCode: 204,
		id:         "000000000000000000000000",
		isErr:      false,
	}}
	for _, d := range data {
		cl.SetHTTPClient(&http.Client{
			Transport: &flute.Transport{
				T: t,
				Services: []flute.Service{
					{
						Endpoint: "http://example.com",
						Routes: []flute.Route{
							{
								Tester: &flute.Tester{
									Method:       "DELETE",
									Path:         "/api/plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/pipeline/" + d.id,
									PartOfHeader: getTestHeader(),
								},
								Response: &flute.Response{
									Base: http.Response{
										StatusCode: d.statusCode,
									},
								},
							},
						},
					},
				},
			},
		})

		_, err := cl.DeletePipeline(ctx, d.id)
		if d.isErr {
			require.NotNil(t, err)
		} else {
			require.Nil(t, err)
		}
	}
}
