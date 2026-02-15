resource "yandex_compute_disk" "boot-disk-nat-vm" {
  name     = local.boot_disk_nat_vm_name
  type     = "network-ssd"
  zone     = "ru-central1-d"
  size     = "20"
  image_id = yandex_compute_image.nat-instance-ubuntu.id
}

resource "yandex_compute_instance" "nat-instance-vm" {
  name        = local.nat_instance_vm_name
  platform_id = "standard-v3"
  zone        = "ru-central1-d"

  resources {
    core_fraction = 20
    cores         = 2
    memory        = 2
  }

  boot_disk {
    disk_id = yandex_compute_disk.boot-disk-nat-vm.id
  }

  network_interface {
    subnet_id          = yandex_vpc_subnet.nat-subnet.id
    security_group_ids = [yandex_vpc_security_group.sg-nat.id]
    nat                = true
  }

  metadata = {
    ssh-keys = "ubuntu:${var.ssh_public_key}"
    user-data = templatefile(var.cloud_init_nat, {
      ssh_public_key = var.ssh_public_key
      sudo_name      = var.sudo_name
    })
  }
}
