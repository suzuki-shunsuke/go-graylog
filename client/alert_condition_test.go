package client_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetAlertConditions(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	_, _, _, err = client.GetAlertConditions()
	if err != nil {
		t.Fatal(err)
	}
}
