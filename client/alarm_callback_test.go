package client_test

import (
	"context"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil/v8"
)

func TestClient_GetAlarmCallbacks(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	_, _, _, err = client.GetAlarmCallbacks(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
