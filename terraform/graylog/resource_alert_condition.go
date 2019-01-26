package graylog

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"

	"github.com/suzuki-shunsuke/go-graylog"
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

	graceKey := "grace"
	backlogKey := "backlog"
	repeatNotificationsKey := "repeat_notifications"
	fieldKey := "field"
	valueKey := "value"
	queryKey := "query"
	thresholdKey := "threshold"
	thresholdTypeKey := "threshold_type"
	timeKey := "time"

	var (
		err error
		ok  bool
	)

	switch d.Get("type").(string) {
	case "field_content_value":
		p := graylog.FieldContentAlertConditionParameters{}
		for k, v := range prms {
			switch k {
			case graceKey:
				if p.Grace, err = convIntfStrToInt(v); err != nil {
					return nil, errors.Wrapf(err, "%s must be int", k)
				}
			case backlogKey:
				if p.Backlog, err = convIntfStrToInt(v); err != nil {
					return nil, errors.Wrapf(err, "%s must be int", k)
				}
			case repeatNotificationsKey:
				if p.RepeatNotifications, err = convIntfStrToBool(v); err != nil {
					return nil, errors.Wrapf(err, "%s must be bool", k)
				}
			case fieldKey:
				if p.Field, ok = v.(string); !ok {
					return nil, fmt.Errorf("%s must be string", k)
				}
			case valueKey:
				if p.Value, ok = v.(string); !ok {
					return nil, fmt.Errorf("%s must be string", k)
				}
			case queryKey:
				if p.Query, ok = v.(string); !ok {
					return nil, fmt.Errorf("%s must be string", k)
				}
			default:
				return nil, fmt.Errorf("invalid attribute for alert condition type `field_content_value`: `%s`", k)
			}
		}
		cond.Parameters = p
		return &cond, nil
	case "field_value":
		p := graylog.FieldAggregationAlertConditionParameters{}
		for k, v := range prms {
			switch k {
			case repeatNotificationsKey:
				if p.RepeatNotifications, err = convIntfStrToBool(v); err != nil {
					return nil, errors.Wrapf(err, "%s must be bool", k)
				}
			case graceKey:
				if p.Grace, err = convIntfStrToInt(v); err != nil {
					return nil, errors.Wrapf(err, "%s must be int", k)
				}
			case backlogKey:
				if p.Backlog, err = convIntfStrToInt(v); err != nil {
					return nil, errors.Wrapf(err, "%s must be int", k)
				}
			case thresholdKey:
				if p.Threshold, err = convIntfStrToInt(v); err != nil {
					return nil, errors.Wrapf(err, "%s must be int", k)
				}
			case timeKey:
				if p.Time, err = convIntfStrToInt(v); err != nil {
					return nil, errors.Wrapf(err, "%s must be int", k)
				}
			case fieldKey:
				if p.Field, ok = v.(string); !ok {
					return nil, fmt.Errorf("%s must be string", k)
				}
			case queryKey:
				if p.Query, ok = v.(string); !ok {
					return nil, fmt.Errorf("%s must be string", k)
				}
			case "type":
				if p.Type, ok = v.(string); !ok {
					return nil, fmt.Errorf("%s must be string", k)
				}
			case thresholdTypeKey:
				if p.ThresholdType, ok = v.(string); !ok {
					return nil, fmt.Errorf("%s must be string", k)
				}
			default:
				return nil, fmt.Errorf("invalid attribute for alert condition type `field_content_value`: `%s`", k)
			}
		}
		cond.Parameters = p
		return &cond, nil
	case "message_count":
		p := graylog.MessageCountAlertConditionParameters{}
		for k, v := range prms {
			switch k {
			case repeatNotificationsKey:
				if p.RepeatNotifications, err = convIntfStrToBool(v); err != nil {
					return nil, errors.Wrapf(err, "%s must be bool", k)
				}
			case graceKey:
				if p.Grace, err = convIntfStrToInt(v); err != nil {
					return nil, errors.Wrapf(err, "%s must be int", k)
				}
			case backlogKey:
				if p.Backlog, err = convIntfStrToInt(v); err != nil {
					return nil, errors.Wrapf(err, "%s must be int", k)
				}
			case thresholdKey:
				if p.Threshold, err = convIntfStrToInt(v); err != nil {
					return nil, errors.Wrapf(err, "%s must be int", k)
				}
			case timeKey:
				if p.Time, err = convIntfStrToInt(v); err != nil {
					return nil, errors.Wrapf(err, "%s must be int", k)
				}
			case queryKey:
				if p.Query, ok = v.(string); !ok {
					return nil, fmt.Errorf("%s must be string", k)
				}
			case thresholdTypeKey:
				if p.ThresholdType, ok = v.(string); !ok {
					return nil, fmt.Errorf("%s must be string", k)
				}
			default:
				return nil, fmt.Errorf("invalid attribute for alert condition type `field_content_value`: `%s`", k)
			}
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
	if err := setStrToRD(d, "type", cond.Type()); err != nil {
		return err
	}
	if err := setStrToRD(d, "title", cond.Title); err != nil {
		return err
	}
	if err := setStrToRD(d, "stream_id", streamID); err != nil {
		return err
	}
	if err := setBoolToRD(d, "in_grace", cond.InGrace); err != nil {
		return err
	}
	if cond.Parameters != nil {
		b, err := json.Marshal(cond.Parameters)
		if err != nil {
			return err
		}
		dest := map[string]interface{}{}
		if err := json.Unmarshal(b, &dest); err != nil {
			return err
		}
		for k, v := range dest {
			d.Set(fmt.Sprintf("parameters.%s", k), v)
		}
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
