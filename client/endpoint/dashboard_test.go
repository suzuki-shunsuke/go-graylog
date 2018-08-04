package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/client/endpoint"
)

func TestDashboards(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s", apiURL, "dashboards")
	act := ep.Dashboards()
	if act != exp {
		t.Fatalf(`ep.Dashboards() = "%s", wanted "%s"`, act, exp)
	}
}

func TestDashboard(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s/%s", apiURL, "dashboards", ID)
	act, err := ep.Dashboard(ID)
	if err != nil {
		t.Fatal(err)
	}
	if act.String() != exp {
		t.Fatalf(`ep.Dashboard("%s") = "%s", wanted "%s"`, ID, act.String(), exp)
	}
}
