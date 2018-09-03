package graylog

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
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
			"system_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"system_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ldap_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"search_base": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"search_pattern": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name_attribute": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_group": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
		},
	}
}

func newLDAPSetting(d *schema.ResourceData) *graylog.LDAPSetting {
	return &graylog.LDAPSetting{
		Enabled:              d.Get("enabled").(bool),
		UseStartTLS:          d.Get("use_start_tls").(bool),
		TrustAllCertificates: d.Get("trust_all_certificates").(bool),
		ActiveDirectory:      d.Get("active_directory").(bool),
		SystemUsername:       d.Get("system_username").(string),
		SystemPassword:       d.Get("system_password").(string),
		LDAPURI:              d.Get("ldap_uri").(string),
		SearchBase:           d.Get("search_base").(string),
		SearchPattern:        d.Get("search_pattern").(string),
		DisplayNameAttribute: d.Get("display_name_attribute").(string),
		DefaultGroup:         d.Get("default_group").(string),
		GroupSearchBase:      d.Get("group_search_base").(string),
		GroupIDAttribute:     d.Get("group_id_attribute").(string),
		GroupSearchPattern:   d.Get("group_search_pattern").(string),
	}
}

func resourceLDAPSettingCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}

	ls := newLDAPSetting(d)
	if _, err = cl.UpdateLDAPSetting(ls.NewUpdateParams()); err != nil {
		return err
	}
	d.SetId(ldapSettingID)
	return nil
}

func resourceLDAPSettingRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	ls, _, err := cl.GetLDAPSetting()
	if err != nil {
		return err
	}

	setBoolToRD(d, "enabled", ls.Enabled)
	setBoolToRD(d, "use_start_tls", ls.UseStartTLS)
	setBoolToRD(d, "trust_all_certificates", ls.TrustAllCertificates)
	setBoolToRD(d, "active_directory", ls.ActiveDirectory)
	setStrToRD(d, "system_username", ls.SystemUsername)
	setStrToRD(d, "system_password", ls.SystemPassword)
	setStrToRD(d, "ldap_uri", ls.LDAPURI)
	setStrToRD(d, "search_base", ls.SearchBase)
	setStrToRD(d, "search_pattern", ls.SearchPattern)
	setStrToRD(d, "display_name_attribute", ls.DisplayNameAttribute)
	setStrToRD(d, "default_group", ls.DefaultGroup)
	setStrToRD(d, "group_search_base", ls.GroupSearchBase)
	setStrToRD(d, "group_id_attribute", ls.GroupIDAttribute)
	setStrToRD(d, "group_search_pattern", ls.GroupSearchPattern)

	return nil
}

func resourceLDAPSettingUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	ls := newLDAPSetting(d)
	_, err = cl.UpdateLDAPSetting(ls.NewUpdateParams())
	return err
}

func resourceLDAPSettingDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	cl, err := client.NewClient(
		config.Endpoint, config.AuthName, config.AuthPassword)
	if err != nil {
		return err
	}
	_, err = cl.DeleteLDAPSetting()
	return err
}
