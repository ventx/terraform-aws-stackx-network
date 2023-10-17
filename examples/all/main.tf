module "network" {
  source = "../../"

  name           = "terratest-all-0-network"
  workspace_name = var.workspace_name
  cluster_name   = var.cluster_name
  tags           = var.tags

  region = var.region

  # VPC
  single_nat_gateway = false

  # Private Subnets
  private             = true
  private_subnet_name = "terratest-all-0-private"
  private_subnet_tags = merge(
    var.tags,
    {
      "Attributes" = "private"
    }
  )

  # Public Subnets
  public             = true
  public_subnet_name = "terratest-all-0-public"
  public_subnet_tags = merge(
    var.tags,
    {
      "Attributes" = "public"
    }
  )

  # Subnets
  internal = true
  cache    = true

  # VPC Endpoints
  private_endpoints = true
  public_endpoints  = true
  db_endpoints      = true
  cache_endpoints   = true
  s3_endpoint       = true
}
