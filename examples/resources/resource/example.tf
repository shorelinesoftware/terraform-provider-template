
resource "shoreline_resource" "books" {
  name        = "books"
  description = "Pods with books app."
  value       = "host | pod | app='bookstore'"
}


resource "shoreline_resource" "params_resource" {
  name        = "params_resource"
  description = "params_resource"
  value       = "host | limit=$limit | az=$az"
  params      = ["limit", "az"]
}