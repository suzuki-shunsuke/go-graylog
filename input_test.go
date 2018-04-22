package graylog_test

import (
	"encoding/json"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestInputUnmarshalJSON(t *testing.T) {
	input := testutil.Input()
	attrs := input.Attributes.(*graylog.InputBeatsAttrs)
	if attrs.BindAddress == "" {
		t.Fatal(`attrs.BindAddress == ""`)
	}
	if err := json.Unmarshal([]byte(`{"id": "fooo"}`), input); err != nil {
		t.Fatal(err)
	}
	attrs = input.Attributes.(*graylog.InputBeatsAttrs)
	if attrs.BindAddress == "" {
		t.Fatal(`attrs.BindAddress == ""`)
	}
}
