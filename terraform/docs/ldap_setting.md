# graylog_ldap_setting

https://github.com/suzuki-shunsuke/terraform-provider-graylog/blob/master/resource_ldap_setting.go

```
resource "graylog_ldap_setting" "foo" {
  enabled = true
  system_username = ""
  system_password = ""
  ldap_uri = "ldap://localhost:389"
  use_start_tls = false
  trust_all_certificates = false
  active_directory = false
  search_base = ""
  search_pattern = ""
  display_name_attribute = ""
  default_group = ""
  group_search_base = ""
  group_id_attribute = ""
  group_search_pattern = ""
}
```

Unlike other resources, LDAP settings has no id,
so when you import the LDAP settings, please specify some string as id.

```
terraform import graylog_ldap_setting.foo bar
```

## Argument Reference

### Required Argument

Nothing.

### Optional Argument

name | default | type | description
--- | --- | --- | ---
description | "" | string |
enabled | false | bool |
system_username | "" | string |
system_password | "" | string |
ldap_uri | "ldap://localhost:389" | string |
use_start_tls | false | bool |
trust_all_certificates | false | bool |
active_directory | false | bool |
search_base | "" | string |
search_pattern | "" | string |
display_name_attribute | "" | string |
group_search_base | "" | string |
group_id_attribute | "" | string |
group_search_pattern | "" | string |
default_group | "" | string | computed
