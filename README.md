# go-graylog

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/go-graylog)
[![Build Status](https://travis-ci.org/suzuki-shunsuke/go-graylog.svg?branch=master)](https://travis-ci.org/suzuki-shunsuke/go-graylog)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/go-graylog/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/go-graylog)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-graylog.svg)](https://github.com/suzuki-shunsuke/go-graylog)
[![GitHub tag](https://img.shields.io/github/tag/suzuki-shunsuke/go-graylog.svg)](https://github.com/suzuki-shunsuke/go-graylog/releases)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-graylog/master/LICENSE)

Graylog API client and mock server for golang.

## Example 1 - Create and update a role

```golang
package main

import (
	"fmt"
	"os"

	"github.com/suzuki-shunsuke/go-graylog"
)

func main() {
	client, err := graylog.NewClient("http://localhost:9000/api", "admin", "admin")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to NewClient", err)
		os.Exit(1)
	}
	params := &graylog.Role{
		Name: "foo", Description: "description",
		Permissions: []string{"users:list"}}
	role, ei, err := client.CreateRole(params)
	fmt.Println(role, ei, err)

	params.Name = "bar"
	role, ei, err = client.UpdateRole("foo", params)
	fmt.Println(role, ei, err)
}
```

## Example 2 - Run a mock server in the program

```golang
server, err := NewMockServer("")
if err != nil {
	t.Error("Failed to Get Mock Server", err)
	return
}
server.Start()
defer server.Close()
client, err := NewClient(server.Endpoint, "admin", "password")
if err != nil {
	t.Error("Failed to NewClient", err)
	return
}
params := &Role{Name: "foo", Permissions: []string{"*"}}
role, ei, err := client.CreateRole(params)
```

## Mock Server CLI tool

Download a binary from [the release page](https://github.com/suzuki-shunsuke/go-graylog/releases).

```
$ graylog-mock-server --help
graylog-mock-server - Run Graylog mock server.

USAGE:
   graylog-mock-server [options]

VERSION:
   0.1.0

OPTIONS:
   --port value       port number. If you don't set this option, a free port is assigned and the assigned port number is outputed to the console when the mock server runs.
   --log-level value  the log level of logrus which the mock server uses internally. (default: "info")
   --data value       data file path. When the server runs data of the file is loaded and when data of the server is changed data is saved at the file. If this option is not set, no data is loaded and saved.
   --help, -h         show help
   --version, -v      print the version
```

## Graylog REST API's Reference

* http://docs.graylog.org/en/2.4/pages/configuration/rest_api.html
* http://docs.graylog.org/en/2.4/pages/users_and_roles/permission_system.html

## Supported Graylog version

We use [the graylog's official Docker Image](https://hub.docker.com/r/graylog/graylog/) .

The version is `2.4.0-1` .

## Contribution

See [CONTRIBUTING.md](CONTRIBUTING.md) .

## See also

* [terraform-provider-graylog](https://github.com/suzuki-shunsuke/terraform-provider-graylog)

## License

[MIT](LICENSE)
