resource "aws_secretsmanager_secret" "credentials" {
  name = "baas-db-credentials"
}
resource "aws_secretsmanager_secret_version" "password" {
  secret_id     = aws_secretsmanager_secret.credentials.id
  secret_string = jsonencode(var.secret_string)
}
