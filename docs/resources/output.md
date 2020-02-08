# graylog_output

* [Example](../../examples/v0.12/output.tf)
* [Source Code](../../graylog/terraform/resource_output.go)

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
title | string |
type | string |
configuration | string | JSON string

`configuration` is a JSON string.
The format of `configuration` depends on the output type.
Please see the [example](https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/example/v0.12/output.tf).
Using the [Graylog's API browser](https://docs.graylog.org/en/3.1/pages/configuration/rest_api.html) you can check the format of `configuration`.

### Optional Argument

None.

## Attrs Reference

None.
