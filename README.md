<h1 align="center">
  <a href="https://github.com/ventx/terraform-aws-stackx-network">
    <!-- Please provide path to your logo here -->
    <img src="https://raw.githubusercontent.com/ventx/terraform-aws-stackx-network/main/docs/images/logo.svg" alt="Logo" width="100" height="100">
  </a>
</h1>

<div align="center">
  ventx/terraform-aws-stackx-network
  <br />
  <a href="#about"><strong>Explore the diagrams »</strong></a>
  <br />
  <br />
  <a href="https://github.com/ventx/terraform-aws-stackx-network/issues/new?assignees=&labels=bug&template=01_BUG_REPORT.md&title=bug%3A+">Report a Bug</a>
  ·
  <a href="https://github.com/ventx/terraform-aws-stackx-network/issues/new?assignees=&labels=enhancement&template=02_FEATURE_REQUEST.md&title=feat%3A+">Request a Feature</a>
  ·
  <a href="https://github.com/ventx/terraform-aws-stackx-network/issues/new?assignees=&labels=question&template=04_SUPPORT_QUESTION.md&title=support%3A+">Ask a Question</a>
</div>

<div align="center">
<br />

[![Project license](https://img.shields.io/github/license/ventx/terraform-aws-stackx-network.svg?style=flat-square)](LICENSE)

[![Pull Requests welcome](https://img.shields.io/badge/PRs-welcome-ff69b4.svg?style=flat-square)](https://github.com/ventx/terraform-aws-stackx-network/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22)
[![code with love by ventx](https://img.shields.io/badge/%3C%2F%3E%20with%20♥%20by-ventx-blue)](https://github.com/ventx)

</div>

<details open>
<summary>Table of Contents</summary>

- [About](#about)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Quickstart](#quickstart)
- [Usage](#usage)
- [Support](#support)
- [Project assistance](#project-assistance)
- [Contributing](#contributing)
- [Authors & contributors](#authors--contributors)
- [Security](#security)
- [License](#license)
- [Acknowledgements](#acknowledgements)
- [Roadmap](#roadmap)

</details>

---

## About

> Terraform networking module which deploys Kubernetes (EKS) optimized VPC networks and subnets to AWS.
> Private, Public, Internal (noinet), Database (RDS) and Cache (ElastiCache)
> subnets supported. CIDR range in 10.x.0.0/16 will be automatically (default) chosen based on Availability Zone.
> Can be used with or without Kubernetes. -- Part of stackx.


<details>
<summary>ℹ️ Architecture Diagrams</summary>
<br>


|                                Diagrams                                 |                                                            Rover                                                            |
|:-------------------------------------------------------------------------------:|:---------------------------------------------------------------------------------------------------------------------------------------:|
| <img src="https://raw.githubusercontent.com/ventx/terraform-aws-stackx-network/main/docs/images/screenshot1.png" title="Diagrams" width="100%"> | <img src="https://raw.githubusercontent.com/ventx/terraform-aws-stackx-network/main/docs/images/screenshot2.png" title="Rover" width="100%"> |

</details>


### Built With

<no value>


## Getting Started

### Prerequisites


* AWS credentials
* Terraform

### Quickstart

To get started, clone the projects, check all configurable [Inputs](#inputs) and deploy everything with `make`.

```shell
git clone https://github.com/ventx/stackx-terraform-aws-network.git
make all # init, validate, plan, apply
```



## Usage


You can run this module in conjunction with other stackx components (recommended) or as single-use (build your own).

### stackx (RECOMMENDED)

This is just a bare minimum example of how to use the module.
See all available stackx modules here: https://github.com/ventx


```hcl
  module "aws-network" {
    source = "ventx/stackx-network/aws"
    version     = "0.2.0" // Pinned and tested version, generated by {x-release-please-version}
  }
```

### Single-Use

```hcl
  module "aws.network" {
    source = "ventx/stackx-network/aws"
    version     = "0.2.0" // Pinned and tested version, generated by {x-release-please-version}
  }
```




## Terraform



### Features


* Simple and easy to use, just the bare minimum
* Optimized for Kubernetes
* Single NAT Gateway (default = `false`)
* IPv6 (default = `false`)
* (Optional) DB Subnet Group ([RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html#USER_VPC.Subnets), default = `false`)
* (Optional) ElastiCache Subnet Group ([ElastiCache](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SubnetGroups.html), default = `false`)
* The minimum VPC CIDR subnet mask is /16, as too small subnet CIDRs might cause trouble to change afterwards and with K8s Pod networking.
* Set `var.k8s` = `true` (default) to get bigger `/18` private subnets (for K8s Pods) and smaller `/22` public, database, cache subnets (if enabled)
* Set `var.k8s` = `false` to get evenly distributed `/20` subnet CIDRs across all subnets

### Resources


* VPC
* Subnets - Public - Private - Internal - Database - Cache
* Subnet Groups
* Route Tables
* Routes
* Internet Gateway
* NAT Gateway/s
* IPv6 Egress Gateway
* S3 VPC Endpoint

### Opinions

Our Terraform modules are are highly opionated:

* Keep modules small, focused, simple and easy to understand
* Prefer simple code over complex code
* Prefer [KISS](https://en.wikipedia.org/wiki/KISS_principle) > [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself)
* Set some sane default values for variables, but do not set a default value if user input is strictly required


These opinions can be seen as some _"soft"_ rules but which are not strictly required.

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 4.49.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 4.49.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_db_subnet_group.db](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_subnet_group) | resource |
| [aws_egress_only_internet_gateway.ipv6](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/egress_only_internet_gateway) | resource |
| [aws_eip.eip](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eip) | resource |
| [aws_elasticache_subnet_group.cache](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticache_subnet_group) | resource |
| [aws_internet_gateway.igw](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/internet_gateway) | resource |
| [aws_nat_gateway.natgw](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/nat_gateway) | resource |
| [aws_route.private_ipv6_egress](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route) | resource |
| [aws_route.private_nat_gateway](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route) | resource |
| [aws_route.public_internet_gateway](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route) | resource |
| [aws_route.public_internet_gateway_ipv6](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route) | resource |
| [aws_route_table.cache](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table) | resource |
| [aws_route_table.db](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table) | resource |
| [aws_route_table.internal](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table) | resource |
| [aws_route_table.private](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table) | resource |
| [aws_route_table.public](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table) | resource |
| [aws_route_table_association.cache](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table_association) | resource |
| [aws_route_table_association.db](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table_association) | resource |
| [aws_route_table_association.internal](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table_association) | resource |
| [aws_route_table_association.private](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table_association) | resource |
| [aws_route_table_association.public](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table_association) | resource |
| [aws_subnet.cache](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/subnet) | resource |
| [aws_subnet.db](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/subnet) | resource |
| [aws_subnet.internal](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/subnet) | resource |
| [aws_subnet.private](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/subnet) | resource |
| [aws_subnet.public](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/subnet) | resource |
| [aws_vpc.vpc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc) | resource |
| [aws_vpc_endpoint.s3](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint) | resource |
| [aws_vpc_endpoint_route_table_association.cache_s3](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint_route_table_association) | resource |
| [aws_vpc_endpoint_route_table_association.db_s3](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint_route_table_association) | resource |
| [aws_vpc_endpoint_route_table_association.internal_s3](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint_route_table_association) | resource |
| [aws_vpc_endpoint_route_table_association.private_s3](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint_route_table_association) | resource |
| [aws_vpc_endpoint_route_table_association.public_s3](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint_route_table_association) | resource |
| [aws_availability_zone.all](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/availability_zone) | data source |
| [aws_availability_zones.available](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/availability_zones) | data source |
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_az_number"></a> [az\_number](#input\_az\_number) | n/a | `map` | <pre>{<br>  "a": 1,<br>  "b": 2,<br>  "c": 3,<br>  "d": 4,<br>  "e": 5,<br>  "f": 6<br>}</pre> | no |
| <a name="input_cache"></a> [cache](#input\_cache) | Enable / Disable Cache Subnets | `bool` | `false` | no |
| <a name="input_cache_endpoints"></a> [cache\_endpoints](#input\_cache\_endpoints) | Enable / Disable VPC Endpoints for Cache Subnets | `bool` | `false` | no |
| <a name="input_cache_subnet_name"></a> [cache\_subnet\_name](#input\_cache\_subnet\_name) | Name for Cache subnets and dependent resources  (will be prefixed with `var.name-`) | `string` | `"cache"` | no |
| <a name="input_cache_subnet_tags"></a> [cache\_subnet\_tags](#input\_cache\_subnet\_tags) | Tags as map for Cache subnets and dependent resources (preferably generated by terraform-null-label) | `map(string)` | <pre>{<br>  "type": "cache"<br>}</pre> | no |
| <a name="input_cluster_name"></a> [cluster\_name](#input\_cluster\_name) | EKS Cluster name to add in tags for ELB and cluster-autoscaler / karpenter discovery of subnets / VPC | `string` | `"stackx"` | no |
| <a name="input_db"></a> [db](#input\_db) | Enable / Disable Database Subnets | `bool` | `false` | no |
| <a name="input_db_endpoints"></a> [db\_endpoints](#input\_db\_endpoints) | Enable / Disable VPC Endpoints for Database Subnets | `bool` | `false` | no |
| <a name="input_db_subnet_name"></a> [db\_subnet\_name](#input\_db\_subnet\_name) | Name for Database subnets and dependent resources  (will be prefixed with `var.name-`) | `string` | `"database"` | no |
| <a name="input_db_subnet_tags"></a> [db\_subnet\_tags](#input\_db\_subnet\_tags) | Tags as map for Database subnets and dependent resources (preferably generated by terraform-null-label) | `map(string)` | <pre>{<br>  "type": "database"<br>}</pre> | no |
| <a name="input_internal"></a> [internal](#input\_internal) | Enable / Disable Internal Subnets | `bool` | `false` | no |
| <a name="input_internal_endpoints"></a> [internal\_endpoints](#input\_internal\_endpoints) | Enable / Disable VPC Endpoints for Internal Subnets | `bool` | `false` | no |
| <a name="input_internal_subnet_name"></a> [internal\_subnet\_name](#input\_internal\_subnet\_name) | Name for Internal subnets and dependent resources  (will be prefixed with `var.name-`) | `string` | `"internal"` | no |
| <a name="input_internal_subnet_tags"></a> [internal\_subnet\_tags](#input\_internal\_subnet\_tags) | Tags as map for Internal subnets and dependent resources (preferably generated by terraform-null-label) | `map(string)` | <pre>{<br>  "type": "internal"<br>}</pre> | no |
| <a name="input_ipv6"></a> [ipv6](#input\_ipv6) | Enable / Disable IPv6 in your VPC and subnets | `bool` | `true` | no |
| <a name="input_k8s"></a> [k8s](#input\_k8s) | Enable / Disable usage for Kubernetes (adding tags to subnets and resources, increase private subnetes size) | `bool` | `true` | no |
| <a name="input_name"></a> [name](#input\_name) | Base Name for all resources (preferably generated by terraform-null-label) | `string` | `"stackx-network"` | no |
| <a name="input_private"></a> [private](#input\_private) | Enable / Disable Private Subnets | `bool` | `true` | no |
| <a name="input_private_endpoints"></a> [private\_endpoints](#input\_private\_endpoints) | Enable / Disable VPC Endpoints for Private Subnets | `bool` | `true` | no |
| <a name="input_private_subnet_name"></a> [private\_subnet\_name](#input\_private\_subnet\_name) | Name for Private subnets and dependent resources (will be prefixed with `var.name-`) | `string` | `"private"` | no |
| <a name="input_private_subnet_tags"></a> [private\_subnet\_tags](#input\_private\_subnet\_tags) | Tags as map for Private subnets and dependent resources (preferably generated by terraform-null-label) | `map(string)` | <pre>{<br>  "type": "private"<br>}</pre> | no |
| <a name="input_public"></a> [public](#input\_public) | Enable / Disable Public Subnets | `bool` | `true` | no |
| <a name="input_public_endpoints"></a> [public\_endpoints](#input\_public\_endpoints) | Enable / Disable VPC Endpoints for Public Subnets | `bool` | `false` | no |
| <a name="input_public_subnet_name"></a> [public\_subnet\_name](#input\_public\_subnet\_name) | Name for Public subnets and dependent resources  (will be prefixed with `var.name-`) | `string` | `"public"` | no |
| <a name="input_public_subnet_tags"></a> [public\_subnet\_tags](#input\_public\_subnet\_tags) | Tags as map for Public subnets and dependent resources (preferably generated by terraform-null-label) | `map(string)` | <pre>{<br>  "type": "public"<br>}</pre> | no |
| <a name="input_region"></a> [region](#input\_region) | AWS Region (e.g. `eu-central-1`) | `string` | `"eu-central-1"` | no |
| <a name="input_region_number"></a> [region\_number](#input\_region\_number) | n/a | `map` | <pre>{<br>  "ap-east-1": 5,<br>  "ap-southeast-1": 6,<br>  "ap-southeast-2": 7,<br>  "eu-central-1": 1,<br>  "eu-west-2": 2,<br>  "us-east-1": 3,<br>  "us-west-2": 4<br>}</pre> | no |
| <a name="input_s3_endpoint"></a> [s3\_endpoint](#input\_s3\_endpoint) | Enable / Disable VPC Endpoint (Gateway) - S3 | `bool` | `true` | no |
| <a name="input_single_nat_gateway"></a> [single\_nat\_gateway](#input\_single\_nat\_gateway) | Should be true if you want to provision a single shared NAT Gateway across all of your private networks (cost efficiency) | `bool` | `false` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | User specific Tags / Labels to attach to resources (will be merged with module tags) | `map(string)` | `{}` | no |
| <a name="input_vpc_cidr_add"></a> [vpc\_cidr\_add](#input\_vpc\_cidr\_add) | Add this number to the VPC CIDR to allow for multiple same region VPCs with different CIDRs | `number` | `0` | no |
| <a name="input_workspace_name"></a> [workspace\_name](#input\_workspace\_name) | Unique Workspace Name for naming / tagging | `string` | n/a | yes |
| <a name="input_zones"></a> [zones](#input\_zones) | Number of AWS Availability Zones to use for every subnet | `number` | `3` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_azs"></a> [azs](#output\_azs) | Number of Availability Zones specified |
| <a name="output_cache_azs"></a> [cache\_azs](#output\_cache\_azs) | Availability Zones for Cache Subnets |
| <a name="output_cache_rt_ids"></a> [cache\_rt\_ids](#output\_cache\_rt\_ids) | Cache Route Table IDs |
| <a name="output_cache_subnet_arns"></a> [cache\_subnet\_arns](#output\_cache\_subnet\_arns) | Cache Subnet ARNs |
| <a name="output_cache_subnet_cidrs"></a> [cache\_subnet\_cidrs](#output\_cache\_subnet\_cidrs) | Cache Subnet CIDRs |
| <a name="output_cache_subnet_ids"></a> [cache\_subnet\_ids](#output\_cache\_subnet\_ids) | Number of Cache Subnet IDs |
| <a name="output_cache_subnet_ipv6_cidr_blocks"></a> [cache\_subnet\_ipv6\_cidr\_blocks](#output\_cache\_subnet\_ipv6\_cidr\_blocks) | Cache Subnet IPv6 CIDR blocks |
| <a name="output_cache_subnets"></a> [cache\_subnets](#output\_cache\_subnets) | Number of Cache Subnet IDs |
| <a name="output_db_azs"></a> [db\_azs](#output\_db\_azs) | Availability Zones for Database Subnets |
| <a name="output_db_rt_ids"></a> [db\_rt\_ids](#output\_db\_rt\_ids) | Database Route Table IDs |
| <a name="output_db_subnet_arns"></a> [db\_subnet\_arns](#output\_db\_subnet\_arns) | Database Subnet ARNs |
| <a name="output_db_subnet_cidrs"></a> [db\_subnet\_cidrs](#output\_db\_subnet\_cidrs) | Database Subnet CIDRs |
| <a name="output_db_subnet_ids"></a> [db\_subnet\_ids](#output\_db\_subnet\_ids) | Number of Database Subnet IDs |
| <a name="output_db_subnet_ipv6_cidr_blocks"></a> [db\_subnet\_ipv6\_cidr\_blocks](#output\_db\_subnet\_ipv6\_cidr\_blocks) | Database Subnet IPv6 CIDR blocks |
| <a name="output_db_subnets"></a> [db\_subnets](#output\_db\_subnets) | Number of Database Subnet IDs |
| <a name="output_egress_only_internet_gateway_id"></a> [egress\_only\_internet\_gateway\_id](#output\_egress\_only\_internet\_gateway\_id) | ID of IPv6 Egress-Only Internet Gateway |
| <a name="output_eip_public_ips"></a> [eip\_public\_ips](#output\_eip\_public\_ips) | Public IPv4 of EIP addresses |
| <a name="output_igw_id"></a> [igw\_id](#output\_igw\_id) | Internet Gateway ID |
| <a name="output_internal_azs"></a> [internal\_azs](#output\_internal\_azs) | Availability Zones for Internal Subnets |
| <a name="output_internal_rt_ids"></a> [internal\_rt\_ids](#output\_internal\_rt\_ids) | Internal Route Table IDs |
| <a name="output_internal_subnet_arns"></a> [internal\_subnet\_arns](#output\_internal\_subnet\_arns) | Internal Subnet ARNs |
| <a name="output_internal_subnet_cidrs"></a> [internal\_subnet\_cidrs](#output\_internal\_subnet\_cidrs) | Internal Subnet CIDRs |
| <a name="output_internal_subnet_ids"></a> [internal\_subnet\_ids](#output\_internal\_subnet\_ids) | Number of Internal Subnet IDs |
| <a name="output_internal_subnet_ipv6_cidr_blocks"></a> [internal\_subnet\_ipv6\_cidr\_blocks](#output\_internal\_subnet\_ipv6\_cidr\_blocks) | Internal Subnet IPv6 CIDR blocks |
| <a name="output_internal_subnets"></a> [internal\_subnets](#output\_internal\_subnets) | Number of Internal Subnet IDs |
| <a name="output_nat_gateway_count"></a> [nat\_gateway\_count](#output\_nat\_gateway\_count) | Number of NAT Gateways |
| <a name="output_nat_gateway_public_ips"></a> [nat\_gateway\_public\_ips](#output\_nat\_gateway\_public\_ips) | Public IPs of NAT Gateways |
| <a name="output_number_azs"></a> [number\_azs](#output\_number\_azs) | Number of Availability Zones specified |
| <a name="output_private_azs"></a> [private\_azs](#output\_private\_azs) | Availability Zones for Private Subnets |
| <a name="output_private_rt_ids"></a> [private\_rt\_ids](#output\_private\_rt\_ids) | Private Route Table IDs |
| <a name="output_private_subnet_arns"></a> [private\_subnet\_arns](#output\_private\_subnet\_arns) | Private Subnet ARNs |
| <a name="output_private_subnet_cidrs"></a> [private\_subnet\_cidrs](#output\_private\_subnet\_cidrs) | Private Subnet CIDRs |
| <a name="output_private_subnet_ids"></a> [private\_subnet\_ids](#output\_private\_subnet\_ids) | Number of Private Subnet IDs |
| <a name="output_private_subnet_ipv6_cidr_blocks"></a> [private\_subnet\_ipv6\_cidr\_blocks](#output\_private\_subnet\_ipv6\_cidr\_blocks) | Private Subnet IPv6 CIDR blocks |
| <a name="output_private_subnets"></a> [private\_subnets](#output\_private\_subnets) | Number of Private Subnet IDs |
| <a name="output_public_azs"></a> [public\_azs](#output\_public\_azs) | Availability Zones for Public Subnets |
| <a name="output_public_rt_ids"></a> [public\_rt\_ids](#output\_public\_rt\_ids) | Public Route Table IDs |
| <a name="output_public_subnet_arns"></a> [public\_subnet\_arns](#output\_public\_subnet\_arns) | Public Subnet ARNs |
| <a name="output_public_subnet_cidrs"></a> [public\_subnet\_cidrs](#output\_public\_subnet\_cidrs) | Public Subnets CIDRs |
| <a name="output_public_subnet_ids"></a> [public\_subnet\_ids](#output\_public\_subnet\_ids) | Number of Public Subnet IDs |
| <a name="output_public_subnet_ipv6_cidr_blocks"></a> [public\_subnet\_ipv6\_cidr\_blocks](#output\_public\_subnet\_ipv6\_cidr\_blocks) | Public Subnet IPv6 CIDR blocks |
| <a name="output_public_subnets"></a> [public\_subnets](#output\_public\_subnets) | Number of Public Subnets |
| <a name="output_vpc_arn"></a> [vpc\_arn](#output\_vpc\_arn) | AWS VPC ID ARN |
| <a name="output_vpc_cidr"></a> [vpc\_cidr](#output\_vpc\_cidr) | AWS VPC IPv4 CIDR |
| <a name="output_vpc_dns"></a> [vpc\_dns](#output\_vpc\_dns) | VPC DNS Server IP |
| <a name="output_vpc_endpoint_s3_id"></a> [vpc\_endpoint\_s3\_id](#output\_vpc\_endpoint\_s3\_id) | ID of VPC endpoint for S3 |
| <a name="output_vpc_endpoint_s3_pl_id"></a> [vpc\_endpoint\_s3\_pl\_id](#output\_vpc\_endpoint\_s3\_pl\_id) | Prefix list ID for the S3 VPC endpoint |
| <a name="output_vpc_id"></a> [vpc\_id](#output\_vpc\_id) | AWS VPC ID |
| <a name="output_vpc_ipv6_association_id"></a> [vpc\_ipv6\_association\_id](#output\_vpc\_ipv6\_association\_id) | AWS VPC aassociation ID for the IPv6 CIDR block |
| <a name="output_vpc_ipv6_cidr"></a> [vpc\_ipv6\_cidr](#output\_vpc\_ipv6\_cidr) | AWS VPC IPv6 CIDR |
| <a name="output_zones"></a> [zones](#output\_zones) | Number of Availability Zones specified |
<!-- END_TF_DOCS -->



## Support

If you need professional support directly by the maintainers of the project, don't hesitate to contact us:
<a href="https://www.ventx.de/kontakt.html">
  <img align="center" src="https://i.imgur.com/OoCRUwz.png" alt="ventx Contact Us Kontakt" />
</a>

- [GitHub issues](https://github.com/ventx/terraform-aws-stackx-network/issues/new?assignees=&labels=question&template=04_SUPPORT_QUESTION.md&title=support%3A+)
- Contact options listed on [this GitHub profile](https://github.com/hajowieland)


## Project assistance

If you want to say **thank you** or/and support active development of terraform-aws-stackx-network:

- Add a [GitHub Star](https://github.com/ventx/terraform-aws-stackx-network) to the project.
- Tweet about the terraform-aws-stackx-network.
- Write interesting articles about the project on [Dev.to](https://dev.to/), [Medium](https://medium.com/) or your personal blog.

Together, we can make terraform-aws-stackx-network **better**!




## Contributing

First off, thanks for taking the time to contribute! Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make will benefit everybody else and are **greatly appreciated**.

Please read [our contribution guidelines](.github/CONTRIBUTING.md), and thank you for being involved!


## Security

terraform-aws-stackx-network follows good practices of security, but 100% security cannot be assured.
terraform-aws-stackx-network is provided **"as is"** without any **warranty**. Use at your own risk.

_For more information and to report security issues, please refer to our [security documentation](.github/SECURITY.md)._


## License

This project is licensed under the **Apache 2.0 license**.

See [LICENSE](LICENSE) for more information.


## Acknowledgements

* All open source contributors who made this possible


## Roadmap

See the [open issues](https://github.com/ventx/terraform-aws-stackx-network/issues) for a list of proposed features (and known issues).

- [Top Feature Requests](https://github.com/ventx/terraform-aws-stackx-network/issues?q=label%3Aenhancement+is%3Aopen+sort%3Areactions-%2B1-desc) (Add your votes using the 👍 reaction)
- [Top Bugs](https://github.com/ventx/terraform-aws-stackx-network/issues?q=is%3Aissue+is%3Aopen+label%3Abug+sort%3Areactions-%2B1-desc) (Add your votes using the 👍 reaction)
- [Newest Bugs](https://github.com/ventx/terraform-aws-stackx-network/issues?q=is%3Aopen+is%3Aissue+label%3Abug)


