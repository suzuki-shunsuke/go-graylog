package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/client/endpoint"
)

func TestStreamRules(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/streams/%s/rules", apiURL, ID)
	act, err := ep.StreamRules(ID)
	if err != nil {
		t.Fatal(err)
	}
	if act.String() != exp {
		t.Fatalf(`ep.StreamRules("%s") = "%s", wanted "%s"`, ID, act.String(), exp)
	}
}

func TestStreamRuleTypes(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/streams/%s/rules/types", apiURL, ID)
	act, err := ep.StreamRuleTypes(ID)
	if err != nil {
		t.Fatal(err)
	}
	if act.String() != exp {
		t.Fatalf(`ep.StreamRuleTypes("%s") = "%s", wanted "%s"`, ID, act.String(), exp)
	}
}

func TestStreamRule(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/streams/%s/rules/%s", apiURL, ID, ID)
	act, err := ep.StreamRule(ID, ID)
	if err != nil {
		t.Fatal(err)
	}
	if act.String() != exp {
		t.Fatalf(`ep.StreamRule("%s", "%s") = "%s", wanted "%s"`, ID, ID, act.String(), exp)
	}
}
