resource "vault_mount" "pki_engine" {
  path = var.ca_name
  type = var.engine
}


resource "vault_pki_secret_backend_root_cert" "ca" {
  backend      = vault_mount.pki_engine.id
  common_name  = var.ca_name
  type         = var.ca_type
  organization = var.org_name
  ttl          = var.ttl
  key_type     = var.key_type
  key_bits     = var.key_bits
}