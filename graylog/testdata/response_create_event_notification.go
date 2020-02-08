package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
)

func EventNotification() *graylog.EventNotification {
	return &graylog.EventNotification{
		ID:          "5de5a365a1de18000cdfdf49",
		Title:       "http",
		Description: "",
		Config: map[string]interface{}{
			"type": "http-notification-v1",
			"url":  "http://example.com",
		},
	}
}
