package graylog

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hashicorp/terraform/terraform"
	"github.com/pkg/errors"
	"github.com/suzuki-shunsuke/go-graylog/client"
	"github.com/suzuki-shunsuke/go-graylog/mockserver"
)

func getStrInt(cfg map[string]interface{}, key string, required bool) (int, error) {
	val, ok := cfg[key]
	if !ok {
		if required {
			return 0, fmt.Errorf("%s is required", key)
		}
		return 0, nil
	}
	v, ok := val.(string)
	if !ok {
		return 0, fmt.Errorf("Failed to convert %s to string: %v", key, val)
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, errors.Wrap(
			err, fmt.Sprintf("%s must be int: %v", key, v))
	}
	return i, nil
}

func getInt(
	cfg map[string]interface{}, key string, required bool,
) (int, error) {
	val, ok := cfg[key]
	if !ok {
		if required {
			return 0, fmt.Errorf("%s is required", key)
		}
		return 0, nil
	}
	v, ok := val.(int)
	if !ok {
		return 0, fmt.Errorf("%s must be int: %v", key, val)
	}
	return v, nil
}

func getString(
	cfg map[string]interface{}, key string, required bool,
) (string, error) {
	val, ok := cfg[key]
	if !ok {
		if required {
			return "", fmt.Errorf("%s is required", key)
		}
		return "", nil
	}
	v, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("%s must be string: %v", key, val)
	}
	return v, nil
}

func getBool(
	cfg map[string]interface{}, key string, required bool,
) (bool, error) {
	val, ok := cfg[key]
	if !ok {
		if required {
			return false, fmt.Errorf("%s is required", key)
		}
		return false, nil
	}
	v, ok := val.(bool)
	if !ok {
		return false, fmt.Errorf("%s must be bool: %v", key, val)
	}
	return v, nil
}

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

type tfConf struct {
	Resource map[string]map[string]interface{} `json:"resource"`
}
