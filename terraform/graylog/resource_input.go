package graylog

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
)

func resourceInput() *schema.Resource {
	return &schema.Resource{
		Create: resourceInputCreate,
		Read:   resourceInputRead,
		Update: resourceInputUpdate,
		Delete: resourceInputDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// required
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"configuration": {
				Type:     schema.TypeMap,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},

			"global": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"node": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"creator_user_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// "attributes": &schema.Schema{
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
			// "context_pack": &schema.Schema{
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
			// "static_fields": &schema.Schema{
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
		},
	}
}

func newInput(d *schema.ResourceData) (*graylog.Input, error) {
	cfg := d.Get("configuration").(map[string]interface{})
	config := &graylog.InputConfiguration{}

	bindAddress, err := getString(cfg, "bind_address", true)
	if err != nil {
		return nil, err
	}
	config.BindAddress = bindAddress

	port, err := getStrInt(cfg, "port", true)
	if err != nil {
		return nil, err
	}
	config.Port = port

	recvBufferSize, err := getStrInt(cfg, "recv_buffer_size", true)
	if err != nil {
		return nil, err
	}
	config.RecvBufferSize = recvBufferSize

	return &graylog.Input{
		Title:         d.Get("title").(string),
		Type:          d.Get("type").(string),
		Global:        d.Get("global").(bool),
		Node:          d.Get("node").(string),
		ID:            d.Id(),
		Configuration: config,
	}, nil
}

func resourceInputCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	input, err := newInput(d)
	if err != nil {
		return err
	}

	if _, err := cl.CreateInput(input); err != nil {
		return err
	}
	d.SetId(input.ID)
	return nil
}

func resourceInputRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	input, _, err := cl.GetInput(d.Id())
	if err != nil {
		return err
	}
	d.Set("title", input.Title)
	d.Set("type", input.Type)
	d.Set("node", input.Node)
	return nil
}

func resourceInputUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	input, err := newInput(d)
	if err != nil {
		return err
	}
	if _, err := cl.UpdateInput(input); err != nil {
		return err
	}
	return nil
}

func resourceInputDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteInput(d.Id()); err != nil {
		return err
	}
	return nil
}
