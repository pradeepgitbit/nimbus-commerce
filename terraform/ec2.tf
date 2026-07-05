resource "aws_instance" "nimbus" {
  ami                    = "ami-01a00762f46d584a1"
  instance_type          = var.instance_type
  vpc_security_group_ids = [aws_security_group.nimbus.id]


  iam_instance_profile = aws_iam_instance_profile.profile.name

  key_name = aws_key_pair.nimbus.key_name


  tags = {
    Name = "Nimbus-Commerce"
  }
}