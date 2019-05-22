# terraform-provider-graylog

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/go-graylog/terraform)
[![Build Status](https://cloud.drone.io/api/badges/suzuki-shunsuke/go-graylog/status.svg)](https://cloud.drone.io/suzuki-shunsuke/go-graylog)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/go-graylog/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/go-graylog)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/go-graylog)](https://goreportcard.com/report/github.com/suzuki-shunsuke/go-graylog)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-graylog.svg)](https://github.com/suzuki-shunsuke/go-graylog)
[![GitHub tag](https://img.shields.io/github/tag/suzuki-shunsuke/go-graylog.svg)](https://github.com/suzuki-shunsuke/go-graylog/releases)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-graylog/master/LICENSE)

terraform provider for [Graylog](https://www.graylog.org/).

This is sub project of [go-graylog](https://github.com/suzuki-shunsuke/go-graylog).

## Motivation

http://docs.graylog.org/en/2.5/pages/users_and_roles/permission_system.html

The Graylog permission system is extremely flexible but you can't utilize this flexibility from Web UI.
By using this provider, you can utilize this flexibility and manage the infrastructure as code.

## Install

[Download binary](https://github.com/suzuki-shunsuke/go-graylog/releases) and install under `~/.terraform.d/plugins`.

https://www.terraform.io/docs/configuration/providers.html#third-party-plugins

```console
$ GO_GRAYLOG_VERSION=0.11.0
$ GO_GRAYLOG_ARCH=darwin_amd64
$ wget https://github.com/suzuki-shunsuke/go-graylog/releases/download/v${GO_GRAYLOG_VERSION}/terraform-provider-graylog_v${GO_GRAYLOG_VERSION}_${GO_GRAYLOG_ARCH}.gz
$ gzip -d terraform-provider-graylog_v${GO_GRAYLOG_VERSION}_${GO_GRAYLOG_ARCH}.gz
$ mkdir -p ~/.terraform.d/plugins
$ mv terraform-provider-graylog_v${GO_GRAYLOG_VERSION}_${GO_GRAYLOG_ARCH} ~/.terraform.d/plugins/terraform-provider-graylog_v${GO_GRAYLOG_VERSION}
$ chmod +x ~/.terraform.d/plugins/terraform-provider-graylog_v${GO_GRAYLOG_VERSION}
```

## Docker Image

https://hub.docker.com/r/suzukishunsuke/terraform-graylog/

Docker image which is installed terraform and terraform-provider-graylog on alpine.

## Example

```hcl
provider "graylog" {
  web_endpoint_uri = "${var.web_endpoint_uri}"
  auth_name = "${var.auth_name}"
  auth_password = "${var.auth_password}"
}

// Role my-role-2
resource "graylog_role" "my-role-2" {
  name = "my-role-2"
  permissions = ["users:edit"]
  description = "Created by terraform"
}
```

And please see https://github.com/suzuki-shunsuke/example/tree/master/graylog-terraform also.

## Variables

### Required

name | Environment variable | description
--- | --- | ---
web_endpoint_uri | GRAYLOG_WEB_ENDPOINT_URI |
auth_name | GRAYLOG_AUTH_NAME |
auth_password | GRAYLOG_AUTH_PASSWORD |

### Optional

name | Environment variable | default | description
--- | --- | --- | ---
x_requested_by | GRAYLOG_X_REQUESTED_BY | terraform-go-graylog | [X-Requested-By Header](https://github.com/Graylog2/graylog2-server/blob/370dd700bc8ada5448bf66459dec9a85fcd22d58/UPGRADING.rst#protecting-against-csrf-http-header-required)
api_version | GRAYLOG_API_VERSION | "v2" | Graylog's API version. The default value is "v2" for compatibility. If you use Graylog v3, please set "v3".

## Resources

* [alarm_callback](docs/alarm_callback.md)
* [alert_condition](docs/alert_condition.md)
* [dashboard](docs/dashboard.md)
* [extractor](docs/extractor.md)
* [index_set](docs/index_set.md)
* [input](docs/input.md)
* [ldap_setting](docs/ldap_setting.md)
* [pipeline](docs/pipeline.md)
* [pipeline_rule](docs/pipeline_rule.md)
* [pipeline_connection](docs/pipeline_connection.md)
* [role](docs/role.md)
* [stream](docs/stream.md)
* [stream_rule](docs/stream_rule.md)
* [user](docs/user.md)

## Unsupported resources

We can't support these resources for some reasons.

### CollectorConfiguration (includes input, output snippet)

We can't support these resources because graylog API doesn't return the created resource id (response body: no content).

The following APIs doesn't return the created resource id (response body: no content).

* POST /plugins/org.graylog.plugins.collector/configurations/{id}/inputs Create a configuration input
* POST /plugins/org.graylog.plugins.collector/configurations/{id}/outputs Create a configuration output
* POST /plugins/org.graylog.plugins.collector/configurations/{id}/snippets Create a configuration snippet
