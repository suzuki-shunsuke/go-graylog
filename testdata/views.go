package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog"
)

var (
	Views = &graylog.Views{
		Total:   1,
		Page:    1,
		PerPage: 50,
		Count:   0,
		Views: []graylog.View{
			{
				ID:          "5d9529c175d97f58f953927a",
				Title:       "test",
				Summary:     "",
				Description: "",
				SearchID:    "5d9529b275d97f58f9539275",
				State: map[string]graylog.ViewState{
					"6971d00a-e605-43fb-b873-e4bca773d286": {
						SelectedFields: []string{
							"source",
							"message",
						},
						Titles: map[string]map[string]string{
							"widget": {
								"038b9bca-4884-496f-b1ba-bc345ad4069e": "Message Count",
								"c8986792-07e0-41fa-aded-cd19c96f2789": "All Messages",
							},
						},
						Widgets: []graylog.ViewWidget{
							{
								ID: "038b9bca-4884-496f-b1ba-bc345ad4069e",
								Config: graylog.AggregationViewWidgetConfig{
									RowPivots: []graylog.ViewWidgetRowPivot{
										{
											Field: "timestamp",
											Type:  "time",
											Config: graylog.ViewWidgetRowPivotConfig{
												Interval: graylog.ViewWidgetRowPivotInterval{
													Type: "auto",
												},
											},
										},
									},
									Series: []graylog.ViewWidgetSeries{
										{
											Config:   graylog.ViewWidgetSeriesConfig{},
											Function: "count()",
										},
									},
									Visualization: "bar",
									Rollup:        true,
								},
							},
							{
								ID: "c8986792-07e0-41fa-aded-cd19c96f2789",
								Config: graylog.MessagesViewWidgetConfig{
									Fields: []string{
										"timestamp",
										"source",
									},
									ShowMessageRow: true,
								},
							},
							{
								ID: "41c694c8-093c-4d67-be42-06390e1c61ba",
								Config: graylog.AggregationViewWidgetConfig{
									RowPivots: []graylog.ViewWidgetRowPivot{},
									Series: []graylog.ViewWidgetSeries{
										{
											Config:   graylog.ViewWidgetSeriesConfig{},
											Function: "count()",
										},
									},
									Visualization: "numeric",
									Rollup:        true,
								},
							},
						},
						WidgetMapping: map[string][]string{
							"038b9bca-4884-496f-b1ba-bc345ad4069e": {
								"9b8a4a7f-9f6e-4032-afa6-e25fe24bab40",
							},
							"41c694c8-093c-4d67-be42-06390e1c61ba": {
								"81419b6e-4d6d-4739-8883-5d34e5267091",
							},
							"c8986792-07e0-41fa-aded-cd19c96f2789": {
								"07599874-f01e-46ae-84c4-cf724e7b0524",
							},
						},
						Positions: map[string]graylog.ViewWidgetPosition{
							"038b9bca-4884-496f-b1ba-bc345ad4069e": {
								Width:  "Infinity",
								Col:    1,
								Row:    5,
								Height: 2,
							},
							"41c694c8-093c-4d67-be42-06390e1c61ba": {
								Width:  4,
								Col:    1,
								Row:    1,
								Height: 4,
							},
							"c8986792-07e0-41fa-aded-cd19c96f2789": {
								Width:  "Infinity",
								Col:    1,
								Row:    7,
								Height: 6,
							},
						},
					},
				},
				DashboardState: graylog.DashboardState{},
				Owner:          "admin",
				CreatedAt:      "2019-10-02T22:49:53.181Z",
			},
		},
	}
)
