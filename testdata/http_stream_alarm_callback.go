package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog"
)

var (
	HTTPStreamAlarmCallback = graylog.AlarmCallback{
		ID:            "5d84c1a92ab79c000d35d6d4",
		StreamID:      "5d84c1a92ab79c000d35d6ca",
		Title:         "test",
		CreatorUserID: "admin",
		CreatedAt:     "2019-09-20T12:10:17.792+0000",
		Configuration: &graylog.HTTPAlarmCallbackConfiguration{
			URL: "https://example.com",
		},
	}
)
