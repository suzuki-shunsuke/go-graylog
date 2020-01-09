package graylog

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog/v9"
)

func resourceStreamRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceStreamRuleCreate,
		Read:   resourceStreamRuleRead,
		Update: resourceStreamRuleUpdate,
		Delete: resourceStreamRuleDelete,

		Importer: &schema.ResourceImporter{
			State: genImport("stream_id", "stream_rule_id"),
		},

		Schema: map[string]*schema.Schema{
			// Required
			"field": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"stream_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inverted": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func newStreamRule(d *schema.ResourceData) (*graylog.StreamRule, error) {
	return &graylog.StreamRule{
		StreamID:    d.Get("stream_id").(string),
		Field:       d.Get("field").(string),
		Value:       d.Get("value").(string),
		Description: d.Get("description").(string),
		Type:        d.Get("type").(int),
		Inverted:    d.Get("inverted").(bool),
		ID:          d.Id(),
	}, nil
}

func resourceStreamRuleCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	rule, err := newStreamRule(d)
	if err != nil {
		return err
	}

	if _, err := cl.CreateStreamRule(ctx, rule); err != nil {
		return err
	}
	d.SetId(rule.ID)
	return nil
}

func resourceStreamRuleRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	rule, ei, err := cl.GetStreamRule(ctx, d.Get("stream_id").(string), d.Id())
	if err != nil {
		return handleGetResourceError(d, ei, err)
	}
	if err := setStrToRD(d, "field", rule.Field); err != nil {
		return err
	}
	if err := setStrToRD(d, "value", rule.Value); err != nil {
		return err
	}
	if err := setStrToRD(d, "description", rule.Description); err != nil {
		return err
	}
	if err := setStrToRD(d, "stream_id", rule.StreamID); err != nil {
		return err
	}
	if err := setIntToRD(d, "type", rule.Type); err != nil {
		return err
	}
	return setBoolToRD(d, "inverted", rule.Inverted)
}

func resourceStreamRuleUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	rule, err := newStreamRule(d)
	if err != nil {
		return err
	}
	if _, err := cl.UpdateStreamRule(ctx, rule); err != nil {
		return err
	}
	return nil
}

func resourceStreamRuleDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteStreamRule(ctx, d.Get("stream_id").(string), d.Id()); err != nil {
		return err
	}
	return nil
}
