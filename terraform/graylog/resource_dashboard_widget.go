package graylog

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-ptr"
)

func resourceDashboardWidget() *schema.Resource {
	return &schema.Resource{
		Create: resourceDashboardWidgetCreate,
		Read:   resourceDashboardWidgetRead,
		Update: resourceDashboardWidgetUpdate,
		Delete: resourceDashboardWidgetDelete,

		Importer: &schema.ResourceImporter{
			State: genImport("dashboard_id", "dashboard_widget_id"),
		},

		Schema: map[string]*schema.Schema{
			// Required
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dashboard_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"cache_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"creator_user_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"quick_values_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"timerange": timeRangeSchema(),
						// Optional
						"stream_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"query": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"interval": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"field": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"sort_order": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"stacked_fields": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"show_data_table": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"show_pie_chart": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"limit": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"data_table_limit": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"quick_values_histogram_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"timerange": timeRangeSchema(),
						// Optional
						"stream_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"query": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"field": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"sort_order": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"stacked_fields": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"limit": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"stats_count_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"timerange": timeRangeSchema(),
						// Optional
						"stream_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"query": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"field": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"stats_function": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"lower_is_better": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"trend": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"search_result_chart_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"timerange": timeRangeSchema(),
						// Optional
						"stream_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"query": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"interval": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"stream_search_result_count_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timerange": timeRangeSchema(),
						// optional
						"stream_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"query": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"lower_is_better": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"trend": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"field_chart_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"timerange": timeRangeSchema(),
						// Optional
						"stream_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"query": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"interval": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"valuetype": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"renderer": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"interpolation": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"range_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"field": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"relative": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func timeRangeSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		MinItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeString,
					Required: true,
				},
				"range": {
					Type:     schema.TypeInt,
					Required: true,
				},
			},
		},
	}
}

func newDashboardWidget(d *schema.ResourceData) (*graylog.Widget, string, error) {
	var config graylog.WidgetConfig
	t := d.Get("type").(string)
	switch t {
	case "QUICKVALUES":
		c, ok := d.GetOk("quick_values_configuration")
		if !ok {
			return nil, "", errors.New(`"quick_values_configuration" must be set`)
		}
		cfg := c.([]interface{})[0].(map[string]interface{})
		timeRange := cfg["timerange"].([]interface{})[0].(map[string]interface{})
		config = &graylog.WidgetConfigQuickValues{
			Timerange: &graylog.Timerange{
				Type:  timeRange["type"].(string),
				Range: timeRange["range"].(int),
			},
			StreamID:       cfg["stream_id"].(string),
			Query:          cfg["query"].(string),
			Interval:       cfg["interval"].(string),
			Field:          cfg["field"].(string),
			SortOrder:      cfg["sort_order"].(string),
			StackedFields:  cfg["stacked_fields"].(string),
			ShowDataTable:  cfg["show_data_table"].(bool),
			ShowPieChart:   cfg["show_pie_chart"].(bool),
			Limit:          cfg["limit"].(int),
			DataTableLimit: cfg["data_table_limit"].(int),
		}
	case "QUICKVALUES_HISTOGRAM":
		c, ok := d.GetOk("quick_values_histogram_configuration")
		if !ok {
			return nil, "", errors.New(`"quick_values_histogram_configuration" must be set`)
		}
		cfg := c.([]interface{})[0].(map[string]interface{})
		timeRange := cfg["timerange"].([]interface{})[0].(map[string]interface{})
		config = &graylog.WidgetConfigQuickValuesHistogram{
			Timerange: &graylog.Timerange{
				Type:  timeRange["type"].(string),
				Range: timeRange["range"].(int),
			},
			StreamID:      cfg["stream_id"].(string),
			Query:         cfg["query"].(string),
			Field:         cfg["field"].(string),
			SortOrder:     cfg["sort_order"].(string),
			StackedFields: cfg["stacked_fields"].(string),
			Limit:         cfg["limit"].(int),
		}
	case "STATS_COUNT":
		c, ok := d.GetOk("stats_count_configuration")
		if !ok {
			return nil, "", errors.New(`"stats_count_configuration" must be set`)
		}
		cfg := c.([]interface{})[0].(map[string]interface{})
		timeRange := cfg["timerange"].([]interface{})[0].(map[string]interface{})
		config = &graylog.WidgetConfigStatsCount{
			Timerange: &graylog.Timerange{
				Type:  timeRange["type"].(string),
				Range: timeRange["range"].(int),
			},
			StreamID:      cfg["stream_id"].(string),
			Query:         cfg["query"].(string),
			Field:         cfg["field"].(string),
			StatsFunction: cfg["stats_function"].(string),
			LowerIsBetter: cfg["lower_is_better"].(bool),
			Trend:         cfg["trend"].(bool),
		}
	case "STREAM_SEARCH_RESULT_COUNT":
		c, ok := d.GetOk("stream_search_result_count_configuration")
		if !ok {
			return nil, "", errors.New(`"stream_search_result_count_configuration" must be set`)
		}
		cfg := c.([]interface{})[0].(map[string]interface{})
		timeRange := cfg["timerange"].([]interface{})[0].(map[string]interface{})
		config = &graylog.WidgetConfigStreamSearchResultCount{
			Timerange: &graylog.Timerange{
				Type:  timeRange["type"].(string),
				Range: timeRange["range"].(int),
			},
			StreamID:      cfg["stream_id"].(string),
			Query:         cfg["query"].(string),
			LowerIsBetter: cfg["lower_is_better"].(bool),
			Trend:         cfg["trend"].(bool),
		}
	case "SEARCH_RESULT_CHART":
		c, ok := d.GetOk("search_result_chart_configuration")
		if !ok {
			return nil, "", errors.New(`"search_result_chart_configuration" must be set`)
		}
		cfg := c.([]interface{})[0].(map[string]interface{})
		timeRange := cfg["timerange"].([]interface{})[0].(map[string]interface{})
		config = &graylog.WidgetConfigSearchResultChart{
			Timerange: &graylog.Timerange{
				Type:  timeRange["type"].(string),
				Range: timeRange["range"].(int),
			},
			StreamID: cfg["stream_id"].(string),
			Query:    cfg["query"].(string),
			Interval: cfg["interval"].(string),
		}
	case "FIELD_CHART":
		c, ok := d.GetOk("field_chart_configuration")
		if !ok {
			return nil, "", errors.New(`"field_chart_configuration" must be set`)
		}
		cfg := c.([]interface{})[0].(map[string]interface{})
		timeRange := cfg["timerange"].([]interface{})[0].(map[string]interface{})
		config = &graylog.WidgetConfigFieldChart{
			Timerange: &graylog.Timerange{
				Type:  timeRange["type"].(string),
				Range: timeRange["range"].(int),
			},
			StreamID:      cfg["stream_id"].(string),
			Query:         cfg["query"].(string),
			Interval:      cfg["interval"].(string),
			ValueType:     cfg["valuetype"].(string),
			Renderer:      cfg["renderer"].(string),
			Interpolation: cfg["interpolation"].(string),
			RangeType:     cfg["range_type"].(string),
			Field:         cfg["field"].(string),
			Relative:      cfg["relative"].(int),
		}
	default:
		return nil, "", errors.New("unsupported type: " + t)
	}
	return &graylog.Widget{
		Description:   d.Get("description").(string),
		CreatorUserID: d.Get("creator_user_id").(string),
		CacheTime:     ptr.PInt(d.Get("cache_time").(int)),
		ID:            d.Id(),
		Config:        config,
	}, d.Get("dashboard_id").(string), nil
}

func resourceDashboardWidgetCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	widget, dashboardID, err := newDashboardWidget(d)
	if err != nil {
		return err
	}

	w, _, err := cl.CreateDashboardWidget(ctx, dashboardID, *widget)
	if err != nil {
		return err
	}
	d.SetId(w.ID)
	return nil
}

func resourceDashboardWidgetRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	widget, _, err := cl.GetDashboardWidget(
		ctx, d.Get("dashboard_id").(string), d.Id())
	if err != nil {
		return err
	}
	if err := setStrToRD(d, "type", widget.Type()); err != nil {
		return err
	}
	if err := setStrToRD(d, "description", widget.Description); err != nil {
		return err
	}
	if widget.CacheTime != nil {
		if err := setIntToRD(d, "cache_time", *widget.CacheTime); err != nil {
			return err
		}
	}
	switch widget.Type() {
	case "STATS_COUNT":
		cfg, ok := widget.Config.(*graylog.WidgetConfigStatsCount)
		if !ok {
			return errors.New("invalid type")
		}
		if err := d.Set("stats_count_configuration", []map[string]interface{}{{
			"timerange": []map[string]interface{}{{
				"type":  cfg.Timerange.Type,
				"range": cfg.Timerange.Range,
			}},
			"stream_id":       cfg.StreamID,
			"query":           cfg.Query,
			"field":           cfg.Field,
			"stats_function":  cfg.StatsFunction,
			"lower_is_better": cfg.LowerIsBetter,
			"trend":           cfg.Trend,
		}}); err != nil {
			return err
		}
	case "QUICKVALUES":
		cfg, ok := widget.Config.(*graylog.WidgetConfigQuickValues)
		if !ok {
			return errors.New("invalid type")
		}
		if err := d.Set("quick_values_configuration", []map[string]interface{}{{
			"timerange": []map[string]interface{}{{
				"type":  cfg.Timerange.Type,
				"range": cfg.Timerange.Range,
			}},
			"stream_id":        cfg.StreamID,
			"query":            cfg.Query,
			"interval":         cfg.Interval,
			"field":            cfg.Field,
			"sort_order":       cfg.SortOrder,
			"stacked_fields":   cfg.StackedFields,
			"show_data_table":  cfg.ShowDataTable,
			"show_pie_chart":   cfg.ShowPieChart,
			"limit":            cfg.Limit,
			"data_table_limit": cfg.DataTableLimit,
		}}); err != nil {
			return err
		}
	case "QUICKVALUES_HISTOGRAM":
		cfg, ok := widget.Config.(*graylog.WidgetConfigQuickValuesHistogram)
		if !ok {
			return errors.New("invalid type")
		}
		if err := d.Set("quick_values_histogram_configuration", []map[string]interface{}{{
			"timerange": []map[string]interface{}{{
				"type":  cfg.Timerange.Type,
				"range": cfg.Timerange.Range,
			}},
			"stream_id":      cfg.StreamID,
			"query":          cfg.Query,
			"field":          cfg.Field,
			"sort_order":     cfg.SortOrder,
			"stacked_fields": cfg.StackedFields,
			"limit":          cfg.Limit,
		}}); err != nil {
			return err
		}
	case "SEARCH_RESULT_CHART":
		cfg, ok := widget.Config.(*graylog.WidgetConfigSearchResultChart)
		if !ok {
			return errors.New("invalid type")
		}
		if err := d.Set("search_result_chart_configuration", []map[string]interface{}{{
			"timerange": []map[string]interface{}{{
				"type":  cfg.Timerange.Type,
				"range": cfg.Timerange.Range,
			}},
			"stream_id": cfg.StreamID,
			"query":     cfg.Query,
			"interval":  cfg.Interval,
		}}); err != nil {
			return err
		}
	case "STREAM_SEARCH_RESULT_COUNT":
		cfg, ok := widget.Config.(*graylog.WidgetConfigStreamSearchResultCount)
		if !ok {
			return errors.New("invalid type")
		}
		if err := d.Set("stream_search_result_count_configuration", []map[string]interface{}{{
			"timerange": []map[string]interface{}{{
				"type":  cfg.Timerange.Type,
				"range": cfg.Timerange.Range,
			}},
			"lower_is_better": cfg.LowerIsBetter,
			"trend":           cfg.Trend,
			"stream_id":       cfg.StreamID,
			"query":           cfg.Query,
		}}); err != nil {
			return err
		}
	case "FIELD_CHART":
		cfg, ok := widget.Config.(*graylog.WidgetConfigFieldChart)
		if !ok {
			return errors.New("invalid type")
		}
		if err := d.Set("field_chart_configuration", []map[string]interface{}{{
			"timerange": []map[string]interface{}{{
				"type":  cfg.Timerange.Type,
				"range": cfg.Timerange.Range,
			}},
			"stream_id":     cfg.StreamID,
			"query":         cfg.Query,
			"interval":      cfg.Interval,
			"valuetype":     cfg.ValueType,
			"renderer":      cfg.Renderer,
			"interpolation": cfg.Interpolation,
			"range_type":    cfg.RangeType,
			"field":         cfg.Field,
			"relative":      cfg.Relative,
		}}); err != nil {
			return err
		}
	default:
		return errors.New("unsupported type: " + widget.Type())
	}
	return setStrToRD(d, "creator_user_id", widget.CreatorUserID)
}

func resourceDashboardWidgetUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	widget, dashboardID, err := newDashboardWidget(d)
	if err != nil {
		return err
	}
	if d.HasChange("cache_time") {
		if _, err := cl.UpdateDashboardWidgetCacheTime(ctx, dashboardID, widget.ID, d.Get("cache_time").(int)); err != nil {
			return err
		}
	}
	if d.HasChange("description") {
		if _, err := cl.UpdateDashboardWidgetDescription(ctx, dashboardID, widget.ID, widget.Description); err != nil {
			return err
		}
	}
	if hasChange(d, "type", "quick_values_histogram_configuration", "quick_values_configuration", "stats_count_configuration", "search_result_chart_configuration", "stream_search_result_count_configuration", "field_chart_configuration") {
		if _, err := cl.UpdateDashboardWidget(ctx, dashboardID, *widget); err != nil {
			return err
		}
	}
	return nil
}

func resourceDashboardWidgetDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteDashboardWidget(ctx, d.Get("dashboard_id").(string), d.Id()); err != nil {
		return err
	}
	return nil
}
