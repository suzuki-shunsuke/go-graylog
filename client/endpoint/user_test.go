package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/client/endpoint"
)

func TestUsers(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s", apiURL, "users")
	if ep.Users() != exp {
		t.Fatalf(`ep.Users() = "%s", wanted "%s"`, ep.Users(), exp)
	}
}

func TestUser(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s", apiURL, "users/foo")
	act, err := ep.User("foo")
	if err != nil {
		t.Fatal(err)
	}
	if act.String() != exp {
		t.Fatalf(`ep.User("foo") = "%s", wanted "%s"`, act.String(), exp)
	}
}
