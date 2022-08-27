module "network" {
  source = "../"

  name           = "stackx-0-network"
  workspace_name = var.workspace_name
  cluster_name   = var.cluster_name
  tags = {
    test     = true,
    deleteme = true
  }

  region = var.region

  # VPC
  single_nat_gateway = false

  # Private Subnets
  private = true
  private_subnet_tags = merge(
    var.tags,
    {
      "Access" = "private"
    }
  )

  # Public Subnets
  public = true
  public_subnet_tags = merge(
    var.tags,
    {
      "Access" = "public"
    }
  )
}
