package graylog

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider returns a terraform resource provider for graylog.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"web_endpoint_uri": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"GRAYLOG_WEB_ENDPOINT_URI"}, nil),
			},
			"auth_name": {
				Type:     schema.TypeString,
				Required: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"GRAYLOG_AUTH_NAME"}, nil),
			},
			"auth_password": {
				Type:     schema.TypeString,
				Required: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"GRAYLOG_AUTH_PASSWORD"}, nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"graylog_role":      resourceRole(),
			"graylog_user":      resourceUser(),
			"graylog_input":     resourceInput(),
			"graylog_index_set": resourceIndexSet(),
			"graylog_stream":    resourceStream(),
			"graylog_dashboard": resourceDashboard(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("web_endpoint_uri").(string)
	authName := d.Get("auth_name").(string)
	authPassword := d.Get("auth_password").(string)
	config := Config{
		Endpoint:     endpoint,
		AuthName:     authName,
		AuthPassword: authPassword,
	}

	if err := config.loadAndValidate(); err != nil {
		return nil, err
	}

	return &config, nil
}
