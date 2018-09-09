package graylog

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
)

func resourceCollectorConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceCollectorConfigurationCreate,
		Read:   resourceCollectorConfigurationRead,
		Update: resourceCollectorConfigurationUpdate,
		Delete: resourceCollectorConfigurationDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func newCollectorConfiguration(d *schema.ResourceData) (*graylog.CollectorConfiguration, error) {
	return &graylog.CollectorConfiguration{
		ID:   d.Id(),
		Name: d.Get("name").(string),
		Tags: set.NewStrSet(getStringArray(d.Get("tags").(*schema.Set).List())...),
	}, nil
}

func resourceCollectorConfigurationCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	cfg, err := newCollectorConfiguration(d)
	if err != nil {
		return err
	}
	if _, err := cl.CreateCollectorConfiguration(cfg); err != nil {
		return err
	}
	d.SetId(cfg.ID)
	return nil
}

func resourceCollectorConfigurationRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	cfg, _, err := cl.GetCollectorConfiguration(d.Id())
	if err != nil {
		return err
	}
	setStrToRD(d, "name", cfg.Name)
	setStrListToRD(d, "tags", cfg.Tags.ToList())
	return nil
}

func resourceCollectorConfigurationUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	cfg, err := newCollectorConfiguration(d)
	if err != nil {
		return err
	}

	if _, _, err = cl.RenameCollectorConfiguration(cfg.ID, cfg.Name); err != nil {
		return err
	}
	return nil
}

func resourceCollectorConfigurationDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteCollectorConfiguration(d.Id()); err != nil {
		return err
	}
	return nil
}
