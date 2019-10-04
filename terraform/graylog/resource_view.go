package graylog

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-graylog"
)

func resourceView() *schema.Resource {
	return &schema.Resource{
		Create: resourceViewCreate,
		Read:   resourceViewRead,
		Update: resourceViewUpdate,
		Delete: resourceViewDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},

			"search_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"state": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"selected_fields": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"titles": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Schema{
								Type:     schema.TypeMap,
								Optional: true,
								Elem: &schema.Schema{
									Type: schema.TypeString,
								},
							},
						},

						"widget_mapping": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeList,
								Elem: &schema.Schema{
									Type: schema.TypeString,
								},
							},
						},

						"widgets": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Optional: true,
									},

									//	"aggregation_config": {
									//		Type:     schema.TypeList,
									//		Optional: true,
									//		MaxItems: 1,
									//		MinItems: 1,
									//		Elem:     &schema.Resource{
									//		},
									//	},
								},
							},
						},

						"positions": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Schema{
								Type:     schema.TypeList,
								Optional: true,
								MaxItems: 1,
								MinItems: 1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										// "width": {
										// 	Type:     schema.TypeInt,
										// 	Optional: true,
										// },
										"col": {
											Type:     schema.TypeInt,
											Optional: true,
										},
										"row": {
											Type:     schema.TypeInt,
											Optional: true,
										},
										"height": {
											Type:     schema.TypeInt,
											Optional: true,
										},
									},
								},
							},
						},
					},
				},
			},

			"summary": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"owner": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"created_at": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			//	"dashboard_state": {
			//		Type:     schema.TypeList,
			//		Optional: true,
			//		MaxItems: 1,
			//		MinItems: 1,
			//		Elem: &schema.Resource{
			//			Schema: map[string]*schema.Schema{},
			//		},
			//	},
		},
	}
}

func newView(d *schema.ResourceData) (*graylog.View, error) {
	return &graylog.View{
		Title:       d.Get("title").(string),
		Summary:     d.Get("summary").(string),
		Description: d.Get("description").(string),
		SearchID:    d.Get("search_id").(string),
		Owner:       d.Get("owner").(string),
		CreatedAt:   d.Get("created_at").(string),
		ID:          d.Id(),
	}, nil
}

func resourceViewCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	view, err := newView(d)
	if err != nil {
		return err
	}

	if _, err := cl.CreateView(ctx, view); err != nil {
		return err
	}
	d.SetId(view.ID)
	return nil
}

func setView(d *schema.ResourceData, view *graylog.View, m interface{}) error {
	if err := setStrToRD(d, "title", view.Title); err != nil {
		return err
	}
	if err := setStrToRD(d, "summary", view.Summary); err != nil {
		return err
	}
	if err := setStrToRD(d, "description", view.Description); err != nil {
		return err
	}
	if err := setStrToRD(d, "search_id", view.SearchID); err != nil {
		return err
	}
	if err := setStrToRD(d, "owner", view.Owner); err != nil {
		return err
	}
	if err := setStrToRD(d, "created_at", view.CreatedAt); err != nil {
		return err
	}

	return nil
}

func resourceViewRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	view, ei, err := cl.GetView(ctx, d.Id())
	if err != nil {
		return handleGetResourceError(d, ei, err)
	}
	return setView(d, view, m.(*Config))
}

func resourceViewUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	view, err := newView(d)
	if err != nil {
		return err
	}
	if _, err := cl.UpdateView(ctx, view); err != nil {
		return err
	}
	return nil
}

func resourceViewDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteView(ctx, d.Id()); err != nil {
		return err
	}
	return nil
}
