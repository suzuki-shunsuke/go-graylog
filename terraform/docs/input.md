# graylog_input

https://github.com/suzuki-shunsuke/terraform-provider-graylog/blob/master/resource_input.go

```
resource "graylog_input" "test" {
  title = "terraform test"
  type = "org.graylog2.inputs.syslog.udp.SyslogUDPInput"
  configuration = {
    bind_address = "0.0.0.0"
    port = 514
    recv_buffer_size = 262144
  }
}
```

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
title | string |
configuration | map[string] |
type | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
global | "" | string |
node | "" | string |

## Attributes Reference

name | type | etc
--- | --- | ---
input_id | string |
