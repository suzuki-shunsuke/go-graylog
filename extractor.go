package graylog

type (
	Extractor struct {
		ID                  string               `json:"id,omitempty"`
		Title               string               `json:"title,omitempty"`
		Type                string               `json:"type,omitempty"`
		Converters          []ExtractorConverter `json:"converters,omitempty"`
		Order               int                  `json:"order,omitempty"`
		Exceptions          int                  `json:"exceptions,omitempty"`
		ConverterExceptions int                  `json:"converter_exceptions,omitempty"`
		Metrics             *ExtractorMetrics    `json:"metrics,omitempty"`
		CursorStrategy      string               `json:"cursor_strategy,omitempty"`
		SourceField         string               `json:"source_field,omitempty"`
		TargetField         string               `json:"target_field,omitempty"`
		ExtractorConfig     *ExtractorConfig     `json:"extractor_config,omitempty"`
		CreatorUserID       string               `json:"creator_user_id,omitempty"`
		ConditionType       string               `json:"condition_type,omitempty"`
		ConditionValue      string               `json:"condition_value,omitempty"`
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

	ExtractorConfig struct {
		ListSeparator            string `json:"list_separator,omitempty"`
		KVSeparator              string `json:"kv_separator,omitempty"`
		KeyPrefix                string `json:"key_prefix,omitempty"`
		KeySeparator             string `json:"key_separator,omitempty"`
		ReplaceKeyWhitespace     bool   `json:"replace_key_whitespace,omitempty"`
		KeyWhitespaceReplacement string `json:"key_whitespace_replacement,omitempty"`
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
		Min            int `json:"min,omitempty"`
		Max            int `json:"max,omitempty"`
		Mean           int `json:"mean,omitempty"`
		StdDev         int `json:"std_dev,omitempty"`
		Percentile95th int `json:"95th_percentile,omitempty"`
		Percentile98th int `json:"98th_percentile,omitempty"`
		Percentile99th int `json:"99th_percentile,omitempty"`
	}

	ExtractorMetricRate struct {
		Total         int     `json:"total,omitempty"`
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
