# ${{values.app_name}} Terraform Module
# Description: Reusable Terraform module for ${{values.app_name}}

terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# Module resources go here
resource "aws_s3_bucket" "module_bucket" {
  bucket = "${var.module_name}-${var.environment}-bucket"

  tags = var.common_tags
}

resource "aws_s3_bucket_versioning" "module_bucket_versioning" {
  bucket = aws_s3_bucket.module_bucket.id
  versioning_configuration {
    status = var.enable_versioning ? "Enabled" : "Suspended"
  }
}