resource "aws_vpc" "baas_vpc" {
  cidr_block = var.vpc_cidr
}
resource "aws_subnet" "baas_vpc_pvt_subnetA" {
  vpc_id            = aws_vpc.baas_vpc.id
  cidr_block        = var.cidr_subnet_a
  availability_zone = var.pvt_subnet_A_az
}
resource "aws_subnet" "baas_vpc_pvt_subnetB" {
  vpc_id            = aws_vpc.baas_vpc.id
  cidr_block        = var.cidr_subnet_b
  availability_zone = var.pvt_subnet_B_az
}
resource "aws_db_subnet_group" "baas_db_subnet_group" {
  name       = var.db_subnet_group
  subnet_ids = [
    aws_subnet.baas_vpc_pvt_subnetA.id,
    aws_subnet.baas_vpc_pvt_subnetB.id]
  tags       = {
    Name = aws_db_subnet_group.baas_db_subnet_group.name
  }
}
resource "aws_db_instance" "baas_db" {
  identifier           = var.baas_db_name
  instance_class       = var.instance_class
  allocated_storage    = var.allocated_db_storage
  engine               = var.engine
  engine_version       = var.engine_version
  username             = var.db_username
  password             = var.db_password
  db_subnet_group_name = aws_db_subnet_group.baas_db_subnet_group.name
  parameter_group_name = aws_db_parameter_group.baas_pg.name
  publicly_accessible  = var.is_db_public
}

resource "aws_db_parameter_group" "baas_pg" {
  name   = var.db_parameter_group
  family = var.parameter_group_family
  parameter {
    name  = "log_connections"
    value = var.log_connections_value
  }
}



