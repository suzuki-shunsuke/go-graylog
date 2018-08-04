package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/client/endpoint"
)

func TestAlerts(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s", apiURL, "streams/alerts")
	act := ep.Alerts()
	if act != exp {
		t.Fatalf(`ep.Alerts() = "%s", wanted "%s"`, act, exp)
	}
}

func TestAlert(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s/%s", apiURL, "streams/alerts", ID)
	act, err := ep.Alert(ID)
	if err != nil {
		t.Fatal(err)
	}
	if act.String() != exp {
		t.Fatalf(`ep.Alert("%s") = "%s", wanted "%s"`, ID, act.String(), exp)
	}
}
