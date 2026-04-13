data "aws_vpc" "default" {
  default = true
}

data "aws_subnet" "public_subnet1" {
  availability_zone = "eu-central-1b"
  default_for_az    = true
}

data "aws_subnet" "public_subnet2" {
  availability_zone = "eu-central-1a"
  default_for_az    = true
}

data "aws_subnet" "public_subnet3" {
  availability_zone = "eu-central-1c"
  default_for_az    = true
}

locals {
  public_subnet_ids = [
    data.aws_subnet.public_subnet1.id,
    data.aws_subnet.public_subnet2.id,
    data.aws_subnet.public_subnet3.id,
  ]
}

resource "aws_ec2_tag" "subnet_elb" {
  for_each    = toset(local.public_subnet_ids)
  resource_id = each.value
  key         = "kubernetes.io/role/elb"
  value       = "1"
}

resource "aws_ec2_tag" "subnet_cluster" {
  for_each    = toset(local.public_subnet_ids)
  resource_id = each.value
  key         = "kubernetes.io/cluster/littlelink"
  value       = "shared"
}
