# go-graylog

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/go-graylog)
[![Build Status](https://cloud.drone.io/api/badges/suzuki-shunsuke/go-graylog/status.svg)](https://cloud.drone.io/suzuki-shunsuke/go-graylog)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/go-graylog/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/go-graylog)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/go-graylog)](https://goreportcard.com/report/github.com/suzuki-shunsuke/go-graylog)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-graylog.svg)](https://github.com/suzuki-shunsuke/go-graylog)
[![GitHub tag](https://img.shields.io/github/tag/suzuki-shunsuke/go-graylog.svg)](https://github.com/suzuki-shunsuke/go-graylog/releases)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-graylog/master/LICENSE)

[Graylog](https://www.graylog.org/) API client for Golang and terraform provider for Graylog.

## Example of API client

* https://godoc.org/github.com/suzuki-shunsuke/go-graylog/client/#example-Client

## Terraform provider

* [terraform-provider-graylog](https://github.com/suzuki-shunsuke/go-graylog/tree/master/terraform)

## Supported Graylog version

We support the following versions.

* v2.5
* v3

We use [the graylog's official Docker Image](https://hub.docker.com/r/graylog/graylog/) for development.

### Support of Graylog v3

Use `client.NewClientV3` instead of `client.NewClient` .

In the terraform provider, please set the variable `api_version` to `v3`.

## Note: Graylog API mock server has been migrated to the other repository

https://github.com/suzuki-shunsuke/graylog-mock-server (deprecated)

## Contribution

See [CONTRIBUTING.md](CONTRIBUTING.md) .

## License

[MIT](LICENSE)
