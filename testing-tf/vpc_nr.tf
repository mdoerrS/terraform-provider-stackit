resource "stackit_vpc_network_range" "network_range" {
  depends_on = [stackit_vpc_region.region]

  project_id  = var.project_id
  vpc_id      = stackit_vpc_region.region.vpc_id
  ip_version  = "ipv4"
  prefix      = "192.168.1.0/24"
  description = "eu01 vpc network range"
}
