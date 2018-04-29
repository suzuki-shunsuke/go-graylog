# graylog_user

https://github.com/suzuki-shunsuke/terraform-provider-graylog/blob/master/resource_user.go

```
resource "graylog_user" "zoo" {
  username = "zoo"
  password = "password"
  email = "zoo@example.com"
  full_name = "zooull"
  permissions = ["users:read:zoo"]
}
```

## Argument Reference

### Required Argument

name | type | etc
--- | --- | ---
username | string |
password | string | sensitive
email | string |
permissions | []string |
full_name | string |

### Optional Argument

name | default | type | etc
--- | --- | --- | ---
roles | [] | []string |
timezone | "" | string | computed
session_timeout_ms | | int | computed

## Attrs Reference

name | type | etc
--- | --- | ---
user_id | string | computed
client_address | | string | computed
external | bool | computed
read_only | bool | computed
session_active | bool | computed
last_activity | string | computed
