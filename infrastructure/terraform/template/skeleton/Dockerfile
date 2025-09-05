# Terraform Docker image for CI/CD
FROM hashicorp/terraform:1.5

WORKDIR /workspace

# Copy Terraform files
COPY *.tf ./
COPY terraform.tfvars.example ./

# Initialize Terraform
RUN terraform init

# Default command
CMD ["terraform", "--help"]

