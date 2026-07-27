resource "stackit_network" "network" {
  depends_on = [stackit_vpc_network_range.network_range]

  project_id         = var.project_id
  name               = "vpc-network-min"
  ipv4_prefix_length = 28
  ipv4_vpc_network_range_id = stackit_vpc_network_range.network_range.network_range_id
  vpc_id                    = stackit_vpc.vpc.vpc_id
}
