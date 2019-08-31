# graylog_grok_pattern

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_grok_pattern.go

```hcl
resource "graylog_grok_pattern" "datestamp" {
  name = "DATESTAMP"
  pattern = "%{DATE}[- ]%{TIME}"
}
```

Note that currently this resource doesn't support content packs.

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
name | string |
pattern | string |

### Optional Argument

Nothing.

## Attrs Reference

Nothing.
