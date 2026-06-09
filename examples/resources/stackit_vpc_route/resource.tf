resource "stackit_vpc_route" "example" {
  project_id       = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  vpc_id           = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  region           = "eu01"
  routing_table_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"

  destination = {
    type  = "cidrv4"
    value = "192.168.0.0/24"
  }
  next_hop = {
    type  = "ipv4"
    value = "192.168.0.10"
  }
  labels = {
    "key" = "value"
  }
}
