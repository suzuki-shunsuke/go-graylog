resource "graylog_user" "test" {
  username  = "test"
  email     = "test@example.com"
  full_name = "test test"
  password  = "password"
  roles = [
    "Reader",
    graylog_role.read-stream-test.name,
  ]
}
