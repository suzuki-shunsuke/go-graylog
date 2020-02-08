package graylog_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/testdata"
)

func TestUserNewUpdateParams(t *testing.T) {
	user := testdata.User()
	prms := user.NewUpdateParams()
	if user.Username != prms.Username {
		t.Fatalf(`prms.Username = "%s", wanted "%s"`, prms.Username, user.Username)
	}
}

func TestUserSetDefaultValues(t *testing.T) {
	user := &graylog.User{}
	user.SetDefaultValues()
	if user.SessionTimeoutMs == 0 {
		t.Fatal("user.SessionTimeoutMs must be set")
	}
	if user.Timezone == "" {
		t.Fatal("user.Timezone must be set")
	}
}
