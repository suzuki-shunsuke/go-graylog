# graylog_dashboard_widget_positions

* [Example](https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/example/v0.12/dashboard.tf)
* [Source code](https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_dashboard_widget_positions.go)

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
