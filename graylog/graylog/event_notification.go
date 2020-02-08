package graylog

type (
	// EventNotification represents an output.
	EventNotification struct {
		ID          string      `json:"id,omitempty"`
		Title       string      `json:"title"`
		Description string      `json:"description"`
		Config      interface{} `json:"config"`
	}

	// EventNotificationsBody represents Get EventNotifications API's response body.
	EventNotificationsBody struct {
		EventNotifications []EventNotification `json:"notifications"`
		Total              int                 `json:"total"`
		Page               int                 `json:"page"`
		PerPage            int                 `json:"per_page"`
		Count              int                 `json:"count"`
		GrandTotal         int                 `json:"grand_total"`
		Query              string              `json:"query"`
	}
)
