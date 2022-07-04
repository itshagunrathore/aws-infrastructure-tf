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
variable "aws_region" {
  default = ""
}
variable "aws_secret_key" {
  default = ""
}
variable "aws_access_key" {
  default = ""
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
