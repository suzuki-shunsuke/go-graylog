# graylog_event_notification

* [Example](https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/example/v0.12/event_notification.tf)
* [Source Code](https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_event_notification.go)

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
title | string |
config | string | JSON string

`config` is a JSON string.
The format of `config` depends on the Event Notification type.
Please see the [example](https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/example/v0.12/event_notification.tf).
Using the [Graylog's API browser](https://docs.graylog.org/en/3.1/pages/configuration/rest_api.html) you can check the format of `config`.

### Optional Argument

name | default | type | description
--- | --- | --- | ---
description | ""| string |

## Attrs Reference

None.
