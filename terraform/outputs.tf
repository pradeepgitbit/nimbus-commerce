output "ec2_public_ip" {
  value = aws_instance.nimbus.public_ip
}

output "ec2_public_dns" {
  value = aws_instance.nimbus.public_dns
}

output "s3_bucket" {
  value = aws_s3_bucket.assets.bucket
}

output "ecr_repositories" {
  value = {
    for name, repo in aws_ecr_repository.repos :
    name => repo.repository_url
  }
}