resource "aws_db_parameter_group" "baas-pg" {
  name   = var.db_parameter_group
  family = var.parameter_group_family
  parameter {
    name  = "log_connections"
    value = var.log_connections_value
  }
}

resource "aws_db_instance" "baas-db" {
  identifier           = var.baas_db_name
  instance_class       = var.instance_class
  allocated_storage    = var.allocated_db_storage
  engine               = var.engine
  engine_version       = var.engine_version
  username             = var.db_username
  password             = var.db_password
  db_subnet_group_name = var.db_subnet_group
  parameter_group_name = aws_db_parameter_group.baas-pg.name
  publicly_accessible  = var.is_db_public
}
resource "random_password" "master"{
  length           = 16
  special          = true
  override_special = "_!%^"
}



