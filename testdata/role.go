package testdata

import (
	"github.com/suzuki-shunsuke/go-set/v6"

	"github.com/suzuki-shunsuke/go-graylog/v10"
)

var (
	Role = graylog.Role{
		Name:        "foo",
		Description: "Allows reading and writing all views and extended searches (built-in)",
		Permissions: set.NewStrSet(
			"extendedsearch:create",
			"extendedsearch:use",
			"view:create",
			"view:edit",
			"view:read",
			"view:use",
		),
		ReadOnly: true,
	}
)
