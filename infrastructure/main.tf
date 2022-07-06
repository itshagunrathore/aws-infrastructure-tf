resource "aws_vpc" "baas_vpc" {
  cidr_block = "10.0.0.0/16"
}
resource "aws_subnet" "baas_vpc_pvt_subnetA" {
  vpc_id            = aws_vpc.baas_vpc.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "us-west-1-a"
}
resource "aws_subnet" "baas_vpc_pvt_subnetB" {
  vpc_id            = aws_vpc.baas_vpc.id
  cidr_block        = "10.0.3.0/24"
  availability_zone = "us-west-2-b"
}


resource "aws_db_subnet_group" "baas_db_subnet_group" {
  name       = "baas_db_subnet_group"
  subnet_ids = [
    aws_subnet.baas_vpc_pvt_subnetA.id,
    aws_subnet.baas_vpc_pvt_subnetB.id]
  tags       = {
    Name = "baas_db"
  }
}


resource "aws_db_instance" "baas_db" {
  identifier           = var.baas_db_name
  instance_class       = "db.t3.micro"
  allocated_storage    = 5
  engine               = var.engine
  engine_version       = var.engine_version
  username             = var.db_username
  password             = var.db_password
  db_subnet_group_name = aws_db_subnet_group.baas_db_subnet_group.name
  parameter_group_name = aws_db_parameter_group.baas_pg.name
  publicly_accessible  = false
}

resource "aws_db_parameter_group" "baas_pg" {
  name   = "baaspg"
  family = "postgres13"
  parameter {
    name  = "log_connections"
    value = "1"
  }
}



