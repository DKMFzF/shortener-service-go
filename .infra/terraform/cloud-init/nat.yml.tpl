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

