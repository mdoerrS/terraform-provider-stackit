data "stackit_vpc_routing_table" "example" {
  project_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  vpc_id     = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  region     = "eu01"

  name        = "example"
  description = "Example description"
  labels = {
    "key" = "value"
  }
  dynamic_routes = true
  system_routes  = true
}
