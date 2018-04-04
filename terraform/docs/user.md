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

### Optional Argument

name | default | type | etc
--- | --- | --- | ---
full_name | "" | string |
roles | [] | []string |
user_id | "" | string | computed
timezone | "" | string |
session_timeout_ms | | int |
external | | bool |
read_only | | bool |
client_address | | string | `0.0.0.0`

## Attributes Reference

name | type | etc
--- | --- | ---
user_id | string |
