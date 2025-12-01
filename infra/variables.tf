variable "aws_region" {
  description = "The AWS region to deploy resources (e.g., eu-west-1)."
  type        = string
  default     = "us-east-1"
}

variable "domain_name" {
  description = "The domain name for the website (e.g., example.com)."
  type        = string
}

variable "bucket_name" {
  description = "The name of the S3 bucket. Must be globally unique. If null, defaults to domain_name."
  type        = string
  default     = null
}

variable "common_tags" {
  description = "Common tags to apply to all resources."
  type        = map(string)
  default = {
    Project     = "StaticWebsite"
    Environment = "Production"
    ManagedBy   = "Terraform"
  }
}
