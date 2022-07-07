variable "dev_db_password" {
  description = "RDS root user password"
  type        = string
  sensitive   = true
}
variable "dev_db_username" {
  description = "RDS root username"
  type        = string
  sensitive   = true
}
variable "dev_aws_region" {
  default = ""
}
variable "dev_aws_secret_key" {
  default = ""
}
variable "dev_aws_access_key" {
  default = ""
}
variable "dev_engine" {
  default = ""
}
variable "dev_engine_version" {
  default = ""
}
variable "dev_baas_db_name" {
  default = ""
}
variable "dev_vpc_cidr" {
  default = "10.0.0.0/16"

}
variable "dev_cidr_subnet_a" {
  default = "10.0.3.0/24"
}
variable "dev_cidr_subnet_b" {
  default = "10.0.2.0/24"
}
