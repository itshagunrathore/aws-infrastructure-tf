resource "random_password" "master" {
  length           = 8
  special          = true
  override_special = "_!%^"
}
