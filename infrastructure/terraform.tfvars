engine                 = "postgres"
engine_version         = "13.1"
db_username            = "admin"
db_password            = "Admin@123"
baas_db_name           = "baasdb"
aws_region             = "us-west-2"
db_subnet_group        = "baas_db_subnet_group"
instance_class         = "db.t3.micro"
allocated_db_storage   = 5
pvt_subnet_A_az        = "us-west-2-a"
pvt_subnet_B_az        = "us-west-2-b"
log_connections_value  = 1
parameter_group_family = "postgres"
is_db_public           = false
db_parameter_group     = "baas_db_parameter_group"
vpc_cidr               = "10.0.3.0/24"
cidr_subnet_a          = "10.0.2.0/24"
cidr_subnet_b          = "10.0.0.0/16"
