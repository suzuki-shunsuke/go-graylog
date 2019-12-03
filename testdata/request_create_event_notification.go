package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v8"
)

var (
	RequestCreateEventNotification = &graylog.EventNotification{
		Title:       "http",
		Description: "",
		Config: map[string]interface{}{
			"type": "http-notification-v1",
			"url":  "http://example.com",
		},
	}
)
