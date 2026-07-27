resource "stackit_vpc" "vpc" {
  project_id  = var.project_id
  name        = "vpc"
  description = "vpc for new network changes"
}
