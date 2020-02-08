package graylog

import (
	"encoding/json"
	"sort"
)

type (
	// Dashboard represents a Graylog's Dashboard.
	// https://docs.graylog.org/en/latest/pages/dashboards.html
	Dashboard struct {
		// required
		Title       string `json:"title,omitempty" v-create:"required"`
		Description string `json:"description,omitempty"`

		// ex. "2018-02-20T11:37:19.305Z"
		CreatedAt     string                    `json:"created_at,omitempty"`
		ID            string                    `json:"id,omitempty" v-create:"isdefault"`
		Widgets       []Widget                  `json:"widgets,omitempty"`
		Positions     []DashboardWidgetPosition `json:"positions"`
		CreatorUserID string                    `json:"creator_user_id,omitempty"`
	}

	DashboardWidgetPosition struct {
		WidgetID string `json:"id"`
		Width    int    `json:"width"`
		Col      int    `json:"col"`
		Row      int    `json:"row"`
		Height   int    `json:"height"`
	}

	// DashboardsBody represents Get Dashboards API's response body.
	// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
	DashboardsBody struct {
		Dashboards []Dashboard `json:"dashboards"`
		Total      int         `json:"total"`
	}
)

func (dashboard *Dashboard) UnmarshalJSON(b []byte) error {
	type alias Dashboard
	a := struct {
		Positions map[string]DashboardWidgetPosition `json:"positions"`
		*alias
	}{
		alias: (*alias)(dashboard),
	}
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	if a.Positions == nil {
		return nil
	}
	positions := []DashboardWidgetPosition{}
	for id, position := range a.Positions {
		position.WidgetID = id
		positions = append(positions, position)
	}

	// sort positions in order to test equality at test.
	// we may improve this behavior.
	sort.Slice(positions, func(i, j int) bool {
		return positions[i].WidgetID > positions[j].WidgetID
	})

	dashboard.Positions = positions
	return nil
}
