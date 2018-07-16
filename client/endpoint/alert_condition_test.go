package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/client/endpoint"
)

func TestAlertConditions(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	act := ep.AlertConditions()
	exp := fmt.Sprintf("%s/%s", apiURL, "alerts/conditions")
	if act != exp {
		t.Fatalf(`ep.AlertConditions() = "%s", wanted "%s"`, act, exp)
	}
}
