# go-graylog

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/suzuki-shunsuke/go-graylog)
[![Build Status](https://cloud.drone.io/api/badges/suzuki-shunsuke/go-graylog/status.svg)](https://cloud.drone.io/suzuki-shunsuke/go-graylog)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/go-graylog/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/go-graylog)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/go-graylog)](https://goreportcard.com/report/github.com/suzuki-shunsuke/go-graylog)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-graylog.svg)](https://github.com/suzuki-shunsuke/go-graylog)
[![GitHub tag](https://img.shields.io/github/tag/suzuki-shunsuke/go-graylog.svg)](https://github.com/suzuki-shunsuke/go-graylog/releases)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-graylog/master/LICENSE)

[Graylog](https://www.graylog.org/) API client for Golang and terraform provider for Graylog.

## Note: Please use terraform-provider-graylog/terraform-provider-graylog

Now we released and developed [terraform-provider-graylog/terraform-provider-graylog](https://github.com/terraform-provider-graylog/terraform-provider-graylog) as the successor of `go-graylog`.
Please see the following document and announcement issue.

* https://terraform-provider-graylog.github.io/
* https://github.com/suzuki-shunsuke/go-graylog/issues/253

## API client

https://pkg.go.dev/github.com/suzuki-shunsuke/go-graylog/client?tab=doc#pkg-examples

If you use Graylog v3, use `client.NewClientV3` instead of `client.NewClient`.

## Terraform provider

Please see [docs/README.md](docs/README.md).

### Docker Image

https://quay.io/repository/suzuki_shunsuke/terraform-graylog

Docker image which is installed terraform and terraform-provider-graylog on Alpine.

## Note: Graylog API mock server has been migrated to the other repository

https://github.com/suzuki-shunsuke/graylog-mock-server (deprecated)

## Contribution

See [CONTRIBUTING.md](CONTRIBUTING.md) .

## License

[MIT](LICENSE)
