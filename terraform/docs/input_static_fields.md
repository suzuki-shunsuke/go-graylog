# graylog_input_static_fields

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_input_static_fields.go

```hcl
resource "graylog_input_static_fields" "test" {
  input_id = "${graylog_input.test.id}"
  fields = {
    foo = "bar"
  }
}
```

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
input_id | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
fields | | map[string]string |
