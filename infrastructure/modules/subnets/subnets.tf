resource "aws_subnet" "baas_vpc_pvt_subnetA" {
  vpc_id            = var.vpc_id
  cidr_block        = var.cidr_subnet_a
  availability_zone = var.pvt_subnet_A_az
}
resource "aws_subnet" "baas_vpc_pvt_subnetB" {
  vpc_id            = var.vpc_id
  cidr_block        = var.cidr_subnet_b
  availability_zone = var.pvt_subnet_B_az
}
resource "aws_db_subnet_group" "baas_db_subnet_group" {
  name       = var.db_subnet_group
  subnet_ids = [
    aws_subnet.baas_vpc_pvt_subnetA.id,
    aws_subnet.baas_vpc_pvt_subnetB.id]
  tags       = {
    Name = var.db_subnet_group
  }
}
