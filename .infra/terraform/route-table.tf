resource "yandex_vpc_route_table" "route-table" {
  name       = local.route_table_name
  network_id = yandex_vpc_network.main-vpc-network.id
  static_route {
    destination_prefix = "0.0.0.0/0"
    next_hop_address   = yandex_compute_instance.nat-instance-vm.network_interface.0.ip_address
  }
}

