package client_test

import (
	"context"
	"fmt"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/go-jsoneq/jsoneq"

	"github.com/suzuki-shunsuke/go-graylog/v9"
	"github.com/suzuki-shunsuke/go-graylog/v9/client"
)

func sampleExtractor1() *graylog.Extractor {
	return &graylog.Extractor{
		ID:         "e9f5e010-4315-11e9-964f-020000000000",
		Title:      "test extractor title",
		Type:       "json",
		Converters: []graylog.ExtractorConverter{},
		Metrics: &graylog.ExtractorMetrics{
			Total: &graylog.ExtractorMetric{
				Time: &graylog.ExtractorMetricTime{
					Min:            2,
					Max:            43,
					Mean:           5,
					StdDev:         4,
					Percentile95th: 11,
					Percentile98th: 12,
					Percentile99th: 13,
				},
				Rate: &graylog.ExtractorMetricRate{
					Total:         203,
					Mean:          15.336248690597982,
					OneMinute:     10.855857873186165,
					FiveMinute:    10.652891347770824,
					FifteenMinute: 10.617728486384332,
				},
				DurationUnit: "microseconds",
				RateUnit:     "events/second",
			},
			Condition: &graylog.ExtractorMetric{
				Time: &graylog.ExtractorMetricTime{
					Max:            40,
					Mean:           2,
					StdDev:         3,
					Percentile95th: 5,
					Percentile98th: 6,
					Percentile99th: 6,
				},
				Rate: &graylog.ExtractorMetricRate{
					Total:         203,
					Mean:          15.33216150262438,
					OneMinute:     10.855857873186165,
					FiveMinute:    10.652891347770824,
					FifteenMinute: 10.617728486384332,
				},
				DurationUnit: "microseconds",
				RateUnit:     "events/second",
			},
			Execution: &graylog.ExtractorMetric{
				Time:         &graylog.ExtractorMetricTime{},
				Rate:         &graylog.ExtractorMetricRate{},
				DurationUnit: "microseconds",
				RateUnit:     "events/second",
			},
			Converters: &graylog.ExtractorMetric{
				Time:         &graylog.ExtractorMetricTime{},
				Rate:         &graylog.ExtractorMetricRate{},
				DurationUnit: "microseconds",
				RateUnit:     "events/second",
			},
			ConditionHits:   0,
			ConditionMisses: 203,
		},
		CursorStrategy: "copy",
		SourceField:    "visit",
		ExtractorConfig: &graylog.ExtractorTypeJSONConfig{
			ListSeparator:            ", ",
			KVSeparator:              "=",
			KeyPrefix:                "visit_",
			KeySeparator:             "_",
			KeyWhitespaceReplacement: "_",
		},
		CreatorUserID: "admin",
		ConditionType: "none",
	}
}

func TestClient_GetExtractors(t *testing.T) {
	ctx := context.Background()
	defer gock.Off()
	client, err := client.NewClient(
		"http://example.com/api", "admin", "password")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		resp       string
		isErr      bool
		total      int
		extractors []graylog.Extractor
	}{
		{
			statusCode: 200,
			resp: `{
  "total": 1,
  "extractors": [
    {
      "id": "e9f5e010-4315-11e9-964f-020000000000",
      "title": "test extractor title",
      "type": "json",
      "converters": [],
      "order": 0,
      "exceptions": 0,
      "metrics": {
        "total": {
          "time": {
            "min": 2,
            "max": 43,
            "mean": 5,
            "std_dev": 4,
            "95th_percentile": 11,
            "98th_percentile": 12,
            "99th_percentile": 13
          },
          "rate": {
            "total": 203,
            "mean": 15.336248690597982,
            "one_minute": 10.855857873186165,
            "five_minute": 10.652891347770824,
            "fifteen_minute": 10.617728486384332
          },
          "duration_unit": "microseconds",
          "rate_unit": "events/second"
        },
        "condition": {
          "time": {
            "min": 0,
            "max": 40,
            "mean": 2,
            "std_dev": 3,
            "95th_percentile": 5,
            "98th_percentile": 6,
            "99th_percentile": 6
          },
          "rate": {
            "total": 203,
            "mean": 15.33216150262438,
            "one_minute": 10.855857873186165,
            "five_minute": 10.652891347770824,
            "fifteen_minute": 10.617728486384332
          },
          "duration_unit": "microseconds",
          "rate_unit": "events/second"
        },
        "execution": {
          "time": {
            "min": 0,
            "max": 0,
            "mean": 0,
            "std_dev": 0,
            "95th_percentile": 0,
            "98th_percentile": 0,
            "99th_percentile": 0
          },
          "rate": {
            "total": 0,
            "mean": 0,
            "one_minute": 0,
            "five_minute": 0,
            "fifteen_minute": 0
          },
          "duration_unit": "microseconds",
          "rate_unit": "events/second"
        },
        "converters": {
          "time": {
            "min": 0,
            "max": 0,
            "mean": 0,
            "std_dev": 0,
            "95th_percentile": 0,
            "98th_percentile": 0,
            "99th_percentile": 0
          },
          "rate": {
            "total": 0,
            "mean": 0,
            "one_minute": 0,
            "five_minute": 0,
            "fifteen_minute": 0
          },
          "duration_unit": "microseconds",
          "rate_unit": "events/second"
        },
        "condition_hits": 0,
        "condition_misses": 203
      },
      "cursor_strategy": "copy",
      "source_field": "visit",
      "target_field": "",
      "extractor_config": {
        "list_separator": ", ",
        "kv_separator": "=",
        "key_prefix": "visit_",
        "key_separator": "_",
        "replace_key_whitespace": false,
        "key_whitespace_replacement": "_"
      },
      "creator_user_id": "admin",
      "condition_type": "none",
      "condition_value": "",
      "converter_exceptions": 0
    }
  ]
}`,
			total: 1,
			extractors: []graylog.Extractor{
				*sampleExtractor1(),
			},
		},
	}
	id := "xxxxx"
	for _, d := range data {
		gock.New("http://example.com").
			Get(fmt.Sprintf("/api/system/inputs/%s/extractors", id)).
			MatchType("json").Reply(d.statusCode).
			BodyString(d.resp)
		extractors, total, _, err := client.GetExtractors(ctx, id)
		if d.isErr {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.Equal(t, d.extractors, extractors)
		require.Equal(t, d.total, total)
	}
}

func TestClient_GetExtractor(t *testing.T) {
	ctx := context.Background()
	defer gock.Off()
	client, err := client.NewClient(
		"http://example.com/api", "admin", "password")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		resp       string
		exp        *graylog.Extractor
		isErr      bool
	}{{
		statusCode: 200,
		resp: `{
      "id": "e9f5e010-4315-11e9-964f-020000000000",
      "title": "test extractor title",
      "type": "json",
      "converters": [],
      "order": 0,
      "exceptions": 0,
      "metrics": {
        "total": {
          "time": {
            "min": 2,
            "max": 43,
            "mean": 5,
            "std_dev": 4,
            "95th_percentile": 11,
            "98th_percentile": 12,
            "99th_percentile": 13
          },
          "rate": {
            "total": 203,
            "mean": 15.336248690597982,
            "one_minute": 10.855857873186165,
            "five_minute": 10.652891347770824,
            "fifteen_minute": 10.617728486384332
          },
          "duration_unit": "microseconds",
          "rate_unit": "events/second"
        },
        "condition": {
          "time": {
            "min": 0,
            "max": 40,
            "mean": 2,
            "std_dev": 3,
            "95th_percentile": 5,
            "98th_percentile": 6,
            "99th_percentile": 6
          },
          "rate": {
            "total": 203,
            "mean": 15.33216150262438,
            "one_minute": 10.855857873186165,
            "five_minute": 10.652891347770824,
            "fifteen_minute": 10.617728486384332
          },
          "duration_unit": "microseconds",
          "rate_unit": "events/second"
        },
        "execution": {
          "time": {
            "min": 0,
            "max": 0,
            "mean": 0,
            "std_dev": 0,
            "95th_percentile": 0,
            "98th_percentile": 0,
            "99th_percentile": 0
          },
          "rate": {
            "total": 0,
            "mean": 0,
            "one_minute": 0,
            "five_minute": 0,
            "fifteen_minute": 0
          },
          "duration_unit": "microseconds",
          "rate_unit": "events/second"
        },
        "converters": {
          "time": {
            "min": 0,
            "max": 0,
            "mean": 0,
            "std_dev": 0,
            "95th_percentile": 0,
            "98th_percentile": 0,
            "99th_percentile": 0
          },
          "rate": {
            "total": 0,
            "mean": 0,
            "one_minute": 0,
            "five_minute": 0,
            "fifteen_minute": 0
          },
          "duration_unit": "microseconds",
          "rate_unit": "events/second"
        },
        "condition_hits": 0,
        "condition_misses": 203
      },
      "cursor_strategy": "copy",
      "source_field": "visit",
      "target_field": "",
      "extractor_config": {
        "list_separator": ", ",
        "kv_separator": "=",
        "key_prefix": "visit_",
        "key_separator": "_",
        "replace_key_whitespace": false,
        "key_whitespace_replacement": "_"
      },
      "creator_user_id": "admin",
      "condition_type": "none",
      "condition_value": "",
      "converter_exceptions": 0
    }`,
		exp: sampleExtractor1(),
	}}
	inputID := "XXX"
	extractorID := "YYY"
	for _, d := range data {
		gock.New("http://example.com").
			Get(fmt.Sprintf("/api/system/inputs/%s/extractors/%s", inputID, extractorID)).
			MatchType("json").Reply(d.statusCode).
			BodyString(d.resp)
		extractor, _, err := client.GetExtractor(ctx, inputID, extractorID)
		if d.isErr {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.Equal(t, d.exp, extractor)
	}
}

func TestClient_CreateExtractor(t *testing.T) {
	ctx := context.Background()
	defer gock.Off()
	client, err := client.NewClient(
		"http://example.com/api", "admin", "password")
	require.Nil(t, err)

	data := []struct {
		statusCode int
		inputID    string
		exp        *graylog.Extractor
		extractor  *graylog.Extractor
		req        string
		resp       string
		isErr      bool
	}{{
		statusCode: 201,
		inputID:    "d3ecb503-b767-4d59-bf6a-e2c000000000",
		extractor: &graylog.Extractor{
			Title:          "test extractor title",
			CursorStrategy: "copy",
			Converters:     []graylog.ExtractorConverter{},
			SourceField:    "visit",
			Type:           "json",
			ExtractorConfig: &graylog.ExtractorTypeJSONConfig{
				ListSeparator:            ", ",
				KVSeparator:              "=",
				KeyPrefix:                "visit_",
				KeySeparator:             "_",
				ReplaceKeyWhitespace:     false,
				KeyWhitespaceReplacement: "_",
			},
			ConditionType: "none",
		},
		exp: &graylog.Extractor{
			ID:             "e9f5e010-4315-11e9-964f-020000000000",
			Title:          "test extractor title",
			CursorStrategy: "copy",
			Converters:     []graylog.ExtractorConverter{},
			SourceField:    "visit",
			Type:           "json",
			ExtractorConfig: &graylog.ExtractorTypeJSONConfig{
				ListSeparator:            ", ",
				KVSeparator:              "=",
				KeyPrefix:                "visit_",
				KeySeparator:             "_",
				ReplaceKeyWhitespace:     false,
				KeyWhitespaceReplacement: "_",
			},
			ConditionType: "none",
		},
		req: `{
      "title": "test extractor title",
			"cut_or_copy": "copy",
      "converters": {},
      "order": 0,
      "source_field": "visit",
      "target_field": "",
      "extractor_type": "json",
      "extractor_config": {
        "list_separator": ", ",
        "kv_separator": "=",
        "key_prefix": "visit_",
        "key_separator": "_",
        "replace_key_whitespace": false,
        "key_whitespace_replacement": "_"
      },
      "condition_type": "none",
      "condition_value": ""
    }`,
		resp: `{"extractor_id": "e9f5e010-4315-11e9-964f-020000000000"}`,
	}}
	for _, d := range data {
		req, err := jsoneq.Convert([]byte(d.req))
		if err != nil {
			t.Fatal(err)
		}
		gock.New("http://example.com").
			Post(fmt.Sprintf("/api/system/inputs/%s/extractors", d.inputID)).
			MatchType("json").JSON(req).Reply(d.statusCode).
			BodyString(d.resp)
		_, err = client.CreateExtractor(ctx, d.inputID, d.extractor)
		if d.isErr {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.Equal(t, d.exp, d.extractor)
	}
}

func TestClient_UpdateExtractor(t *testing.T) {
	ctx := context.Background()
	defer gock.Off()
	client, err := client.NewClient(
		"http://example.com/api", "admin", "password")
	require.Nil(t, err)

	data := []struct {
		statusCode  int
		inputID     string
		extractorID string
		extractor   *graylog.Extractor
		exp         *graylog.Extractor
		req         string
		resp        string
		isErr       bool
	}{{
		statusCode:  204,
		inputID:     "d3ecb503-b767-4d59-bf6a-e2c000000000",
		extractorID: "e9f5e010-4315-11e9-964f-020000000000",
		extractor: &graylog.Extractor{
			ID:             "e9f5e010-4315-11e9-964f-020000000000",
			Title:          "test extractor title",
			CursorStrategy: "copy",
			Converters:     []graylog.ExtractorConverter{},
			SourceField:    "visit",
			Type:           "json",
			ExtractorConfig: &graylog.ExtractorTypeJSONConfig{
				ListSeparator:            ", ",
				KVSeparator:              "=",
				KeyPrefix:                "visit_",
				KeySeparator:             "_",
				KeyWhitespaceReplacement: "_",
			},
			ConditionType: "none",
		},
		exp: sampleExtractor1(),
		req: `{
      "title": "test extractor title",
			"cut_or_copy": "copy",
      "converters": {},
      "order": 0,
      "source_field": "visit",
      "target_field": "",
      "extractor_type": "json",
      "extractor_config": {
        "list_separator": ", ",
        "kv_separator": "=",
        "key_prefix": "visit_",
        "key_separator": "_",
        "replace_key_whitespace": false,
        "key_whitespace_replacement": "_"
      },
      "condition_type": "none",
      "condition_value": ""
    }`,
		resp: `{
      "id": "e9f5e010-4315-11e9-964f-020000000000",
      "title": "test extractor title",
      "type": "json",
      "converters": [],
      "order": 0,
      "exceptions": 0,
      "metrics": {
        "total": {
          "time": {
            "min": 2,
            "max": 43,
            "mean": 5,
            "std_dev": 4,
            "95th_percentile": 11,
            "98th_percentile": 12,
            "99th_percentile": 13
          },
          "rate": {
            "total": 203,
            "mean": 15.336248690597982,
            "one_minute": 10.855857873186165,
            "five_minute": 10.652891347770824,
            "fifteen_minute": 10.617728486384332
          },
          "duration_unit": "microseconds",
          "rate_unit": "events/second"
        },
        "condition": {
          "time": {
            "min": 0,
            "max": 40,
            "mean": 2,
            "std_dev": 3,
            "95th_percentile": 5,
            "98th_percentile": 6,
            "99th_percentile": 6
          },
          "rate": {
            "total": 203,
            "mean": 15.33216150262438,
            "one_minute": 10.855857873186165,
            "five_minute": 10.652891347770824,
            "fifteen_minute": 10.617728486384332
          },
          "duration_unit": "microseconds",
          "rate_unit": "events/second"
        },
        "execution": {
          "time": {
            "min": 0,
            "max": 0,
            "mean": 0,
            "std_dev": 0,
            "95th_percentile": 0,
            "98th_percentile": 0,
            "99th_percentile": 0
          },
          "rate": {
            "total": 0,
            "mean": 0,
            "one_minute": 0,
            "five_minute": 0,
            "fifteen_minute": 0
          },
          "duration_unit": "microseconds",
          "rate_unit": "events/second"
        },
        "converters": {
          "time": {
            "min": 0,
            "max": 0,
            "mean": 0,
            "std_dev": 0,
            "95th_percentile": 0,
            "98th_percentile": 0,
            "99th_percentile": 0
          },
          "rate": {
            "total": 0,
            "mean": 0,
            "one_minute": 0,
            "five_minute": 0,
            "fifteen_minute": 0
          },
          "duration_unit": "microseconds",
          "rate_unit": "events/second"
        },
        "condition_hits": 0,
        "condition_misses": 203
      },
      "cursor_strategy": "copy",
      "source_field": "visit",
      "target_field": "",
      "extractor_config": {
        "list_separator": ", ",
        "kv_separator": "=",
        "key_prefix": "visit_",
        "key_separator": "_",
        "replace_key_whitespace": false,
        "key_whitespace_replacement": "_"
      },
      "creator_user_id": "admin",
      "condition_type": "none",
      "condition_value": "",
      "converter_exceptions": 0
    }`,
	}}
	for _, d := range data {
		req, err := jsoneq.Convert([]byte(d.req))
		if err != nil {
			t.Fatal(err)
		}
		gock.New("http://example.com").
			Put(fmt.Sprintf("/api/system/inputs/%s/extractors/%s", d.inputID, d.extractorID)).
			MatchType("json").JSON(req).Reply(d.statusCode).BodyString(d.resp)
		_, err = client.UpdateExtractor(ctx, d.inputID, d.extractor)
		if d.isErr {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.Equal(t, d.exp, d.extractor)
	}
}

func TestClient_DeleteExtractor(t *testing.T) {
	ctx := context.Background()
	defer gock.Off()
	client, err := client.NewClient(
		"http://example.com/api", "admin", "password")
	require.Nil(t, err)
	data := []struct {
		statusCode  int
		inputID     string
		extractorID string
		isErr       bool
	}{{
		statusCode:  204,
		inputID:     "d3ecb503-b767-4d59-bf6a-e2c000000000",
		extractorID: "e9f5e010-4315-11e9-964f-020000000000",
	}}
	for _, d := range data {
		gock.New("http://example.com").
			Delete(fmt.Sprintf("/api/system/inputs/%s/extractors/%s", d.inputID, d.extractorID)).
			MatchType("json").Reply(d.statusCode)
		_, err := client.DeleteExtractor(ctx, d.inputID, d.extractorID)
		if d.isErr {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
	}
}
