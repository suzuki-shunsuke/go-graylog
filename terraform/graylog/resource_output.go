package graylog

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-jsoneq/jsoneq"

	"github.com/suzuki-shunsuke/go-graylog/v10"
)

func resourceOutput() *schema.Resource {
	return &schema.Resource{
		Create: resourceOutputCreate,
		Read:   resourceOutputRead,
		Update: resourceOutputUpdate,
		Delete: resourceOutputDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"configuration": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: schemaDiffSuppressJSONString,
			},
		},
	}
}

func newOutput(d *schema.ResourceData) (*graylog.Output, error) {
	cfgS := d.Get("configuration").(string)
	cfg, err := jsoneq.ConvertByte([]byte(cfgS))
	if err != nil {
		return nil, fmt.Errorf("failed to parse the 'configuration'. 'configuration' must be a JSON string '%s': %w", cfgS, err)
	}
	return &graylog.Output{
		ID:            d.Id(),
		Title:         d.Get("title").(string),
		Type:          d.Get("type").(string),
		Configuration: cfg,
	}, nil
}

func resourceOutputCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	output, err := newOutput(d)
	if err != nil {
		return err
	}
	if _, err := cl.CreateOutput(ctx, output); err != nil {
		return err
	}
	d.SetId(output.ID)
	return nil
}

func resourceOutputRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	output, ei, err := cl.GetOutput(ctx, d.Id())
	if err != nil {
		return handleGetResourceError(d, ei, err)
	}
	if err := setStrToRD(d, "title", output.Title); err != nil {
		return err
	}
	if err := setStrToRD(d, "type", output.Type); err != nil {
		return err
	}
	b, err := json.Marshal(output.Configuration)
	if err != nil {
		return err
	}
	return setStrToRD(d, "configuration", string(b))
}

func resourceOutputUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	output, err := newOutput(d)
	if err != nil {
		return err
	}
	_, err = cl.UpdateOutput(ctx, output)
	return err
}

func resourceOutputDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteOutput(ctx, d.Id()); err != nil {
		return err
	}
	return nil
}
