package graylog

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-graylog"
)

func resourceStream() *schema.Resource {
	return &schema.Resource{
		Create: resourceStreamCreate,
		Read:   resourceStreamRead,
		Update: resourceStreamUpdate,
		Delete: resourceStreamDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"index_set_id": {
				Type:     schema.TypeString,
				Required: true,
				// Not ForceNew
			},

			// Optional
			// rules
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// content_pack
			"matching_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remove_matches_from_default_stream": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			// attributes
			"creator_user_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// outputs
			"created_at": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_default": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			// alert_conditions
			// alert_receivers
		},
	}
}

func newStream(d *schema.ResourceData) (*graylog.Stream, error) {
	return &graylog.Stream{
		IndexSetID:   d.Get("index_set_id").(string),
		Title:        d.Get("title").(string),
		Description:  d.Get("description").(string),
		MatchingType: d.Get("matching_type").(string),
		RemoveMatchesFromDefaultStream: d.Get(
			"remove_matches_from_default_stream").(bool),
		ID: d.Id(),
	}, nil
}

func resourceStreamCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	stream, err := newStream(d)
	if err != nil {
		return err
	}

	if _, err := cl.CreateStream(ctx, stream); err != nil {
		return err
	}
	d.SetId(stream.ID)
	// resume if needed
	disabled := d.Get("disabled").(bool)
	if !disabled {
		if _, err := cl.ResumeStream(ctx, stream.ID); err != nil {
			return err
		}
	}
	return nil
}

func resourceStreamRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	stream, ei, err := cl.GetStream(ctx, d.Id())
	if err != nil {
		return handleGetResourceError(d, ei, err)
	}
	return setStream(d, stream, m.(*Config))

	// alert_receivers
	// alert_conditions
}

func resourceStreamUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	stream, err := newStream(d)
	if err != nil {
		return err
	}
	if _, err := cl.UpdateStream(ctx, stream); err != nil {
		return err
	}
	return nil
}

func resourceStreamDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteStream(ctx, d.Id()); err != nil {
		return err
	}
	return nil
}
