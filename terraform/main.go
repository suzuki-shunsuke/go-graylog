package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"

	"github.com/suzuki-shunsuke/go-graylog/terraform/graylog/v8"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return graylog.Provider()
		},
	})
}
