package graylog_test

import (
	"encoding/json"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
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

func TestNewInputAttrsByType(t *testing.T) {
	attrs := graylog.NewInputAttrsByType("hoge")
	if attrs.InputType() != "hoge" {
		t.Fatalf(`attrs.InputType() = "%s", wanted "hoge"`, attrs.InputType())
	}
}

func TestInputNewUpdateParams(t *testing.T) {
	input := testutil.Input()
	prms := input.NewUpdateParams()
	if input.ID != prms.ID {
		t.Fatalf(`prms.ID = "%s", wanted "%s"`, prms.ID, input.ID)
	}
}

func TestInputUpdatePramsDataToInputUpdateParams(t *testing.T) {
	data := &graylog.InputUpdateParamsData{}
	prms := &graylog.InputUpdateParams{}
	if err := data.ToInputUpdateParams(prms); err != nil {
		t.Fatal(err)
	}
	data = &graylog.InputUpdateParamsData{
		Type: graylog.InputTypeBeats,
	}
	if err := data.ToInputUpdateParams(prms); err != nil {
		t.Fatal(err)
	}
}

func TestInputDataToInput(t *testing.T) {
	input := &graylog.Input{}
	data := &graylog.InputData{
		Type: "hoge",
		Attrs: map[string]interface{}{
			"bind_address": "0.0.0.0",
		}}
	// if err := data.ToInput(input); err != nil {
	// 	t.Fatal(err)
	// }
	data.Type = graylog.InputTypeBeats
	if err := data.ToInput(input); err != nil {
		t.Fatal(err)
	}
	attrs, ok := input.Attrs.(*graylog.InputBeatsAttrs)
	if !ok {
		t.Fatal("attrs must be beats attrs")
	}
	if attrs.BindAddress != "0.0.0.0" {
		t.Fatalf(`bind_address = "%s", wanted "%s"`, attrs.BindAddress, "0.0.0.0")
	}
}
