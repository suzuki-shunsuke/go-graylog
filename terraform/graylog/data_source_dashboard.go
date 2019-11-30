package graylog

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog/v8"
)

func dataSourceDashboard() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDashboardRead,

		Schema: map[string]*schema.Schema{
			"dashboard_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"title": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Optional
			"created_at": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceDashboardRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}

	if id, ok := d.GetOk("dashboard_id"); ok {
		if _, ok := d.GetOk("title"); ok {
			return errors.New("only one of dashboard_id or title must be set")
		}
		db, _, err := cl.GetDashboard(ctx, id.(string))
		if err != nil {
			return err
		}
		return setDashboard(d, db)
	}

	if t, ok := d.GetOk("title"); ok {
		title := t.(string)
		dashboards, _, _, err := cl.GetDashboards(ctx)
		if err != nil {
			return err
		}
		dbs := []graylog.Dashboard{}
		for _, db := range dashboards {
			if db.Title == title {
				dbs = append(dbs, db)
			}
		}
		switch len(dbs) {
		case 0:
			return errors.New("matched dashboard is not found")
		case 1:
			return setDashboard(d, &dbs[0])
		}
		return errors.New("title isn't unique")
	}
	return errors.New("one of dashboard_id or title must be set")
}
