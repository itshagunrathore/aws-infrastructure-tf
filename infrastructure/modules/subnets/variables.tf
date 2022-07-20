variable "db_subnet_group" {
  default = ""
}
variable "pvt_subnet_A_az" {
  default = ""
}
variable "pvt_subnet_B_az" {
  default = ""
}
variable "vpc_id" {
}

variable "cidr_subnet_a" {
  default = "10.0.3.0/24"
}
variable "cidr_subnet_b" {
  default = "10.0.2.0/24"
}
