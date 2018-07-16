package graylog

// Alert represents an Alert.
// http://docs.graylog.org/en/2.4/pages/streams/alerts.html
type Alert struct {
	ID            string                    `json:"id"`
	Type          string                    `json:"type"`
	CreatorUserID string                    `json:"creator_user_id"`
	CreatedAt     string                    `json:"created_at"`
	Parameters    *AlertConditionParameters `json:"parameters"`
	InGrace       bool                      `json:"in_grace"`
	Title         string                    `json:"title"`
}

// AlertsBody represents Get Alerts API's response body.
// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
type AlertsBody struct {
	Alerts []Alert `json:"alerts"`
	Total  int     `json:"total"`
}
