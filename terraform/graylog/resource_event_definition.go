package graylog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-jsoneq/jsoneq"

	"github.com/suzuki-shunsuke/go-graylog/v8"
)

func resourceEventDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventDefinitionCreate,
		Read:   resourceEventDefinitionRead,
		Update: resourceEventDefinitionUpdate,
		Delete: resourceEventDefinitionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
				ValidateFunc: wrapValidateFunc(func(v interface{}, k string) error {
					priority := v.(int)
					if priority < 1 || priority > 3 {
						return errors.New("'priority' should be either 1, 2, and 3")
					}
					return nil
				}),
			},
			"config": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: schemaDiffSuppressJSONString,
				ValidateFunc:     validateFuncEventDefinitionConfig,
			},
			"notification_settings": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"grace_period_ms": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"backlog_size": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"alert": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"field_spec": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: schemaDiffSuppressJSONString,
				ValidateFunc:     wrapValidateFunc(validateFuncEventDefinitionFieldSpec),
			},
			"notifications": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notification_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			// "key_spec": {
			// 	Type:     schema.TypeString,
			// 	Required: true,
			// },
			// optional
			//	"storage": {
			//		Type:     schema.TypeString,
			//		Optional: true,
			//	},
		},
	}
}

func validateFuncEventDefinitionFieldSpec(v interface{}, k string) error {
	a := strings.TrimSpace(v.(string))
	if len(a) == 0 {
		return nil
	}
	spec := map[string]graylog.EventDefinitionFieldSpec{}
	if err := json.Unmarshal([]byte(a), &spec); err != nil {
		return fmt.Errorf(
			"failed to parse the 'field_spec'. 'field_spec' must be a JSON string: %w", err)
	}
	return nil
}

func getFieldSpec(d *schema.ResourceData) (map[string]graylog.EventDefinitionFieldSpec, error) {
	s := strings.TrimSpace(d.Get("field_spec").(string))
	if len(s) == 0 {
		return nil, nil
	}
	spec := map[string]graylog.EventDefinitionFieldSpec{}
	if err := json.Unmarshal([]byte(s), &spec); err != nil {
		return nil, fmt.Errorf("failed to parse the 'field_spec'. 'field_spec' must be a JSON string: %w", err)
	}
	return spec, nil
}

func validateFuncEventDefinitionConfig(v interface{}, k string) (s []string, es []error) {
	c, err := jsoneq.ConvertByte([]byte(v.(string)))
	if err != nil {
		es = append(es, fmt.Errorf("failed to parse the 'config'. 'config' must be a JSON string: %w", err))
		return
	}
	cfg, ok := c.(map[string]interface{})
	if !ok {
		es = append(es, errors.New("'config' should be a JSON string which represents object"))
		return
	}
	t, ok := cfg["type"]
	if !ok {
		es = append(es, errors.New("'type' of 'config' is required"))
		return
	}
	switch t {
	case "aggregation-v1":
		requiredKeys := []string{"query", "streams", "search_within_ms", "execute_every_ms"}
		for _, key := range requiredKeys {
			if _, ok := cfg[key]; !ok {
				es = append(es, errors.New("'"+key+"' of 'config' is required"))
			}
		}
	}
	return
}

func getDefinitionCfg(d *schema.ResourceData) (map[string]interface{}, error) {
	cfgS := d.Get("config").(string)
	c, err := jsoneq.ConvertByte([]byte(cfgS))
	if err != nil {
		return nil, fmt.Errorf("failed to parse the 'config'. 'config' must be a JSON string '%s': %w", cfgS, err)
	}
	cfg, ok := c.(map[string]interface{})
	if !ok {
		return nil, errors.New("'config' should be a JSON string which represents object")
	}
	t, ok := cfg["type"]
	if !ok {
		return nil, errors.New("'type' of 'config' is required")
	}
	switch t {
	case "aggregation-v1":
		if _, ok := cfg["series"]; !ok {
			cfg["series"] = []interface{}{}
		}
		if _, ok := cfg["group_by"]; !ok {
			cfg["group_by"] = []interface{}{}
		}
		requiredKeys := []string{"query", "streams", "search_within_ms", "execute_every_ms"}
		for _, key := range requiredKeys {
			if _, ok := cfg[key]; !ok {
				return nil, errors.New("'" + key + "' of 'config' is required")
			}
		}
	}
	return cfg, nil
}

func getDefinitionSettings(d *schema.ResourceData) graylog.EventDefinitionNotificationSettings {
	s := d.Get("notification_settings").([]interface{})
	if len(s) == 0 {
		return graylog.EventDefinitionNotificationSettings{}
	}
	settings := s[0].(map[string]interface{})
	gracePeriodMS := 0
	if a, ok := settings["grace_period_ms"]; ok {
		gracePeriodMS = a.(int)
	}
	backlogSize := 0
	if a, ok := settings["backlog_size"]; ok {
		backlogSize = a.(int)
	}
	return graylog.EventDefinitionNotificationSettings{
		GracePeriodMS: gracePeriodMS,
		BacklogSize:   backlogSize,
	}
}

func newEventDefinition(d *schema.ResourceData) (*graylog.EventDefinition, error) {
	cfg, err := getDefinitionCfg(d)
	if err != nil {
		return nil, err
	}

	fieldSpec, err := getFieldSpec(d)
	if err != nil {
		return nil, err
	}

	a := d.Get("notifications").([]interface{})
	notifs := make([]graylog.EventDefinitionNotification, len(a))
	for i, b := range a {
		c := b.(map[string]interface{})
		notifs[i] = graylog.EventDefinitionNotification{
			NotificationID: c["notification_id"].(string),
		}
	}

	return &graylog.EventDefinition{
		ID:                   d.Id(),
		Title:                d.Get("title").(string),
		Description:          d.Get("description").(string),
		Priority:             d.Get("priority").(int),
		Alert:                d.Get("alert").(bool),
		NotificationSettings: getDefinitionSettings(d),
		Config:               cfg,
		FieldSpec:            fieldSpec,
		Notifications:        notifs,
	}, nil
}

func resourceEventDefinitionCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	notif, err := newEventDefinition(d)
	if err != nil {
		return err
	}
	if _, err := cl.CreateEventDefinition(ctx, notif); err != nil {
		return err
	}
	d.SetId(notif.ID)
	return nil
}

func resourceEventDefinitionRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	notif, ei, err := cl.GetEventDefinition(ctx, d.Id())
	if err != nil {
		return handleGetResourceError(d, ei, err)
	}
	if err := setStrToRD(d, "title", notif.Title); err != nil {
		return err
	}
	if err := setStrToRD(d, "description", notif.Description); err != nil {
		return err
	}
	if err := setIntToRD(d, "priority", notif.Priority); err != nil {
		return err
	}
	if err := setBoolToRD(d, "alert", notif.Alert); err != nil {
		return err
	}
	if err := d.Set("notification_settings", []map[string]interface{}{
		{
			"grace_period_ms": notif.NotificationSettings.GracePeriodMS,
			"backlog_size":    notif.NotificationSettings.BacklogSize,
		},
	}); err != nil {
		return err
	}
	b, err := json.Marshal(notif.FieldSpec)
	if err != nil {
		return err
	}
	if err := setStrToRD(d, "field_spec", string(b)); err != nil {
		return err
	}
	b, err = json.Marshal(notif.Config)
	if err != nil {
		return err
	}
	return setStrToRD(d, "config", string(b))
}

func resourceEventDefinitionUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	notif, err := newEventDefinition(d)
	if err != nil {
		return err
	}
	_, err = cl.UpdateEventDefinition(ctx, notif)
	return err
}

func resourceEventDefinitionDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteEventDefinition(ctx, d.Id()); err != nil {
		return err
	}
	return nil
}
