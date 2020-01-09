package graylog_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/v9"
)

func TestExtractor_MarshalJSON(t *testing.T) {
	extractor := &graylog.Extractor{}
	b := []byte(`{
  "id": "4306a4f0-7c69-11e9-b7ff-0242ac1d0004",
  "title": "test",
  "type": "grok",
  "converters": [],
  "order": 0,
  "exceptions": 0,
  "metrics": {
    "total": {
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
    "condition": {
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
    "condition_misses": 0
  },
  "cursor_strategy": "copy",
  "source_field": "message",
  "target_field": "none",
  "extractor_config": {
    "grok_pattern": "%{DATA}"
  },
  "creator_user_id": "admin",
  "condition_type": "none",
  "condition_value": "",
  "converter_exceptions": 0
}`)
	err := json.Unmarshal(b, extractor)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.NewDecoder(bytes.NewBuffer(b)).Decode(extractor); err != nil {
		t.Fatal(err)
	}
}
