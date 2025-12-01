output "s3_bucket_name" {
  description = "The name of the S3 bucket where you should upload your files."
  value       = aws_s3_bucket.website.id
}

output "cloudfront_domain_name" {
  description = "The internal CloudFront domain name."
  value       = aws_cloudfront_distribution.s3_distribution.domain_name
}

output "website_url" {
  description = "The full URL of your website."
  value       = "https://${var.domain_name}"
}
