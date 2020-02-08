package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v10"
)

func StreamAlertConditions() *graylog.AlertConditionsBody {
	return &graylog.AlertConditionsBody{
		AlertConditions: []graylog.AlertCondition{
			{
				ID:            "56f9f507-601d-4a54-a2f4-4bda93bb8492",
				CreatorUserID: "admin",
				CreatedAt:     "2019-09-20T12:10:17.792+0000",
				Title:         "test",
				InGrace:       false,
				Parameters: graylog.FieldContentAlertConditionParameters{
					Grace:               0,
					Backlog:             2,
					RepeatNotifications: false,
					Field:               "message",
					Value:               "hoge hoge",
					Query:               "*",
				},
			},
		},
		Total: 1,
	}
}
