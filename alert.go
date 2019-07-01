package graylog

// Alert represents an Alert.
// https://docs.graylog.org/en/2.4/pages/streams/alerts.html
type Alert struct {
	ID                  string                    `json:"id"`
	Description         string                    `json:"description"`
	ConditionID         string                    `json:"condition_id"`
	StreamID            string                    `json:"stream_id"`
	TriggeredAt         string                    `json:"triggered_at"`
	ResolvedAt          string                    `json:"resolved_at"`
	IsInterval          bool                      `json:"is_interval"`
	ConditionParameters *AlertConditionParameters `json:"condition_parameters"`
}

// AlertsBody represents Get Alerts API's response body.
// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
type AlertsBody struct {
	Alerts []Alert `json:"alerts"`
	Total  int     `json:"total"`
}
