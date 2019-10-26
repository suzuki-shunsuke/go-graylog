package graylog

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/suzuki-shunsuke/go-jsoneq/jsoneq"

	"github.com/suzuki-shunsuke/go-graylog"
)

func resourceExtractor() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtractorCreate,
		Read:   resourceExtractorRead,
		Update: resourceExtractorUpdate,
		Delete: resourceExtractorDelete,

		Importer: &schema.ResourceImporter{
			State: genImport("input_id", "extractor_id"),
		},

		Schema: map[string]*schema.Schema{
			"input_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cursor_strategy": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_field": {
				Type:     schema.TypeString,
				Required: true,
			},
			"condition_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"order": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"condition_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_field": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"general_string_extractor_config": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"general_int_extractor_config": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"general_float_extractor_config": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
			},
			"general_bool_extractor_config": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},

			"json_type_extractor_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"list_separator": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"kv_separator": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"key_prefix": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"key_separator": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"replace_key_whitespace": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"key_whitespace_replacement": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"grok_type_extractor_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"grok_pattern": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"regex_type_extractor_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"regex_value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"converters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"config": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"date_format": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"time_zone": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"locale": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func newExtractorConfig(d *schema.ResourceData, t string) interface{} {
	switch t {
	case "json":
		a := d.Get("json_type_extractor_config").([]interface{})[0].(map[string]interface{})
		return &graylog.ExtractorTypeJSONConfig{
			ListSeparator:            a["list_separator"].(string),
			KVSeparator:              a["kv_separator"].(string),
			KeyPrefix:                a["key_prefix"].(string),
			KeySeparator:             a["key_separator"].(string),
			ReplaceKeyWhitespace:     a["replace_key_whitespace"].(bool),
			KeyWhitespaceReplacement: a["key_whitespace_replacement"].(string),
		}
	case "grok":
		a := d.Get("grok_type_extractor_config").([]interface{})[0].(map[string]interface{})
		return &graylog.ExtractorTypeGrokConfig{
			GrokPattern: a["grok_pattern"].(string),
		}
	case "regex":
		a := d.Get("regex_type_extractor_config").([]interface{})[0].(map[string]interface{})
		return &graylog.ExtractorTypeRegexConfig{
			RegexValue: a["regex_value"].(string),
		}
	default:
		cfg := map[string]interface{}{}
		for _, k := range []string{"bool", "int", "string", "float"} {
			if c := d.Get(fmt.Sprintf("general_%s_extractor_config", k)); c != nil {
				for k, v := range c.(map[string]interface{}) {
					cfg[k] = v
				}
			}
		}
		return cfg
	}
}

func newExtractor(d *schema.ResourceData) (*graylog.Extractor, string, error) {
	t := d.Get("type").(string)
	cfg := newExtractorConfig(d, t)
	list := d.Get("converters").([]interface{})
	converters := make([]graylog.ExtractorConverter, len(list))
	for i, a := range list {
		b := a.(map[string]interface{})
		c := b["config"].([]interface{})
		cfg := &graylog.ExtractorConverterConfig{}
		if len(c) > 0 {
			if d, ok := c[0].(map[string]interface{}); ok {
				cfg.DateFormat = d["date_format"].(string)
				cfg.TimeZone = d["time_zone"].(string)
				cfg.Locale = d["locale"].(string)
			}
		}
		converters[i] = graylog.ExtractorConverter{
			Type:   b["type"].(string),
			Config: cfg,
		}
	}
	return &graylog.Extractor{
		Title:           d.Get("title").(string),
		Type:            d.Get("type").(string),
		Converters:      converters,
		Order:           d.Get("order").(int),
		CursorStrategy:  d.Get("cursor_strategy").(string),
		SourceField:     d.Get("source_field").(string),
		TargetField:     d.Get("target_field").(string),
		ConditionType:   d.Get("condition_type").(string),
		ConditionValue:  d.Get("condition_value").(string),
		ExtractorConfig: cfg,
		ID:              d.Id(),
	}, d.Get("input_id").(string), nil
}

func resourceExtractorCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	extractor, inputID, err := newExtractor(d)
	if err != nil {
		return err
	}

	if _, err := cl.CreateExtractor(ctx, inputID, extractor); err != nil {
		return err
	}
	d.SetId(extractor.ID)
	return nil
}

func resourceExtractorRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	extractor, ei, err := cl.GetExtractor(ctx, d.Get("input_id").(string), d.Id())
	if err != nil {
		return handleGetResourceError(d, ei, err)
	}
	if err := setStrToRD(d, "title", extractor.Title); err != nil {
		return err
	}
	if err := setStrToRD(d, "type", extractor.Type); err != nil {
		return err
	}
	if err := setStrToRD(d, "source_field", extractor.SourceField); err != nil {
		return err
	}
	if err := setStrToRD(d, "target_field", extractor.TargetField); err != nil {
		return err
	}
	if err := setStrToRD(d, "condition_type", extractor.ConditionType); err != nil {
		return err
	}
	if err := setStrToRD(d, "cursor_strategy", extractor.CursorStrategy); err != nil {
		return err
	}
	if err := setStrToRD(d, "condition_value", extractor.ConditionValue); err != nil {
		return err
	}
	a, err := jsoneq.Convert(extractor.ExtractorConfig)
	if err != nil {
		return errors.Wrap(err, "failed to convert extractor_config")
	}
	b, ok := a.(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed to convert extractor_config to map[string]interface{}")
	}
	switch extractor.Type {
	case "json":
		if err := d.Set("json_type_extractor_config", []map[string]interface{}{b}); err != nil {
			return err
		}
	case "grok":
		if err := d.Set("grok_type_extractor_config", []map[string]interface{}{b}); err != nil {
			return err
		}
	case "regex":
		if err := d.Set("regex_type_extractor_config", []map[string]interface{}{b}); err != nil {
			return err
		}
	default:
		intMap := map[string]int{}
		floatMap := map[string]float64{}
		strMap := map[string]string{}
		boolMap := map[string]bool{}
		for k, v := range b {
			switch a := v.(type) {
			case int:
				intMap[k] = a
			case float64:
				floatMap[k] = a
			case string:
				strMap[k] = a
			case bool:
				boolMap[k] = a
			default:
				return fmt.Errorf("%s is invalid type", k)
			}
		}
		if err := d.Set("general_bool_extractor_config", boolMap); err != nil {
			return err
		}
		if err := d.Set("general_int_extractor_config", intMap); err != nil {
			return err
		}
		if err := d.Set("general_float_extractor_config", floatMap); err != nil {
			return err
		}
		if err := d.Set("general_string_extractor_config", strMap); err != nil {
			return err
		}
	}
	converters := make([]map[string]interface{}, len(extractor.Converters))
	for i, converter := range extractor.Converters {
		converters[i] = map[string]interface{}{
			"type": converter.Type,
			"config": []map[string]interface{}{
				{
					"date_format": converter.Config.DateFormat,
					"time_zone":   converter.Config.TimeZone,
					"locale":      converter.Config.Locale,
				},
			},
		}
	}
	if err := d.Set("converters", converters); err != nil {
		return err
	}
	return setIntToRD(d, "order", extractor.Order)
}

func resourceExtractorUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	extractor, inputID, err := newExtractor(d)
	if err != nil {
		return err
	}
	if _, err := cl.UpdateExtractor(ctx, inputID, extractor); err != nil {
		return err
	}
	return nil
}

func resourceExtractorDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteExtractor(ctx, d.Get("input_id").(string), d.Id()); err != nil {
		return err
	}
	return nil
}
