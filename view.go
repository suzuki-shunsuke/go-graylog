package graylog

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type (
	// Stream represents a steram.
	View struct {
		ID          string `json:"id,omitempty"`
		Title       string `json:"title"`
		Summary     string `json:"summary"`
		Description string `json:"description"`
		SearchID    string `json:"search_id"`
		Owner       string `json:"owner"`
		CreatedAt   string `json:"created_at,omitempty"`
		// Properties []interface{} `json:"properties"`
		// Requires map[string]interface{} `json:"requires"`
		DashboardState DashboardState       `json:"dashboard_state"`
		State          map[string]ViewState `json:"state"`
	}

	ViewState struct {
		SelectedFields []string                      `json:"selected_fields"`
		Titles         map[string]map[string]string  `json:"titles"`
		WidgetMapping  map[string][]string           `json:"widget_mapping"`
		Widgets        []ViewWidget                  `json:"widgets"`
		Positions      map[string]ViewWidgetPosition `json:"positions"`
		// StatecMessageListID interface{} `json:"state_message_list_id"`
		// Formatting interface{} `json:"formatting"`
	}

	ViewWidget struct {
		ID string `json:"id"`
		// Filter interface{} `json:"filter"`
		Config ViewWidgetConfig
	}

	ViewWidgetConfig interface {
		Type() string
	}

	AggregationViewWidgetConfig struct {
		RowPivots     []ViewWidgetRowPivot `json:"row_pivots"`
		Series        []ViewWidgetSeries   `json:"series"`
		Visualization string               `json:"visualization"`
		Rollup        bool                 `json:"rollup"`
		// Sort []interface{} `json:"sort"`
		// ColumnPivots []interface{} `json:"column_pivots"`
		// VisualizationConfig interface{} `json:"visualization_config"`
		// FormattingSettings interface{} `json:"formatting_settings"`
	}

	MessagesViewWidgetConfig struct {
		Fields         []string `json:"fields"`
		ShowMessageRow bool     `json:"show_message_row"`
	}

	ViewWidgetSeries struct {
		Config   ViewWidgetSeriesConfig `json:"config"`
		Function string                 `json:"function"`
	}

	ViewWidgetSeriesConfig struct {
		// Name interface{}
	}

	ViewWidgetRowPivot struct {
		Field  string                   `json:"field"`
		Type   string                   `json:"type"`
		Config ViewWidgetRowPivotConfig `json:"config"`
	}

	ViewWidgetRowPivotConfig struct {
		Interval ViewWidgetRowPivotInterval
	}

	ViewWidgetRowPivotInterval struct {
		Type string `json:"type"`
		// Scaling   interface{} `json:"scalling"
	}

	ViewWidgetPosition struct {
		Width  interface{} `json:"width"` // int or "Infinity"
		Col    int         `json:"col"`
		Row    int         `json:"row"`
		Height int         `json:"height"`
	}

	DashboardState struct {
		// Widgets interface{} `json:"widgets"`
		// Positions interface{} `json:"positions"`
	}

	// Views represents Get View API's response body.
	Views struct {
		Total   int    `json:"total"`
		Page    int    `json:"page"`
		PerPage int    `json:"per_page"`
		Count   int    `json:"count"`
		Views   []View `json:"views"`
	}
)

func (widget ViewWidget) Type() string {
	return widget.Config.Type()
}

func (widget AggregationViewWidgetConfig) Type() string {
	return "aggregation"
}

func (widget MessagesViewWidgetConfig) Type() string {
	return "messages"
}

// UnmarshalJSON unmarshals JSON into an alert condition.
func (widget *ViewWidget) UnmarshalJSON(b []byte) error {
	errMsg := "failed to unmarshal JSON to view widget"
	if widget == nil {
		return fmt.Errorf("%s: view widget is nil", errMsg)
	}
	type alias ViewWidget
	a := struct {
		Type   string          `json:"type"`
		Config json.RawMessage `json:"config"`
		*alias
	}{
		alias: (*alias)(widget),
	}
	if err := json.Unmarshal(b, &a); err != nil {
		return errors.Wrap(err, errMsg)
	}
	switch a.Type {
	case "aggregation":
		p := AggregationViewWidgetConfig{}
		if err := json.Unmarshal(a.Config, &p); err != nil {
			return errors.Wrap(err, errMsg)
		}
		widget.Config = p
		return nil
	case "messages":
		p := MessagesViewWidgetConfig{}
		if err := json.Unmarshal(a.Config, &p); err != nil {
			return errors.Wrap(err, errMsg)
		}
		widget.Config = p
		return nil
	}
	// TODO
	return nil
}

// UnmarshalJSON unmarshals JSON into an alert condition.
func (position *ViewWidgetPosition) UnmarshalJSON(b []byte) error {
	errMsg := "failed to unmarshal JSON to view widget position"
	if position == nil {
		return fmt.Errorf("%s: view widget position is nil", errMsg)
	}
	type alias ViewWidgetPosition

	a := struct {
		*alias
		Width string `json:"width"`
	}{
		alias: (*alias)(position),
	}
	if err := json.Unmarshal(b, &a); err == nil {
		position.Width = a.Width
		return nil
	}

	c := struct {
		*alias
		Width json.Number `json:"width"`
	}{
		alias: (*alias)(position),
	}
	if err := json.Unmarshal(b, &c); err != nil {
		return errors.Wrap(err, errMsg)
	}
	if i, err := c.Width.Int64(); err == nil {
		position.Width = int(i)
		return nil
	}
	if f, err := c.Width.Float64(); err == nil {
		position.Width = f
		return nil
	}
	// TODO
	return nil
}
