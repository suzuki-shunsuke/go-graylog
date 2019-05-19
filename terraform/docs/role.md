# graylog_role

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_role.go

```hcl
resource "graylog_role" "foo" {
  name = "foo"
  description = "user foo"
  permissions = ["*"]
}
```

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
