package graylog

import (
	"errors"
	"os"
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
	d *schema.ResourceData, ei *client.ErrorInfo, err error, codes ...int,
) error {
	if ei != nil && ei.Response != nil {
		if ei.Response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		for _, code := range codes {
			if ei.Response.StatusCode == code {
				d.SetId("")
				return nil
			}
		}
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

func setEnv() {
	os.Setenv("TF_ACC", "true")
	os.Setenv("GRAYLOG_WEB_ENDPOINT_URI", "http://example.com/api")
	os.Setenv("GRAYLOG_AUTH_NAME", "admin")
	os.Setenv("GRAYLOG_AUTH_PASSWORD", "admin")
	os.Setenv("GRAYLOG_API_VERSION", "v3")
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

func hasChange(d *schema.ResourceData, keys ...string) bool {
	for _, k := range keys {
		if d.HasChange(k) {
			return true
		}
	}
	return false
}
