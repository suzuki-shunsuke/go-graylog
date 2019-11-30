package graylog

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog/v8"
)

func resourceDashboardWidgetPositions() *schema.Resource {
	return &schema.Resource{
		Create: resourceDashboardWidgetPositionsCreate,
		Read:   resourceDashboardWidgetPositionsRead,
		Update: resourceDashboardWidgetPositionsUpdate,
		Delete: resourceDashboardWidgetPositionsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"dashboard_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"positions": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"widget_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						// Optional
						"width": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"col": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"row": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"height": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func newDashboardWidgetPositions(d *schema.ResourceData) (
	[]graylog.DashboardWidgetPosition, string, error,
) {
	arr := d.Get("positions").(*schema.Set).List()
	positions := make([]graylog.DashboardWidgetPosition, len(arr))
	for i, a := range arr {
		p := a.(map[string]interface{})
		positions[i] = graylog.DashboardWidgetPosition{
			WidgetID: p["widget_id"].(string),
			Width:    p["width"].(int),
			Col:      p["col"].(int),
			Row:      p["row"].(int),
			Height:   p["height"].(int),
		}
	}
	return positions, d.Get("dashboard_id").(string), nil
}

func resourceDashboardWidgetPositionsCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	positions, dID, err := newDashboardWidgetPositions(d)
	if err != nil {
		return err
	}
	if _, err := cl.UpdateDashboardWidgetPositions(ctx, dID, positions); err != nil {
		return err
	}
	d.SetId(dID)
	return nil
}

func resourceDashboardWidgetPositionsRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	db, _, err := cl.GetDashboard(ctx, d.Id())
	if err != nil {
		return err
	}
	if err := setStrToRD(d, "dashboard_id", d.Id()); err != nil {
		return err
	}
	positions := make([]map[string]interface{}, len(db.Positions))
	for i, p := range db.Positions {
		positions[i] = map[string]interface{}{
			"widget_id": p.WidgetID,
			"col":       p.Col,
			"width":     p.Width,
			"row":       p.Row,
			"height":    p.Height,
		}
	}
	return d.Set("positions", positions)
}

func resourceDashboardWidgetPositionsUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	positions, dID, err := newDashboardWidgetPositions(d)
	if err != nil {
		return err
	}
	if _, err := cl.UpdateDashboardWidgetPositions(ctx, dID, positions); err != nil {
		return err
	}
	return nil
}

func resourceDashboardWidgetPositionsDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.UpdateDashboardWidgetPositions(ctx, d.Id(), []graylog.DashboardWidgetPosition{}); err != nil {
		return err
	}
	return nil
}
