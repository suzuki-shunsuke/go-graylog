package graylog

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceInputStaticFields() *schema.Resource {
	return &schema.Resource{
		Create: resourceInputStaticFieldsCreate,
		Read:   resourceInputStaticFieldsRead,
		Update: resourceInputStaticFieldsUpdate,
		Delete: resourceInputStaticFieldsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"input_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// Optional
			"fields": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func newInputStaticFields(d *schema.ResourceData) (string, map[string]string, error) {
	f := d.Get("fields").(map[string]interface{})
	fields := make(map[string]string, len(f))
	for k, v := range f {
		fields[k] = v.(string)
	}
	return d.Get("input_id").(string), fields, nil
}

func resourceInputStaticFieldsCreate(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	inputID, fields, err := newInputStaticFields(d)
	if err != nil {
		return err
	}
	input, _, err := cl.GetInput(inputID)
	if err != nil {
		return err
	}
	staticFields := input.StaticFields
	if staticFields == nil {
		staticFields = map[string]string{}
	}

	for k, v := range fields {
		if oldV, ok := input.StaticFields[k]; ok {
			if v == oldV {
				continue
			}
			if _, err := cl.CreateInputStaticField(inputID, k, v); err != nil {
				return err
			}
			continue
		}
		if _, err := cl.CreateInputStaticField(inputID, k, v); err != nil {
			return err
		}
	}
	for k := range staticFields {
		if _, ok := fields[k]; ok {
			continue
		}
		if _, err := cl.DeleteInputStaticField(inputID, k); err != nil {
			return err
		}
	}
	d.SetId(inputID)
	return nil
}

func resourceInputStaticFieldsRead(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	input, _, err := cl.GetInput(d.Get("input_id").(string))
	if err != nil {
		return err
	}
	return d.Set("fields", input.StaticFields)
}

func resourceInputStaticFieldsUpdate(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}

	o, n := d.GetChange("fields")
	oldF := o.(map[string]interface{})
	newF := n.(map[string]interface{})
	oldFields := make(map[string]string, len(oldF))
	for k, v := range oldF {
		oldFields[k] = v.(string)
	}
	newFields := make(map[string]string, len(newF))
	for k, v := range newF {
		newFields[k] = v.(string)
	}
	inputID := d.Get("input_id").(string)
	for k, v := range oldFields {
		if newV, ok := newFields[k]; ok {
			if v == newV {
				continue
			}
			// update
			if _, err := cl.CreateInputStaticField(inputID, k, newV); err != nil {
				return err
			}
			continue
		}
		// delete
		if _, err := cl.DeleteInputStaticField(inputID, k); err != nil {
			return err
		}
	}
	for k, newV := range newFields {
		if _, ok := oldFields[k]; ok {
			continue
		}
		// create
		if _, err := cl.CreateInputStaticField(inputID, k, newV); err != nil {
			return err
		}
	}
	return nil
}

func resourceInputStaticFieldsDelete(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	inputID, fields, err := newInputStaticFields(d)
	if err != nil {
		return err
	}
	for k := range fields {
		if _, err := cl.DeleteInputStaticField(inputID, k); err != nil {
			return err
		}
	}
	return nil
}
