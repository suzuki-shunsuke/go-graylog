package graylog

import (
	"errors"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-jsoneq/jsoneq"

	"github.com/suzuki-shunsuke/go-graylog/v9/client"
)

func schemaDiffSuppressJSONString(k, oldV, newV string, d *schema.ResourceData) bool {
	b, err := jsoneq.Equal([]byte(oldV), []byte(newV))
	if err != nil {
		return false
	}
	return b
}

func wrapValidateFunc(f func(v interface{}, k string) error) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (s []string, es []error) {
		if err := f(v, k); err != nil {
			es = append(es, err)
		}
		return
	}
}

func genImport(keys ...string) schema.StateFunc {
	return func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
		a := strings.Split(d.Id(), "/")
		size := len(keys)
		if len(a) != size {
			return nil, errors.New("format of import argument should be " + strings.Join(keys, "/"))
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

func handleGetResourceError(
	d *schema.ResourceData, ei *client.ErrorInfo, err error,
) error {
	if ei != nil && ei.Response != nil && ei.Response.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	return err
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

// func setEnv() (*client.Client, error) {
// 	_, ok := os.LookupEnv("TF_ACC")
// 	if !ok {
// 		if err := os.Setenv("TF_ACC", "true"); err != nil {
// 			return nil, err
// 		}
// 	}
// 	authName, ok := os.LookupEnv("GRAYLOG_AUTH_NAME")
// 	if !ok {
// 		authName = "admin"
// 		if err := os.Setenv("GRAYLOG_AUTH_NAME", authName); err != nil {
// 			return nil, err
// 		}
// 	}
// 	authPass, ok := os.LookupEnv("GRAYLOG_AUTH_PASSWORD")
// 	if !ok {
// 		authPass = "admin"
// 		if err := os.Setenv("GRAYLOG_AUTH_PASSWORD", "admin"); err != nil {
// 			return nil, err
// 		}
// 	}
// 	var (
// 		err error
// 		cl  *client.Client
// 	)
// 	endpoint := os.Getenv("GRAYLOG_WEB_ENDPOINT_URI")
// 	if endpoint == "" {
// 		return nil, errors.New("GRAYLOG_WEB_ENDPOINT_URI is required")
// 	}
// 	if os.Getenv("GRAYLOG_API_VERSION") == "v3" {
// 		cl, err = client.NewClientV3(endpoint, authName, authPass)
// 	} else {
// 		cl, err = client.NewClient(endpoint, authName, authPass)
// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	return cl, nil
// }

// func getIDFromTfState(tfState *terraform.State, key string) (string, error) {
// 	rs, ok := tfState.RootModule().Resources[key]
// 	if !ok {
// 		return "", errors.New("not found: " + key)
// 	}
// 	id := rs.Primary.ID
// 	if id == "" {
// 		return "", errors.New("no ID is set")
// 	}
// 	return id, nil
// }

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

func hasChange(d *schema.ResourceData, keys ...string) bool {
	for _, k := range keys {
		if d.HasChange(k) {
			return true
		}
	}
	return false
}
