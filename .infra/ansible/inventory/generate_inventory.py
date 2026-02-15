#!/usr/bin/env python3
import json
import subprocess

def get_terraform_output():
    cmd = ["terraform", "output", "-json"]
    result = subprocess.run(cmd, capture_output=True, text=True, cwd="../terraform")
    return json.loads(result.stdout)

def main():
    output = get_terraform_output()
    
    inventory = {
        "_meta": {
            "hostvars": {}
        },
        "docker_hosts": {
            "hosts": [],
            "vars": {
                "ansible_user": "ubuntu"
            }
        }
    }
    
    for i, ip in enumerate(output["vm_public_ips"]["value"]):
        hostname = output["vm_names"]["value"][i]
        inventory["docker_hosts"]["hosts"].append(hostname)
        inventory["_meta"]["hostvars"][hostname] = {
            "ansible_host": ip,
            "ansible_user": "ubuntu"
        }
    
    print(json.dumps(inventory, indent=2))

if __name__ == "__main__":
    main()

