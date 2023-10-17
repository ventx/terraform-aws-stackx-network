provider "aws" {
  region = "us-east-1"
}

variables {
    name           = "terratest-all-0-network"
    workspace_name = "terratest"
    cluster_name   = "test"
    tags           = {
        "owner"     = "terraform-aws-network",
        "managedby" = "terratest",
        "project"   = "stackx",
        "workspace" = "terratest"
      }

    region = "us-east-1"

    # VPC
    single_nat_gateway = true

    # Private Subnets
    private             = true
    private_subnet_name = "terratest-all-0-private"

    # Public Subnets
    public             = true
    public_subnet_name = "terratest-all-0-public"

    # Subnets
    internal = true
    db       = true

    # VPC Endpoints
    private_endpoints = true
    public_endpoints  = true
    s3_endpoint       = true
}

run "valid_config" {

  command = plan

  assert {
    condition     = aws_vpc.vpc.cidr_block == "10.3.0.0/16"
    error_message = "VPC CIDR did not match expected"
  }

  assert {
    condition = aws_vpc.vpc.enable_dns_support == true
    error_message = "VPC has not DNS enabled, did not match expected"
  }

  assert {
    condition     = length(aws_subnet.private) == 3
    error_message = "Private Subnet count did not match expected"
  }

  assert {
    condition     = aws_vpc.vpc.tags["Name"] == "terratest-all-0-network"
    error_message = "VPC tag 'Name' did not match expected"
  }
}

run "execute" {
  # Test apply configuration
}
