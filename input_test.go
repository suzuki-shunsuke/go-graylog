package graylog_test

import (
	"encoding/json"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/v9/testdata"
)

func TestInputUnmarshalJSON(t *testing.T) {
	input := testdata.Input()
	if err := json.Unmarshal([]byte(`{"id": "fooo"}`), input); err != nil {
		t.Fatal(err)
	}
}

func TestInputNewUpdateParams(t *testing.T) {
	input := testdata.Input()
	prms := input.NewUpdateParams()
	if input.ID != prms.ID {
		t.Fatalf(`prms.ID = "%s", wanted "%s"`, prms.ID, input.ID)
	}
}
