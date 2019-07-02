package graylog

import (
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
				Optional: true,
			},
			"config": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"timerange": {
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
						},
						// optional
						"stream_id": {
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
						"query": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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
		},
	}
}

func newDashboardWidget(d *schema.ResourceData) (*graylog.Widget, string, error) {
	cfg := d.Get("config").([]interface{})[0].(map[string]interface{})
	timeRange := cfg["timerange"].([]interface{})[0].(map[string]interface{})
	return &graylog.Widget{
		Type:          d.Get("type").(string),
		Description:   d.Get("description").(string),
		CreatorUserID: d.Get("creator_user_id").(string),
		CacheTime:     ptr.PInt(d.Get("cache_time").(int)),
		ID:            d.Id(),
		Config: &graylog.WidgetConfig{
			Timerange: &graylog.Timerange{
				Type:  timeRange["type"].(string),
				Range: timeRange["range"].(int),
			},
			LowerIsBetter: cfg["lower_is_better"].(bool),
			Trend:         cfg["trend"].(bool),
			StreamID:      cfg["stream_id"].(string),
			Query:         cfg["query"].(string),
		},
	}, d.Get("dashboard_id").(string), nil
}

func resourceDashboardWidgetCreate(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	widget, dashboardID, err := newDashboardWidget(d)
	if err != nil {
		return err
	}

	w, _, err := cl.CreateDashboardWidget(dashboardID, *widget)
	if err != nil {
		return err
	}
	d.SetId(w.ID)
	return nil
}

func resourceDashboardWidgetRead(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	widget, _, err := cl.GetDashboardWidget(
		d.Get("dashboard_id").(string), d.Id())
	if err != nil {
		return err
	}
	if err := setStrToRD(d, "type", widget.Type); err != nil {
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
	if err := d.Set("config", []map[string]interface{}{{
		"timerange": []map[string]interface{}{{
			"type":  widget.Config.Timerange.Type,
			"range": widget.Config.Timerange.Range,
		}},
		"lower_is_better": widget.Config.LowerIsBetter,
		"trend":           widget.Config.Trend,
		"stream_id":       widget.Config.StreamID,
		"query":           widget.Config.Query,
	}}); err != nil {
		return err
	}
	return setStrToRD(d, "creator_user_id", widget.CreatorUserID)
}

func resourceDashboardWidgetUpdate(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	widget, dashboardID, err := newDashboardWidget(d)
	if err != nil {
		return err
	}
	if d.HasChange("cache_time") {
		if _, err := cl.UpdateDashboardWidgetCacheTime(dashboardID, widget.ID, d.Get("cache_time").(int)); err != nil {
			return err
		}
	}
	if d.HasChange("description") {
		if _, err := cl.UpdateDashboardWidgetDescription(dashboardID, widget.ID, widget.Description); err != nil {
			return err
		}
	}
	if hasChange(d, "type", "config") {
		if _, err := cl.UpdateDashboardWidget(dashboardID, *widget); err != nil {
			return err
		}
	}
	return nil
}

func resourceDashboardWidgetDelete(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteDashboardWidget(d.Get("dashboard_id").(string), d.Id()); err != nil {
		return err
	}
	return nil
}
