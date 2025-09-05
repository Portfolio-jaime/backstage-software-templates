variable "module_name" {
  description = "Name of the module"
  type        = string
  default     = "${{values.app_name}}"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "${{values.app_env}}"
}

variable "enable_versioning" {
  description = "Enable S3 bucket versioning"
  type        = bool
  default     = true
}

variable "common_tags" {
  description = "Common tags to apply to all resources"
  type        = map(string)
  default = {
    Module      = "${{values.app_name}}"
    Environment = "${{values.app_env}}"
    ManagedBy   = "Terraform"
  }
}