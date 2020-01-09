package graylog

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-set/v6"

	"github.com/suzuki-shunsuke/go-graylog/v8"
)

const (
	ldapSettingID = "ldap_setting_id"
)

func resourceLDAPSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLDAPSettingCreate,
		Read:   resourceLDAPSettingRead,
		Update: resourceLDAPSettingUpdate,
		Delete: resourceLDAPSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// required
			"system_username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ldap_uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"search_base": {
				Type:     schema.TypeString,
				Required: true,
			},
			"search_pattern": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name_attribute": {
				Type:     schema.TypeString,
				Required: true,
			},
			"default_group": {
				Type:     schema.TypeString,
				Required: true,
			},

			// optional
			// system_password is required to create the resource
			"system_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_start_tls": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"trust_all_certificates": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"active_directory": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"group_search_base": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_id_attribute": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_search_pattern": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_mapping": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"additional_default_groups": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func newLDAPSetting(d *schema.ResourceData) (*graylog.LDAPSetting, error) {
	setting := &graylog.LDAPSetting{
		Enabled:                 d.Get("enabled").(bool),
		UseStartTLS:             d.Get("use_start_tls").(bool),
		TrustAllCertificates:    d.Get("trust_all_certificates").(bool),
		ActiveDirectory:         d.Get("active_directory").(bool),
		SystemUsername:          d.Get("system_username").(string),
		SystemPassword:          d.Get("system_password").(string),
		LDAPURI:                 d.Get("ldap_uri").(string),
		SearchBase:              d.Get("search_base").(string),
		SearchPattern:           d.Get("search_pattern").(string),
		DisplayNameAttribute:    d.Get("display_name_attribute").(string),
		DefaultGroup:            d.Get("default_group").(string),
		GroupSearchBase:         d.Get("group_search_base").(string),
		GroupIDAttribute:        d.Get("group_id_attribute").(string),
		GroupSearchPattern:      d.Get("group_search_pattern").(string),
		AdditionalDefaultGroups: set.NewStrSet(getStringArray(d.Get("additional_default_groups").(*schema.Set).List())...),
	}
	mapping := map[string]string{}
	for k, v := range d.Get("group_mapping").(map[string]interface{}) {
		s, ok := v.(string)
		if !ok {
			return nil, errors.New("group_mapping's value must be string")
		}
		mapping[k] = s
	}
	setting.GroupMapping = mapping
	return setting, nil
}

func resourceLDAPSettingCreate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}

	ls, err := newLDAPSetting(d)
	if err != nil {
		return err
	}
	if _, err = cl.UpdateLDAPSetting(ctx, ls); err != nil {
		return err
	}
	d.SetId(ldapSettingID)
	return nil
}

func resourceLDAPSettingRead(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	ls, _, err := cl.GetLDAPSetting(ctx)
	if err != nil {
		return err
	}

	if err := setBoolToRD(d, "enabled", ls.Enabled); err != nil {
		return err
	}
	if err := setBoolToRD(d, "use_start_tls", ls.UseStartTLS); err != nil {
		return err
	}
	if err := setBoolToRD(d, "trust_all_certificates", ls.TrustAllCertificates); err != nil {
		return err
	}
	if err := setBoolToRD(d, "active_directory", ls.ActiveDirectory); err != nil {
		return err
	}
	if err := setStrToRD(d, "system_username", ls.SystemUsername); err != nil {
		return err
	}
	if err := setStrToRD(d, "system_password", ls.SystemPassword); err != nil {
		return err
	}
	if err := setStrToRD(d, "ldap_uri", ls.LDAPURI); err != nil {
		return err
	}
	if err := setStrToRD(d, "search_base", ls.SearchBase); err != nil {
		return err
	}
	if err := setStrToRD(d, "search_pattern", ls.SearchPattern); err != nil {
		return err
	}
	if err := setStrToRD(d, "display_name_attribute", ls.DisplayNameAttribute); err != nil {
		return err
	}
	if err := setStrToRD(d, "default_group", ls.DefaultGroup); err != nil {
		return err
	}
	if err := setStrToRD(d, "group_search_base", ls.GroupSearchBase); err != nil {
		return err
	}
	if err := setStrToRD(d, "group_id_attribute", ls.GroupIDAttribute); err != nil {
		return err
	}
	if err := setStrToRD(d, "group_search_pattern", ls.GroupSearchPattern); err != nil {
		return err
	}
	if err := setStrListToRD(d, "additional_default_groups", ls.AdditionalDefaultGroups.ToList()); err != nil {
		return err
	}
	return setMapStrToStrToRD(d, "group_mapping", ls.GroupMapping)
}

func resourceLDAPSettingUpdate(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	ls, err := newLDAPSetting(d)
	if err != nil {
		return err
	}
	_, err = cl.UpdateLDAPSetting(ctx, ls)
	return err
}

func resourceLDAPSettingDelete(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := newClient(m)
	if err != nil {
		return err
	}
	_, err = cl.DeleteLDAPSetting(ctx)
	return err
}
