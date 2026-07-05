resource "random_id" "bucket" {

  byte_length = 4

}

resource "aws_s3_bucket" "assets" {

  bucket = "nimbus-assets-${random_id.bucket.hex}"

}