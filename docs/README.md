# terraform-provider-graylog

terraform provider for [Graylog](https://www.graylog.org/).

## Install

[Download binary](https://github.com/suzuki-shunsuke/go-graylog/releases) and install it.

https://www.terraform.io/docs/configuration/providers.html#third-party-plugins

## Docker Image

https://quay.io/repository/suzuki_shunsuke/terraform-graylog

Docker image which is installed terraform and terraform-provider-graylog on Alpine.

## Example

```hcl
provider "graylog" {
  web_endpoint_uri = "${var.web_endpoint_uri}"
  auth_name = "${var.auth_name}"
  auth_password = "${var.auth_password}"
}

// Role my-role-2
resource "graylog_role" "my-role-2" {
  name = "my-role-2"
  permissions = ["users:edit"]
  description = "Created by terraform"
}
```

And please see [example v0.11](../examples/v0.11) and [example v0.12](../examples/v0.12) also.

## Variables

### Required

name | Environment variable | description
--- | --- | ---
web_endpoint_uri | GRAYLOG_WEB_ENDPOINT_URI | API endpoint, for example https://graylog.example.com/api
auth_name | GRAYLOG_AUTH_NAME | Username or API token or Session Token
auth_password | GRAYLOG_AUTH_PASSWORD | Password or the literal `"token"` or `"session"`

About `auth_name` and `auth_password`, please see the [Graylog's Documentation](https://docs.graylog.org/en/latest/pages/configuration/rest_api.html).

You can authenticate with either password or access token or session token.

password

```
auth_name = "<user name>"
auth_password = "<password>"
```

access token

```
auth_name = "<access token>"
auth_password = "token"
```

session token

```
auth_name = "<session token>"
auth_password = "session"
```

### Optional

name | Environment variable | default | description
--- | --- | --- | ---
x_requested_by | GRAYLOG_X_REQUESTED_BY | terraform-go-graylog | [X-Requested-By Header](https://github.com/Graylog2/graylog2-server/blob/370dd700bc8ada5448bf66459dec9a85fcd22d58/UPGRADING.rst#protecting-against-csrf-http-header-required)
api_version | GRAYLOG_API_VERSION | "v2" | Graylog's API version. The default value is "v2" for compatibility. If you use Graylog v3, please set "v3".

## Resources

* [alarm_callback](resources/alarm_callback.md)
* [alert_condition](resources/alert_condition.md)
* [dashboard](resources/dashboard.md)
* [dashboard_widget](resources/dashboard_widget.md)
* [dashboard_widget_positions](resources/dashboard_widget_positions.md)
* [extractor](resources/extractor.md)
* [event_definition](resources/event_definition.md)
* [event_notification](resources/event_notification.md)
* [grok_pattern](resources/grok_pattern.md)
* [index_set](resources/index_set.md)
* [input](resources/input.md)
* [input_static_fields](resources/input_static_fields.md)
* [ldap_setting](resources/ldap_setting.md)
* [output](resources/output.md)
* [pipeline](resources/pipeline.md)
* [pipeline_rule](resources/pipeline_rule.md)
* [pipeline_connection](resources/pipeline_connection.md)
* [role](resources/role.md)
* [stream](resources/stream.md)
* [stream_output](resources/stream_output.md)
* [stream_rule](resources/stream_rule.md)
* [user](resources/user.md)

## Data sources

* [dashboard](data-sources/dashboard.md)
* [index_set](data-sources/index_set.md)
* [stream](data-sources/stream.md)

## Unsupported resources

We can't support these resources for some reasons.

### CollectorConfiguration (includes input, output snippet)

We can't support these resources because graylog API doesn't return the created resource id (response body: no content).

The following APIs doesn't return the created resource id (response body: no content).

* POST /plugins/org.graylog.plugins.collector/configurations/{id}/inputs Create a configuration input
* POST /plugins/org.graylog.plugins.collector/configurations/{id}/outputs Create a configuration output
* POST /plugins/org.graylog.plugins.collector/configurations/{id}/snippets Create a configuration snippet
