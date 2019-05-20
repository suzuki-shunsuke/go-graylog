package graylog

import (
	"github.com/hashicorp/terraform/helper/schema"

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
			// Required
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
			"target_field": {
				Type:     schema.TypeString,
				Required: true,
			},
			"condition_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			// converters
			// extractor_config

			// Optional
			"order": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"condition_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func newExtractor(d *schema.ResourceData) (*graylog.Extractor, string, error) {
	return &graylog.Extractor{
		Title:          d.Get("title").(string),
		Type:           d.Get("type").(string),
		Order:          d.Get("order").(int),
		SourceField:    d.Get("source_field").(string),
		TargetField:    d.Get("target_field").(string),
		ConditionType:  d.Get("condition_type").(string),
		ConditionValue: d.Get("condition_value").(string),
		// converters
		// extractor_config
		ID: d.Id(),
	}, d.Get("input_id").(string), nil
}

func resourceExtractorCreate(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	extractor, inputID, err := newExtractor(d)
	if err != nil {
		return err
	}

	if _, err := cl.CreateExtractor(inputID, extractor); err != nil {
		return err
	}
	d.SetId(extractor.ID)
	return nil
}

func resourceExtractorRead(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	extractor, _, err := cl.GetExtractor(d.Get("input_id").(string), d.Id())
	if err != nil {
		return err
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
	if err := setStrToRD(d, "condition_value", extractor.ConditionValue); err != nil {
		return err
	}
	return setIntToRD(d, "order", extractor.Order)
}

func resourceExtractorUpdate(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	extractor, inputID, err := newExtractor(d)
	if err != nil {
		return err
	}
	if _, err := cl.UpdateExtractor(inputID, extractor); err != nil {
		return err
	}
	return nil
}

func resourceExtractorDelete(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteExtractor(d.Get("input_id").(string), d.Id()); err != nil {
		return err
	}
	return nil
}
