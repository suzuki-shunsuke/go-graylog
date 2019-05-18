# graylog_pipeline

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_pipeline.go

```
resource "graylog_pipeline" "test" {
  source      = "source": "pipeline \"test\"\nstage 0 match either\nend"
  description = "description"
}
```

In HCL, you can use here document.

https://github.com/hashicorp/hcl#syntax

```
resource "graylog_pipeline" "test" {
  source = <<EOF
pipeline "test4"
  stage 0 match either
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
