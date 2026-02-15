#cloud-config

users:
  - name: ${sudo_name}
    groups: sudo
    shell: /bin/bash
    sudo: 'ALL=(ALL) NOPASSWD:ALL'
    ssh_authorized_keys:
      - ${ssh_public_key}

runcmd:
  - |
    #!/bin/bash
    set -e

    # Проверка подключения к сети
    until ping -c 1 8.8.8.8 >/dev/null 2>&1; do
      echo "Waiting for internet connectivity..."
      sleep 3
    done

    apt-get update

    # Установка утилит
    echo "Install utilities: fish, kitty, curl, make"
    apt-get install -y fish kitty curl make gnupg lsb-release software-properties-common

    # Установка Docker
    echo "Install Docker"
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
    echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null
    apt-get update
    apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

    systemctl enable docker
    systemctl start docker
    docker --version
    docker compose version

