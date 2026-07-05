locals {

  repositories = [

    "nimbus-product",

    "nimbus-inventory",

    "nimbus-gateway"

  ]

}

resource "aws_ecr_repository" "repos" {

  for_each = toset(local.repositories)

  name = each.value

  image_scanning_configuration {

    scan_on_push = true

  }

}