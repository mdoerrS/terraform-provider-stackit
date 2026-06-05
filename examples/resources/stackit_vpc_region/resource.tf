resource "stackit_vpc_region" "example" {
  project_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  vpc_id     = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  region     = "eu01"

  default_routing_table = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  ipv4 = {
    default_nameservers = [
      "1.1.1.1",
      "8.8.8.8",
      "9.9.9.9",
    ],
    default_network_range = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  }
}
