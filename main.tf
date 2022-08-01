module "vpc" {
  source   = "./modules/vpc"
  vpc_cidr = var.vpc_cidr
}

module "subnets" {
  source          = "./modules/subnets"
  vpc_id          = module.vpc.vpc_id
  cidr_subnet_a   = var.cidr_subnet_a
  cidr_subnet_b   = var.cidr_subnet_b
  pvt_subnet_A_az = var.pvt_subnet_A_az
  pvt_subnet_B_az = var.pvt_subnet_B_az
  db_subnet_group = var.db_subnet_group
}
module "secrets" {
  source = "./modules/secrets"
  db_username = var.db_username
}
module "rds" {
  source               = "./modules/rds"
  baas_db_name         = var.baas_db_name
  allocated_db_storage = var.allocated_db_storage
  instance_class       = var.instance_class
  engine               = var.engine
  engine_version       = var.engine_version
  db_subnet_group      = var.db_subnet_group
  db_password          = module.secrets.rds_db_password
  db_username          = var.db_username
  db_parameter_group   = var.db_parameter_group
}
