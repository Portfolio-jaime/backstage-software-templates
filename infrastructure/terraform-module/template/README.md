# ${{values.module_name}}

${{values.description}}

## Usage

```hcl
module "${{values.module_name}}" {
  source = "."
  
  environment  = "dev"
  project_name = "${{values.module_name}}"
}
```

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 1.0 |
| aws | ~> 5.0 |

## Providers

| Name | Version |
|------|---------|
| aws | ~> 5.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| environment | Environment name | `string` | `"dev"` | no |
| project_name | Name of the project | `string` | `"${{values.module_name}}"` | no |

## Outputs

| Name | Description |
|------|-------------|
| bucket_name | Name of the S3 bucket |
| bucket_arn | ARN of the S3 bucket |