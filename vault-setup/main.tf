terraform {
  backend "local" {
    path = "terraform.tfstate"
  }
}

provider "vault" {
  address = "http://127.0.0.1:8200"
  token   = "myroot"
}

locals {
  ca_name = "${var.org_name}CA"
}

module "Base" {
  source   = "./modules/pki-base"
  ca_name  = local.ca_name
  org_name = var.org_name
  key_bits = var.key_bits
  key_type = var.key_type
}

module "Test_certs" {
  source     = "./modules/pki-roles"
  backend_id = module.Base.engine_id
  org_name   = var.org_name
  key_bits   = var.key_bits
  key_type   = var.key_type
}
