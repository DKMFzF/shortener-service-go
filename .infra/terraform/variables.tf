variable "ip_jumbox" {
  description = "Ip address jumbox machine"
  type        = string
}

variable "yc_token_id" {
  description = "Yandex account tocken in yc"
  type        = string
}

variable "ssh_public_key" {
  description = "SSH public key to use for VM access"
  type        = string
}

variable "folder_id" {
  description = "Folder in yandex cloud"
  type        = string
}

variable "cloud_id" {
  description = "Yandex cloud id"
  type        = string
}

variable "sudo_name" {
  description = "Username Admin account"
  type        = string
}

variable "cloud_init_nat" {
  description = "Path to nat preset-config"
  default     = "./cloud-init/nat.yml.tpl"
  type        = string
}

variable "cloud_init_privat_vm" {
  description = "Path to privat-vm preset-config"
  default     = "./cloud-init/meta.yml.tpl"
  type        = string
}

variable "cloud_init_prod_vm" {
  description = "Path to prod-vm preset-config"
  default     = "./cloud-init/prod.yml.tpl"
  type        = string
}

locals {
  main_vpc_network_name   = "main-vpc-network"
  nat_subnet_name         = "nat-subnet"
  route_table_prod_name = "route-table-prod"
}

