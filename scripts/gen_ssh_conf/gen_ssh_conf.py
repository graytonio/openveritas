#!/usr/bin/env python
import requests
import configparser
import argparse
from urllib.parse import urljoin
from pathlib import Path

def parse_arguments():
    parser = argparse.ArgumentParser()
    parser.add_argument('--confirm', action="store_true", help="Write config to ~/.ssh/config")
    parser.add_argument('--list', action="store_true", help="Output config to stdout")
    return parser.parse_args()

def fetch_config(path):
    config = configparser.ConfigParser()
    config.read(path)
    return config

def fetch_nodes(host):
    response = requests.get(urljoin(host, "/prop/ip_address_0"))
    nodes = response.json()
    return nodes

def generate_ssh_config_file(nodes, user):
    content = ""
    for i in range(len(nodes)):
        content += "Host " + nodes[i]["node_name"] + "\n"
        content += "\tHostName " + nodes[i]["property_value"] + "\n"
        content += "\tUser " + user + "\n\n"
    return content

def write_to_ssh_config(content):
    f = open(Path.home().joinpath(".ssh/config"), "w")
    f.write(content)
    f.close()

def main():
    args = parse_arguments()

    if not args.list: 
        print("Generating new ssh conf")

    config = fetch_config(Path.home().joinpath(".veritasrc"))
    host = config["default"]["host"]
    nodes = fetch_nodes(host)
    content = generate_ssh_config_file(nodes, "ansible")

    if not args.confirm or args.list:
        print(content)
        return
    
    print("Writing to ssh config")
    write_to_ssh_config(content)

if __name__ == "__main__":
    main()