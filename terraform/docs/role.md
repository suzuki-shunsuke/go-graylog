# graylog_role

* [Example](https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/example/v0.12/role.tf)
* [Source Code](https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_role.go)

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
name | string |
permissions | []string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
description | "" | string |

## Attrs Reference

name | type | etc
--- | --- | ---
read_only | bool | computed
