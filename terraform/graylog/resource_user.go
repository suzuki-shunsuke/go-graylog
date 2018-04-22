package graylog

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
	"github.com/suzuki-shunsuke/go-set"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"permissions": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"full_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"roles": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_timeout_ms": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"user_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"read_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"client_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_active": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"last_activity": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Computed: true,
			},
		},
	}
}

func newUser(d *schema.ResourceData) *graylog.User {
	permissions := set.NewStrSet(getStringArray(d.Get("permissions").([]interface{}))...)
	roles := set.NewStrSet(getStringArray(d.Get("roles").([]interface{}))...)
	return &graylog.User{
		Username:         d.Get("username").(string),
		Roles:            roles,
		Permissions:      permissions,
		Email:            d.Get("email").(string),
		FullName:         d.Get("full_name").(string),
		Timezone:         d.Get("timezone").(string),
		SessionTimeoutMs: d.Get("session_timeout_ms").(int),
		External:         d.Get("external").(bool),
		ClientAddress:    d.Get("client_address").(string),
		Password:         d.Get("password").(string),
		ReadOnly:         d.Get("read_only").(bool),
		// SessionActive:    d.Get("session_active").(bool),
	}
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	user := newUser(d)
	if _, err = cl.CreateUser(user); err != nil {
		return err
	}
	d.SetId(user.Username)
	return nil
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	user, _, err := cl.GetUser(d.Get("username").(string))
	if err != nil {
		return err
	}
	d.Set("username", user.Username)
	d.Set("email", user.Email)
	d.Set("permissions", user.Permissions)
	d.Set("timezone", user.Timezone)
	d.Set("session_timeout_ms", user.SessionTimeoutMs)
	d.Set("external", user.External)
	d.Set("read_only", user.ReadOnly)
	return nil
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	user := newUser(d)
	if _, err = cl.UpdateUser(user.NewUpdateParams()); err != nil {
		return err
	}
	return nil
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteUser(d.Get("username").(string)); err != nil {
		return err
	}
	return nil
}
