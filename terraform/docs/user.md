# graylog_user

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_user.go

```hcl
resource "graylog_user" "zoo" {
  username = "zoo"
  email = "zoo@example.com"
  full_name = "zooull"
  roles = ["Reader"]
}
```

## Argument Reference

### Required Argument

name | type | etc
--- | --- | ---
username | string | force_new
email | string |
full_name | string |

### Optional Argument

name | default | type | etc
--- | --- | --- | ---
password | string | sensitive
permissions | string set | computed
roles | [] | string set |
timezone | "" | string | computed
session_timeout_ms | | int | computed

## Attrs Reference

name | type | etc
--- | --- | ---
user_id | string | computed
external | bool | computed
read_only | bool | computed
client_address | | string | computed
session_active | bool | computed
last_activity | string | computed
