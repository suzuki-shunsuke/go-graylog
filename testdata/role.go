package testdata

import (
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog/v8"
)

var (
	Role = &graylog.Role{
		Name:        "Views Manager",
		Description: "Allows reading and writing all views and extended searches (built-in)",
		Permissions: set.StrSet{
			"extendedsearch:create": struct{}{},
			"extendedsearch:use":    struct{}{},
			"view:create":           struct{}{},
			"view:edit":             struct{}{},
			"view:read":             struct{}{},
			"view:use":              struct{}{},
		},
		ReadOnly: true,
	}
)
