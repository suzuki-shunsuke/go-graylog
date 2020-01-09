package graylog

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-set/v6"
)

func resourceStreamOutput() *schema.Resource {
	return &schema.Resource{
		Create: resourceStreamOutputCreate,
		Read:   resourceStreamOutputRead,
		Update: resourceStreamOutputUpdate,
		Delete: resourceStreamOutputDelete,

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
			"output_ids": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func newStreamOutput(d *schema.ResourceData) (string, []string) {
	return d.Get("stream_id").(string), getStringArray(d.Get("output_ids").(*schema.Set).List())
}

func resourceStreamOutputCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	streamID, outputIDs := newStreamOutput(d)
	if _, err := cl.CreateStreamOutputs(ctx, streamID, outputIDs); err != nil {
		return err
	}
	d.SetId(streamID)
	return nil
}

func resourceStreamOutputRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	outputIDs := []string{}
	outputs, _, ei, err := cl.GetStreamOutputs(ctx, d.Id())
	if err != nil {
		if ei == nil || ei.Response == nil || ei.Response.StatusCode != 404 {
			return err
		}
	} else {
		outputIDs = make([]string, len(outputs))
		for i, output := range outputs {
			outputIDs[i] = output.ID
		}
	}
	if err := setStrToRD(d, "stream_id", d.Id()); err != nil {
		return err
	}
	return setStrListToRD(d, "output_ids", outputIDs)
}

func resourceStreamOutputUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	streamID := d.Get("stream_id").(string)
	oldV, newV := d.GetChange("output_ids")

	oldVSet := set.NewStrSet(getStringArray(oldV.(*schema.Set).List())...)
	newVSet := set.NewStrSet(getStringArray(newV.(*schema.Set).List())...)
	for k := range oldVSet.ToMap(false) {
		if newVSet.Has(k) {
			continue
		}
		if _, err := cl.DeleteStreamOutput(ctx, streamID, k); err != nil {
			return err
		}
	}
	if oldVSet.HasAll(newVSet.ToList()...) {
		return nil
	}

	if _, err := cl.CreateStreamOutputs(ctx, streamID, newVSet.ToList()); err != nil {
		return err
	}
	return nil
}

func resourceStreamOutputDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	streamID, outputIDs := newStreamOutput(d)
	for _, outputID := range outputIDs {
		if _, err := cl.DeleteStreamOutput(ctx, streamID, outputID); err != nil {
			return err
		}
	}
	return nil
}
