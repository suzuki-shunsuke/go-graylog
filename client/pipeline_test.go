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

func TestClient_GetPipelines(t *testing.T) {
	ctx := context.Background()
	defer gock.Off()
	client, err := client.NewClient("http://example.com/api", "admin", "password")
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
		gock.New("http://example.com").
			Get("/api/plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/pipeline").
			MatchType("json").Reply(d.statusCode).
			BodyString(d.resp)
		pipes, _, err := client.GetPipelines(ctx)
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
	defer gock.Off()
	client, err := client.NewClient("http://example.com/api", "admin", "password")
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
		gock.New("http://example.com").
			Get(fmt.Sprintf("/api/plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/pipeline/%s", d.exp.ID)).
			MatchType("json").Reply(d.statusCode).BodyString(d.resp)
		pipe, _, err := client.GetPipeline(ctx, d.exp.ID)
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
	defer gock.Off()
	client, err := client.NewClient("http://example.com/api", "admin", "password")
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
		gock.New("http://example.com").
			Post("/api/plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/pipeline").
			MatchType("json").JSON(d.pipeline).Reply(d.statusCode).BodyString(d.resp)
		_, err := client.CreatePipeline(ctx, d.pipeline)
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
	defer gock.Off()
	client, err := client.NewClient("http://example.com/api", "admin", "password")
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
		gock.New("http://example.com").
			Put(fmt.Sprintf("/api/plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/pipeline/%s", d.exp.ID)).
			MatchType("json").Reply(d.statusCode).BodyString(d.resp)
		_, err := client.UpdatePipeline(ctx, d.pipeline)
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
	defer gock.Off()
	client, err := client.NewClient("http://example.com/api", "admin", "password")
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
		gock.New("http://example.com").
			Delete(fmt.Sprintf("/api/plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/pipeline/%s", d.id)).
			MatchType("json").Reply(d.statusCode)
		_, err := client.DeletePipeline(ctx, d.id)
		if d.isErr {
			require.NotNil(t, err)
		} else {
			require.Nil(t, err)
		}
	}
}
