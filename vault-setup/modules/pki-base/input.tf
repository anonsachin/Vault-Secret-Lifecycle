
variable "key_type" {
  description = "The family of the encryption algorithm."
  type        = string
  default     = "ec"
}
variable "ca_type" {
  description = "The type of ca for the PKI"
  type        = string
  default     = "internal"
}

variable "key_bits" {
  description = "The size of the encryption algorithm"
  type        = number
  default     = 256
}
variable "ttl" {
  description = "The life span of the root CA cert."
  type        = string
  default     = "87600h"
}

variable "ca_name" {
  description = "The name of the ca."
  type        = string
}
variable "engine" {
  description = "The type of the vault engine being created."
  type        = string
  default     = "pki"
}