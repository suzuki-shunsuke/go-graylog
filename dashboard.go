package graylog

// Dashboard represents a Graylog's Dashboard.
// http://docs.graylog.org/en/latest/pages/dashboards.html
type Dashboard struct {
	// required
	Title       string `json:"title,omitempty" v-create:"required"`
	Description string `json:"description,omitempty"`

	// ex. "2018-02-20T11:37:19.305Z"
	CreatedAt string `json:"created_at,omitempty"`
	ID        string `json:"id,omitempty" v-create:"isdefault"`
	// TODO support positions
	Widgets []Widget `json:"widgets,omitempty"`
}

// Widget represents a Graylog's Dashboard Widget.
type Widget struct {
	// ex. "STREAM_SEARCH_RESULT_COUNT"
	Type          string        `json:"type,omitempty"`
	CreatorUserID string        `json:"creator_user_id,omitempty" v-create:"isdefault"`
	ID            string        `json:"id,omitempty" v-create:"isdefault"`
	CacheTime     int           `json:"cache_time,omitempty" v-create:"isdefault"`
	Config        *WidgetConfig `json:"config,omitempty"`
}

// WidgetConfig represents a Graylog's Dashboard Widget configuration.
type WidgetConfig struct {
	Timerange     *Timerange `json:"timerange"`
	LowerIsBetter bool       `json:"lower_is_better"`
	Trend         bool       `json:"trend"`
	StreamID      string     `json:"stream_id,omitempty" v-create:"isdefault"`
	Query         string     `json:"query,omitempty" v-create:"isdefault"`
}

// Timerange represents a timerange.
type Timerange struct {
	Type  string `json:"timerange"`
	Range int    `json:"range"`
}

// DashboardsBody represents Get Dashboards API's response body.
// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
type DashboardsBody struct {
	Dashboards []Dashboard `json:"dashboards"`
	Total      int         `json:"total"`
}
