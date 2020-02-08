package terraform

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
)

func resourcePipelineRule() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineRuleCreate,
		Read:   resourcePipelineRuleRead,
		Update: resourcePipelineRuleUpdate,
		Delete: resourcePipelineRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"source": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// We don't define the attribute "title",
			// because the request parameter "title" is ignored in create and update API.
		},
	}
}

func newPipelineRule(d *schema.ResourceData) *graylog.PipelineRule {
	return &graylog.PipelineRule{
		ID:          d.Id(),
		Source:      d.Get("source").(string),
		Description: d.Get("description").(string),
	}
}

func resourcePipelineRuleCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	rule := newPipelineRule(d)
	if rule.Source == "" {
		return errors.New("source is required to create a pipeline rule")
	}
	if _, err = cl.CreatePipelineRule(ctx, rule); err != nil {
		return err
	}
	d.SetId(rule.ID)
	return nil
}

func resourcePipelineRuleRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	rule, ei, err := cl.GetPipelineRule(ctx, d.Id())
	if err != nil {
		return handleGetResourceError(d, ei, err)
	}
	if err := setStrToRD(d, "source", rule.Source); err != nil {
		return err
	}
	return setStrToRD(d, "description", rule.Description)
}

func resourcePipelineRuleUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	rule := newPipelineRule(d)
	_, err = cl.UpdatePipelineRule(ctx, rule)
	return err
}

func resourcePipelineRuleDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeletePipelineRule(ctx, d.Id()); err != nil {
		return err
	}
	return nil
}
