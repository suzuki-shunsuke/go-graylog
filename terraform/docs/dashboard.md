# graylog_dashboard

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_dashboard.go

```hcl
resource "graylog_dashboard" "test-dashboard" {
  title = "test-dashboard"
  description = "test dashboard"
}
```

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
title | string |
description | string |

### Optional Argument

None

## Attrs Reference

name | type | etc
--- | --- | ---
created_at | string | computed
