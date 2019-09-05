# graylog_grok_pattern

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_grok_pattern.go

```hcl
resource "graylog_grok_pattern" "datestamp" {
  name = "DATESTAMP"
  pattern = "%%{DATE}[- ]%%{TIME}"
}
```

Note that currently this resource doesn't support content packs.

And if you use the sequence `%{`, you have to escape it.

https://github.com/hashicorp/hcl2/blob/57bd5f374f26cdb7ae1b1c92fd6eb71335b9805b/hcl/hclsyntax/spec.md#template-literals

> The interpolation and directive introductions are escaped by doubling their leading characters.
> The ${ sequence is escaped as $${ and the %{ sequence is escaped as %%{.

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
