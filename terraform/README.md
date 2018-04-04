# terraform-provider-graylog

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/terraform-provider-graylog)
[![Build Status](https://travis-ci.org/suzuki-shunsuke/terraform-provider-graylog.svg?branch=master)](https://travis-ci.org/suzuki-shunsuke/terraform-provider-graylog)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/terraform-provider-graylog/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/terraform-provider-graylog)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/terraform-provider-graylog.svg)](https://github.com/suzuki-shunsuke/terraform-provider-graylog)
[![GitHub tag](https://img.shields.io/github/tag/suzuki-shunsuke/terraform-provider-graylog.svg)](https://github.com/suzuki-shunsuke/terraform-provider-graylog/releases)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/terraform-provider-graylog/master/LICENSE)

terraform provider for graylog

## Status

This provider is still in beta.

## Motivation

http://docs.graylog.org/en/2.4/pages/users_and_roles/permission_system.html

The Graylog permission system is extremely flexible but you can't utilize this flexibility from Web UI.
By using this provider, you can utilize this flexibility and manage the infrastructure as code.

## Install

```
$ go get github.com/suzuki-shunsuke/terraform-provider-graylog
```

Or download binary.

```
$ curl 
```

## Example

```
// Role my-role-2
resource "graylog_role" "my-role-2" {
  name = "my-role-2"
  permissions = ["users:edit"]
  description = "Created by terraform"
}
```

## Variables

name | Environment variable | description
--- | --- | ---
web_endpoint_uri | GRAYLOG_WEB_ENDPOINT_URI |
auth_name | GRAYLOG_AUTH_NAME |
auth_password | GRAYLOG_AUTH_PASSWORD |

## Resources

* [role](docs/role.md)
* [user](docs/user.md)
* [input](docs/input.md)
* [index_set](docs/index_set.md)
* [stream](docs/stream.md)

## Supported Graylog version

We use [the graylog's official Docker Image](https://hub.docker.com/r/graylog/graylog/) .

The version is `2.4.0-1` .

## Contribution

See [CONTRIBUTING.md](CONTRIBUTING.md) .

## See also

* [go-graylog](https://github.com/suzuki-shunsuke/go-graylog): Graylog API client and simple mock server for golang

## License

[MIT](LICENSE)
