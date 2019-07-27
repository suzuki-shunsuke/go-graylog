package graylog

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/suzuki-shunsuke/go-graylog"
)

func resourceAlertCondition() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlertConditionCreate,
		Read:   resourceAlertConditionRead,
		Update: resourceAlertConditionUpdate,
		Delete: resourceAlertConditionDelete,

		Importer: &schema.ResourceImporter{
			State: genImport("stream_id", "alert_condition_id"),
		},

		SchemaVersion: 1,
		MigrateState:  alertConditionStateMigrateFunc,

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
				ForceNew: true,
			},
			"field_content_value_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"grace": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"backlog": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"query": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"repeat_notifications": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"field_value_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"threshold_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"grace": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"backlog": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"query": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"threshold": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"time": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"repeat_notifications": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"message_count_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threshold_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"grace": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"backlog": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"query": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"threshold": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"time": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"repeat_notifications": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"general_string_parameters": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"general_int_parameters": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"general_float_parameters": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
			},
			"general_bool_parameters": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},

			"in_grace": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func alertConditionStateMigrateFunc(
	v int, is *terraform.InstanceState, meta interface{},
) (*terraform.InstanceState, error) {
	if is.Empty() {
		return is, nil
	}

	switch v {
	case 0:
		return migrateAlertConditionStateV0toV1(is)
	default:
		return is, fmt.Errorf("unexpected schema version: %d", v)
	}
}

func migrateAlertConditionStateV0toV1(is *terraform.InstanceState) (*terraform.InstanceState, error) {
	// parameters.backlog ->
	prefix := ""
	switch is.Attributes["type"] {
	case "field_content_value":
		is.Attributes["field_content_value_parameters.#"] = "1"
		delete(is.Attributes, "parameters.%")
		prefix = "field_content_value_parameters.0"
	case "field_value":
		is.Attributes["field_value_parameters.#"] = "1"
		delete(is.Attributes, "parameters.%")
		prefix = "field_value_parameters.0"
	case "message_count":
		is.Attributes["message_count_parameters.#"] = "1"
		delete(is.Attributes, "parameters.%")
		prefix = "message_count_parameters.0"
	default:
		prefix = "general_string_parameters"
	}

	for k, v := range is.Attributes {
		if !strings.HasPrefix(k, "parameters.") {
			continue
		}
		is.Attributes[strings.Replace(k, "parameters", prefix, 1)] = v
	}
	return is, nil
}

func newAlertCondition(d *schema.ResourceData) (*graylog.AlertCondition, error) {
	cond := graylog.AlertCondition{
		Title:   d.Get("title").(string),
		InGrace: d.Get("in_grace").(bool),
		ID:      d.Id(),
	}
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
		ok bool
	)

	switch d.Get("type").(string) {
	case "field_content_value":
		p := graylog.FieldContentAlertConditionParameters{}
		prms := d.Get("field_content_value_parameters")
		if prms == nil {
			return nil, fmt.Errorf("field_content_value is required")
		}
		for k, v := range prms.([]interface{})[0].(map[string]interface{}) {
			switch k {
			case graceKey:
				if p.Grace, ok = v.(int); !ok {
					return nil, fmt.Errorf("%s must be int: %v", k, v)
				}
			case backlogKey:
				if p.Backlog, ok = v.(int); !ok {
					return nil, fmt.Errorf("%s must be int: %v", k, v)
				}
			case repeatNotificationsKey:
				if p.RepeatNotifications, ok = v.(bool); !ok {
					return nil, fmt.Errorf("%s must be bool: %v", k, v)
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
		prms := d.Get("field_value_parameters")
		if prms == nil {
			return nil, fmt.Errorf("field_value_parameters is required")
		}
		for k, v := range prms.([]interface{})[0].(map[string]interface{}) {
			switch k {
			case repeatNotificationsKey:
				if p.RepeatNotifications, ok = v.(bool); !ok {
					return nil, fmt.Errorf("%s must be bool: %v", k, v)
				}
			case graceKey:
				if p.Grace, ok = v.(int); !ok {
					return nil, fmt.Errorf("%s must be int: %v", k, v)
				}
			case backlogKey:
				if p.Backlog, ok = v.(int); !ok {
					return nil, fmt.Errorf("%s must be int: %v", k, v)
				}
			case thresholdKey:
				if p.Threshold, ok = v.(int); !ok {
					return nil, fmt.Errorf("%s must be int: %v", k, v)
				}
			case timeKey:
				if p.Time, ok = v.(int); !ok {
					return nil, fmt.Errorf("%s must be int: %v", k, v)
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
				return nil, fmt.Errorf("invalid attribute for alert condition type `field_value`: `%s`", k)
			}
		}
		cond.Parameters = p
		return &cond, nil
	case "message_count":
		p := graylog.MessageCountAlertConditionParameters{}
		prms := d.Get("message_count_parameters")
		if prms == nil {
			return nil, fmt.Errorf("message_count_parameters is required")
		}
		for k, v := range prms.([]interface{})[0].(map[string]interface{}) {
			switch k {
			case repeatNotificationsKey:
				if p.RepeatNotifications, ok = v.(bool); !ok {
					return nil, fmt.Errorf("%s must be bool: %v", k, v)
				}
			case graceKey:
				if p.Grace, ok = v.(int); !ok {
					return nil, fmt.Errorf("%s must be int: %v", k, v)
				}
			case backlogKey:
				if p.Backlog, ok = v.(int); !ok {
					return nil, fmt.Errorf("%s must be int: %v", k, v)
				}
			case thresholdKey:
				if p.Threshold, ok = v.(int); !ok {
					return nil, fmt.Errorf("%s must be int: %v", k, v)
				}
			case timeKey:
				if p.Time, ok = v.(int); !ok {
					return nil, fmt.Errorf("%s must be int: %v", k, v)
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
				return nil, fmt.Errorf("invalid attribute for alert condition type `message_count`: `%s`", k)
			}
		}
		cond.Parameters = p
		return &cond, nil
	}

	gap := graylog.GeneralAlertConditionParameters{
		Type:       d.Get("type").(string),
		Parameters: map[string]interface{}{},
	}
	for _, k := range []string{"bool", "int", "string", "float"} {
		if prms := d.Get(fmt.Sprintf("general_%s_parameters", k)); prms != nil {
			for k, v := range prms.(map[string]interface{}) {
				gap.Parameters[k] = v
			}
		}
	}
	cond.Parameters = &gap
	return &cond, nil
}

func resourceAlertConditionCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	cond, err := newAlertCondition(d)
	if err != nil {
		return err
	}

	if _, err := cl.CreateStreamAlertCondition(ctx, d.Get("stream_id").(string), cond); err != nil {
		return err
	}
	d.SetId(cond.ID)
	return nil
}

func resourceAlertConditionRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	streamID := d.Get("stream_id").(string)
	cond, _, err := cl.GetStreamAlertCondition(ctx, streamID, d.Id())
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
	if cond.Parameters == nil {
		return nil
	}
	switch cond.Type() {
	case "field_content_value":
		prms, ok := cond.Parameters.(graylog.FieldContentAlertConditionParameters)
		if !ok {
			return fmt.Errorf("parameters is invalid type as field_content_value: %v", cond.Parameters)
		}
		return d.Set(
			"field_content_value_parameters",
			[]map[string]interface{}{{
				"grace":   prms.Grace,
				"backlog": prms.Backlog,
				"field":   prms.Field,
				"value":   prms.Value,
				"query":   prms.Query,
			}})
	case "field_value":
		prms, ok := cond.Parameters.(graylog.FieldAggregationAlertConditionParameters)
		if !ok {
			return fmt.Errorf("parameters is invalid type as field_value")
		}
		return d.Set(
			"field_value_parameters",
			[]map[string]interface{}{{
				"grace":                prms.Grace,
				"backlog":              prms.Backlog,
				"threshold":            prms.Threshold,
				"time":                 prms.Time,
				"repeat_notifications": prms.RepeatNotifications,
				"field":                prms.Field,
				"query":                prms.Query,
				"threshold_type":       prms.ThresholdType,
				"type":                 prms.Type,
			}})
	case "message_count":
		prms, ok := cond.Parameters.(graylog.MessageCountAlertConditionParameters)
		if !ok {
			return fmt.Errorf("parameters is invalid type as message_count")
		}
		return d.Set(
			"message_count_parameters",
			[]map[string]interface{}{{
				"grace":                prms.Grace,
				"backlog":              prms.Backlog,
				"threshold":            prms.Threshold,
				"time":                 prms.Time,
				"repeat_notifications": prms.RepeatNotifications,
				"query":                prms.Query,
				"threshold_type":       prms.ThresholdType,
			}})
	}
	prms, ok := cond.Parameters.(graylog.GeneralAlertConditionParameters)
	if !ok {
		return fmt.Errorf("parameters is invalid type as GeneralAlertConditionParameters")
	}
	intM := map[string]interface{}{}
	strM := map[string]interface{}{}
	floatM := map[string]interface{}{}
	boolM := map[string]interface{}{}
	for k, v := range prms.Parameters {
		switch v.(type) {
		case int:
			intM[k] = v
		case bool:
			boolM[k] = v
		case float64:
			floatM[k] = v
		case float32:
			floatM[k] = v
		case string:
			strM[k] = v
		default:
			return fmt.Errorf("%s is invalid type", k)
		}
	}
	if len(intM) != 0 {
		if err := d.Set("general_int_parameters", intM); err != nil {
			return err
		}
	}
	if len(strM) != 0 {
		if err := d.Set("general_string_parameters", strM); err != nil {
			return err
		}
	}
	if len(floatM) != 0 {
		if err := d.Set("general_float_parameters", floatM); err != nil {
			return err
		}
	}
	if len(boolM) != 0 {
		if err := d.Set("general_bool_parameters", boolM); err != nil {
			return err
		}
	}
	return nil
}

func resourceAlertConditionUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	cond, err := newAlertCondition(d)
	if err != nil {
		return err
	}
	if _, err := cl.UpdateStreamAlertCondition(ctx, d.Get("stream_id").(string), cond); err != nil {
		return err
	}
	return nil
}

func resourceAlertConditionDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteStreamAlertCondition(ctx, d.Get("stream_id").(string), d.Id()); err != nil {
		return err
	}
	return nil
}
