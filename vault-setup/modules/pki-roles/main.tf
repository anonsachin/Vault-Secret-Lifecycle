resource "vault_pki_secret_backend_role" "fabric-client" {
  backend                            = var.backend_id
  name                               = var.role_name
  generate_lease                     = true
  allow_localhost                    = true
  allow_subdomains                   = true
  allowed_domains                    = var.base_domain
  server_flag                        = var.server_flag
  client_flag                        = var.client_flag
  key_usage                          = var.key_usage
  key_type                           = var.key_type
  key_bits                           = var.key_bits
  ou                                 = var.OU
  organization                       = [var.org_name]
  basic_constraints_valid_for_non_ca = var.non_ca_flag
}