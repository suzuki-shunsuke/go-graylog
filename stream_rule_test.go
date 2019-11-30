package graylog_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil/v8"
)

func TestStreamRuleNewUpdateParams(t *testing.T) {
	rule := testutil.StreamRule()
	prms := rule.NewUpdateParams()
	if rule.ID != prms.ID {
		t.Fatalf(`prms.ID = "%s", wanted "%s"`, prms.ID, rule.ID)
	}
}
