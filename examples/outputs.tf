output "public_subnet_ids" {
  description = "Number of Public Subnet IDs"
  value       = module.network.public_subnet_ids
}

output "private_subnet_ids" {
  description = "Number of Private Subnet IDs"
  value       = module.network.private_subnet_ids
}

output "vpc_id" {
  description = "AWS VPC ID"
  value       = module.network.vpc_id
}

output "zones" {
  description = "Number of Availability Zones used"
  value       = module.network.zones
}
