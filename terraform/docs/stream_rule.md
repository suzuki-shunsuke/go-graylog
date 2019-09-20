# graylog_stream_rule

* [Example](https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/example/v0.12/stream_rule.tf)
* [Source code](https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_stream_rule.go)

## How to import

Specify `<stream id>/<stream rule id>` as ID.

```console
$ terraform import graylog_stream_rule.test 5bb1b4b5c9e77bbbbbbbbbbb/5c4acaefc9e77bbbbbbbbbbb
```

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
field | string |
value | string |
description | string |
type | int |
stream_id | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
inverted | | bool |
