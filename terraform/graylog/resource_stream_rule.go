package graylog

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-graylog"
)

func resourceStreamRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceStreamRuleCreate,
		Read:   resourceStreamRuleRead,
		Update: resourceStreamRuleUpdate,
		Delete: resourceStreamRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	rule, err := newStreamRule(d)
	if err != nil {
		return err
	}

	if _, err := cl.CreateStreamRule(rule); err != nil {
		return err
	}
	d.SetId(rule.ID)
	return nil
}

func resourceStreamRuleRead(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	rule, _, err := cl.GetStreamRule(d.Get("stream_id").(string), d.Id())
	if err != nil {
		return err
	}
	setStrToRD(d, "field", rule.Field)
	setStrToRD(d, "value", rule.Value)
	setStrToRD(d, "description", rule.Description)
	setStrToRD(d, "stream_id", rule.StreamID)
	setIntToRD(d, "type", rule.Type)
	setBoolToRD(d, "inverted", rule.Inverted)
	return nil
}

func resourceStreamRuleUpdate(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	rule, err := newStreamRule(d)
	if err != nil {
		return err
	}
	if _, err := cl.UpdateStreamRule(rule); err != nil {
		return err
	}
	return nil
}

func resourceStreamRuleDelete(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteStreamRule(d.Get("stream_id").(string), d.Id()); err != nil {
		return err
	}
	return nil
}
