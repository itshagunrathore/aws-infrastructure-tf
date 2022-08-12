
output "rds_db_password" {
  value     = random_password.master.result
  sensitive = true
}
