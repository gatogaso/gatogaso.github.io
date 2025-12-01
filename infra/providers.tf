terraform {
  required_version = ">= 1.5.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# Primary Provider (Your default region)
provider "aws" {
  region = var.aws_region

  default_tags {
    tags = var.common_tags
  }
}

# Secondary Provider (MANDATORY for ACM Certificates)
# CloudFront certificates *must* be in us-east-1
provider "aws" {
  alias  = "us_east_1"
  region = "us-east-1"

  default_tags {
    tags = var.common_tags
  }
}
