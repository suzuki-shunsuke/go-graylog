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
			"x_requested_by": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"GRAYLOG_X_REQUESTED_BY"}, "terraform-go-graylog"),
			},
			"api_version": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"GRAYLOG_API_VERSION"}, "v2"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"graylog_alert_condition":     resourceAlertCondition(),
			"graylog_alarm_callback":      resourceAlarmCallback(),
			"graylog_dashboard":           resourceDashboard(),
			"graylog_extractor":           resourceExtractor(),
			"graylog_index_set":           resourceIndexSet(),
			"graylog_input":               resourceInput(),
			"graylog_ldap_setting":        resourceLDAPSetting(),
			"graylog_pipeline":            resourcePipeline(),
			"graylog_pipeline_rule":       resourcePipelineRule(),
			"graylog_pipeline_connection": resourcePipelineConnection(),
			"graylog_role":                resourceRole(),
			"graylog_stream":              resourceStream(),
			"graylog_stream_rule":         resourceStreamRule(),
			"graylog_user":                resourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"graylog_index_set": dataSourceIndexSet(),
			"graylog_stream":    dataSourceStream(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Endpoint:     d.Get("web_endpoint_uri").(string),
		AuthName:     d.Get("auth_name").(string),
		AuthPassword: d.Get("auth_password").(string),
		XRequestedBy: d.Get("x_requested_by").(string),
		APIVersion:   d.Get("api_version").(string),
	}

	if err := config.loadAndValidate(); err != nil {
		return nil, err
	}

	return &config, nil
}
