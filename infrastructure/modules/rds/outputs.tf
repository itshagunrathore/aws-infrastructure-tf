output "rds_hostname" {
  description = "RDS instance hostname"
  value       = aws_db_instance.baas-db.address
  sensitive   = true
}

output "rds_port" {
  description = "RDS instance port"
  value       = aws_db_instance.baas-db.port
  sensitive   = true
}

output "rds_username" {
  description = "RDS instance root username"
  value       = aws_db_instance.baas-db.username
  sensitive   = true
}
output "aws_db_parameter_group" {
  description = "PG name"
  value       = aws_db_parameter_group.baas-pg.name
}


