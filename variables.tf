variable "aws_region" {
  default   = ""
  sensitive = true
}
variable "aws_secret_key" {
  default   = ""
  sensitive = true
}
variable "aws_access_key" {
  default   = ""
  sensitive = true
}
variable "db_password" {
  description = "RDS root user password"
  type        = string
  sensitive   = true
}
variable "db_username" {
  description = "RDS root username"
  type        = string
  sensitive   = true
}

variable "db_parameter_group" {
  default = ""
}
variable "db_subnet_group" {

}
variable "engine" {
  default = ""
}
variable "engine_version" {
  default = ""
}
variable "db_name" {
}
variable "log_connections_value" {
  default = 1
}
variable "parameter_group_family" {
  default = "postgres"
}
variable "allocated_db_storage" {
  default = 5
}
variable "pvt_subnet_A_az" {
  default = ""
}
variable "pvt_subnet_B_az" {
  default = ""
}
variable "instance_class" {
  default = ""
}
variable "is_db_public" {
  default = false
}
variable "vpc_cidr" {
  default = "10.0.0.0/16"
}
variable "cidr_subnet_a" {
  default = "10.0.3.0/24"
}
variable "cidr_subnet_b" {
  default = "10.0.2.0/24"
}

variable "subnet_a_id" {

}
variable "subnet_b_id" {
}
variable "secret_string" {
  type = map(string)
}
