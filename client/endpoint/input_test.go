package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/client/endpoint"
)

func TestInputs(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s", apiURL, "system/inputs")
	act := ep.Inputs()
	if act != exp {
		t.Fatalf(`ep.Inputs() = "%s", wanted "%s"`, act, exp)
	}
}

func TestInput(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s/%s", apiURL, "system/inputs", ID)
	act, err := ep.Input(ID)
	if err != nil {
		t.Fatal(err)
	}
	if act.String() != exp {
		t.Fatalf(`ep.Input("%s") = "%s", wanted "%s"`, ID, act.String(), exp)
	}
}
