package graylog_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/testdata"
)

func TestStreamRuleNewUpdateParams(t *testing.T) {
	rule := testdata.StreamRule()
	prms := rule.NewUpdateParams()
	if rule.ID != prms.ID {
		t.Fatalf(`prms.ID = "%s", wanted "%s"`, prms.ID, rule.ID)
	}
}
