# graylog_collector_configuration

https://github.com/suzuki-shunsuke/terraform-provider-graylog/blob/master/resource_collector_configuration.go

```
resource "graylog_collector_configuration" "test" {
  tags = ["test"]
  name = "terraform test"
}
```

Note that collector configuration's inputs, outputs and snippets are defined as other resources.

* [input](collector_configuration_input.md)
* [output](collector_configuration_output.md)
* [snippet](collector_configuration_snippet.md)

Note that graylog doesn't provide an API to update collector configuration's tags,
so it is impossible to update tags.

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
name | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
tags | [] | []string |

## Attrs Reference

name | type | etc
--- | --- | ---
id | string |
