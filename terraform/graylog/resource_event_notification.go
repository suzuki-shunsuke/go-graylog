package graylog

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-jsoneq/jsoneq"

	"github.com/suzuki-shunsuke/go-graylog/v8"
)

func resourceEventNotification() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventNotificationCreate,
		Read:   resourceEventNotificationRead,
		Update: resourceEventNotificationUpdate,
		Delete: resourceEventNotificationDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"config": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: cfgSchemaDiffSuppress,
			},
		},
	}
}

func newEventNotification(d *schema.ResourceData) (*graylog.EventNotification, error) {
	cfgS := d.Get("config").(string)
	cfg, err := jsoneq.ConvertByte([]byte(cfgS))
	if err != nil {
		return nil, fmt.Errorf("failed to parse the 'config'. 'config' must be a JSON string '%s': %w", cfgS, err)
	}
	return &graylog.EventNotification{
		ID:          d.Id(),
		Title:       d.Get("title").(string),
		Description: d.Get("description").(string),
		Config:      cfg,
	}, nil
}

func resourceEventNotificationCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	notif, err := newEventNotification(d)
	if err != nil {
		return err
	}
	if _, err := cl.CreateEventNotification(ctx, notif); err != nil {
		return err
	}
	d.SetId(notif.ID)
	return nil
}

func resourceEventNotificationRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	notif, ei, err := cl.GetEventNotification(ctx, d.Id())
	if err != nil {
		return handleGetResourceError(d, ei, err)
	}
	if err := setStrToRD(d, "title", notif.Title); err != nil {
		return err
	}
	if err := setStrToRD(d, "description", notif.Description); err != nil {
		return err
	}
	b, err := json.Marshal(notif.Config)
	if err != nil {
		return err
	}
	return setStrToRD(d, "config", string(b))
}

func resourceEventNotificationUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	notif, err := newEventNotification(d)
	if err != nil {
		return err
	}
	_, err = cl.UpdateEventNotification(ctx, notif)
	return err
}

func resourceEventNotificationDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteEventNotification(ctx, d.Id()); err != nil {
		return err
	}
	return nil
}
