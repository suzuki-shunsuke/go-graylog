package graylog

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-graylog"
)

func resourceStream() *schema.Resource {
	return &schema.Resource{
		Create: resourceStreamCreate,
		Read:   resourceStreamRead,
		Update: resourceStreamUpdate,
		Delete: resourceStreamDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"index_set_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			// rules
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// content_pack
			"matching_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remove_matches_from_default_stream": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			// attributes
			"creator_user_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// outputs
			"created_at": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_default": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			// alert_conditions
			// alert_receivers

			"pipelines": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func newStream(d *schema.ResourceData) (*graylog.Stream, []string, error) {
	return &graylog.Stream{
		IndexSetID:   d.Get("index_set_id").(string),
		Title:        d.Get("title").(string),
		Description:  d.Get("description").(string),
		MatchingType: d.Get("matching_type").(string),
		RemoveMatchesFromDefaultStream: d.Get(
			"remove_matches_from_default_stream").(bool),
		ID: d.Id(),
	}, getStringArray(d.Get("pipelines").(*schema.Set).List()), nil
}

func resourceStreamCreate(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	stream, pipelines, err := newStream(d)
	if err != nil {
		return err
	}

	if _, err := cl.CreateStream(stream); err != nil {
		return err
	}
	d.SetId(stream.ID)
	// resume if needed
	disabled := d.Get("disabled").(bool)
	if !disabled {
		if _, err := cl.ResumeStream(stream.ID); err != nil {
			return err
		}
	}
	// create pipeline connection
	if _, ok := d.GetOk("pipelines"); ok {
		if _, err := cl.ConnectPipelinesToStream(&graylog.PipelineConnection{
			PipelineIDs: pipelines,
			StreamID:    stream.ID,
		}); err != nil {
			return err
		}
	}

	return nil
}

func resourceStreamRead(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	stream, _, err := cl.GetStream(d.Id())
	if err != nil {
		return err
	}
	if err := setStrToRD(d, "index_set_id", stream.IndexSetID); err != nil {
		return err
	}
	if err := setStrToRD(d, "title", stream.Title); err != nil {
		return err
	}
	if err := setStrToRD(d, "description", stream.Description); err != nil {
		return err
	}
	if err := setStrToRD(d, "matching_type", stream.MatchingType); err != nil {
		return err
	}
	if err := setBoolToRD(d, "remove_matches_from_default_stream", stream.RemoveMatchesFromDefaultStream); err != nil {
		return err
	}
	// rules
	// content_pack
	if err := setStrToRD(d, "creator_user_id", stream.CreatorUserID); err != nil {
		return err
	}
	if err := setStrToRD(d, "created_at", stream.CreatedAt); err != nil {
		return err
	}
	if err := setBoolToRD(d, "disabled", stream.Disabled); err != nil {
		return err
	}
	if err := setBoolToRD(d, "is_default", stream.IsDefault); err != nil {
		return err
	}
	if _, ok := d.GetOk("pipelines"); ok {
		pipelines := []string{}
		conn, ei, err := cl.GetPipelineConnectionsOfStream(d.Id())
		if err != nil {
			if ei == nil || ei.Response == nil || ei.Response.StatusCode != 404 {
				return err
			}
		} else {
			pipelines = conn.PipelineIDs
		}
		if err := setStrListToRD(d, "pipelines", pipelines); err != nil {
			return err
		}
	}
	return nil

	// alert_receivers
	// alert_conditions
}

func resourceStreamUpdate(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	stream, pipelines, err := newStream(d)
	if err != nil {
		return err
	}
	if _, err := cl.UpdateStream(stream); err != nil {
		return err
	}
	if _, ok := d.GetOk("pipelines"); ok {
		if _, err := cl.ConnectPipelinesToStream(&graylog.PipelineConnection{
			PipelineIDs: pipelines,
			StreamID:    stream.ID,
		}); err != nil {
			return err
		}
	}
	return nil
}

func resourceStreamDelete(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteStream(d.Id()); err != nil {
		return err
	}
	return nil
}
