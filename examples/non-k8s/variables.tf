variable "tags" {
  default = {
    "owner"     = "terraform-aws-network",
    "managedby" = "terratest",
    "project"   = "stackx",
    "workspace" = "terratest"
  }
}

variable "workspace_name" {
  default = "terratest"
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
