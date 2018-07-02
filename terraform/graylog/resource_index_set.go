package graylog

import (
	"encoding/json"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
	"github.com/suzuki-shunsuke/go-graylog/util"
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
	}, nil
}

func resourceIndexSetCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	is, err := newIndexSet(d)
	if err != nil {
		return err
	}
	if _, err := cl.CreateIndexSet(is); err != nil {
		return err
	}
	d.SetId(is.ID)
	return nil
}

func resourceIndexSetRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	is, _, err := cl.GetIndexSet(d.Id())
	if err != nil {
		return err
	}
	if is.RotationStrategy != nil {
		b, err := json.Marshal(is.RotationStrategy)
		if err != nil {
			return err
		}
		dest := map[string]interface{}{}
		if err := json.Unmarshal(b, &dest); err != nil {
			return err
		}
		d.Set("rotation_strategy", []map[string]interface{}{dest})
	}
	if is.RetentionStrategy != nil {
		b, err := json.Marshal(is.RetentionStrategy)
		if err != nil {
			return err
		}
		dest := map[string]interface{}{}
		if err := json.Unmarshal(b, &dest); err != nil {
			return err
		}
		d.Set("retention_strategy", []map[string]interface{}{dest})
	}

	setStrToRD(d, "title", is.Title)
	setStrToRD(d, "index_prefix", is.IndexPrefix)
	setStrToRD(d, "description", is.Description)
	setIntToRD(d, "shards", is.Shards)
	setIntToRD(d, "replicas", is.Replicas)
	setStrToRD(d, "rotation_strategy_class", is.RotationStrategyClass)
	setStrToRD(d, "retention_strategy_class", is.RetentionStrategyClass)
	setStrToRD(d, "index_analyzer", is.IndexAnalyzer)
	setIntToRD(d, "index_optimization_max_num_segments", is.IndexOptimizationMaxNumSegments)
	setBoolToRD(d, "index_optimization_disabled", is.IndexOptimizationDisabled)
	setBoolToRD(d, "writable", is.Writable)
	setBoolToRD(d, "default", is.Default)
	return nil
}

func resourceIndexSetUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	is, err := newIndexSet(d)
	if err != nil {
		return err
	}

	if _, _, err = cl.UpdateIndexSet(is.NewUpdateParams()); err != nil {
		return err
	}
	return nil
}

func resourceIndexSetDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteIndexSet(d.Id()); err != nil {
		return err
	}
	return nil
}
