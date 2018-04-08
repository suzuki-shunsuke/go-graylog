package graylog

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
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
				Type:     schema.TypeMap,
				Required: true,
			},
			"retention_strategy_class": {
				Type:     schema.TypeString,
				Required: true,
			},
			// type required
			"retention_strategy": {
				Type:     schema.TypeMap,
				Required: true,
			},
			"creation_date": {
				Type:     schema.TypeString,
				Required: true,
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

	ros := d.Get("rotation_strategy").(map[string]interface{})
	res := d.Get("retention_strategy").(map[string]interface{})

	rosType, err := getString(ros, "type", true)
	if err != nil {
		return nil, err
	}
	rotationStrategy.Type = rosType

	maxDocsPerIndex, err := getStrInt(ros, "max_docs_per_index", false)
	if err != nil {
		return nil, err
	}
	rotationStrategy.MaxDocsPerIndex = maxDocsPerIndex

	resType, err := getString(res, "type", true)
	if err != nil {
		return nil, err
	}
	retentionStrategy.Type = resType

	maxNumberOfIndices, err := getStrInt(ros, "max_number_of_indices", false)
	if err != nil {
		return nil, err
	}
	retentionStrategy.MaxNumberOfIndices = maxNumberOfIndices

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
	indexSet, _, err := cl.GetIndexSet(d.Id())
	if err != nil {
		return err
	}
	d.Set("title", indexSet.Title)
	d.Set("description", indexSet.Description)
	d.Set("shards", indexSet.Shards)
	d.Set("replicas", indexSet.Replicas)
	d.Set("rotation_strategy_class", indexSet.RotationStrategyClass)
	d.Set("retention_strategy_class", indexSet.RetentionStrategyClass)
	d.Set("index_analyzer", indexSet.IndexAnalyzer)
	d.Set(
		"index_optimization_max_num_segments",
		indexSet.IndexOptimizationMaxNumSegments)
	d.Set("index_optimization_disabled", indexSet.IndexOptimizationDisabled)
	d.Set("writable", indexSet.Writable)
	d.Set("default", indexSet.Default)
	return nil
}

func resourceIndexSetUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	indexSet, err := newIndexSet(d)
	if err != nil {
		return err
	}

	if _, err = cl.UpdateIndexSet(indexSet); err != nil {
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
