package graylog

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog"
)

func resourcePipeline() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineCreate,
		Read:   resourcePipelineRead,
		Update: resourcePipelineUpdate,
		Delete: resourcePipelineDelete,

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

func newPipeline(d *schema.ResourceData) *graylog.Pipeline {
	return &graylog.Pipeline{
		ID:          d.Id(),
		Source:      d.Get("source").(string),
		Description: d.Get("description").(string),
	}
}

func resourcePipelineCreate(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	pipe := newPipeline(d)
	if pipe.Source == "" {
		return fmt.Errorf("source is required to create a pipeline")
	}
	if _, err = cl.CreatePipeline(pipe); err != nil {
		return err
	}
	d.SetId(pipe.ID)
	return nil
}

func resourcePipelineRead(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	pipe, _, err := cl.GetPipeline(d.Id())
	if err != nil {
		return err
	}
	if err := setStrToRD(d, "source", pipe.Source); err != nil {
		return err
	}
	return setStrToRD(d, "description", pipe.Description)
}

func resourcePipelineUpdate(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	pipe := newPipeline(d)
	_, err = cl.UpdatePipeline(pipe)
	return err
}

func resourcePipelineDelete(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeletePipeline(d.Id()); err != nil {
		return err
	}
	return nil
}
