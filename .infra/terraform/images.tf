resource "yandex_compute_image" "ubuntu-2404-lts" {
  source_family = "ubuntu-2404-lts"
}

resource "yandex_compute_image" "nat-instance-ubuntu" {
  source_family = "nat-instance-ubuntu"
}
