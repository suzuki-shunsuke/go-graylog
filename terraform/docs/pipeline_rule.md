# graylog_pipeline_rule

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_pipeline_rule.go

```
resource "graylog_pipeline_rule" "test" {
  source      = "rule \"test\"\nwhen\n    to_long($message.status) \u003c 500\nthen\n    set_field(\"status_01\", 1);\nend"
  description = "description"
}
```

In HCL, you can use here document.

https://github.com/hashicorp/hcl#syntax

```
resource "graylog_pipeline_rule" "test" {
  source = <<EOF
rule "test"
when
    to_long($message.status) < 500
then
    set_field("status_01", 1);
end
EOF

  description = "test"
}
```

## Argument Reference

### Required Argument

name | type | etc
--- | --- | ---
source | string |

### Optional Argument

name | default | type | etc
--- | --- | --- | ---
description | string |
