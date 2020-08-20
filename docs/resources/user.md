# graylog_user

* [Example](../../examples/v0.12/user.tf)
* [Source code](../../graylog/terraform/resource_user.go)

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
session_timeout_ms | 3600000 | int |

## Attrs Reference

name | type | etc
--- | --- | ---
user_id | string | computed
external | bool | computed
read_only | bool | computed
client_address | | string | computed
session_active | bool | computed
last_activity | string | computed
