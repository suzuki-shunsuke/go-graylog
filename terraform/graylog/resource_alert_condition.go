package graylog

import (
	"encoding/json"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/util"
)

func resourceAlertCondition() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlertConditionCreate,
		Read:   resourceAlertConditionRead,
		Update: resourceAlertConditionUpdate,
		Delete: resourceAlertConditionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"stream_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parameters": {
				// we can't resrict attributes of third party alert condition plugin, so parameters is schema.TypeMap .
				Type:     schema.TypeMap,
				Required: true,
			},

			"in_grace": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func newAlertCondition(d *schema.ResourceData) (*graylog.AlertCondition, error) {
	cond := graylog.AlertCondition{
		Title:   d.Get("title").(string),
		InGrace: d.Get("in_grace").(bool),
		ID:      d.Id(),
	}
	prms := d.Get("parameters").(map[string]interface{})
	switch d.Get("type").(string) {
	case "field_content_value":
		p := graylog.FieldContentAlertConditionParameters{}
		if err := util.MSDecode(prms, &p); err != nil {
			return nil, err
		}
		cond.Parameters = p
		return &cond, nil
	case "field_value":
		p := graylog.FieldAggregationAlertConditionParameters{}
		if err := util.MSDecode(prms, &p); err != nil {
			return nil, err
		}
		cond.Parameters = p
		return &cond, nil
	case "message_count":
		p := graylog.MessageCountAlertConditionParameters{}
		if err := util.MSDecode(prms, &p); err != nil {
			return nil, err
		}
		cond.Parameters = p
		return &cond, nil
	}
	return &cond, nil
}

func resourceAlertConditionCreate(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	cond, err := newAlertCondition(d)
	if err != nil {
		return err
	}

	if _, err := cl.CreateStreamAlertCondition(d.Get("stream_id").(string), cond); err != nil {
		return err
	}
	d.SetId(cond.ID)
	return nil
}

func resourceAlertConditionRead(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	streamID := d.Get("stream_id").(string)
	cond, _, err := cl.GetStreamAlertCondition(streamID, d.Id())
	if err != nil {
		return err
	}
	setStrToRD(d, "type", cond.Type())
	setStrToRD(d, "title", cond.Title)
	setStrToRD(d, "stream_id", streamID)
	setBoolToRD(d, "in_grace", cond.InGrace)
	if cond.Parameters != nil {
		b, err := json.Marshal(cond.Parameters)
		if err != nil {
			return err
		}
		dest := map[string]interface{}{}
		if err := json.Unmarshal(b, &dest); err != nil {
			return err
		}
		d.Set("parameters", []map[string]interface{}{dest})
	}
	return nil
}

func resourceAlertConditionUpdate(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	cond, err := newAlertCondition(d)
	if err != nil {
		return err
	}
	if _, err := cl.UpdateStreamAlertCondition(d.Get("stream_id").(string), cond); err != nil {
		return err
	}
	return nil
}

func resourceAlertConditionDelete(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteStreamAlertCondition(d.Get("stream_id").(string), d.Id()); err != nil {
		return err
	}
	return nil
}
