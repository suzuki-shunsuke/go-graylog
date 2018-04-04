package graylog

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
	"github.com/suzuki-shunsuke/go-set"
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"permissions": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"read_only": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func newRole(d *schema.ResourceData) *graylog.Role {
	return &graylog.Role{
		Name:        d.Get("name").(string),
		Permissions: set.NewStrSet(getStringArray(d.Get("permissions").([]interface{}))...),
		Description: d.Get("description").(string),
		ReadOnly:    d.Get("read_only").(bool),
	}
}

func resourceRoleCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	role := newRole(d)
	if _, err := cl.CreateRole(role); err != nil {
		return err
	}
	d.SetId(role.Name)
	return nil
}

func resourceRoleRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	role, _, err := cl.GetRole(d.Get("name").(string))
	if err != nil {
		return err
	}
	d.Set("name", role.Name)
	d.Set("permissions", role.Permissions)
	d.Set("description", role.Description)
	d.Set("read_only", role.ReadOnly)
	return nil
}

func resourceRoleUpdate(d *schema.ResourceData, m interface{}) error {
	o, n := d.GetChange("name")
	oldName := o.(string)
	newName := n.(string)
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	role := newRole(d)
	role.Name = newName
	if _, err := cl.UpdateRole(oldName, role); err != nil {
		return err
	}
	return nil
}

func resourceRoleDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteRole(d.Get("name").(string)); err != nil {
		return err
	}
	return nil
}
