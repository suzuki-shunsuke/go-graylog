# Data source graylog_dashboard

* [Example](https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/example/v0.12/dashboard.tf)
* [Source code](https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/data_source_dashboard.go)

## Required Argument

One of `dashboard_id` or `title` must be set.

## Attributes

name | type | description
--- | --- | ---
title | string |
dashboard_id | string |
description | string |
