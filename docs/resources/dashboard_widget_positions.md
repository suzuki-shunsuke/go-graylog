# graylog_dashboard_widget_positions

* [Example](../../examples/v0.12/dashboard.tf)
* [Source code](../../graylog/terraform/resource_dashboard_widget_positions.go)

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
