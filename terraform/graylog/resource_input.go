package graylog

import (
	"encoding/json"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
	"github.com/suzuki-shunsuke/go-graylog/util"
	"github.com/suzuki-shunsuke/go-ptr"
)

func resourceInput() *schema.Resource {
	cfgSchema := map[string]*schema.Schema{}
	for _, s := range graylog.InputAttributesStrFields {
		cfgSchema[s] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		}
	}
	for _, s := range graylog.InputAttributesIntFields {
		cfgSchema[s] = &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		}
	}
	for _, s := range graylog.InputAttributesBoolFields {
		cfgSchema[s] = &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		}
	}
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
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"attributes": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: cfgSchema,
				},
				MaxItems: 1,
				MinItems: 1,
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
				Computed: true,
			},
			"creator_user_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
	cfg := d.Get("attributes").([]interface{})[0].(map[string]interface{})
	config := &graylog.InputAttributes{}
	if err := util.MSDecode(cfg, config); err != nil {
		return nil, err
	}
	return &graylog.Input{
		Title:         d.Get("title").(string),
		Type:          d.Get("type").(string),
		Global:        ptr.PBool(d.Get("global").(bool)),
		Node:          d.Get("node").(string),
		ID:            d.Id(),
		Attributes:    config,
		CreatorUserID: d.Get("creator_user_id").(string),
		CreatedAt:     d.Get("created_at").(string),
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
	input, ei, err := cl.GetInput(d.Id())
	if err != nil {
		if ei != nil && ei.Response != nil && ei.Response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return err
	}
	if input.Attributes != nil {
		b, err := json.Marshal(input.Attributes)
		if err != nil {
			return err
		}
		dest := map[string]interface{}{}
		if err := json.Unmarshal(b, &dest); err != nil {
			return err
		}
		d.Set("attributes", dest)
	}
	d.Set("title", input.Title)
	d.Set("type", input.Type)
	d.Set("node", input.Node)
	d.Set("creator_user_id", input.CreatorUserID)
	d.Set("created_at", input.CreatedAt)
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
