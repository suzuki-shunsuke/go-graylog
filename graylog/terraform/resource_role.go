package terraform

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-set/v6"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
)

func resourceRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoleCreate,
		Read:   resourceRoleRead,
		Update: resourceRoleUpdate,
		Delete: resourceRoleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"permissions": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"read_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func newRole(d *schema.ResourceData) *graylog.Role {
	return &graylog.Role{
		Name:        d.Get("name").(string),
		Permissions: set.NewStrSet(getStringArray(d.Get("permissions").(*schema.Set).List())...),
		Description: d.Get("description").(string),
		ReadOnly:    d.Get("read_only").(bool),
	}
}

func resourceRoleCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	role := newRole(d)
	if _, err := cl.CreateRole(ctx, role); err != nil {
		return err
	}
	d.SetId(role.Name)
	return nil
}

func resourceRoleRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	role, ei, err := cl.GetRole(ctx, d.Id())
	if err != nil {
		return handleGetResourceError(d, ei, err)
	}
	if err := setStrToRD(d, "name", role.Name); err != nil {
		return err
	}
	if err := setStrListToRD(d, "permissions", role.Permissions.ToList()); err != nil {
		return err
	}
	if err := setStrToRD(d, "description", role.Description); err != nil {
		return err
	}
	return setBoolToRD(d, "read_only", role.ReadOnly)
}

func resourceRoleUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	o, n := d.GetChange("name")
	oldName := o.(string)
	newName := n.(string)
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	role := newRole(d)
	role.Name = newName
	_, _, err = cl.UpdateRole(ctx, oldName, role.NewUpdateParams())
	return err
}

func resourceRoleDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteRole(ctx, d.Id()); err != nil {
		return err
	}
	return nil
}
