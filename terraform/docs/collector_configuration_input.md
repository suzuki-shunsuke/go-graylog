# graylog_collector_configuration_input

https://github.com/suzuki-shunsuke/terraform-provider-graylog/blob/master/resource_collector_configuration_input.go

```
resource "graylog_collector_configuration_input" "test" {
  backend = "winlogbeat"
  type = ""
  name = ""
  properties = {
    event = "[{'name':'Application'},{'name':'System'},{'name':'Security'}]"
  }
  forward_to = "${graylog_collector_configuration_output.test-terraform.id}"
  collector_configuration_id = "${graylog_collector_configuration.test-terraform.id}"
}
```

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
backend | string |
type | string |
name | string |
forward_to | string |
collector_configuration_id | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
properties | null | object |

## Attrs Reference

name | type | etc
--- | --- | ---
input_id | string |
