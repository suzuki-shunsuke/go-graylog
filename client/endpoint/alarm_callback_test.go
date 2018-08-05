package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/client/endpoint"
)

func TestAlarmCallbacks(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s", apiURL, "alerts/callbacks")
	act := ep.AlarmCallbacks()
	if act != exp {
		t.Fatalf(`ep.AlarmCallbacks() = "%s", wanted "%s"`, act, exp)
	}
}
