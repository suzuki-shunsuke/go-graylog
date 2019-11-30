package graylog_test

import (
	"encoding/json"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/v8"
	"github.com/suzuki-shunsuke/go-graylog/testutil/v8"
)

func TestInputUnmarshalJSON(t *testing.T) {
	input := testutil.Input()
	attrs := input.Attrs.(*graylog.InputBeatsAttrs)
	if attrs.BindAddress == "" {
		t.Fatal(`attrs.BindAddress == ""`)
	}
	if err := json.Unmarshal([]byte(`{"id": "fooo"}`), input); err != nil {
		t.Fatal(err)
	}
	attrs = input.Attrs.(*graylog.InputBeatsAttrs)
	if attrs.BindAddress == "" {
		t.Fatal(`attrs.BindAddress == ""`)
	}
}

func TestInputNewUpdateParams(t *testing.T) {
	input := testutil.Input()
	prms := input.NewUpdateParams()
	if input.ID != prms.ID {
		t.Fatalf(`prms.ID = "%s", wanted "%s"`, prms.ID, input.ID)
	}
}
