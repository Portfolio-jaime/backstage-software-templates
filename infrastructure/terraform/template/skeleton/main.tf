terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "s3" {
    bucket = "ba-terraform-state-${{values.app_env}}"
    key    = "${{values.app_name}}/terraform.tfstate"
    region = "eu-west-1"
  }
}

provider "aws" {
  region = var.aws_region

  default_tags {
    tags = {
      Project     = "${{values.app_name}}"
      Environment = "${{values.app_env}}"
      Owner       = "BA-DevOps"
      ManagedBy   = "Terraform"
    }
  }
}