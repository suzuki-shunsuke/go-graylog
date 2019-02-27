package client_test

import (
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
)

func TestGetPipelineRules(t *testing.T) {
	defer gock.Off()
	client, err := client.NewClient(
		"http://example.com/api", "admin", "password")
	require.Nil(t, err)
	data := []struct {
		statusCode int
		resp       string
		rules      []graylog.PipelineRule
		isErr      bool
	}{{
		statusCode: 200,
		resp: `[
  {
    "title": "has request_time",
    "description": "nginx request_time",
    "source": "rule \"has request_time\"\nwhen\n    true\nthen\n    set_field(\"request_time_double\", to_double($message.request_time));\n    set_field(\"request_time\", to_double($message.request_time));\nend",
    "created_at": "2018-01-01T23:00:00.000Z",
    "modified_at": "2018-01-02T00:00:00.000Z",
    "errors": null,
    "id": "5c732c6dc9e77c0000000000"
  },
  {
    "title": "convert status not 5xx to 1",
    "description": "description convert status not 5xx to 1",
    "source": "rule \"convert status not 5xx to 1\"\nwhen\n    to_long($message.status) < 500\nthen\n    set_field(\"status_01\", 1);\nend",
    "created_at": "2018-01-01T23:00:00.000Z",
    "modified_at": "2018-01-02T00:00:00.000Z",
    "errors": null,
    "id": "5c7360000000000000000000"
  }]`,
		rules: []graylog.PipelineRule{
			{
				ID:          "5c732c6dc9e77c0000000000",
				Title:       "has request_time",
				Description: "nginx request_time",
				Source:      "rule \"has request_time\"\nwhen\n    true\nthen\n    set_field(\"request_time_double\", to_double($message.request_time));\n    set_field(\"request_time\", to_double($message.request_time));\nend",
			}, {
				ID:          "5c7360000000000000000000",
				Title:       "convert status not 5xx to 1",
				Description: "description convert status not 5xx to 1",
				Source:      "rule \"convert status not 5xx to 1\"\nwhen\n    to_long($message.status) < 500\nthen\n    set_field(\"status_01\", 1);\nend",
			},
		},
		isErr: false,
	}}
	for _, d := range data {
		gock.New("http://example.com").
			Get("/api/plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/rule").
			MatchType("json").Reply(d.statusCode).
			BodyString(d.resp)
		rules, _, err := client.GetPipelineRules()
		if d.isErr {
			require.NotNil(t, err)
		} else {
			require.Nil(t, err)
			require.Equal(t, d.rules, rules)
		}
	}
}
