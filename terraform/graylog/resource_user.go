package graylog

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog/v8"
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
				ForceNew: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"full_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			// password is required to create but not required to update
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"permissions": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"roles": {
				Type:     schema.TypeSet,
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
			"session_active": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"last_activity": {
				Type:     schema.TypeString,
				Required: false,
				Computed: true,
			},
		},
	}
}

func newUser(d *schema.ResourceData) *graylog.User {
	return &graylog.User{
		Username:         d.Get("username").(string),
		Password:         d.Get("password").(string),
		Email:            d.Get("email").(string),
		Permissions:      set.NewStrSet(getStringArray(d.Get("permissions").(*schema.Set).List())...),
		FullName:         d.Get("full_name").(string),
		Roles:            set.NewStrSet(getStringArray(d.Get("roles").(*schema.Set).List())...),
		Timezone:         d.Get("timezone").(string),
		SessionTimeoutMs: d.Get("session_timeout_ms").(int),
		ID:               d.Get("user_id").(string),
		External:         d.Get("external").(bool),
		ReadOnly:         d.Get("read_only").(bool),
		ClientAddress:    d.Get("client_address").(string),
		SessionActive:    d.Get("session_active").(bool),
		LastActivity:     d.Get("last_activity").(string),
	}
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	user := newUser(d)
	if user.Password == "" {
		return errors.New("password is required to create a user")
	}
	if _, err = cl.CreateUser(ctx, user); err != nil {
		return err
	}
	d.SetId(user.Username)
	return setStrToRD(d, "user_id", user.ID)
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	user, ei, err := cl.GetUser(ctx, d.Id())
	if err != nil {
		return handleGetResourceError(d, ei, err)
	}
	if err := setStrToRD(d, "username", user.Username); err != nil {
		return err
	}
	if err := setStrToRD(d, "email", user.Email); err != nil {
		return err
	}
	if err := setStrListToRD(d, "permissions", user.Permissions.ToList()); err != nil {
		return err
	}
	if err := setStrToRD(d, "full_name", user.FullName); err != nil {
		return err
	}
	if err := setStrListToRD(d, "roles", user.Roles.ToList()); err != nil {
		return err
	}
	if err := setStrToRD(d, "timezone", user.Timezone); err != nil {
		return err
	}
	if err := setIntToRD(d, "session_timeout_ms", user.SessionTimeoutMs); err != nil {
		return err
	}
	if err := setStrToRD(d, "user_id", user.ID); err != nil {
		return err
	}
	if err := setBoolToRD(d, "external", user.External); err != nil {
		return err
	}
	if err := setBoolToRD(d, "read_only", user.ReadOnly); err != nil {
		return err
	}
	if err := setStrToRD(d, "client_address", user.ClientAddress); err != nil {
		return err
	}
	if err := setBoolToRD(d, "session_active", user.SessionActive); err != nil {
		return err
	}
	return setStrToRD(d, "last_activity", user.LastActivity)
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	user := newUser(d)
	_, err = cl.UpdateUser(ctx, user.NewUpdateParams())
	return err
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	if _, err := cl.DeleteUser(ctx, d.Id()); err != nil {
		return err
	}
	return nil
}
