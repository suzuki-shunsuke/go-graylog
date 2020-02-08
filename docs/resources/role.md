# graylog_role

* [Example](../../examples/v0.12/role.tf)
* [Source Code](../../graylog/terraform/resource_role.go)

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
