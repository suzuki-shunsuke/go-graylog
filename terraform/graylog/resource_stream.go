package graylog

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
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
			"rule": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"field": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"inverted": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
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
		},
	}
}

func newStream(d *schema.ResourceData) *graylog.Stream {
	return &graylog.Stream{
		IndexSetID:   d.Get("index_set_id").(string),
		Title:        d.Get("title").(string),
		Description:  d.Get("description").(string),
		MatchingType: d.Get("matching_type").(string),
		Rules:        expandRules(d.Get("rule").(*schema.Set)),
		RemoveMatchesFromDefaultStream: d.Get(
			"remove_matches_from_default_stream").(bool),
		ID: d.Id(),
	}
}

func expandRules(data *schema.Set) []graylog.StreamRule {
	rules := []graylog.StreamRule{}
	for _, v := range data.List() {
		rules = append(rules, expandStreamRule(v.(map[string]interface{})))
	}
	return rules
}

func expandStreamRule(m map[string]interface{}) graylog.StreamRule {
	return graylog.StreamRule{
		Field:       m["field"].(string),
		Value:       m["value"].(string),
		Description: m["description"].(string),
		Type:        m["type"].(int),
		Inverted:    m["inverted"].(bool),
	}
}

func resourceStreamCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return fmt.Errorf("unable to create http client: %v", err)
	}

	stream := newStream(d)
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
	return nil
}

func resourceStreamRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return fmt.Errorf("unable to create http client: %v", err)
	}
	stream, _, err := cl.GetStream(d.Id())
	if err != nil {
		return err
	}
	setStrToRD(d, "index_set_id", stream.IndexSetID)
	setStrToRD(d, "title", stream.Title)
	setStrToRD(d, "description", stream.Description)
	setStrToRD(d, "matching_type", stream.MatchingType)
	setBoolToRD(
		d, "remove_matches_from_default_stream",
		stream.RemoveMatchesFromDefaultStream)
	d.Set("rule", stream.Rules)
	// content_pack
	setStrToRD(d, "creator_user_id", stream.CreatorUserID)
	setStrToRD(d, "created_at", stream.CreatedAt)
	setBoolToRD(d, "disabled", stream.Disabled)
	setBoolToRD(d, "is_default", stream.IsDefault)
	// alert_receivers
	// alert_conditions
	return nil
}

func resourceStreamUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return fmt.Errorf("unable to create http client: %v", err)
	}

	stream := newStream(d)
	if err := updateStreamRule(cl, stream); err != nil {
		return err
	}

	if _, err = cl.UpdateStream(stream); err != nil {
		return err
	}
	return nil
}

func updateStreamRule(cl *client.Client, stream *graylog.Stream) error {
	streamRules, _, _, err := cl.GetStreamRules(stream.ID)
	if err != nil {
		return fmt.Errorf("failed to get stream rules: %v", err)
	}

	// check if the stream rule was removed from the tf config
	for _, sr := range streamRules {
		exists := false
		for _, r := range stream.Rules {
			if r.ID == sr.ID {
				exists = true
				break
			}
		}

		if !exists {
			if _, err := cl.DeleteStreamRule(stream.ID, sr.ID); err != nil {
				return fmt.Errorf("failed to delete stream rule: %v", err)
			}
		}
	}

	for _, r := range stream.Rules {
		r.StreamID = stream.ID
		// create directly when either there are no stream rules in graylog server or the stream rule does not have id
		if len(streamRules) == 0 || r.ID == "" {
			if _, err := cl.CreateStreamRule(&r); err != nil {
				return fmt.Errorf("failed to create stream rule: %v", err)
			}
		} else {
			_, res, err := cl.GetStreamRule(stream.ID, r.ID)
			if err != nil {
				return fmt.Errorf("failed to get stream rule: %v", err)
			}

			switch res.Response.StatusCode {
			case http.StatusNotFound:
				if _, err = cl.CreateStreamRule(&r); err != nil {
					return fmt.Errorf("failed to create stream rule: %v", err)
				}
			case http.StatusOK:
				if _, err = cl.UpdateStreamRule(&r); err != nil {
					return fmt.Errorf("failed to update stream rule: %v", err)
				}
			}
		}
	}

	return nil
}

func resourceStreamDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteStream(d.Id()); err != nil {
		return err
	}
	return nil
}
