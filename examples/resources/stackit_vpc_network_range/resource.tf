data "stackit_vpc_network_range" "example" {
  project_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  vpc_id     = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  region     = "eu01"

  ipv4 = {
    description = "Example description"
    labels = {
      "key" = "value"
    }
    prefix                = "10.10.0.0/25"
    default_prefix_length = "25"
    max_prefix_length     = "29"
    min_prefix_length     = "24"
    nameservers = [
      "1.1.1.1",
      "8.8.8.8",
      "9.9.9.9",
    ]
  }
}
