package graylog

type (
	// Dashboard represents a Graylog's Dashboard.
	// https://docs.graylog.org/en/latest/pages/dashboards.html
	Dashboard struct {
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
	Widget struct {
		// ex. "STREAM_SEARCH_RESULT_COUNT"
		Type          string        `json:"type,omitempty" v-create:"required"`
		Description   string        `json:"description,omitempty" v-create:"required"`
		CreatorUserID string        `json:"creator_user_id,omitempty" v-create:"isdefault"`
		ID            string        `json:"id,omitempty" v-create:"isdefault"`
		CacheTime     *int          `json:"cache_time,omitempty" v-create:"isdefault"`
		Config        *WidgetConfig `json:"config,omitempty" v-create:"required"`
	}

	// WidgetConfig represents a Graylog's Dashboard Widget configuration.
	WidgetConfig struct {
		Timerange     *Timerange `json:"timerange" v-create:"required"`
		LowerIsBetter bool       `json:"lower_is_better,omitempty"`
		Trend         bool       `json:"trend,omitempty"`
		StreamID      string     `json:"stream_id,omitempty" v-create:"isdefault"`
		Query         string     `json:"query,omitempty" v-create:"isdefault"`
	}

	// Timerange represents a timerange.
	Timerange struct {
		Type  string `json:"type" v-create:"required"`
		Range int    `json:"range" v-create:"required"`
	}

	// DashboardsBody represents Get Dashboards API's response body.
	// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
	DashboardsBody struct {
		Dashboards []Dashboard `json:"dashboards"`
		Total      int         `json:"total"`
	}
)
