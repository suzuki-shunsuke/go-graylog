package terraform

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
)

func resourceDashboard() *schema.Resource {
	return &schema.Resource{
		Create: resourceDashboardCreate,
		Read:   resourceDashboardRead,
		Update: resourceDashboardUpdate,
		Delete: resourceDashboardDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"created_at": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// TODO support widget
			// "widgets": {
			// 	Type:     schema.TypeSet,
			// 	Optional: true,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"type": {
			// 				Type:     schema.TypeString,
			// 				Optional: true,
			// 			},
			// 			"creator_user_id": {
			// 				Type:     schema.TypeString,
			// 				Optional: true,
			// 				Computed: true,
			// 			},
			// 			"cache_time": {
			// 				Type:     schema.TypeInt,
			// 				Optional: true,
			// 			},
			// 			// TODO add config
			// 		},
			// 	},
			// },
		},
	}
}

func setDashboard(d *schema.ResourceData, db *graylog.Dashboard) error {
	d.SetId(db.ID)
	if err := setStrToRD(d, "title", db.Title); err != nil {
		return err
	}
	if err := setStrToRD(d, "description", db.Description); err != nil {
		return err
	}
	return setStrToRD(d, "created_at", db.CreatedAt)
}

func newDashboard(d *schema.ResourceData) (*graylog.Dashboard, error) {
	return &graylog.Dashboard{
		ID:          d.Id(),
		Title:       d.Get("title").(string),
		Description: d.Get("description").(string),
		CreatedAt:   d.Get("created_at").(string),
		// Widgets:     d.Get("widgets").(string),
	}, nil
}

func resourceDashboardCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	db, err := newDashboard(d)
	if err != nil {
		return err
	}
	if _, err := cl.CreateDashboard(ctx, db); err != nil {
		return err
	}
	d.SetId(db.ID)
	return nil
}

func resourceDashboardRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	db, ei, err := cl.GetDashboard(ctx, d.Id())
	if err != nil {
		return handleGetResourceError(d, ei, err)
	}
	return setDashboard(d, db)
}

func resourceDashboardUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	db, err := newDashboard(d)
	if err != nil {
		return err
	}

	if _, err = cl.UpdateDashboard(ctx, db); err != nil {
		return err
	}
	return nil
}

func resourceDashboardDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteDashboard(ctx, d.Id()); err != nil {
		return err
	}
	return nil
}
