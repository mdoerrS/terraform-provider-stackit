resource "stackit_vpc_region" "region" {
  project_id = var.project_id
  vpc_id     = stackit_vpc.vpc.vpc_id
  region     = "eu01"
}
