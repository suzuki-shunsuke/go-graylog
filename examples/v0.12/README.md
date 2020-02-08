# example

## Requirement

* Terraform v0.12
* terraform-provider-graylog
* [graylog-plugin-slack](https://github.com/graylog-labs/graylog-plugin-slack)
  * install a jar file on the `plugin` directory
* Docker Engine
* Docker Compose

## Getting Started

```console
# It takes some time to launch Graylog.
$ docker-compose up -d
$ terraform init
$ terraform plan
$ terraform apply
```
