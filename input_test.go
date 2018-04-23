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

func TestNewInputAttrs(t *testing.T) {
	attrs, err := graylog.NewInputAttrs("hoge")
	if err != nil {
		t.Fatal(err)
	}
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

func TestInputUpdatePramsDataToInput(t *testing.T) {
	data := graylog.InputUpdateParamsData{}
	prms := &graylog.InputUpdateParams{}
	if err := data.ToInput(prms); err != nil {
		t.Fatal(err)
	}
	data = graylog.InputUpdateParamsData{
		Type: graylog.InputTypeBeats,
	}
	if err := data.ToInput(prms); err != nil {
		t.Fatal(err)
	}
}

func TestInputDataToInput(t *testing.T) {
	input := &graylog.Input{}
	data := &graylog.InputData{Type: "hoge"}
	if err := data.ToInput(input); err != nil {
		t.Fatal(err)
	}
}
