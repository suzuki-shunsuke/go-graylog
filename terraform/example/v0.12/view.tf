resource "graylog_view" "test" {
  title        = "test"
  description = "description"
  summary = "summary"
  # set appropriate search_id
  # search_id = "5d9529b275d97f58f9539279"
# state = {
#   "6971d00a-e605-43fb-b873-e4bca773d286" = {
#     selected_fields = ["source", "message"]
#   }
# }
}
