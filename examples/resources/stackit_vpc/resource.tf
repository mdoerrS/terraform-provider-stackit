data "stackit_vpc" "example_vpc" {
  project_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"

  name        = "example"
  description = "Example description"
  labels = {
    "key" = "value"
  }
}
