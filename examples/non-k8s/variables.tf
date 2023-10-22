variable "tags" {
  default = {
    "owner"     = "terraform-aws-network",
    "managedby" = "terratest",
    "project"   = "stackx",
  }
}

variable "cluster_name" {
  default = "test"
}

variable "cluster_number" {
  default = 0
}

variable "region" {
  default = "eu-central-1"
}
