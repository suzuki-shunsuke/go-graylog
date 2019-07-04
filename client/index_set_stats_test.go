package client_test

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetIndexSetStats(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	iss, _, _, _, err := client.GetIndexSets(ctx, 0, 0, false)
	if err != nil {
		t.Fatal(err)
	}
	u, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet(u.String())
	if len(iss) == 0 {
		if _, err := client.CreateIndexSet(ctx, is); err != nil {
			t.Fatal(err)
		}
		testutil.WaitAfterCreateIndexSet(server)
		// clean
		defer func(id string) {
			if _, err := client.DeleteIndexSet(ctx, id); err != nil {
				t.Fatal(err)
			}
			testutil.WaitAfterDeleteIndexSet(server)
		}(is.ID)
	} else {
		is = &(iss[0])
	}

	if _, _, err := client.GetIndexSetStats(ctx, is.ID); err != nil {
		t.Fatal(err)
	}
	if _, _, err := client.GetIndexSetStats(ctx, ""); err == nil {
		t.Fatal("index set id is required")
	}
	// if _, _, err := client.GetIndexSetStats("h"); err == nil {
	// 	t.Fatal(`no index set whose id is "h"`)
	// }
}

func TestGetTotalIndexSetsStats(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	u, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	is, f, err := testutil.GetIndexSet(ctx, client, server, u.String())
	if err != nil {
		t.Fatal(err)
	}
	if f != nil {
		defer f(is.ID)
	}
	if _, _, err := client.GetTotalIndexSetsStats(ctx); err != nil {
		t.Fatal(err)
	}
}
