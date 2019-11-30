package graylog

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog/v8"
	"github.com/suzuki-shunsuke/go-graylog/util/v8"
)

func resourceIndexSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceIndexSetCreate,
		Read:   resourceIndexSetRead,
		Update: resourceIndexSetUpdate,
		Delete: resourceIndexSetDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"index_prefix": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rotation_strategy_class": {
				Type:     schema.TypeString,
				Required: true,
			},
			// type required
			"rotation_strategy": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"max_docs_per_index": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_size": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rotation_period": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"retention_strategy_class": {
				Type:     schema.TypeString,
				Required: true,
			},
			// type required
			"retention_strategy": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"max_number_of_indices": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"index_analyzer": {
				Type:     schema.TypeString,
				Required: true,
			},
			// >= 1
			"shards": {
				Type:     schema.TypeInt,
				Required: true,
			},
			// >= 1
			"index_optimization_max_num_segments": {
				Type:     schema.TypeInt,
				Required: true,
			},

			// Optional
			"creation_date": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"replicas": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"index_optimization_disabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"writable": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"default": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			// field_type_refresh_interval is added from Graylog API v3
			"field_type_refresh_interval": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func newIndexSet(d *schema.ResourceData) (*graylog.IndexSet, error) {
	rotationStrategy := &graylog.RotationStrategy{}
	retentionStrategy := &graylog.RetentionStrategy{}

	ros := d.Get("rotation_strategy").([]interface{})[0].(map[string]interface{})
	res := d.Get("retention_strategy").([]interface{})[0].(map[string]interface{})
	if err := util.MSDecode(ros, rotationStrategy); err != nil {
		return nil, err
	}
	if err := util.MSDecode(res, retentionStrategy); err != nil {
		return nil, err
	}

	return &graylog.IndexSet{
		ID:                              d.Id(),
		Title:                           d.Get("title").(string),
		IndexPrefix:                     d.Get("index_prefix").(string),
		Description:                     d.Get("description").(string),
		Shards:                          d.Get("shards").(int),
		Replicas:                        d.Get("replicas").(int),
		RotationStrategyClass:           d.Get("rotation_strategy_class").(string),
		RotationStrategy:                rotationStrategy,
		RetentionStrategyClass:          d.Get("retention_strategy_class").(string),
		RetentionStrategy:               retentionStrategy,
		IndexAnalyzer:                   d.Get("index_analyzer").(string),
		IndexOptimizationMaxNumSegments: d.Get("index_optimization_max_num_segments").(int),
		IndexOptimizationDisabled:       d.Get("index_optimization_disabled").(bool),
		Writable:                        d.Get("writable").(bool),
		Default:                         d.Get("default").(bool),
		CreationDate:                    d.Get("creation_date").(string),
		FieldTypeRefreshInterval:        d.Get("field_type_refresh_interval").(int),
	}, nil
}

func resourceIndexSetCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	is, err := newIndexSet(d)
	if err != nil {
		return err
	}
	if _, err := cl.CreateIndexSet(ctx, is); err != nil {
		return err
	}
	d.SetId(is.ID)
	return nil
}

func resourceIndexSetRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	cfg := m.(*Config)
	if err != nil {
		return err
	}
	is, ei, err := cl.GetIndexSet(ctx, d.Id())
	if err != nil {
		return handleGetResourceError(d, ei, err)
	}
	return setIndexSet(d, is, cfg)
}

func resourceIndexSetUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	is, err := newIndexSet(d)
	if err != nil {
		return err
	}

	if _, _, err = cl.UpdateIndexSet(ctx, is.NewUpdateParams()); err != nil {
		return err
	}
	return nil
}

func resourceIndexSetDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteIndexSet(ctx, d.Id()); err != nil {
		return err
	}
	return nil
}
