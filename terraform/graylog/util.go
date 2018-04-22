package graylog

import (
	"fmt"
	"os"

	"github.com/hashicorp/terraform/terraform"
	"github.com/suzuki-shunsuke/go-graylog/client"
	"github.com/suzuki-shunsuke/go-graylog/mockserver"
)

func getStringArray(src []interface{}) []string {
	dest := make([]string, len(src))
	for i, p := range src {
		dest[i] = p.(string)
	}
	return dest
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
	)
	endpoint, ok := os.LookupEnv("GRAYLOG_WEB_ENDPOINT_URI")
	if !ok {
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
	cl, err := client.NewClient(endpoint, authName, authPass)
	if err != nil {
		return nil, nil, err
	}
	return cl, server, nil
}

func getIDFromTfState(tfState *terraform.State, key string) (string, error) {
	rs, ok := tfState.RootModule().Resources[key]
	if !ok {
		return "", fmt.Errorf("Not found: %s", key)
	}
	id := rs.Primary.ID
	if id == "" {
		return "", fmt.Errorf("No ID is set")
	}
	return id, nil
}
