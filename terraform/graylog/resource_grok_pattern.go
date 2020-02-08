package graylog

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog/v10"
)

func resourceGrokPattern() *schema.Resource {
	return &schema.Resource{
		Create: resourceGrokPatternCreate,
		Read:   resourceGrokPatternRead,
		Update: resourceGrokPatternUpdate,
		Delete: resourceGrokPatternDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pattern": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func newGrokPattern(d *schema.ResourceData) *graylog.GrokPattern {
	return &graylog.GrokPattern{
		Name:    d.Get("name").(string),
		Pattern: d.Get("pattern").(string),
		ID:      d.Id(),
	}
}

func resourceGrokPatternCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	grokPattern := newGrokPattern(d)
	if _, err := cl.CreateGrokPattern(ctx, grokPattern); err != nil {
		return err
	}
	d.SetId(grokPattern.ID)
	return nil
}

func resourceGrokPatternRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	grokPattern, ei, err := cl.GetGrokPattern(ctx, d.Id())
	if err != nil {
		return handleGetResourceError(d, ei, err)
	}
	if err := setStrToRD(d, "name", grokPattern.Name); err != nil {
		return err
	}
	return setStrToRD(d, "pattern", grokPattern.Pattern)
}

func resourceGrokPatternUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	grokPattern := newGrokPattern(d)
	_, err = cl.UpdateGrokPattern(ctx, grokPattern)
	return err
}

func resourceGrokPatternDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteGrokPattern(ctx, d.Id()); err != nil {
		return err
	}
	return nil
}
