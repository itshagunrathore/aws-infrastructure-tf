#module "vpc" {
#  source   = "./modules/vpc"
#  vpc_cidr = var.vpc_cidr
#}

module "subnets" {
  source          = "./modules/subnets"
  #  vpc_id          = module.vpc.vpc_id
  subnet_a_id     = var.subnet_a_id
  subnet_b_id     = var.subnet_b_id
  db_subnet_group = var.db_subnet_group
  #  pvt_subnet_A_az = var.pvt_subnet_A_az
  #  pvt_subnet_B_az = var.pvt_subnet_B_az
  #  db_subnet_group = var.db_subnet_group
  #  subnet_ids = [
  #    var.cidr_subnet_a,
  #    var.cidr_subnet_b
  #  ]
  #  vpc_id =
}
module "utils" {
  source = "./modules/utils"
}

module "rds" {
  source               = "./modules/rds"
  db_name              = var.db_name
  allocated_db_storage = var.allocated_db_storage
  instance_class       = var.instance_class
  engine               = var.engine
  engine_version       = var.engine_version
  db_subnet_group      = module.subnets.db_subnet_group
  db_password          = module.utils.rds_db_password
  db_username          = var.db_username
  db_parameter_group   = var.db_parameter_group
}
module "secrets" {
  source        = "./modules/secrets"
  secret_string = {
    username = var.db_username
    password = module.utils.rds_db_password
  }
}
