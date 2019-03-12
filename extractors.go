package graylog

// Extractor represents a Graylog's Input Extractor.
// http://docs.graylog.org/en/3.0/pages/extractors.html
type Extractor struct {
	ID                  string                   `json:"id,omitempty" v-create:"isdefault"`
	Title               string                   `json:"title,omitempty" v-create:"required"`
	Type                string                   `json:"type,omitempty" v-create:"required"`
	CursorStrategy      string                   `json:"cursor_strategy,omitempty" v-create:"required"`
	SourceField         string                   `json:"source_field,omitempty" v-create:"required"`
	TargetField         string                   `json:"target_field,omitempty" v-create:"required"`
	ExtractorConfig     map[string]interface{}   `json:"extractor_config,omitempty" v-create:"required"`
	CreatorUserID       string                   `json:"creator_user_id,omitempty" v-create:"isdefault"`
	Converters          []map[string]interface{} `json:"converters,omitempty" v-create:"required"`
	ConditionType       string                   `json:"condition_type,omitempty" v-create:"required"`
	ConditionValue      string                   `json:"condition_value,omitempty" v-create:"required"`
	Order               int                      `json:"order,omitempty" v-create:"required"`
	Exceptions          int                      `json:"exceptions,omitempty" v-create:"required"`
	ConverterExceptions int                      `json:"converter_exceptions,omitempty" v-create:"required"`
	Metrics             map[string]interface{}   `json:"metrics,omitempty" v-create:"required"`
}

// ExtractorsBody represents Get Extractors API's response body.
// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
type ExtractorsBody struct {
	Extractors []Extractor `json:"extractors"`
	Total      int         `json:"total"`
}
