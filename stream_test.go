package graylog_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil/v8"
)

func TestStreamNewUpdateParams(t *testing.T) {
	stream := testutil.Stream()
	prms := stream.NewUpdateParams()
	if stream.ID != prms.ID {
		t.Fatalf(`prms.ID = "%s", wanted "%s"`, prms.ID, stream.ID)
	}
}
