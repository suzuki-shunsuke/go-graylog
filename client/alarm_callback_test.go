package client_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetAlarmCallbacks(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	_, _, _, err = client.GetAlarmCallbacks()
	if err != nil {
		t.Fatal(err)
	}
}
