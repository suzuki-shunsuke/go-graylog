# graylog_ldap_setting

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_ldap_setting.go

```
resource "graylog_ldap_setting" "foo" {
  enabled = true
  system_username = "admin"
  system_password = "password"
  ldap_uri = "ldap://localhost:389"
  use_start_tls = false
  trust_all_certificates = false
  active_directory = false
  search_base = "OU=user,OU=foo,DC=example,DC=com"
  search_pattern = "(cn={0})"
  display_name_attribute = "displayname"
  default_group = "Reader"
  group_search_base = ""
  group_id_attribute = ""
  group_search_pattern = ""
  group_mapping = {
    foo = "Reader"
  }
}
```

Unlike other resources, LDAP settings has no id,
so when you import the LDAP settings, please specify some string as id.

```
terraform import graylog_ldap_setting.foo bar
```

## Argument Reference

### Required Argument

name | default | type | description
--- | --- | --- | ---
system_username | "" | string |
ldap_uri | "ldap://localhost:389" | string |
search_base | "" | string |
search_pattern | "" | string |
display_name_attribute | "" | string |
system_password | "" | string | sensitive
default_group | "" | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
description | "" | string |
enabled | false | bool |
use_start_tls | false | bool |
trust_all_certificates | false | bool |
active_directory | false | bool |
group_search_base | "" | string |
group_id_attribute | "" | string |
group_search_pattern | "" | string |
group_mapping | | map[string]string |
