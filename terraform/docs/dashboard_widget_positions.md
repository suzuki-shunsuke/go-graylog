# graylog_dashboard_widget_positions

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_dashboard_widget_positions.go

```hcl
resource "graylog_dashboard_widget_positions" "test" {
  dashboard_id = graylog_dashboard.test.id
  positions {
    widget_id = graylog_dashboard_widget.test1.id
    width = 2
    col = 6
    row = 1
    height = 6
  }

  positions {
    widget_id = graylog_dashboard_widget.test2.id
    width = 4
    col = 1
    row = 1
    height = 2
  }
}
```

## Required arguments

name | type | description
--- | --- | ---
dashboard_id | string |
positions | array |
positions[].widget_id | string |

### Optional arguments

name | type | default | description
--- | --- | --- | ---
positions[].width | int | |
positions[].col | int | |
positions[].row | int | |
positions[].height | int | |
