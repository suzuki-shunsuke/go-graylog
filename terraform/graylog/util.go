package graylog

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/suzuki-shunsuke/go-graylog/client"
	"github.com/suzuki-shunsuke/go-graylog/mockserver"
)

func genImport(keys ...string) schema.StateFunc {
	return func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
		a := strings.Split(d.Id(), "/")
		size := len(keys)
		if len(a) != size {
			return nil, fmt.Errorf("format of import argument should be %s", strings.Join(keys, "/"))
		}
		for i, k := range keys[:size-1] {
			if err := setStrToRD(d, k, a[i]); err != nil {
				return nil, err
			}
		}
		d.SetId(a[size-1])
		return []*schema.ResourceData{d}, nil
	}
}

func getStringArray(src []interface{}) []string {
	dest := make([]string, len(src))
	for i, p := range src {
		dest[i] = p.(string)
	}
	return dest
}

func newClient(m interface{}) (*client.Client, error) {
	config := m.(*Config)
	var (
		cl  *client.Client
		err error
	)
	if config.APIVersion == "v3" {
		cl, err = client.NewClientV3(
			config.Endpoint, config.AuthName, config.AuthPassword)
	} else {
		cl, err = client.NewClient(
			config.Endpoint, config.AuthName, config.AuthPassword)
	}
	if err != nil {
		return cl, err
	}
	if config.XRequestedBy != "" {
		cl.SetXRequestedBy(config.XRequestedBy)
	}
	return cl, nil
}

func setEnv() (*client.Client, *mockserver.Server, error) {
	_, ok := os.LookupEnv("TF_ACC")
	if !ok {
		if err := os.Setenv("TF_ACC", "true"); err != nil {
			return nil, nil, err
		}
	}
	authName, ok := os.LookupEnv("GRAYLOG_AUTH_NAME")
	if !ok {
		authName = "admin"
		if err := os.Setenv("GRAYLOG_AUTH_NAME", authName); err != nil {
			return nil, nil, err
		}
	}
	authPass, ok := os.LookupEnv("GRAYLOG_AUTH_PASSWORD")
	if !ok {
		authPass = "admin"
		if err := os.Setenv("GRAYLOG_AUTH_PASSWORD", "admin"); err != nil {
			return nil, nil, err
		}
	}
	var (
		server *mockserver.Server
		err    error
		cl     *client.Client
	)
	endpoint := os.Getenv("GRAYLOG_WEB_ENDPOINT_URI")
	if endpoint == "" {
		server, err = mockserver.NewServer("", nil)
		if err != nil {
			return nil, nil, err
		}
		server.SetAuth(true)
		endpoint = server.Endpoint()
		if err := os.Setenv("GRAYLOG_WEB_ENDPOINT_URI", endpoint); err != nil {
			return nil, nil, err
		}
	}
	if os.Getenv("GRAYLOG_API_VERSION") == "v3" {
		cl, err = client.NewClientV3(endpoint, authName, authPass)
	} else {
		cl, err = client.NewClient(endpoint, authName, authPass)
	}
	if err != nil {
		return nil, nil, err
	}
	return cl, server, nil
}

func getIDFromTfState(tfState *terraform.State, key string) (string, error) {
	rs, ok := tfState.RootModule().Resources[key]
	if !ok {
		return "", fmt.Errorf("not found: %s", key)
	}
	id := rs.Primary.ID
	if id == "" {
		return "", fmt.Errorf("no ID is set")
	}
	return id, nil
}

func setStrListToRD(d *schema.ResourceData, key string, val []string) error {
	return d.Set(key, val)
}

func setMapStrToStrToRD(d *schema.ResourceData, key string, val map[string]string) error {
	return d.Set(key, val)
}

func setStrToRD(d *schema.ResourceData, key, val string) error {
	return d.Set(key, val)
}

func setIntToRD(d *schema.ResourceData, key string, val int) error {
	return d.Set(key, val)
}

func setBoolToRD(d *schema.ResourceData, key string, val bool) error {
	return d.Set(key, val)
}
