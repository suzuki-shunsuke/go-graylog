package graylog

// AlarmCallback represents an Alarm Callback.
// http://docs.graylog.org/en/latest/pages/streams/alerts.html#alert-notifications-types-explained
type AlarmCallback struct {
	// ex. "org.graylog2.alarmcallbacks.HTTPAlarmCallback"
	ID            string `json:"id"`
	CreatorUserID string `json:"creator_user_id"`
	CreatedAt     string `json:"created_at"`
	Title         string `json:"title"`
	// Type          string `json:"type"`
	// TODO add Configuration
	// Configuration AlarmCallbackConfiguration `json:"configuration"`
}

// HTTPAlarmCallbackConfiguration represents a configuration of HTTPAlarmCallback.
type HTTPAlarmCallbackConfiguration struct {
	URL string `json:"url"`
}

// EmailAlarmCallbackConfiguration represents a configuration of EmailAlarmCallback.
type EmailAlarmCallbackConfiguration struct {
	Body    string `json:"body"`
	Sender  string `json:"sender"`
	Subject string `json:"subject"`
	// UserReceivers: `json:"user_receivers"`
	// EmailReceivers: `json:"email_receivers"`
}

// AlarmCallbacksBody represents Get Alarm Callbacks API's response body.
// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
type AlarmCallbacksBody struct {
	AlarmCallbacks []AlarmCallback `json:"alarmcallbacks"`
	Total          int             `json:"total"`
}
