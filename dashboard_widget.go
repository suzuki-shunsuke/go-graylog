package graylog

import (
	"encoding/json"
)

type (
	// Widget represents a Graylog's Dashboard Widget.
	Widget struct {
		// ex. "STREAM_SEARCH_RESULT_COUNT"
		Description   string       `json:"description,omitempty" v-create:"required"`
		CreatorUserID string       `json:"creator_user_id,omitempty" v-create:"isdefault"`
		ID            string       `json:"id,omitempty" v-create:"isdefault"`
		CacheTime     *int         `json:"cache_time,omitempty" v-create:"isdefault"`
		Config        WidgetConfig `json:"config,omitempty" v-create:"required"`
	}

	// WidgetConfig represents a Graylog's Dashboard Widget configuration.
	WidgetConfig interface {
		Type() string
	}

	WidgetConfigStreamSearchResultCount struct {
		Timerange     *Timerange `json:"timerange" v-create:"required"`
		StreamID      string     `json:"stream_id,omitempty" v-create:"isdefault"`
		Query         string     `json:"query" v-create:"isdefault"`
		LowerIsBetter bool       `json:"lower_is_better"`
		Trend         bool       `json:"trend"`
	}

	WidgetConfigSearchResultChart struct {
		Timerange *Timerange `json:"timerange" v-create:"required"`
		StreamID  string     `json:"stream_id,omitempty" v-create:"isdefault"`
		Query     string     `json:"query" v-create:"isdefault"`
		Interval  string     `json:"interval,omitempty"`
	}

	WidgetConfigQuickValues struct {
		Timerange      *Timerange `json:"timerange" v-create:"required"`
		StreamID       string     `json:"stream_id,omitempty" v-create:"isdefault"`
		Query          string     `json:"query" v-create:"isdefault"`
		Interval       string     `json:"interval,omitempty"`
		Field          string     `json:"field,omitempty"`
		SortOrder      string     `json:"sort_order,omitempty"`
		StackedFields  string     `json:"stacked_fields,omitempty"`
		ShowDataTable  bool       `json:"show_data_table,omitempty"`
		ShowPieChart   bool       `json:"show_pie_chart,omitempty"`
		Limit          int        `json:"limit,omitempty"`
		DataTableLimit int        `json:"data_table_limit,omitempty"`
	}

	WidgetConfigQuickValuesHistogram struct {
		Timerange     *Timerange `json:"timerange" v-create:"required"`
		StreamID      string     `json:"stream_id,omitempty" v-create:"isdefault"`
		Query         string     `json:"query" v-create:"isdefault"`
		Field         string     `json:"field,omitempty"`
		SortOrder     string     `json:"sort_order,omitempty"`
		StackedFields string     `json:"stacked_fields,omitempty"`
		Limit         int        `json:"limit,omitempty"`
	}

	WidgetConfigFieldChart struct {
		Timerange     *Timerange `json:"timerange"`
		StreamID      string     `json:"stream_id,omitempty"`
		Query         string     `json:"query"`
		Interval      string     `json:"interval,omitempty"`
		ValueType     string     `json:"valuetype,omitempty"`
		Renderer      string     `json:"renderer,omitempty"`
		Interpolation string     `json:"interpolation,omitempty"`
		RangeType     string     `json:"rangeType,omitempty"`
		Field         string     `json:"field,omitempty"`
		Relative      int        `json:"relative,omitempty"`
	}

	WidgetConfigStatsCount struct {
		Timerange     *Timerange `json:"timerange"`
		StreamID      string     `json:"stream_id,omitempty"`
		Query         string     `json:"query"`
		Field         string     `json:"field,omitempty"`
		StatsFunction string     `json:"stats_function,omitempty"`
		LowerIsBetter bool       `json:"lower_is_better,omitempty"`
		Trend         bool       `json:"trend,omitempty"`
	}

	WidgetConfigUnknownType struct {
		T      string
		Fields map[string]interface{}
	}

	// Timerange represents a timerange.
	Timerange struct {
		Type  string `json:"type" v-create:"required"`
		Range int    `json:"range" v-create:"required"`
	}
)

func (widget *Widget) Type() string {
	return widget.Config.Type()
}

func setConfig(widget *Widget, src json.RawMessage, cfg WidgetConfig) error {
	if err := json.Unmarshal(src, cfg); err != nil {
		return err
	}
	widget.Config = cfg
	return nil
}

func (widget *Widget) UnmarshalJSON(b []byte) error {
	type alias Widget
	a := struct {
		Type   string          `json:"type"`
		Config json.RawMessage `json:"config"`
		*alias
	}{
		alias: (*alias)(widget),
	}
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	switch a.Type {
	case "STATS_COUNT":
		cfg := &WidgetConfigStatsCount{}
		return setConfig(widget, a.Config, cfg)
	case "QUICKVALUES":
		cfg := &WidgetConfigQuickValues{}
		return setConfig(widget, a.Config, cfg)
	case "QUICKVALUES_HISTOGRAM":
		cfg := &WidgetConfigQuickValuesHistogram{}
		return setConfig(widget, a.Config, cfg)
	case "FIELD_CHART":
		cfg := &WidgetConfigFieldChart{}
		return setConfig(widget, a.Config, cfg)
	case "STREAM_SEARCH_RESULT_COUNT":
		cfg := &WidgetConfigStreamSearchResultCount{}
		return setConfig(widget, a.Config, cfg)
	case "SEARCH_RESULT_CHART":
		cfg := &WidgetConfigSearchResultChart{}
		return setConfig(widget, a.Config, cfg)
		// case "STACKED_CHART":
		// case "SEARCH_RESULT_COUNT":
	}
	cfg := &WidgetConfigUnknownType{
		T: a.Type,
	}
	return setConfig(widget, a.Config, cfg)
}

func (widget *Widget) MarshalJSON() ([]byte, error) {
	type alias Widget
	return json.Marshal(struct {
		Type string `json:"type"`
		*alias
	}{
		Type:  widget.Type(),
		alias: (*alias)(widget),
	})
}

func (cfg *WidgetConfigQuickValues) Type() string {
	return "QUICKVALUES"
}

func (cfg *WidgetConfigQuickValuesHistogram) Type() string {
	return "QUICKVALUES_HISTOGRAM"
}

func (cfg *WidgetConfigStatsCount) Type() string {
	return "STATS_COUNT"
}

func (cfg *WidgetConfigFieldChart) Type() string {
	return "FIELD_CHART"
}

func (cfg *WidgetConfigStreamSearchResultCount) Type() string {
	return "STREAM_SEARCH_RESULT_COUNT"
}

func (cfg *WidgetConfigSearchResultChart) Type() string {
	return "SEARCH_RESULT_CHART"
}

func (cfg *WidgetConfigUnknownType) Type() string {
	return cfg.T
}

func (cfg *WidgetConfigUnknownType) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &cfg.Fields)
}
