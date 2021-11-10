variable "backend_id" {
  description = "The id of the pki engine to link it to."
  type        = string
}

variable "role_name" {
  description = "The name of the vault role."
  type        = string
  default     = "client"
}

variable "base_domain" {
  description = "The base domain the certificate is to be linked to."
  type        = list(string)
  default     = ["service.consul"]
}

variable "key_type" {
  description = "The family of the encryption algorithm."
  type        = string
  default     = "ec"
}

variable "org_name" {
  description = "The name of the org to whom the CA belongs to."
  type        = string
}

variable "key_bits" {
  description = "The size of the encryption algorithm"
  type        = number
  default     = 256
}

variable "server_flag" {
  description = "Flag enabling TLS server."
  type        = bool
  default     = true
}
variable "client_flag" {
  description = "Flag enabling TLS client."
  type        = bool
  default     = true
}
variable "OU" {
  description = "The organisational unit the code belongs to."
  type        = list(string)
  default     = ["client"]
}
variable "non_ca_flag" {
  description = "The flag telling if the cert to be generated is CA or not."
  default     = true
}
variable "key_usage" {
  description = "The list of functions the certificate can do."
  type        = list(string)
  default     = ["DigitalSignature", "KeyAgreement", "KeyEncipherment"]
}