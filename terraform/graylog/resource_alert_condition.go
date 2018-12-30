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

	switch d.Get("type").(string) {
	case "field_content_value":
		p := graylog.FieldContentAlertConditionParameters{}
		for k, v := range prms {
			switch k {
			case graceKey:
				grace, err := convIntfStrToInt(v)
				if err != nil {
					return nil, errors.Wrap(err, "grace must be int")
				}
				p.Grace = grace
			case backlogKey:
				backlog, err := convIntfStrToInt(v)
				if err != nil {
					return nil, errors.Wrap(err, "backlog must be int")
				}
				p.Backlog = backlog
			case repeatNotificationsKey:
				rn, err := convIntfStrToBool(v)
				if err != nil {
					return nil, errors.Wrap(err, "repeat_notifications must be bool")
				}
				p.RepeatNotifications = rn
			case fieldKey:
				field, ok := v.(string)
				if !ok {
					return nil, fmt.Errorf("field must be string")
				}
				p.Field = field
			case valueKey:
				value, ok := v.(string)
				if !ok {
					return nil, fmt.Errorf("value must be string")
				}
				p.Value = value
			case queryKey:
				query, ok := v.(string)
				if !ok {
					return nil, fmt.Errorf("query must be string")
				}
				p.Query = query
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
				rn, err := convIntfStrToBool(v)
				if err != nil {
					return nil, errors.Wrap(err, "repeat_notifications must be bool")
				}
				p.RepeatNotifications = rn
			case graceKey:
				grace, err := convIntfStrToInt(v)
				if err != nil {
					return nil, errors.Wrap(err, "grace must be int")
				}
				p.Grace = grace
			case backlogKey:
				backlog, err := convIntfStrToInt(v)
				if err != nil {
					return nil, errors.Wrap(err, "backlog must be int")
				}
				p.Backlog = backlog
			case fieldKey:
				field, ok := v.(string)
				if !ok {
					return nil, fmt.Errorf("field must be string")
				}
				p.Field = field
			case queryKey:
				query, ok := v.(string)
				if !ok {
					return nil, fmt.Errorf("query must be string")
				}
				p.Query = query
			case "type":
				t, ok := v.(string)
				if !ok {
					return nil, fmt.Errorf("type must be string")
				}
				p.Type = t
			case thresholdKey:
				threshold, err := convIntfStrToInt(v)
				if err != nil {
					return nil, errors.Wrap(err, "threshold must be int")
				}
				p.Threshold = threshold
			case timeKey:
				time, err := convIntfStrToInt(v)
				if err != nil {
					return nil, errors.Wrap(err, "time must be int")
				}
				p.Time = time
			case thresholdTypeKey:
				thresholdType, ok := v.(string)
				if !ok {
					return nil, fmt.Errorf("threshold_type must be string")
				}
				p.ThresholdType = thresholdType
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
				rn, err := convIntfStrToBool(v)
				if err != nil {
					return nil, errors.Wrap(err, "repeat_notifications must be bool")
				}
				p.RepeatNotifications = rn
			case graceKey:
				grace, err := convIntfStrToInt(v)
				if err != nil {
					return nil, errors.Wrap(err, "grace must be int")
				}
				p.Grace = grace
			case backlogKey:
				backlog, err := convIntfStrToInt(v)
				if err != nil {
					return nil, errors.Wrap(err, "backlog must be int")
				}
				p.Backlog = backlog
			case queryKey:
				query, ok := v.(string)
				if !ok {
					return nil, fmt.Errorf("query must be string")
				}
				p.Query = query
			case thresholdKey:
				threshold, err := convIntfStrToInt(v)
				if err != nil {
					return nil, errors.Wrap(err, "threshold must be int")
				}
				p.Threshold = threshold
			case timeKey:
				time, err := convIntfStrToInt(v)
				if err != nil {
					return nil, errors.Wrap(err, "time must be int")
				}
				p.Time = time
			case thresholdTypeKey:
				thresholdType, ok := v.(string)
				if !ok {
					return nil, fmt.Errorf("threshold_type must be string")
				}
				p.ThresholdType = thresholdType
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
