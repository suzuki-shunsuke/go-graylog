package graylog

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog/v8"
)

func resourcePipelineConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineConnectionCreate,
		Read:   resourcePipelineConnectionRead,
		Update: resourcePipelineConnectionUpdate,
		Delete: resourcePipelineConnectionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"stream_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pipeline_ids": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func newPipelineConnection(d *schema.ResourceData) *graylog.PipelineConnection {
	return &graylog.PipelineConnection{
		StreamID:    d.Get("stream_id").(string),
		PipelineIDs: getStringArray(d.Get("pipeline_ids").(*schema.Set).List()),
	}
}

func resourcePipelineConnectionCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	conn := newPipelineConnection(d)
	if _, err := cl.ConnectPipelinesToStream(ctx, conn); err != nil {
		return err
	}
	d.SetId(conn.StreamID)
	return nil
}

func resourcePipelineConnectionRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	pipelines := []string{}
	conn, ei, err := cl.GetPipelineConnectionsOfStream(ctx, d.Id())
	if err != nil {
		if ei == nil || ei.Response == nil || ei.Response.StatusCode != 404 {
			return err
		}
	} else {
		pipelines = conn.PipelineIDs
	}
	if err := setStrToRD(d, "stream_id", d.Id()); err != nil {
		return err
	}
	return setStrListToRD(d, "pipeline_ids", pipelines)
}

func resourcePipelineConnectionUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	conn := newPipelineConnection(d)
	if _, err := cl.ConnectPipelinesToStream(ctx, conn); err != nil {
		return err
	}
	return nil
}

func resourcePipelineConnectionDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	conn := newPipelineConnection(d)
	if _, err := cl.ConnectPipelinesToStream(ctx, &graylog.PipelineConnection{
		StreamID:    conn.StreamID,
		PipelineIDs: []string{},
	}); err != nil {
		return err
	}
	return nil
}
