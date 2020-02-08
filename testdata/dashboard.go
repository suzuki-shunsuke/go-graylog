package testdata

import (
	"github.com/suzuki-shunsuke/go-ptr"

	"github.com/suzuki-shunsuke/go-graylog/v10"
)

func Dashboard() graylog.Dashboard {
	return graylog.Dashboard{
		Title:       "test",
		Description: "test",
		CreatedAt:   "2019-09-20T12:10:17.486Z",
		ID:          "5d84c1a92ab79c000d35d6c7",
		Widgets: []graylog.Widget{
			{
				Description:   "Quick values",
				CreatorUserID: "admin",
				ID:            "78ae7029-0eb4-4064-b3a0-c51306093877",
				CacheTime:     ptr.PInt(10),
				Config: &graylog.WidgetConfigQuickValues{
					Timerange: &graylog.Timerange{
						Type:  "relative",
						Range: 300,
					},
					StreamID:       "5d84c1a92ab79c000d35d6ca",
					Query:          "",
					Interval:       "",
					Field:          "status",
					SortOrder:      "desc",
					StackedFields:  "",
					ShowDataTable:  true,
					ShowPieChart:   true,
					Limit:          5,
					DataTableLimit: 60,
				},
			},
			{
				Description:   "Stream search result count change",
				CreatorUserID: "admin",
				ID:            "ede5fd51-6286-40ee-9b82-249207808344",
				CacheTime:     ptr.PInt(10),
				Config: &graylog.WidgetConfigStreamSearchResultCount{
					Timerange: &graylog.Timerange{
						Type:  "relative",
						Range: 400,
					},
					StreamID:      "5d84c1a92ab79c000d35d6ca",
					Query:         "",
					LowerIsBetter: true,
					Trend:         true,
				},
			},
		},
		Positions: []graylog.DashboardWidgetPosition{
			{
				WidgetID: "ede5fd51-6286-40ee-9b82-249207808344",
				Width:    1,
				Col:      0,
				Row:      0,
				Height:   1,
			},
			{
				WidgetID: "78ae7029-0eb4-4064-b3a0-c51306093877",
				Width:    2,
				Col:      1,
				Row:      0,
				Height:   2,
			},
		},
		CreatorUserID: "admin",
	}
}
