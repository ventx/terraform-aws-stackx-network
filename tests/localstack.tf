//
// Provider configuration for Localstack
// If running against real cloud provider, please delete/disable this file
//
provider "aws" {

  access_key = "mock_access_key"
  secret_key = "mock_secret_key"

  s3_use_path_style           = false
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  endpoints {
    ec2         = "http://localhost:4566"
    elasticache = "http://localhost:4566"
    rds         = "http://localhost:4566"
  }
}
