resource "yandex_vpc_network" "main-vpc-network" {
  name = local.main_vpc_network_name
}

resource "yandex_vpc_subnet" "prod-subnet" {
  name = local.prod_subnet_name
  zone = "ru-central1-d"

  network_id = yandex_vpc_network.main-vpc-network.id

  v4_cidr_blocks = ["192.168.3.0/24"]
}

resource "yandex_vpc_subnet" "nat-subnet" {
  name = local.nat_subnet_name
  zone = "ru-central1-d"

  network_id = yandex_vpc_network.main-vpc-network.id

  v4_cidr_blocks = ["192.168.1.0/24"]
}

