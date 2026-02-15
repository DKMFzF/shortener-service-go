resource "yandex_vpc_security_group" "sg-nat" {
  name       = local.sg_nat_name
  network_id = yandex_vpc_network.main-vpc-network.id

  egress {
    protocol       = "any"
    description    = "any"
    v4_cidr_blocks = ["0.0.0.0/0"]
  }

  # dev endpoint
  ingress {
    protocol       = "tcp"
    description    = "ssh"
    v4_cidr_blocks = [var.ip_jumbox]
    port           = 22
  }

  #ingress {
  #protocol       = "tcp"
  #description    = "ext-http"
  #v4_cidr_blocks = [var.ip_jumbox]
  #port           = 80
  #}

  #ingress {
  #protocol       = "tcp"
  #description    = "ext-https"
  #v4_cidr_blocks = [var.ip_jumbox]
  #port           = 443
  #}

  ingress {
    protocol       = "tcp"
    description    = "Return traffic from private subnet"
    v4_cidr_blocks = ["192.168.2.0/24"]
  }
}
