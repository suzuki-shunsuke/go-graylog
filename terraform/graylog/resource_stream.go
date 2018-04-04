package graylog

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
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
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"index_set_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			// rules
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			// content_pack
			"matching_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remove_matches_from_default_stream": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},

			// attributes
			"creator_user_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// outputs
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_default": &schema.Schema{
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
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	stream, err := newStream(d)
	if err != nil {
		return err
	}

	if _, err := cl.CreateStream(stream); err != nil {
		return err
	}
	d.SetId(stream.ID)
	return nil
}

func resourceStreamRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	stream, _, err := cl.GetStream(d.Id())
	if err != nil {
		return err
	}
	d.Set("index_set_id", stream.IndexSetID)
	d.Set("title", stream.Title)
	d.Set("description", stream.Description)
	d.Set("matching_type", stream.MatchingType)
	d.Set(
		"remove_matches_from_default_stream",
		stream.RemoveMatchesFromDefaultStream)
	// rules
	// content_pack
	d.Set("creator_user_id", stream.CreatorUserID)
	d.Set("created_at", stream.CreatedAt)
	d.Set("disabled", stream.Disabled)
	d.Set("is_default", stream.IsDefault)
	// alert_receivers
	// alert_conditions
	return nil
}

func resourceStreamUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	stream, err := newStream(d)
	if err != nil {
		return err
	}
	if _, err := cl.UpdateStream(stream); err != nil {
		return err
	}
	return nil
}

func resourceStreamDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteStream(d.Id()); err != nil {
		return err
	}
	return nil
}
