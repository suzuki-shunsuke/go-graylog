package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v10"
)

func StdoutOutput() *graylog.Output {
	return &graylog.Output{
		ID:            "5de32c3edf46c6001233473f",
		Title:         "test",
		Type:          "org.graylog2.outputs.LoggingOutput",
		CreatorUserID: "admin",
		CreatedAt:     "2019-12-01T02:58:06.801Z",
		Configuration: map[string]interface{}{
			"prefix": "Writing message: ",
		},
	}
}
