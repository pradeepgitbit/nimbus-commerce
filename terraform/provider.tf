provider "aws" {

  region = var.aws_region

  default_tags {

    tags = {

      Project = "Nimbus-Commerce"

      ManagedBy = "Terraform"

      Environment = "Dev"

    }

  }

}