resource "aws_s3_bucket" "example" {
  bucket = "${var.module_name}-bucket"
}
