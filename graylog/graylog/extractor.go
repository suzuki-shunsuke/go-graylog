package graylog

import (
	"encoding/json"
)

type (
	Extractor struct {
		ID                  string               `json:"id,omitempty"`
		Title               string               `json:"title,omitempty"`
		Type                string               `json:"type,omitempty"`
		Converters          []ExtractorConverter `json:"converters,omitempty"`
		Order               int                  `json:"order"`
		Exceptions          int                  `json:"exceptions,omitempty"`
		ConverterExceptions int                  `json:"converter_exceptions,omitempty"`
		Metrics             *ExtractorMetrics    `json:"metrics,omitempty"`
		CursorStrategy      string               `json:"cursor_strategy,omitempty"`
		SourceField         string               `json:"source_field,omitempty"`
		TargetField         string               `json:"target_field"`
		ExtractorConfig     interface{}          `json:"extractor_config,omitempty"`
		CreatorUserID       string               `json:"creator_user_id,omitempty"`
		ConditionType       string               `json:"condition_type,omitempty"`
		ConditionValue      string               `json:"condition_value"`
	}

	ExtractorConverter struct {
		Type   string                    `json:"type"`
		Config *ExtractorConverterConfig `json:"config"`
	}

	ExtractorConverterConfig struct {
		DateFormat string `json:"date_format,omitempty"`
		TimeZone   string `json:"time_zone,omitempty"`
		Locale     string `json:"locale,omitempty"`
	}

	ExtractorTypeJSONConfig struct {
		ListSeparator            string `json:"list_separator,omitempty"`
		KVSeparator              string `json:"kv_separator,omitempty"`
		KeyPrefix                string `json:"key_prefix,omitempty"`
		KeySeparator             string `json:"key_separator,omitempty"`
		ReplaceKeyWhitespace     bool   `json:"replace_key_whitespace"`
		KeyWhitespaceReplacement string `json:"key_whitespace_replacement,omitempty"`
	}

	ExtractorTypeGrokConfig struct {
		GrokPattern string `json:"grok_pattern"`
	}

	ExtractorTypeRegexConfig struct {
		RegexValue string `json:"regex_value"`
	}

	ExtractorTypeSplitAndIndexConfig struct {
		SplitBy string `json:"split_by"`
		Index   int    `json:"index"`
	}

	ExtractorMetrics struct {
		Total           *ExtractorMetric `json:"total,omitempty"`
		Condition       *ExtractorMetric `json:"condition,omitempty"`
		Execution       *ExtractorMetric `json:"execution,omitempty"`
		Converters      *ExtractorMetric `json:"converters,omitempty"`
		ConditionHits   int              `json:"condition_hits,omitempty"`
		ConditionMisses int              `json:"condition_misses,omitempty"`
	}

	ExtractorMetric struct {
		Time         *ExtractorMetricTime `json:"time,omitempty"`
		Rate         *ExtractorMetricRate `json:"rate,omitempty"`
		DurationUnit string               `json:"duration_unit,omitempty"`
		RateUnit     string               `json:"rate_unit,omitempty"`
	}

	ExtractorMetricTime struct {
		// in Graylog, the type of these fields are double.
		// https://github.com/Graylog2/graylog2-server/blob/484fdf0940718cb6bf23f812d000fd07f165eb92/graylog2-server/src/main/java/org/graylog2/rest/models/metrics/responses/TimerMetricsResponse.java#L25-L39
		Min            float64 `json:"min,omitempty"`
		Max            float64 `json:"max,omitempty"`
		Mean           float64 `json:"mean,omitempty"`
		StdDev         float64 `json:"std_dev,omitempty"`
		Percentile95th float64 `json:"95th_percentile,omitempty"`
		Percentile98th float64 `json:"98th_percentile,omitempty"`
		Percentile99th float64 `json:"99th_percentile,omitempty"`
	}

	ExtractorMetricRate struct {
		// in Graylog, the type of these fields are double.
		// https://github.com/Graylog2/graylog2-server/blob/484fdf0940718cb6bf23f812d000fd07f165eb92/graylog2-server/src/main/java/org/graylog2/rest/models/metrics/responses/RateMetricsResponse.java#L25-L35
		Total         float64 `json:"total,omitempty"`
		Mean          float64 `json:"mean,omitempty"`
		OneMinute     float64 `json:"one_minute,omitempty"`
		FiveMinute    float64 `json:"five_minute,omitempty"`
		FifteenMinute float64 `json:"fifteen_minute,omitempty"`
	}

	// ExtractorsBody represents Get Extractors API's response body.
	// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
	ExtractorsBody struct {
		Total      int         `json:"total"`
		Extractors []Extractor `json:"extractors"`
	}
)

func (extractor *Extractor) UnmarshalJSON(b []byte) error {
	type alias Extractor
	a := &struct {
		*alias
		ExtractorConfig json.RawMessage `json:"extractor_config,omitempty"`
	}{
		alias: (*alias)(extractor),
	}
	if err := json.Unmarshal(b, a); err != nil {
		return err
	}
	cfgs := map[string]interface{}{
		"json":            &ExtractorTypeJSONConfig{},
		"grok":            &ExtractorTypeGrokConfig{},
		"regex":           &ExtractorTypeRegexConfig{},
		"split_and_index": &ExtractorTypeSplitAndIndexConfig{},
	}
	if cfg, ok := cfgs[a.Type]; ok {
		if err := json.Unmarshal(a.ExtractorConfig, cfg); err != nil {
			return err
		}
		extractor.ExtractorConfig = cfg
		return nil
	}
	cfg := map[string]interface{}{}
	if err := json.Unmarshal(a.ExtractorConfig, &cfg); err != nil {
		return err
	}
	extractor.ExtractorConfig = cfg
	return nil
}
