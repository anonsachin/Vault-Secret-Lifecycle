variable "org_name" {
  description = "The name of the org for which we are setting the PKI"
  type        = string
  default     = "NewOrg"
}

variable "base_domain" {
  description = "The base domain for peers"
  type        = list(string)
  default     = ["service.consul"]
}

variable "key_type" {
  description = "The family of the encryption algorithm."
  type        = string
  default     = "ec"
}

variable "key_bits" {
  description = "The size of the encryption algorithm"
  type        = number
  default     = 256
}