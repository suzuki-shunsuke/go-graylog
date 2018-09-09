# graylog_collector_configuration_output

https://github.com/suzuki-shunsuke/terraform-provider-graylog/blob/master/resource_collector_configuration_output.go

```
resource "graylog_collector_configuration_output" "test" {
  backend = "filebeat"
  type = "logstash"
  name = "test"
  properties = {
    hosts = "['localhost:5044']"
  }
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
collector_configuration_id | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
properties | null | object |

## Attrs Reference

name | type | etc
--- | --- | ---
output_id | string |
