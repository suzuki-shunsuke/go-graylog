package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/client/endpoint"
)

func TestStreams(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/streams", apiURL)
	act := ep.Streams()
	if act != exp {
		t.Fatalf(`ep.Streams() = "%s", wanted "%s"`, act, exp)
	}
}

func TestStream(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/streams/%s", apiURL, ID)
	act, err := ep.Stream(ID)
	if err != nil {
		t.Fatal(err)
	}
	if act.String() != exp {
		t.Fatalf(`ep.Stream("%s") = "%s", wanted "%s"`, ID, act.String(), exp)
	}
}

func TestPauseStream(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/streams/%s/pause", apiURL, ID)
	act, err := ep.PauseStream(ID)
	if err != nil {
		t.Fatal(err)
	}
	if act.String() != exp {
		t.Fatalf(`ep.PauseStream("%s") = "%s", wanted "%s"`, ID, act.String(), exp)
	}
}

func TestResumeStream(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/streams/%s/resume", apiURL, ID)
	act, err := ep.ResumeStream(ID)
	if err != nil {
		t.Fatal(err)
	}
	if act.String() != exp {
		t.Fatalf(`ep.ResumeStream("%s") = "%s", wanted "%s"`, ID, act.String(), exp)
	}
}

func TestEnabledStreams(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/streams/enabled", apiURL)
	act := ep.EnabledStreams()
	if act != exp {
		t.Fatalf(`ep.EnabledStreams() = "%s", wanted "%s"`, act, exp)
	}
}
