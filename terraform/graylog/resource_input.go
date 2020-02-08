package graylog

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog/v10"
)

func resourceInput() *schema.Resource {
	cfgSchema := map[string]*schema.Schema{}
	for s := range graylog.InputAttrsStrFieldSet.ToMap(false) {
		cfgSchema[s] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		}
	}
	for s := range graylog.InputAttrsIntFieldSet.ToMap(false) {
		cfgSchema[s] = &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		}
	}
	for s := range graylog.InputAttrsBoolFieldSet.ToMap(false) {
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
			"static_fields": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			// "context_pack": &schema.Schema{
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
		},
	}
}

func newInput(d *schema.ResourceData) (*graylog.Input, error) {
	sf := d.Get("static_fields").(map[string]interface{})
	staticFields := make(map[string]string, len(sf))
	for k, v := range sf {
		staticFields[k] = v.(string)
	}

	data := &graylog.InputData{
		Title:         d.Get("title").(string),
		Type:          d.Get("type").(string),
		Global:        d.Get("global").(bool),
		Node:          d.Get("node").(string),
		ID:            d.Id(),
		CreatorUserID: d.Get("creator_user_id").(string),
		CreatedAt:     d.Get("created_at").(string),
		Attrs:         d.Get("attributes").([]interface{})[0].(map[string]interface{}),
		StaticFields:  staticFields,
	}
	input := &graylog.Input{}
	if err := data.ToInput(input); err != nil {
		return nil, err
	}
	return input, nil
}

func resourceInputCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	input, err := newInput(d)
	if err != nil {
		return err
	}

	if _, err := cl.CreateInput(ctx, input); err != nil {
		return err
	}
	d.SetId(input.ID)
	return nil
}

func resourceInputRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	input, ei, err := cl.GetInput(ctx, d.Id())
	if err != nil {
		return handleGetResourceError(d, ei, err)
	}
	if input.Attrs != nil {
		b, err := json.Marshal(input.Attrs)
		if err != nil {
			return err
		}
		dest := map[string]interface{}{}
		if err := json.Unmarshal(b, &dest); err != nil {
			return err
		}
		if err := d.Set("attributes", []map[string]interface{}{dest}); err != nil {
			return err
		}
	}
	if err := setStrToRD(d, "title", input.Title); err != nil {
		return err
	}
	if err := setStrToRD(d, "type", input.Type()); err != nil {
		return err
	}
	if err := setStrToRD(d, "node", input.Node); err != nil {
		return err
	}
	if err := setBoolToRD(d, "global", input.Global); err != nil {
		return err
	}
	if err := setStrToRD(d, "creator_user_id", input.CreatorUserID); err != nil {
		return err
	}
	if err := d.Set("static_fields", input.StaticFields); err != nil {
		return err
	}
	return setStrToRD(d, "created_at", input.CreatedAt)
}

func resourceInputUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	input, err := newInput(d)
	if err != nil {
		return err
	}
	if _, _, err := cl.UpdateInput(ctx, input.NewUpdateParams()); err != nil {
		return err
	}
	return nil
}

func resourceInputDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteInput(ctx, d.Id()); err != nil {
		return err
	}
	return nil
}
