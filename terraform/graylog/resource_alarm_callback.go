package graylog

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-set/v6"

	"github.com/suzuki-shunsuke/go-graylog/v9"
)

func resourceAlarmCallback() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlarmCallbackCreate,
		Read:   resourceAlarmCallbackRead,
		Update: resourceAlarmCallbackUpdate,
		Delete: resourceAlarmCallbackDelete,

		Importer: &schema.ResourceImporter{
			State: genImport("stream_id", "alarm_callback_id"),
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
				ForceNew: true,
			},

			"general_string_configuration": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"general_int_configuration": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"general_float_configuration": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
			},
			"general_bool_configuration": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},

			"email_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"sender": {
							Type:     schema.TypeString,
							Required: true,
						},
						"subject": {
							Type:     schema.TypeString,
							Required: true,
						},
						// Optional
						"body": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_receivers": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"email_receivers": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"http_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"slack_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"color": {
							Type:     schema.TypeString,
							Required: true,
						},
						"webhook_url": {
							Type:     schema.TypeString,
							Required: true,
						},
						"channel": {
							Type:     schema.TypeString,
							Required: true,
						},
						// Optional
						"icon_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"graylog2_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"icon_emoji": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"proxy_address": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"custom_message": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"backlog_items": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"link_names": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"notify_channel": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func newAlarmCallback(d *schema.ResourceData) (*graylog.AlarmCallback, error) {
	ac := graylog.AlarmCallback{
		Title:    d.Get("title").(string),
		StreamID: d.Get("stream_id").(string),
		ID:       d.Id(),
	}
	switch d.Get("type").(string) {
	case graylog.HTTPAlarmCallbackType:
		p := graylog.HTTPAlarmCallbackConfiguration{}
		hc := d.Get("http_configuration")
		if hc == nil {
			return nil, errors.New("http_configuration is required")
		}
		p.URL = hc.([]interface{})[0].(map[string]interface{})["url"].(string)
		ac.Configuration = &p
		return &ac, nil
	case graylog.EmailAlarmCallbackType:
		p := graylog.EmailAlarmCallbackConfiguration{}
		ec := d.Get("email_configuration")
		if ec == nil {
			return nil, errors.New("email_configuration is required")
		}
		emailCfg := ec.([]interface{})[0].(map[string]interface{})
		for k, v := range emailCfg {
			switch k {
			case "sender":
				p.Sender = v.(string)
			case "subject":
				p.Subject = v.(string)
			case "body":
				p.Body = v.(string)
			case "email_receivers":
				arr := set.NewStrSet()
				for _, a := range v.(*schema.Set).List() {
					arr.Add(a.(string))
				}
				p.EmailReceivers = arr
			case "user_receivers":
				arr := set.NewStrSet()
				for _, a := range v.(*schema.Set).List() {
					arr.Add(a.(string))
				}
				p.UserReceivers = arr
			default:
				return nil, fmt.Errorf("invalid attribute for alarm callback type `%s`: `%s`", graylog.EmailAlarmCallbackType, k)
			}
		}
		ac.Configuration = &p
		return &ac, nil
	case graylog.SlackAlarmCallbackType:
		p := graylog.SlackAlarmCallbackConfiguration{}
		sc := d.Get("slack_configuration")
		if sc == nil {
			return nil, errors.New("slack_configuration is required")
		}
		slackCfg := sc.([]interface{})[0].(map[string]interface{})
		for k, v := range slackCfg {
			switch k {
			case "color":
				p.Color = v.(string)
			case "webhook_url":
				p.WebhookURL = v.(string)
			case "channel":
				p.Channel = v.(string)
			case "icon_url":
				p.IconURL = v.(string)
			case "icon_emoji":
				p.IconEmoji = v.(string)
			case "graylog2_url":
				p.Graylog2URL = v.(string)
			case "user_name":
				p.UserName = v.(string)
			case "proxy_address":
				p.ProxyAddress = v.(string)
			case "custom_message":
				p.CustomMessage = v.(string)
			case "backlog_items":
				p.BacklogItems = v.(int)
			case "link_names":
				p.LinkNames = v.(bool)
			case "notify_channel":
				p.NotifyChannel = v.(bool)
			default:
				return nil, fmt.Errorf("invalid attribute for alarm callback type `%s`: `%s`", graylog.SlackAlarmCallbackType, k)
			}
		}
		ac.Configuration = &p
		return &ac, nil
	}
	gc := graylog.GeneralAlarmCallbackConfiguration{
		Type:          d.Get("type").(string),
		Configuration: map[string]interface{}{},
	}
	for _, k := range []string{"bool", "int", "string", "float"} {
		if cfg := d.Get(fmt.Sprintf("general_%s_configuration", k)); cfg != nil {
			for k, v := range cfg.(map[string]interface{}) {
				gc.Configuration[k] = v
			}
		}
	}
	ac.Configuration = &gc
	return &ac, nil
}

func resourceAlarmCallbackCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	ac, err := newAlarmCallback(d)
	if err != nil {
		return err
	}

	if _, err := cl.CreateStreamAlarmCallback(ctx, ac); err != nil {
		return err
	}
	d.SetId(ac.ID)
	return nil
}

func resourceAlarmCallbackRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	streamID := d.Get("stream_id").(string)
	ac, ei, err := cl.GetStreamAlarmCallback(ctx, streamID, d.Id())
	if err != nil {
		return handleGetResourceError(d, ei, err)
	}
	if err := setStrToRD(d, "type", ac.Type()); err != nil {
		return err
	}
	if err := setStrToRD(d, "title", ac.Title); err != nil {
		return err
	}
	if err := setStrToRD(d, "stream_id", streamID); err != nil {
		return err
	}
	if ac.Configuration != nil {
		switch ac.Type() {
		case graylog.HTTPAlarmCallbackType:
			cfg, ok := ac.Configuration.(*graylog.HTTPAlarmCallbackConfiguration)
			if !ok {
				return errors.New("configuration is invalid type")
			}
			return d.Set("http_configuration", []map[string]interface{}{
				{"url": cfg.URL},
			})
		case graylog.EmailAlarmCallbackType:
			cfg, ok := ac.Configuration.(*graylog.EmailAlarmCallbackConfiguration)
			if !ok {
				return errors.New("configuration is invalid type")
			}
			return d.Set("email_configuration", []map[string]interface{}{{
				"sender":          cfg.Sender,
				"subject":         cfg.Subject,
				"body":            cfg.Body,
				"user_receivers":  cfg.UserReceivers.ToList(),
				"email_receivers": cfg.EmailReceivers.ToList(),
			}})
		case graylog.SlackAlarmCallbackType:
			cfg, ok := ac.Configuration.(*graylog.SlackAlarmCallbackConfiguration)
			if !ok {
				return errors.New("configuration is invalid type")
			}
			return d.Set("slack_configuration", []map[string]interface{}{{
				"color":          cfg.Color,
				"webhook_url":    cfg.WebhookURL,
				"channel":        cfg.Channel,
				"icon_url":       cfg.IconURL,
				"graylog2_url":   cfg.Graylog2URL,
				"icon_emoji":     cfg.IconEmoji,
				"user_name":      cfg.UserName,
				"proxy_address":  cfg.ProxyAddress,
				"custom_message": cfg.CustomMessage,
				"backlog_items":  cfg.BacklogItems,
				"link_names":     cfg.LinkNames,
				"notify_channel": cfg.NotifyChannel,
			}})
		}
		cfg, ok := ac.Configuration.(*graylog.GeneralAlarmCallbackConfiguration)
		if !ok {
			return errors.New("configuration is invalid type")
		}
		intM := map[string]interface{}{}
		strM := map[string]interface{}{}
		floatM := map[string]interface{}{}
		boolM := map[string]interface{}{}
		for k, v := range cfg.Configuration {
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
				return errors.New(k + " is invalid type")
			}
		}
		if err := d.Set("general_int_configuration", intM); err != nil {
			return err
		}
		if err := d.Set("general_string_configuration", strM); err != nil {
			return err
		}
		if err := d.Set("general_float_configuration", floatM); err != nil {
			return err
		}
		if err := d.Set("general_bool_configuration", boolM); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func resourceAlarmCallbackUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	ac, err := newAlarmCallback(d)
	if err != nil {
		return err
	}
	if _, err := cl.UpdateStreamAlarmCallback(ctx, ac); err != nil {
		return err
	}
	return nil
}

func resourceAlarmCallbackDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteStreamAlarmCallback(ctx, d.Get("stream_id").(string), d.Id()); err != nil {
		return err
	}
	return nil
}
