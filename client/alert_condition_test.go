package client_test

import (
	"context"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestClient_GetAlertConditions(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	_, _, _, err = client.GetAlertConditions(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
