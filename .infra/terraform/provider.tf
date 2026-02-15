terraform {
  required_providers {
    yandex = {
      source = "yandex-cloud/yandex"
    }
  }
  required_version = ">= 0.13"
}

provider "yandex" {
  zone      = "ru-central1-d"
  token     = var.yc_token_id
  cloud_id  = var.cloud_id
  folder_id = var.folder_id
}


