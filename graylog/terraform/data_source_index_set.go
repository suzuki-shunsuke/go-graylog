package terraform

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
)

func dataSourceIndexSet() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIndexSetRead,

		Schema: map[string]*schema.Schema{
			"index_set_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"title": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"index_prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// computed
			"rotation_strategy_class": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rotation_strategy": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_docs_per_index": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rotation_period": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"retention_strategy_class": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retention_strategy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_number_of_indices": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"index_analyzer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"shards": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"index_optimization_max_num_segments": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"creation_date": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replicas": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"index_optimization_disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"writable": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"default": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			// field_type_refresh_interval is added from Graylog API v3
			"field_type_refresh_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func setIndexSet(d *schema.ResourceData, is *graylog.IndexSet, cfg *Config) error {
	if is.RotationStrategy != nil {
		b, err := json.Marshal(is.RotationStrategy)
		if err != nil {
			return err
		}
		dest := map[string]interface{}{}
		if err := json.Unmarshal(b, &dest); err != nil {
			return err
		}
		if err := d.Set("rotation_strategy", []map[string]interface{}{dest}); err != nil {
			return err
		}
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
		if err := d.Set("retention_strategy", []map[string]interface{}{dest}); err != nil {
			return err
		}
	}

	if err := setStrToRD(d, "title", is.Title); err != nil {
		return err
	}
	if err := setStrToRD(d, "index_prefix", is.IndexPrefix); err != nil {
		return err
	}
	if err := setStrToRD(d, "description", is.Description); err != nil {
		return err
	}
	if err := setIntToRD(d, "shards", is.Shards); err != nil {
		return err
	}
	if err := setIntToRD(d, "replicas", is.Replicas); err != nil {
		return err
	}
	if err := setStrToRD(d, "rotation_strategy_class", is.RotationStrategyClass); err != nil {
		return err
	}
	if err := setStrToRD(d, "retention_strategy_class", is.RetentionStrategyClass); err != nil {
		return err
	}
	if err := setStrToRD(d, "index_analyzer", is.IndexAnalyzer); err != nil {
		return err
	}
	if err := setIntToRD(d, "index_optimization_max_num_segments", is.IndexOptimizationMaxNumSegments); err != nil {
		return err
	}
	if err := setBoolToRD(d, "index_optimization_disabled", is.IndexOptimizationDisabled); err != nil {
		return err
	}
	if err := setBoolToRD(d, "writable", is.Writable); err != nil {
		return err
	}

	if cfg.APIVersion == "v3" {
		if err := setIntToRD(d, "field_type_refresh_interval", is.FieldTypeRefreshInterval); err != nil {
			return err
		}
	}
	d.SetId(is.ID)
	return setBoolToRD(d, "default", is.Default)
}

func dataSourceIndexSetRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	cfg := m.(*Config)

	if id, ok := d.GetOk("index_set_id"); ok {
		if _, ok := d.GetOk("title"); ok {
			return errors.New("only one of index_set_id or title or index_prefix must be set")
		}
		if _, ok := d.GetOk("index_prefix"); ok {
			return errors.New("only one of index_set_id or title or index_prefix must be set")
		}
		is, _, err := cl.GetIndexSet(ctx, id.(string))
		if err != nil {
			return err
		}
		return setIndexSet(d, is, cfg)
	}

	if t, ok := d.GetOk("title"); ok {
		if _, ok := d.GetOk("index_prefix"); ok {
			return errors.New("only one of index_set_id or title or index_prefix must be set")
		}
		title := t.(string)
		indexSets, _, _, _, err := cl.GetIndexSets(ctx, 0, 0, false)
		if err != nil {
			return err
		}
		iss := []graylog.IndexSet{}
		for _, is := range indexSets {
			if is.Title == title {
				iss = append(iss, is)
			}
		}
		switch len(iss) {
		case 0:
			return errors.New("matched index set is not found")
		case 1:
			return setIndexSet(d, &iss[0], cfg)
		}
		return errors.New("title isn't unique")
	}

	if p, ok := d.GetOk("index_prefix"); ok {
		prefix := p.(string)
		indexSets, _, _, _, err := cl.GetIndexSets(ctx, 0, 0, false)
		if err != nil {
			return err
		}
		for _, is := range indexSets {
			if is.IndexPrefix == prefix {
				return setIndexSet(d, &is, cfg)
			}
		}
		return errors.New("matched index prefix is not found")
	}
	return errors.New("one of index_set_id or title or prefix must be set")
}
