package graylog

// AlertCondition represents an Alert Condition.
// http://docs.graylog.org/en/2.4/pages/streams/alerts.html#conditions
type AlertCondition struct {
	Type          string                    `json:"type"`
	ID            string                    `json:"id"`
	CreatorUserID string                    `json:"creator_user_id"`
	CreatedAt     string                    `json:"created_at"`
	Title         string                    `json:"title"`
	InGrace       bool                      `json:"in_grace"`
	Parameters    *AlertConditionParameters `json:"parameters"`
}

// AlertConditionParameters represents Alert Condition's parameters.
type AlertConditionParameters struct {
	Grace               int    `json:"grace"`
	Backlog             int    `json:"backlog"`
	RepeatNotifications bool   `json:"repeat_notifications"`
	Field               string `json:"field"`
	Value               string `json:"value"`
}

// AlertConditionsBody represents Get Alert Conditions API's response body.
// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
type AlertConditionsBody struct {
	AlertConditions []AlertCondition `json:"conditions"`
	Total           int              `json:"total"`
}
