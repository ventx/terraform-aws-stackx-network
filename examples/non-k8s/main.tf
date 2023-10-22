module "network" {
  source = "../../"

  k8s = false

  name         = "test-0-network"
  cluster_name = var.cluster_name
  tags         = var.tags

  region = var.region

  # VPC
  single_nat_gateway = true

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
