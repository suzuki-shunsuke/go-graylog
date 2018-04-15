# graylog_role

https://github.com/suzuki-shunsuke/terraform-provider-graylog/blob/master/resource_role.go

```
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

## Attributes Reference

name | type | etc
--- | --- | ---
read_only | bool | computed
