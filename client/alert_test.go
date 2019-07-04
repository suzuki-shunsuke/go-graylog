package client_test

import (
	"context"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetAlerts(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	_, _, _, err = client.GetAlerts(ctx, 0, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAlert(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, _, err := client.GetAlert(ctx, ""); err == nil {
		t.Fatal("alert id is required")
	}
	if _, _, err := client.GetAlert(ctx, "h"); err == nil {
		t.Fatal("alert should not be found")
	}
}
