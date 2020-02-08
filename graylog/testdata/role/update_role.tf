resource "graylog_role" "foo" {
  name        = "foo"
  description = "updated description"

  permissions = [
    "view:edit",
    "extendedsearch:use",
    "extendedsearch:create",
    "view:read",
    "view:use"
  ]

  read_only = false
}
