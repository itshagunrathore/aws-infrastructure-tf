output "aws_secretsmanager_secret_id" {
  value = aws_secretsmanager_secret.password.id
}
output "rds_db_password" {
  value     = random_password.master.result
  sensitive = true
}
