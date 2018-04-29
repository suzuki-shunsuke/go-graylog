package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestHandleNotFound(t *testing.T) {
	server, _, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	endpoint := fmt.Sprintf("%s/dummy", server.Endpoint())
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatal(err)
	}
	hc := &http.Client{}
	if _, err := hc.Do(req); err != nil {
		t.Fatal(err)
	}
}
