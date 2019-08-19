resource "graylog_role" "terraform" {
  name        = "terraform"
  description = "terraform"

  permissions = [
    "dashboards:*",
    "indexsets:*",
    "inputs:*",
    "roles:*",
    "streams:*",
    "users:*",
    "pipeline_rule:*",
  ]
}

resource "graylog_role" "terraform-read" {
  name        = "terraform-read"
  description = "terraform-read"

  permissions = [
    "dashboards:read",
    "indexsets:read",
    "inputs:read",
    "roles:read",
    "streams:read",
    "users:list",
    "users:edit",
    "users:tokenlist",
    "users:tokencreate",
    "pipeline_rule:read",
  ]
}
