package graylog

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog"
)

func dataSourceStream() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStreamRead,

		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stream_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"index_set_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// content_pack
			"matching_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remove_matches_from_default_stream": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"is_default": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func setStream(d *schema.ResourceData, stream *graylog.Stream, cfg *Config) error {
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
	d.SetId(stream.ID)
	return setBoolToRD(d, "is_default", stream.IsDefault)
}

func dataSourceStreamRead(d *schema.ResourceData, m interface{}) error {
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	cfg := m.(*Config)

	if id, ok := d.GetOk("stream_id"); ok {
		if _, ok := d.GetOk("title"); ok {
			return errors.New("both stream_id and title must not be set")
		}
		stream, _, err := cl.GetStream(id.(string))
		if err != nil {
			return err
		}
		return setStream(d, stream, cfg)
	}

	if t, ok := d.GetOk("title"); ok {
		title := t.(string)
		streams, _, _, err := cl.GetStreams()
		if err != nil {
			return err
		}
		arr := []graylog.Stream{}
		for _, stream := range streams {
			if stream.Title == title {
				arr = append(arr, stream)
			}
		}
		switch len(arr) {
		case 0:
			return errors.New("matched stream is not found")
		case 1:
			return setStream(d, &arr[0], cfg)
		}
		return errors.New("title isn't unique")
	}
	return errors.New("one of stream_id or title must be set")
}
