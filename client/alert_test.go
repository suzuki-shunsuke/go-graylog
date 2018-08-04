package client_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetAlerts(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	_, _, _, err = client.GetAlerts(0, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAlert(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, _, err := client.GetAlert(""); err == nil {
		t.Fatal("alert id is required")
	}
	if _, _, err := client.GetAlert("h"); err == nil {
		t.Fatal("alert should not be found")
	}
}
