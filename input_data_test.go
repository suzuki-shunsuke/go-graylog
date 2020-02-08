package graylog_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/v10"
)

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
