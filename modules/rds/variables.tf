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

}

variable "engine" {
  default = ""
}
variable "engine_version" {
  default = ""
}
variable "baas_db_name" {
  default = ""
}
variable "log_connections_value" {
  default = 1
}
variable "parameter_group_family" {
  default = "postgres"
}
variable "db_subnet_group" {}
variable "allocated_db_storage" {
  default = 5
}
variable "instance_class" {
  default = ""
}
variable "is_db_public" {
  default = false
}
