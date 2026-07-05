data "aws_iam_policy_document" "ec2_assume_role" {

  statement {

    actions = ["sts:AssumeRole"]

    principals {

      type = "Service"

      identifiers = ["ec2.amazonaws.com"]

    }

  }

}

resource "aws_iam_role" "ec2_role" {

  name = "nimbus-ec2-role"

  assume_role_policy = data.aws_iam_policy_document.ec2_assume_role.json

}

resource "aws_iam_role_policy_attachment" "ecr" {

  role = aws_iam_role.ec2_role.name

  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"

}

resource "aws_iam_role_policy_attachment" "s3" {

  role = aws_iam_role.ec2_role.name

  policy_arn = "arn:aws:iam::aws:policy/AmazonS3FullAccess"

}

resource "aws_iam_instance_profile" "profile" {

  name = "nimbus-instance-profile"

  role = aws_iam_role.ec2_role.name

}

resource "aws_key_pair" "nimbus" {

  key_name = "nimbus-key"

  public_key = file("~/.ssh/nimbus-key.pub")

}