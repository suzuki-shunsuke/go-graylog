package graylog

// AlertCondition represents an Alert Condition.
// http://docs.graylog.org/en/2.4/pages/streams/alerts.html#conditions
type AlertCondition struct {
	Type          string                    `json:"type" v-create:"required" v-update:"required"`
	ID            string                    `json:"id"`
	CreatorUserID string                    `json:"creator_user_id"`
	CreatedAt     string                    `json:"created_at"`
	Title         string                    `json:"title" v-create:"required" v-update:"required"`
	InGrace       bool                      `json:"in_grace"`
	Parameters    *AlertConditionParameters `json:"parameters" v-create:"reqired" v-update:"required"`
}

// AlertConditionParameters represents Alert Condition's parameters.
type AlertConditionParameters struct {
	Grace               int    `json:"grace" v-create:"required"`
	Backlog             int    `json:"backlog" v-create:"required"`
	RepeatNotifications bool   `json:"repeat_notifications"`
	Field               string `json:"field" v-create:"required"`
	Value               string `json:"value" v-create:"required"`
	Query               string `json:"query"`
}

// AlertConditionsBody represents Get Alert Conditions API's response body.
// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
type AlertConditionsBody struct {
	AlertConditions []AlertCondition `json:"conditions"`
	Total           int              `json:"total"`
}
