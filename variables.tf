# Define config variables
variable "labelPrefix" {
  type        = string
  description = "Your college username. This will form the beginning of various resource names."
}

variable "region" {
  default = "canadacentral"
}

variable "admin_username" {
  type        = string
  default     = "azureadmin"
  description = "The username for the local user account on the VM."
}

variable "private_key_path" {
  description = "The path to the private SSH key used for VM authentication"
  type        = string
  default     = "~/.ssh/id_rsa"  # You can adjust the default path if necessary
}
