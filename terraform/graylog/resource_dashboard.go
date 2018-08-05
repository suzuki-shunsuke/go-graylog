package graylog

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
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
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	db, err := newDashboard(d)
	if err != nil {
		return err
	}
	if _, err := cl.CreateDashboard(db); err != nil {
		return err
	}
	d.SetId(db.ID)
	return nil
}

func resourceDashboardRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	db, _, err := cl.GetDashboard(d.Id())
	if err != nil {
		return err
	}
	setStrToRD(d, "title", db.Title)
	setStrToRD(d, "description", db.Description)
	setStrToRD(d, "created_at", db.CreatedAt)
	return nil
}

func resourceDashboardUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	db, err := newDashboard(d)
	if err != nil {
		return err
	}

	if _, err = cl.UpdateDashboard(db); err != nil {
		return err
	}
	return nil
}

func resourceDashboardDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteDashboard(d.Id()); err != nil {
		return err
	}
	return nil
}
